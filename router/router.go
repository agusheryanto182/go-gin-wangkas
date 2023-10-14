package router

import (
	"path/filepath"

	"github.com/agusheryanto182/go-wangkas/auth"
	"github.com/agusheryanto182/go-wangkas/handler"
	handler2 "github.com/agusheryanto182/go-wangkas/web/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(handlers handler.Handler, transactionWebHandler handler2.TransactionHandlerWeb, sessionWebHandler handler2.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))

	router.Use(sessions.Sessions("wangkas", cookieStore))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")

	api.GET("/data", handlers.GetAllData)
	api.GET("/data/weeks/:id", handlers.GetDataByWeekID)

	router.GET("/transactions", auth.AuthAdminMiddleware(), transactionWebHandler.Index)
	router.GET("/transactions/new", auth.AuthAdminMiddleware(), transactionWebHandler.New)
	router.POST("/transactions", auth.AuthAdminMiddleware(), transactionWebHandler.Create)
	router.GET("/transactions/edit/:id", auth.AuthAdminMiddleware(), transactionWebHandler.Edit)
	router.POST("/transactions/delete/:id", auth.AuthAdminMiddleware(), transactionWebHandler.Delete)
	router.POST("/transactions/update/:id", auth.AuthAdminMiddleware(), transactionWebHandler.Update)
	router.GET("/transactions/search", auth.AuthAdminMiddleware(), transactionWebHandler.SearchByWeek)

	router.GET("/login", sessionWebHandler.New)
	router.POST("/session", sessionWebHandler.Create)
	router.GET("/logout", sessionWebHandler.Destroy)

	return router
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
