package cmd

import (
	"finances/internal/infra/db"
	"finances/internal/infra/server"
	"finances/internal/infra/settings"
)

func Execute() {
	settings.LoadEnvs()
	db.ConnectDatabase()
	server.StartServer()
}
