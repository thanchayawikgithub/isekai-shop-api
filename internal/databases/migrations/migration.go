package migrations

import (
	"log"

	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	"gorm.io/gorm"
)

type Migration struct {
	db   *gorm.DB
	conf *config.Database
}

// NewMigration creates a new Migration instance.
func NewMigration(db *gorm.DB, conf *config.Database) *Migration {
	return &Migration{db, conf}
}

// Migrate handles the schema and data migrations.
func (m *Migration) Migrate() {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Migration failed: %v", r) // Log panic
		}
	}()

	if err := m.schemaMigrate(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Schema migration failed: %v", err)
	}

	if err := m.dataMigrate(tx); err != nil {
		tx.Rollback()
		log.Fatalf("Data migration failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatalf("Transaction commit failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}

// schemaMigrate creates the necessary tables in the database.
func (m *Migration) schemaMigrate(tx *gorm.DB) error {
	tables := []interface{}{
		&entities.Player{},
		&entities.Admin{},
		&entities.Item{},
		&entities.PlayerCoin{},
		&entities.Inventory{},
		&entities.PurchaseHistory{},
	}

	for _, table := range tables {
		if err := tx.Migrator().CreateTable(table); err != nil {
			return err
		}
	}

	return nil
}

// dataMigrate inserts initial data into the database.
func (m *Migration) dataMigrate(tx *gorm.DB) error {
	items := []entities.Item{
		{
			Name:        "Sword",
			Description: "A sword that can be used to fight enemies.",
			Price:       100,
			Picture:     "https://i.pinimg.com/736x/73/cc/79/73cc79391b764ec40a5c77052bb846b9.jpg",
		},
		{
			Name:        "Shield",
			Description: "A shield that can be used to block enemy attacks.",
			Price:       50,
			Picture:     "https://i.pinimg.com/736x/fe/83/27/fe832717d33f05c2dbd845809ce877b8.jpg",
		},
		{
			Name:        "Potion",
			Description: "A potion that can be used to heal wounds.",
			Price:       30,
			Picture:     "https://i.pinimg.com/564x/14/7e/7d/147e7d58fa2becce0045f3aadf1808b1.jpg",
		},
		{
			Name:        "Bow",
			Description: "A bow that can be used to shoot enemies from afar.",
			Price:       80,
			Picture:     "https://i.pinimg.com/564x/1f/91/72/1f9172f5bc27094c4e167e55f8cce2f2.jpg",
		},
		{
			Name:        "Arrow",
			Description: "An arrow that can be used with a bow to shoot enemies from afar.",
			Price:       10,
			Picture:     "https://i.pinimg.com/564x/3f/25/84/3f25842cb4a8ad53a19575cc3d25c844.jpg",
		},
	}

	if err := tx.CreateInBatches(items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
