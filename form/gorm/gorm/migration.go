package gorm

func init() {
	DB.AutoMigrate(&Teacher{})
	//DB.Migrator().
}
