package console

import (
	"bot-oleg/intermal/app"
	"bot-oleg/intermal/config"
	"fmt"
	"github.com/tkcrm/mx/cfg"
	"github.com/tkcrm/mx/logger"
	"github.com/urfave/cli/v2"
)

const (
	defaultConfigPath = "configs/config.yaml"
)

func InitCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "start",
			Description: "Bot oleg",
			Flags:       []cli.Flag{cfgPathsFlag()},
			Action: func(c *cli.Context) error {
				conf, err := loadConfig(c.Args().Slice(), c.StringSlice("configs"))
				if err != nil {
					return fmt.Errorf("load config: %v", err)
				}
				l := logger.New()
				l.Info("Logger init")
				app.Run(c.Context, conf, l)
				return nil
			},
		},
	}
}

func cfgPathsFlag() *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:    "configs",
		Aliases: []string{"c"},
		Value:   cli.NewStringSlice(defaultConfigPath),
		Usage:   "allows you to use your own paths to configuration files, separated by commas (config.yaml,config.prod.yml,.env)",
	}
}

func loadConfig(args, configPaths []string) (*config.Config, error) {
	conf := new(config.Config)
	if err := cfg.Load(conf,
		cfg.WithLoaderConfig(cfg.Config{
			Args:       args,
			Files:      configPaths,
			MergeFiles: true,
		}),
	); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return conf, nil
}
