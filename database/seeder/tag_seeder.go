package seeder

import (
	"database/sql"
	"fmt"
)

func SeedTag(db *sql.DB) {
	_, err := db.Exec(`
        INSERT INTO tag (name, descriptions, slug)
        VALUES 
            ('Dangdut', 'Paragi danggutan', 'dangdut'), 
            ('Reage Musik', 'Paragi musik reage', 'reage-musik'),
            ('Pop Class Indie', 'Tah anu kitu', 'pop-class-indie')
`)
	if err != nil {
		fmt.Println("Seed tag gagal:", err)
		return
	}
	fmt.Println("Seed tag berhasil.")
}
