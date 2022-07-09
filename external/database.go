package external

import (
	"github.com/smoothbronco/gqlgen_tutorial/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	dns := "root:password@tcp(127.0.0.1:3306)/gqlgen_tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Todo{})
	db.Logger = db.Logger.LogMode(logger.Info)
	return db, nil
}
