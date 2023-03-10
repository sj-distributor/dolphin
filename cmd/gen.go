package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sj-distributor/dolphin/model"
	"github.com/sj-distributor/dolphin/templates"
	"github.com/urfave/cli"
)

var fileName string = ""

var genCmd = cli.Command{
	Name:  "generate",
	Usage: "generate contents",
	Action: func(ctx *cli.Context) error {
		if err := generate("*.graphql", "."); err != nil {
			if err := DownDolphinPackage(); err != nil {
				return cli.NewExitError(err, 1)
			}
			return cli.NewExitError(err, 1)
		}

		return DownDolphinPackage()
	},
}

func DownDolphinPackage() error {
	cmd := exec.Command("sh", "-c", "go get -d github.com/sj-distributor/dolphin")
	return cmd.Run()
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

	modelSource := ""
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

	genPath := path.Join(p, "gen")
	ensureDir(genPath)

	err = model.EnrichModelObjects(&m)
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

	if err := ioutil.WriteFile(path.Join(p, "gen/schema.graphql"), []byte(schema), 0644); err != nil {
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

	cmd := exec.Command("sh", "-c", "go get github.com/99designs/gqlgen && go generate ./...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func generateFiles(p string, m *model.Model, c *model.Config) error {
	data := templates.TemplateData{Model: m, Config: c}

	if err := templates.WriteTemplate(templates.Model, path.Join(p, "gen/models.go"), data); err != nil {
		return err
	}
	return nil
}

func ensureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}
