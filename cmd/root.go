package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Execute 初始化并运行 dolphin CLI 应用程序。
// 它设置命令行解析器、注册子命令，并将 generate 作为默认操作。
func Execute() {
	app := cli.NewApp()
	app.Name = "dolphin"
	app.Usage = "This tool is for generating GraphQL API using gqlgen and gorm"
	app.Version = "0.0.1"

	app.Action = genCmd.Action
	app.Usage = genCmd.Usage
	app.Flags = genCmd.Flags

	app.Commands = []cli.Command{
		initCmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("dolphin: %s", err.Error())
	}
}
