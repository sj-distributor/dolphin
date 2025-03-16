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

	"github.com/iancoleman/strcase"
	"github.com/sj-distributor/dolphin/model"
	"github.com/sj-distributor/dolphin/templates"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var initCmd = cli.Command{
	Name:  "init",
	Usage: "initialize new project",
	Action: func(ctx *cli.Context) error {
		p := "."

		fmt.Printf("Initializing project in %s\n", p)

		if !fileExists(path.Join(p, model.YmlFileName)) {
			if err := createConfigFile(p, ctx.Args().First()); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if !fileExists(path.Join(p, "model.graphql")) {
			if err := createDummyModelFile(p); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if err := createGitignoreFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createMainFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createAuthFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createSrcFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createMiddlewareFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createEnumsFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		if err := createMakeFile(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		// if err := createDockerFile(p); err != nil {
		// 	return cli.NewExitError(err, 1)
		// }

		if !fileExists(path.Join(p, "go.mod")) {
			if err := initModules(p); err != nil {
				return cli.NewExitError(err, 1)
			}
		}

		if err := runGenerate(p); err != nil {
			return cli.NewExitError(err, 1)
		}

		cmd := exec.Command("sh", "-c", "go mod tidy")
		if err := cmd.Run(); err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	},
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	}
	return false
}

func createConfigFile(p, isAuto string) error {
	defaultPackagep, err := getDefaultPackageName(p)
	if err != nil {
		return err
	}

	if isAuto != "auto" {
		packagep := templates.Prompt(fmt.Sprintf("Package p (default: %s)", defaultPackagep))
		if packagep != "" {
			defaultPackagep = packagep
		}
	}

	c := model.Config{Package: defaultPackagep}

	content, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(p, model.YmlFileName), content, 0644)
	return err
}

func createGitignoreFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	return templates.WriteTemplate(templates.Gitignore, path.Join(p, ".gitignore"), templates.TemplateData{Config: &c})
}

func createMainFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	return templates.WriteTemplate(templates.Main, path.Join(p, "main.go"), templates.TemplateData{Config: &c})
}

func createDummyModelFile(p string) error {
	data := templates.TemplateData{Model: nil, Config: nil}
	ensureDir(path.Join(p, "model"))

	if err := templates.WriteTemplate(templates.DummyModel, path.Join(p, "model/test.graphql"), data); err != nil {
		return err
	}

	// if err := templates.WriteTemplate(templates.UploadModel, path.Join(p, "model/upload.graphql"), data); err != nil {
	// 	return err
	// }

	return nil
}

func createMakeFile(p string) error {
	defaultPackagep, err := getDefaultPackageName(p)
	if err != nil {
		return err
	}
	databaseName := extractRepoName(defaultPackagep)

	data := templates.TemplateData{Model: nil, Config: nil}
	return templates.WriteTemplate(templates.Makefile(databaseName), path.Join(p, "makefile"), data)
}

// func createDockerFile(p string) error {
// 	c, err := model.LoadConfigFromPath(p)
// 	if err != nil {
// 		return err
// 	}
// 	data := templates.TemplateData{Model: nil, Config: &c}
// 	return templates.WriteTemplate(templates.Dockerfile, path.Join(p, "Dockerfile"), data)
// }

func initModules(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	return templates.RunInteractiveInDir(fmt.Sprintf("go mod init %s", c.Package), p)
}

func createResolverFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}

	ensureDir(path.Join(p, "src"))

	if err := templates.WriteTemplate(templates.ResolverSrc, path.Join(p, "src/resolver.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func createAuthFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "auth"))

	if err := templates.WriteTemplate(templates.AuthRouter, path.Join(p, "auth/auth_router.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	if err := templates.WriteTemplate(templates.AuthOpenRouters, path.Join(p, "auth/open_routes.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func createSrcFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "src"))

	// if err := templates.WriteTemplate(templates.UpLoad, path.Join(p, "src/upload.go"), templates.TemplateData{Config: &c}); err != nil {
	// 	return err
	// }

	if err := templates.WriteTemplate(templates.ResolverSrc, path.Join(p, "src/resolver.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func createMiddlewareFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "src/middleware"))

	if err := templates.WriteTemplate(templates.MiddlewareHandler, path.Join(p, "src/middleware/handler.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func createEnumsFile(p string) error {
	c, err := model.LoadConfigFromPath(p)
	if err != nil {
		return err
	}
	ensureDir(path.Join(p, "enums"))

	if err := templates.WriteTemplate(templates.EnumsConst, path.Join(p, "enums/constants.go"), templates.TemplateData{Config: &c}); err != nil {
		return err
	}

	return nil
}

func runGenerate(p string) error {
	return generate("*.graphql", p)
}

func getDefaultPackageName(p string) (string, error) {
	configSource, err := ioutil.ReadFile(path.Join(p, "go.mod"))
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("module .*")
	defaultPackagep := strings.Replace(re.FindString(string(configSource)), "module ", "", -1)

	if os.Getenv("GOp") != "" {
		cw, _ := os.Getwd()
		defaultPackagep, _ = filepath.Rel(os.Getenv("GOp")+"/src", cw)
	}

	if defaultPackagep == "" {
		defaultPackagep = "github.com/dolphin/graphql-test"
	}

	return defaultPackagep, nil
}

func extractRepoName(url string) string {
	index := strings.LastIndex(url, "/")

	if index == -1 || index == len(url)-1 {
		return ""
	}

	repoName := url[index+1:]
	return strcase.ToSnake(repoName)
}
