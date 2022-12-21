package ibm_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	APIKey  string   `yaml:"api_key,omitempty" mapstructure:"api_key"`
	Regions []string `yaml:"regions,omitempty" mapstructure:"regions"`
}
