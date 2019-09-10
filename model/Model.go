package model
import ("time")
type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreateBy  string
	UpdateBy  string
	CreateTime time.Time
	UpdateTime time.Time
}