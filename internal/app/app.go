package app

func GetApp() (app *App, err error) {
	env, err := loadEnvironment()

	if err != nil {
		return
	}

	db, err := loadDB(env.DATABASE_URL)

	if err != nil {
		return
	}

	app = &App{
		DB:  db,
		ENV: env,
	}

	return
}
