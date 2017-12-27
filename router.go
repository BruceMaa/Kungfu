package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// 初始化路由配置
func initRouter(middle ...gin.HandlerFunc) http.Handler {
	router := gin.Default()
	router.Use(middle...)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Panda!"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/person", Add)
	router.GET("/person")
	router.PUT("/person")
	router.DELETE("person")

	wechatMpServer(router)

	return router
}

var mysql *sql.DB

func init() {

	mysql, err := sql.Open("mysql", "rq8hxb1ru3:0A3r059F7FQd@tcp(rm-bp16gxb8lj09v0u4j.mysql.rds.aliyuncs.com:3306)/rq8hxb1ru3")
	//defer mysql.Close()
	if err != nil {
		Logger.Fatal(err)
	}

	mysql.SetMaxIdleConns(20)
	mysql.SetMaxOpenConns(20)

	if err = mysql.Ping(); err != nil {
		Logger.Panic(err)
	}
	Logger.Info("database")
}

func Add(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	Logger.Infof("mysql.Ping() %+v\n", mysql.Ping())

	//rs, err := mysql.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", firstName, lastName)
	//if err != nil {
	//	Logger.Fatal(err)
	//}
	//
	//id, err := rs.LastInsertId()
	//if err != nil {
	//	Logger.Fatal(err)
	//}
	//Logger.Info("insert person id {}", id)
	//c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("insert successful %d", id)})
	c.JSON(http.StatusOK, gin.H{"firstName": firstName, "lastName": lastName})
}
