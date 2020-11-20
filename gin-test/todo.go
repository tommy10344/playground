package main

import "github.com/jinzhu/gorm"

// Status タスクの実行状態
type Status string

const (
	// New 未実行
	New Status = "New"
	// Active 実行中
	Active = "Active"
	// Done 完了
	Done = "Done"
)

// Todo タスク情報
type Todo struct {
	gorm.Model
	Text   string
	Status Status
}
