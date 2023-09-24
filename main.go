package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/agusheryanto182/go-wangkas/handler"
	"github.com/agusheryanto182/go-wangkas/transaction"
	webHandler "github.com/agusheryanto182/go-wangkas/web/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
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

	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := handler.NewTransaksiHandler(transactionService)

	transactionWebHandler := webHandler.NewTransactionsHandler(transactionService)

	router := gin.Default()

	router.Use(cors.Default())

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")

	api.GET("/data", transactionHandler.GetAllData)
	api.GET("/data/weeks/:id", transactionHandler.GetDataByWeekID)

	router.GET("/transactions", transactionWebHandler.Index)
	router.GET("/transactions/new", transactionWebHandler.New)
	router.POST("/transactions", transactionWebHandler.Create)
	router.GET("/transactions/edit/:id", transactionWebHandler.Edit)
	router.POST("/transactions/update/:id", transactionWebHandler.Update)
	router.GET("/transactions/search", transactionWebHandler.SearchByWeek)

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
