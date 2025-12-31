package config

type (
	// Config is the base config.yaml file
	Config struct {
		Init     *Init    `yaml:"init"`
		Searches []Search `yaml:"searches"`
	}

	Init struct {
		HTTP        *HTTP    `yaml:"http"`
		File        *File    `yaml:"file"`
		NamedSearch string   `yaml:"named_search"`
		Arguments   []string `yaml:"arguments"`
	}

	Database struct {
		SQLite   *SQLite   `yaml:"sqlite"`
		Postgres *Postgres `yaml:"postgres"`
	}

	SQLite struct {
		File  string `yaml:"file"`
		Query string `yaml:"query"`
	}

	Postgres struct {
		Host string `yaml:"host"`
	}

	Header struct {
		BearerEnvVar string `yaml:"bearer-env-var"`
	}

	Auth struct {
		Header *Header `yaml:"header"`
	}

	Search struct {
		Name string `yaml:"name"`

		Init    *SearchInit `yaml:"init"`
		Results *Results    `yaml:"results"`
		Select  *Select     `yaml:"select"`
	}

	Results struct {
		Table *Table `yaml:"table"`
		List  *List  `yaml:"list"`
	}

	List struct {
	}

	Select struct {
		NamedSearch string         `yaml:"named_search"`
		Search      *Search        `yaml:"search"`
		Modal       *Modal         `yaml:"modal"`
		Datastore   map[string]int `yaml:"datastore"`
	}

	Modal struct{}

	Table struct {
		JSON *JSON `yaml:"json"`
		CSV  *CSV  `yaml:"csv"`
		SQL  *SQL  `yaml:"sql"`
	}

	SQL struct {
		Columns []string `yaml:"columns"`
	}

	JSON struct {
		Rows    string       `yaml:"rows"`
		Columns []JSONColumn `yaml:"columns"`
	}

	CSV struct {
	}

	JSONColumn struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}

	SearchInit struct {
		// SearchInit always needs one type of driver
		HTTP      *HTTP    `yaml:"http"`
		File      *File    `yaml:"file"`
		SQLite    *SQLite  `yaml:"sqlite"`
		Arguments []string `yaml:"arguments"`
	}

	File struct {
		Path string `yaml:"path"`
	}
)
