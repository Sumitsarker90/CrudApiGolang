package model


type Book struct{
	Id int `gorm:"type:int;primary_key"`
	Name string  `gorm:"type:varchar(255)"`
}