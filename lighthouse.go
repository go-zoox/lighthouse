package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-zoox/config"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/lighthouse/admin"
	"github.com/go-zoox/lighthouse/constants"
	"github.com/go-zoox/lighthouse/core"
	"github.com/go-zoox/logger"
	"github.com/urfave/cli/v2"
)

// //go:embed admin/static
// var StaticFS embed.FS

func main() {
	app := &cli.App{
		Name:        "lighthouse",
		Usage:       "DNS Server",
		Description: "An Easy Self Hosted DNS Server",
		Version:     fmt.Sprintf("%s (%s %s)", constants.Version, constants.CommitHash, constants.BuildTime),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				// Value:   "conf/lighthouse.yaml",
				Usage:   "The path to the configuration file",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "port",
				Value:   "53",
				Usage:   "The port to listen on (Only works when config file is not specified)",
				Aliases: []string{"p"},
			},
		},
		Action: func(c *cli.Context) error {
			configFilePath := c.String("config")
			// if configFilePath == "" {
			// 	logger.Error("config file is required")
			// 	os.Exit(1)
			// }

			var cfg core.Config

			if configFilePath != "" {
				if !fs.IsExist(configFilePath) {
					logger.Error("config file(%s) not found", configFilePath)
					os.Exit(1)
				}

				if err := config.Load(&cfg, &config.LoadOptions{
					FilePath: configFilePath,
				}); err != nil {
					logger.Error("failed to read config file", err)
					os.Exit(1)
				}

				// j, _ := json.MarshalIndent(cfg, "", "  ")
				// fmt.Println(string(j))
				// os.Exit(0)
			} else {
				port, err := strconv.Atoi(c.String("port"))
				if err != nil {
					logger.Error("failed to parse port", err)
					os.Exit(1)
				}

				cfg.Server.Port = int64(port)
			}

			// @TODO
			if os.Getenv("DEBUG") == "true" {
				logger.Debug("config: %v", cfg)
			}

			go admin.Start(&cfg.Web, &cfg)

			core.Serve(&cfg)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal("%s", err.Error())
	}
}
