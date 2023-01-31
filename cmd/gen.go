package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/sj-distributor/dolphin/model"
	"github.com/urfave/cli"
)

var genCmd = cli.Command{
	Name:  "generate",
	Usage: "generate contents",
	Action: func(ctx *cli.Context) error {
		if err := generate("*.graphql", "."); err != nil {
			return cli.NewExitError(err, 1)
		}
		return nil
	},
}

var fileName string = ""

func generate(fileDirPath, p string) error {
	fileDirPath = path.Join(p, "./model/"+fileDirPath)
	matches, err := filepath.Glob(fileDirPath)
	if err != nil {
		return err
	}

	if len(matches) > 0 {
		name := matches[0]
		name = strings.Replace(name, ".graphql", "", -1)
		arr := strings.Split(name, "-")
		if len(arr) > 1 {
			fileName = arr[1]
		}
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

	schema, err := model.PrintSchema(m)
	if err != nil {
		return err
	}

	schema = "# This schema is generated, please don't update it manually\n\n" + schema

	if err := ioutil.WriteFile(path.Join(p, "gen/schema.graphql"), []byte(schema), 0644); err != nil {
		return err
	}

	fmt.Printf("Running gqlgen generator in %s ...\n", path.Join(p, "gen"))

	packageArray := []string{
		"go run github.com/99designs/gqlgen",
	}
	packageToStr := strings.Join(packageArray, " && ")
	cmd := exec.Command("sh", "-c", "cd "+path.Join(p, "gen")+" && "+packageToStr)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func generateFiles(p string, m *model.Model, c *model.Config) error {
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
