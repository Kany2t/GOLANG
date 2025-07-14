package gorm

func Query() {
	//指定表或者model
	//查询单条数据
	//查询第一条
	t := Teacher{}
	res := DB.First(&t)
	println(res.RowsAffected, res.Error, t)

	//查询最后一条
	t = Teacher{}
	res = DB.Last(&t)
	println(res.RowsAffected, res.Error, t)

	//无排序，取第一条
	t = Teacher{}
	res = DB.Take(&t)
	println(res.RowsAffected, res.Error, t)

	//将结果填充到集合不支持特殊类型处理，无法完成类型转换
	result := map[string]interface{}{}
	res = DB.Model(&Teacher{}).Omit("	Birthday", "Roles", "JobInfo2").First(&result)
	println(res.RowsAffected, res.Error, result)

	result = map[string]interface{}{}
	res = DB.Table("teachers").Omit("Birthday", "Roles", "JobInfo2").First(&result)
	println(res.RowsAffected, res.Error, result)

	result = map[string]interface{}{}
	res = DB.Table("teachers").Take(&result)
	println(res.RowsAffected, res.Error, result)

	var teachers []Teacher
	res = DB.Where("name=?", "张三").Or("name=?", "king").Order("id desc").Limit(10).Find(&teachers)
	println(res.RowsAffected, res.Error, teachers)

}
