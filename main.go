package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/agusheryanto182/go-wangkas/auth"
	"github.com/agusheryanto182/go-wangkas/handler"
	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/agusheryanto182/go-wangkas/user"
	webHandler "github.com/agusheryanto182/go-wangkas/web/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "root"
	dbName := "go_wangkas"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := handler.NewTransaksiHandler(transactionService)

	userWebHandler := webHandler.NewUserHandler(userService)
	transactionWebHandler := webHandler.NewTransactionsHandler(transactionService)
	sessionWebHandler := webHandler.NewSessionHandler(userService)

	router := gin.Default()

	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))

	router.Use(sessions.Sessions("webcrowdfunding", cookieStore))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")

	api.GET("/data", transactionHandler.GetAllData)
	api.GET("/data/weeks/:id", transactionHandler.GetDataByWeekID)

	router.POST("/users", userWebHandler.Create)
	router.GET("/users/edit/:id", userWebHandler.Edit)
	router.POST("/users/update/:id", userWebHandler.Update)

	router.GET("/transactions", transactionWebHandler.Index)
	router.GET("/transactions/edit/:id", transactionWebHandler.Index)
	router.GET("/transactions/update/:id", transactionWebHandler.Index)

	router.GET("/login", sessionWebHandler.New)
	router.POST("/session", sessionWebHandler.Create)
	router.GET("/logout", sessionWebHandler.Destroy)

	router.Run()
}

func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)

		files := append(layoutCopy, include)

		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}
