package main

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases/migrations"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server"
)

func main() {
	conf := config.LoadConfig()
	db := databases.NewPostgresDatabase(conf.Database)

	migration := migrations.NewMigration(db.Connect(), conf.Database)
	migration.Migrate()

	server := server.NewEchoServer(conf, db.Connect())
	server.Start()
}
