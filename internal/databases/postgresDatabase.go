package databases

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/thanchayawikgithub/isekai-shop-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstace *postgresDatabase
	once                    sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
			conf.Schema,
		)

		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", conf.DBName)

		postgresDatabaseInstace = &postgresDatabase{db}

	})

	return postgresDatabaseInstace
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return postgresDatabaseInstace.DB
}
