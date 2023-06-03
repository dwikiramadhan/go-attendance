package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database structure
type DbInstance struct {
	Db *gorm.DB
}

var (
	_ = godotenv.Load(".env")
	//DB variable
	DB DbInstance
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() {
	// dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	var err error

	// db, err := gorm.Open("mysql", "root:root@localhost/go_api_shop_gonc?charset=utf8")
	//databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
		//fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	// db.DB().SetMaxIdleConns(10)
	// DB.LogMode(true)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	// log.Println("Connected database: ")

	DB = DbInstance{Db: db}

	// return DB

}

// RemoveDb Delete the database after running testing cases.
// func RemoveDb(db *gorm.DB) error {
// 	sqlDB, _ := db.DB()
// 	sqlDB.Close()
// 	err := os.Remove(path.Join(".", "app.db"))
// 	return err
// }

// // GetDb Using this function to get a connection, you can create your connection pool here.
// func GetDb() *gorm.DB {
// 	return DB
// }
