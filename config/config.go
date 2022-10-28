package config

var cfg *Config

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		Sqs  `yaml:"sqs"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Sqs struct {
		SqsHost                  string `yaml:"proxy-host"`
		SqsRegion                string `yaml:"region"`
		Queues                   `yaml:"queues"`
		Receivers                int   `yaml:"receivers"`
		MaxNumberOfMessages      int64 `yaml:"max-number-of-messages"`
		MessageVisibilityTimeout int64 `yaml:"message-visibility-timeout"`
		PollDelayInMilliseconds  int   `yaml:"poll-delay-in-milliseconds"`
	}

	Queues struct {
		QueueList string `yaml:"queue-list"`
		Prefix    string `yaml:"prefix"`
	}
)

func ExportConfig(config *Config) {
	cfg = config
}

func GetConfig() *Config {
	return cfg
}
