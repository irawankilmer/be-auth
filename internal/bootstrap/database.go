package bootstrap

import (
	"be-blog/database/seeder"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func MigrateAndSeed(db *sql.DB) {
	doMigration(db)
	doSeeding(db)
}

func doMigration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Driver error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migrasi gagal!: %v", err)
	}

	fmt.Println("Migrasi database berhasil!")
}

func doSeeding(db *sql.DB) {
	fmt.Println("Menjalankan seeder...")
	seeder.SeedTag(db)
	fmt.Println("Seeder berhasil...")
}
