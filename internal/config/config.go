package config

type (
	Config struct {
		Init
		Searches []*Search `yaml:"searches"`
	}

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
		Name string `yaml:"name"`

		Init    *Init    `yaml:"init"`
		Results *Results `yaml:"results"`
		Select  *Select  `yaml:"select"`
	}

	Results struct {
		Table *Table `yaml:"table"`
	}

	Select struct {
		NamedSearch string  `yaml:"named_search"`
		Search      *Search `yaml:"search"`
		Modal       *Modal  `yaml:"modal"`
	}

	Modal struct{}

	Table struct {
		Columns []*Column `yaml:"columns"`
	}

	Column struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}

	Init struct {
		NamedSearch string `yaml:"named_search"`
		HTTP        HTTP   `yaml:"http"`
	}

	HTTP struct {
		URL    string `yaml:"url"`
		Method string `yaml:"method"` // defaults to GET
		Auth   Auth   `yaml:"auth"`
	}
)
