package db

import "gorm.io/gorm"

var DB *gorm.DB

func Connect() *gorm.DB {
	var err error
	
	dns := "host=localhost port=5432 user=admin password=admin dbname=postgress sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{Logger; logger.Default.LogMode(logger.Info)})
	if err != nil{
		log.Println(err)
	}
	return DB
}

