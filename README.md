## BE Auth

### Install
```bash
go get github.com/irawankilmer/be-auth@v1.0.3
```

### Penggunaan
Di main
```go
	config.LoadENV()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	app := container.InitApp()
	routes.InitRouter(r, app)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
```