package main

import "github.com/jinzhu/gorm"

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	db.AutoMigrate(&Todo{})
}

func dbInsert(text string, status Status) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Database opening error")
	}
	defer db.Close()

	db.Create(&Todo{Text: text, Status: status})
}

func dbUpdate(id int, text string, status Status) {
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
