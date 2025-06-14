## BE Auth

### Install
```bash
go get github.com/irawankilmer/be-auth@v1.0.0
```

### Penggunaan
Di main
```go
// 1. Init App (load DB, service, repo, dll)
app := bootstrap.InitApp()

// Setup Router
router := gin.Default()

// Register Routes
authRoutes.RegisterAuthRoutes(router, app.AuthService)
```