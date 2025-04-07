package models

type User struct {
	BaseModel
	Age uint64 `gorm:"column:age"`
}
