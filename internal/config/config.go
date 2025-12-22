package config

type (
	Header struct {
		BearerEnvVar string `yaml:"bearer-env-var"`
	}

	Auth struct {
		Header *Header `yaml:"header"`
	}

	Data struct {
		HTTP string `yaml:"http"`
		Auth *Auth  `yaml:"auth"`
	}

	Search struct {
		Init    *Init    `yaml:"init"`
		Results *Results `yaml:"results"`
	}

	Results struct {
		Table *Table `yaml:"table"`
	}

	Table struct {
		Columns []string `yaml:"columns"`
	}

	Init struct {
		Data *Data `yaml:"data"`
	}

	Config struct {
		Search *Search `yaml:"search"`
	}
)
