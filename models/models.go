package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         uint
	Openid     string
	Mobile     string
	Email      string
	Nickname   string
	Avatar     string
	CityText   string
	lat        string
	Lon        string
	Set        uint
	CreateTime uint
}
