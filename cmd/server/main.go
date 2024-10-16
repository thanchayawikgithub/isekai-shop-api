package main

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/server"
)

func main() {
	conf := config.LoadConfig()
	db := databases.NewPostgresDatabase(conf.Database)

	server := server.NewEchoServer(conf, db)
	server.Start()
}
