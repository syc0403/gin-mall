package dao

import "fmt"

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8m64").AutoMigrate()
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
