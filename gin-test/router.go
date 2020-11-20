package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/assets", "assets")

	// Index
	router.GET("/", func(c *gin.Context) {
		todos := dbGetAll()
		c.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create
	router.POST("/new", func(c *gin.Context) {
		text := c.PostForm("text")
		status := Status(c.PostForm("status"))
		dbInsert(text, status)
		c.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		c.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// Update
	router.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("id is not integer: %s", n))
		}
		text := c.PostForm("text")
		status := Status(c.PostForm("status"))
		dbUpdate(id, text, status)
		c.Redirect(302, "/")
	})

	// 削除確認
	router.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("id is not integer: %s", n))
		}
		todo := dbGetOne(id)
		c.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	// Delete
	router.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("id is not integer: %s", n))
		}
		dbDelete(id)
		c.Redirect(302, "/")
	})

	// Delete All
	router.POST("/delete_all", func(c *gin.Context) {
		dbDeleteAll()
		c.Redirect(302, "/")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/json/hoge", func(c *gin.Context) {
		c.File("assets/json/hoge.json")
	})

	return router
}
