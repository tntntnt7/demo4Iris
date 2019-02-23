package models

import "time"

type Base struct {
	Id			string
	CreateTime	time.Time
	UpdateTime	time.Time
}
