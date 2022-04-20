package main

import (
	"os"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/fs/type/yaml"
	lighthouse "github.com/go-zoox/lighthouse/src"
	"github.com/go-zoox/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "lighthouse",
		Description: "A simple dns server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   "conf/lighthouse.yaml",
				Usage:   "The path to the configuration file",
				Aliases: []string{"c"},
			},
		},
		Action: func(c *cli.Context) error {
			configFilePath := c.String("config")
			if configFilePath == "" {
				logger.Error("config file is required")
				os.Exit(1)
			}

			if !fs.IsExist(configFilePath) {
				logger.Error("config file(%s) not found", configFilePath)
				os.Exit(1)
			}

			var config lighthouse.Config
			if err := yaml.Read(configFilePath, &config); err != nil {
				logger.Error("failed to read config file", err)
				os.Exit(1)
			}

			// fmt.Println(config)

			lighthouse.Serve(&config)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal("%s", err.Error())
	}
}
