package app

import (
	"github.com/bandanascripts/phantompay/pkg/service/app/structs"
	"github.com/bandanascripts/phantompay/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	utils.LoadEnv()
}

func GetAddr() string {
	return os.Getenv("DB_ADDRESS")
}

var Db *gorm.DB

func Connect() {

	database, err := gorm.Open(mysql.Open(GetAddr()), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database : %v", err)
	}

	if err := database.AutoMigrate(
		&structs.PhantomPayUserData{},
		structs.PhantomPayTransaction{},
		&structs.PhantomPayWallet{},
	); err != nil {
		log.Fatalf("failed to auto migrate : %v", err)
	}

	Db = database
}

func GormDb() *gorm.DB {
	return Db
}
