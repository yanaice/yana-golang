package main

import (
	"yana-golang/configs"
	"yana-golang/dbmysql"
	"yana-golang/handler"
	"yana-golang/model"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()
	// logs.Init()

	// dbMock := dbmock.Init()
	// dbMongo := dbmongo.Init()
	dbmysql := dbmysql.Init()
	// dsn := viper.GetString("mysql.username")

	// getUser, err := dbmysql.User.GetUser(2)
	// if err != nil {
	// 	fmt.Printf("getUser err: %+v", err)
	// } else {
	// 	fmt.Printf("getUser: %+v", getUser)
	// }

	// byCitizen, err := dbmysql.User.GetUserByCitizenId("1101433333333")
	// if err != nil {
	// 	fmt.Printf("byCitizen err: %+v", err)
	// } else {
	// 	fmt.Printf("byCitizen: %+v", byCitizen)
	// }
	// dbmysql.User.UpdateUser("xx", model.User{
	// 	FirstName: "Yana1234",
	// })
	// logs.Info(dsn)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		handler.HandlerReturnData(c, gin.H{"message": "pong"})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		userId := c.Param("id")
		user, err := dbmysql.User.GetUser(userId)

		if err != nil {
			handler.HandleReturnError(c, err)
		}

		handler.HandlerReturnData(c, user)
	})
	r.GET("/users/line/:lineId", func(c *gin.Context) {
		lineId := c.Param("lineId")
		user, err := dbmysql.User.GetUserByLineId(lineId)

		if err != nil {
			handler.HandleReturnError(c, err)
		}

		handler.HandlerReturnData(c, user)
	})
	r.GET("/users", func(c *gin.Context) {
		search := c.Query("search")
		var users []model.User
		if search != "" {
			users, _ = dbmysql.User.SearchUser(search)
		} else {
			users, _ = dbmysql.User.ListUser()
		}

		handler.HandlerReturnData(c, users)
	})
	r.POST("/users", func(c *gin.Context) {
		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			handler.HandleReturnError(c, err)
		}

		user, err := dbmysql.User.CreateUser(req)
		if err != nil {
			handler.HandleReturnError(c, err)
		}
		handler.HandlerReturnData(c, user)
	})
	r.PUT("/users/line/:lineId", func(c *gin.Context) {
		lineId := c.Param("lineId")

		var req model.User
		if err := c.ShouldBindJSON(&req); err != nil {
			handler.HandleReturnError(c, err)
		}

		user, err := dbmysql.User.UpdateUserByLineId(lineId, req)
		if err != nil {
			handler.HandleReturnError(c, err)
		}
		handler.HandlerReturnData(c, user)
	})
	r.Run()
}
