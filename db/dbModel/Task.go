package dbModel

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	Path string `gorm:"type:varchar(255);not null"`
	Port int    `gorm:"type:int;not null"`
	Pid  int
}
