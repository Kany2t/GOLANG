package gorm

import (
	"log"
	"time"
)

var teacherTemp = Teacher{
	Name:      "test",
	Email:     "test@test.com",
	Age:       30,
	WorkYears: 10,
	Birthday:  time.Now().Unix(),
	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},
	Roles: []string{"普通用户", "讲师"},
	JobInfo: Job{
		Title:    "teacher1",
		Location: "beijing",
	},
	JobInfo2: Job{
		Title:    "teacher2",
		Location: "beijing",
	},
}

func CreateRecord() {
	t := teacherTemp
	res := DB.Create(&t)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	println(res.RowsAffected, res.Error, t)
	//正向选择
	t1 := teacherTemp
	res = DB.Select("Name", "Age").Create(&t1)
	println(res.RowsAffected, res.Error, t1)
	//反向选择
	t2 := teacherTemp
	res = DB.Omit("Email", "Birthday").Create(&t2)
	println(res.RowsAffected, res.Error, t2)

	//批量传入
	var teachers1 = []Teacher{{Name: "king", Age: 40}, {Name: "jack", Age: 41}, {Name: "nick", Age: 35}}
	DB.CreateInBatches(teachers1, 2)
	for _, t := range teachers1 {
		println(t.ID)
	}
}
