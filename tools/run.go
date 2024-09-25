package tools

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// RunInteractiveInDir ...
func RunInteractiveInDir(cmd, dir string) error {
	if os.Getenv("DEBUG") != "" {
		log.Println(cmd)
	}
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Dir = dir
	err := command.Run()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
