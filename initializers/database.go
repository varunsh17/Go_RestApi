package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdb() {
	var err error
	// postgres://idyfxnuc:9EN9R8oe2J2EfgjBbJ18339EBFch7uYW@bubble.db.elephantsql.com/
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ")

	}
}
