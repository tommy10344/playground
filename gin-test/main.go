package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

// Todo データ型
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	db.AutoMigrate(&Todo{})
}

func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	db.Create(&Todo{Text: text, Status: status})
}

func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
}

func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}

func dbDeleteAll() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	for _, todo := range todos {
		db.Delete(todo)
	}
}

func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	return todos
}

func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	var todo Todo
	db.First(&todo, id)
	return todo
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

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
		status := c.PostForm("status")
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
		status := c.PostForm("status")
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
		c.File("json/hoge.json")
	})

	router.Run()
}
