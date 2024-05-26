package main

import (
	"qayyuum/go_fintech/api"
	"qayyuum/go_fintech/migrations"
)

func main() {
	migrations.Migrate()
	api.StartAPI()
}
