package utils

func LoadConfig() (*Config, error) {
	env, err := LoadEnv()

	if err != nil {
		return nil, err
	}

	config := Config{
		Env: env,
	}

	return &config, nil
}
