package gorm

import (
	"database/sql"
	"gorm.io/gorm"
)

type Roles []string
type Teacher struct {
	gorm.Model
	Name      string `gorm:"size: 256"`
	Email     string `gorm:"uniqueIndex"`
	Age       uint   `gorm:"check:age>30"`
	WorkYears uint8
	Birthday  int64 `gorm:"serializer:unixtime;type:time"`
	StuNumber sql.NullString
	Roles     Roles `gorm:"serializer:json"`
	JobInfo   Job   `gorm:"embeddedPrefix:job_"`
	JobInfo2  Job   `gorm:"type:bytes;serializer:gob"`
}
type Job struct {
	Title    string
	Location string
}
