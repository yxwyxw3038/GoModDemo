package model
import ("time")
type Model struct {
	ID        int64 
	CreateBy  string
	UpdateBy  string
	CreateTime time.Time
	UpdateTime time.Time
}