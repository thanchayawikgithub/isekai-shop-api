package main

import (
	"fmt"

	"github.com/thanchayawikgithub/isekai-shop-api/config"
	"github.com/thanchayawikgithub/isekai-shop-api/databases"
)

func main() {
	conf := config.LoadCofig()
	db := databases.NewPostgresDatabase(conf.Database)

	fmt.Println(db.Connect())
}
