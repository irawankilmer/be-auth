### Library
```bash
go get github.com/labstack/echo/v4
go get github.com/go-sql-driver/mysql
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go install entgo.io/ent/cmd/ent@latest
```

### Generate ent base
```bash
go run entgo.io/ent/cmd/ent init User Role RefreshToken
```