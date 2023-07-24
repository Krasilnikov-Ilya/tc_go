package db

import (
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	// count of entities per table when creating new base
	dbScale := uint16(100)

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	fake := faker.New()
	db = conn

	Migrate()
	Fill(fake, dbScale)
}

func Migrate() {
	GetDB().Debug().AutoMigrate(&Account{}, &Car{}, &Person{}, &House{}, &ParkingPlace{})
	GetDB().Debug().Model(&Car{}).AddForeignKey(
		"person_id", "people(id)",
		"CASCADE", "CASCADE",
	)
	GetDB().Debug().Model(&ParkingPlace{}).AddForeignKey(
		"house_id", "houses(id)",
		"CASCADE", "CASCADE",
	)
	GetDB().Debug().Model(&Person{}).AddForeignKey(
		"house_id", "houses(id)",
		"CASCADE", "CASCADE",
	)
}

func Fill(fake faker.Faker, dbScale uint16) {
	if db.Table("houses").Find(&[]House{}).RowsAffected <= 0 {
		r := GetHouseListForMigration(fake, dbScale)
		for _, house := range r {
			db.Create(&house)
		}
	}

	if db.Table("people").Find(&[]Person{}).RowsAffected <= 0 {
		for _, person := range GetPersonListForMigration(fake, dbScale) {
			db.Create(&person)
		}
	}

	if db.Table("cars").Find(&[]Car{}).RowsAffected <= 0 {
		r := GetCarListForMigration(fake, dbScale)
		for _, car := range r {
			db.Create(&car)
		}
	}

	if db.Table("parking_places").Find(&[]ParkingPlace{}).RowsAffected <= 0 {
		r := GetParkingPlacesListForMigration(fake, dbScale)
		for _, pp := range r {
			db.Create(&pp)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
