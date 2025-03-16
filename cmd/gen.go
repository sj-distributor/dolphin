package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sj-distributor/dolphin/model"
	"github.com/sj-distributor/dolphin/templates"
	"github.com/sj-distributor/dolphin/tools"
	"github.com/urfave/cli"
)

var fileName string = ""

var genCmd = cli.Command{
	Name:  "generate",
	Usage: "generate contents",
	Action: func(ctx *cli.Context) error {
		return generate("*.graphql", ".")
	},
}

func generate(fileDirPath, p string) error {
	fileDirPath = path.Join(p, "./model/"+fileDirPath)
	matches, err := filepath.Glob(fileDirPath)
	if err != nil {
		return err
	}

	if len(matches) <= 0 {
		return fmt.Errorf("model files is empty")
	}

	name := matches[0]
	name = strings.Replace(name, ".graphql", "", -1)
	arr := strings.Split(name, "-")
	if len(arr) > 1 {
		fileName = arr[1]
	}

	fmt.Println("Generating contents from", matches, "...")
	modelSource := `
		directive @format on FIELD_DEFINITION
		directive @validator(required: String, immutable: String, type: String, minLength: Int, maxLength: Int, minValue: Int, maxValue: Int) on INPUT_FIELD_DEFINITION
		input FileField {
			hash: String!
			file: Upload!
		}
	`
	for _, file := range matches {
		fmt.Println("Appending content from model file", file)
		source, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		modelSource += string(source) + "\n"
	}

	m, err := model.Parse(modelSource)
	if err != nil {
		return err
	}

	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}

	ensureDir(path.Join(p, "gen"))

	err = model.EnrichModelObjects(&m)
	if err != nil {
		return err
	}

	if err := createUtilsFile(p); err != nil {
		return cli.NewExitError(err, 1)
	}

	// 接口
	err = generateInterface(p, &m, &c)
	if err != nil {
		return err
	}

	// 接口文档
	err = generateInterfaceDocument(p, &m, &c)
	if err != nil {
		return err
	}

	err = generateFiles(p, &m, &c)
	if err != nil {
		return err
	}

	err = model.EnrichModel(&m)
	if err != nil {
		return err
	}

	schemaSDL, err := model.PrintSchema(m)
	if err != nil {
		return err
	}

	// err = model.BuildFederatedModel(&m)
	// if err != nil {
	// 	return err
	// }

	schema, err := model.PrintSchema(m)
	if err != nil {
		return err
	}

	schema = "# This schema is generated, please don't update it manually\n\n" + schema

	if err := ioutil.WriteFile(path.Join(p, "gen/schema.graphqls"), []byte(schema), 0644); err != nil {
		return err
	}

	var re = regexp.MustCompile(`(?sm)schema {[^}]+}`)
	schemaSDL = re.ReplaceAllString(schemaSDL, ``)
	var re2 = regexp.MustCompile(`(?sm)type _Service {[^}]+}`)
	schemaSDL = re2.ReplaceAllString(schemaSDL, ``)
	schemaSDL = strings.Replace(schemaSDL, "\n  _service: _Service!", "", 1)
	schemaSDL = strings.Replace(schemaSDL, "\n  _entities(representations: [_Any!]!): [_Entity]!", "", 1)
	schemaSDL = strings.Replace(schemaSDL, "\nscalar _Any", "", 1)
	var re3 = regexp.MustCompile(`(?sm)[\n]{3,}`)
	schemaSDL = re3.ReplaceAllString(schemaSDL, "\n\n")
	schemaSDL = strings.Trim(schemaSDL, "\n")
	constants := map[string]interface{}{
		"SchemaSDL": schemaSDL,
	}

	if err := templates.WriteTemplateRaw(templates.Constants, path.Join(p, "gen/constants.go"), constants); err != nil {
		return err
	}

	fmt.Printf("Running gqlgen generator in %s ...\n", path.Join(p, "gen"))

	if err := tools.RunInteractiveInDir("go mod tidy && go run github.com/99designs/gqlgen", path.Join(p, "gen")); err != nil {
		return err
	}

	return nil
}

// 生成前端接口接口
func generateInterface(p string, m *model.Model, c *model.Config) error {
	data := templates.TemplateData{Model: m, Config: c}
	return templates.WriteInterfaceTemplate(templates.Graphql, path.Join(p, "docs/api.gql"), data)
}

// 生成前端接口接口文档
func generateInterfaceDocument(p string, m *model.Model, c *model.Config) error {
	data := templates.TemplateData{Model: m, Config: c}
	return templates.WriteInterfaceTemplate(templates.GraphqlApi, path.Join(p, "docs/api.json"), data)
}

func generateFiles(p string, m *model.Model, c *model.Config) error {
	data := templates.TemplateData{Model: m, Config: c}
	if err := templates.WriteTemplate(templates.GQLGen, path.Join(p, "gqlgen.yml"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Database, path.Join(p, "gen/database.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Model, path.Join(p, "gen/models.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.QueryFilters, path.Join(p, "gen/query-filters.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Sorting, path.Join(p, "gen/sorting.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Filters, path.Join(p, "gen/filters.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Loaders, path.Join(p, "gen/loaders.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.HttpHandler, path.Join(p, "gen/http-handler.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverCore, path.Join(p, "gen/resolver.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverQueries, path.Join(p, "gen/resolver-queries.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverMutations, path.Join(p, "gen/resolver-mutations.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverSubscriptions, path.Join(p, "gen/resolver-subscriptions.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResultType, path.Join(p, "gen/result-type.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverUtils, path.Join(p, "gen/resolver-utils.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverEvents, path.Join(p, "gen/events.go"), data); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.ResolverSockets, path.Join(p, "gen/resolver-socket.go"), data); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.ResolverSrcGen, path.Join(p, "src/resolver_gen.go"), data); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.ResolverSrcContext, path.Join(p, "src/context.go"), data); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.Validator, path.Join(p, "utils/validator.go"), data); err != nil {
		return err
	}

	// html
	if err := templates.WriterOriginalFile(templates.HandlerHtml, path.Join(p, "gen/html.go")); err != nil {
		return err
	}

	return nil
}

func createUtilsFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "utils"))
	if err := templates.WriteTemplate(templates.ResolverSrcUtils, path.Join(p, "utils/utils.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}
	if err := templates.WriteTemplate(templates.Rule, path.Join(p, "utils/rule.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.Encrypt, path.Join(p, "utils/encrypt.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.Rsa, path.Join(p, "utils/rsa.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func ensureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}
