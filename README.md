___
## **Migrasi**
- #### A. Buat file migrasi
```go
migrate create -ext sql -dir database/migration -seq create_tag_table
```

- #### B. Contoh file migrasi
1. Up
```sql
CREATE TABLE IF NOT EXISTS tag (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```
 2. Down
```sql
DROP TABLE IF EXISTS tag;
```

- #### C. Contoh file seeder
```go
func SeedTag(db *sql.DB) {
	_, err := db.Exec(`
        INSERT INTO tag (name)
        VALUES 
            ('Coba1'), 
            ('Coba2')`)
	if err != nil {
		fmt.Println("Seed tag gagal:", err)
		return
	}
	fmt.Println("Seed tag berhasil.")
}
```

- #### D. Rollback migrasi
```go
migrate -database "mysql://user:pass@tcp(localhost:3306)/dbname" -path database/migration down 1
```
___