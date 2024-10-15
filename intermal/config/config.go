package config

type (
	Config struct {
		Telegram TelegramConfig `yaml:"telegram"`
	}
	TelegramConfig struct {
		Token string `yaml:"token"`
	}
)
