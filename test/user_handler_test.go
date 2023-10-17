package test

import (
	"fmt"
	"log"

	handler1 "github.com/agusheryanto182/go-wangkas/handler"
	"github.com/agusheryanto182/go-wangkas/router"
	"github.com/agusheryanto182/go-wangkas/transaction"
	"github.com/agusheryanto182/go-wangkas/user"
	"github.com/agusheryanto182/go-wangkas/web/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setUpDB() *gorm.DB {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "Root1234!"
	dbName := "go_wangkas"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func setUpRouter(db *gorm.DB) *gin.Engine {
	userRepository := user.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	transactionService := transaction.NewService(transactionRepository)

	transactionHandler := handler1.NewTransactionHandler(transactionService)
	transactionWebHandler := handler.NewTransactionsHandler2(transactionService)
	sessionHandler := handler.NewSessionHandler(userService)

	router := router.NewRouter(transactionHandler, transactionWebHandler, sessionHandler)

	return router
}

// func TestSessionSuccess(t *testing.T) {
// 	db := setUpDB()
// 	router := setUpRouter(db)

// 	requestBody := strings.NewReader(`{"email" : "admin@gmail.com", "password" : "admin"}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/session", requestBody)

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// }
