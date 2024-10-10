package migrations

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/databases"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"gorm.io/gorm"
)

func MigrateSchema() {
	conf := config.LoadConfig()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
