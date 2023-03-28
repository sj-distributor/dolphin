package templates

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"

	"github.com/sj-distributor/dolphin/model"
	"github.com/urfave/cli"
)

type TemplateData struct {
	Model     *model.Model
	Config    *model.Config
	RawSchema *string
}

func WriteTemplate(t, filename string, data TemplateData) error {
	return WriteTemplateRaw(t, filename, data)
}

func WriteTemplateRaw(t, filename string, data interface{}) error {
	temp, err := template.New(filename).Parse(t)
	if err != nil {
		return err
	}

	var content bytes.Buffer
	writer := io.Writer(&content)

	err = temp.Execute(writer, &data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, content.Bytes(), 0777)
	if err != nil {
		return err
	}

	if path.Ext(filename) == ".go" {
		err := RunInteractive(fmt.Sprintf("$GOPATH/bin/goimports -w %s", filename))
		if err != nil {
			fmt.Println("goimports is error", fmt.Sprintf("$GOPATH/bin/goimports -w %s", filename))
		}
		return err
	}
	return nil
}

// RunInteractiveInDir ...
func RunInteractiveInDir(cmd, dir string) error {
	if os.Getenv("DEBUG") != "" {
		log.Println(cmd)
	}

	// command := exec.Command("sh", "-c", "set -o pipefail && " + cmd)
	command := exec.Command("sh", "-c", cmd)
	if err := command.Run(); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

// RunInteractive ...
func RunInteractive(cmd string) error {
	return RunInteractiveInDir(cmd, "")
}
