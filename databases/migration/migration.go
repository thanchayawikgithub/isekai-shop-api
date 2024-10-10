package main

import (
	"github.com/thanchayawikgithub/isekai-shop-api/config"
	"github.com/thanchayawikgithub/isekai-shop-api/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/entities"
)

func main() {
	conf := config.LoadCofig()
	db := databases.NewPostgresDatabase(conf.Database)

	playerMigration(db)
	adminMigration(db)
	itemMigration(db)
	playerCoinMigration(db)
	inventoryMigration(db)
	purchaseHistoryMigration(db)
}

func playerMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.Player{})
}

func adminMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.PurchaseHistory{})
}
