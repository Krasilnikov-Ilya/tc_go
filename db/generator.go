package db

import (
	"github.com/golang-sql/civil"
	"github.com/jaswdr/faker"
	"strings"
	"tc_go/customtypes"
	"time"
)

func GetPersonListForMigration(f faker.Faker, limit uint16) []Person {
	r := make([]Person, limit)
	for i := range r {
		r[i] = GetRandomPersonDataWithSpecifiedId(f, uint16(i+1), limit)
	}
	return r
}

func GetCarListForMigration(f faker.Faker, limit uint16) []Car {
	r := make([]Car, limit)
	for i := range r {
		r[i] = GetRandomCarDataWithSpecifiedId(f, uint16(i+1), limit)
	}
	return r
}

func GetHouseListForMigration(f faker.Faker, limit uint16) []House {
	r := make([]House, limit)
	for i := range r {
		r[i] = GetRandomHouseDataWithSpecifiedId(f, uint16(i+1), limit)
	}
	return r
}

func GetParkingPlacesListForMigration(f faker.Faker, limit uint16) []ParkingPlace {
	r := make([]ParkingPlace, limit)
	for i := range r {
		r[i] = GetRandomParkingPlaceDataWithSpecifiedId(f, uint16(i+1), limit)
	}
	return r
}

func GetRandomPersonDataWithSpecifiedId(f faker.Faker, i uint16, limit uint16) Person {
	id := f.UInt16Between(1, limit)
	return Person{
		ID:         i,
		FirstName:  f.Person().FirstNameMale(),
		SecondName: f.Person().LastName(),
		BirthDay:   RandomBirthDayInLimits(f),
		Sex:        strings.ToLower(f.Person().Gender()),
		Money:      f.RandomFloat(2, 0, 10000),
		HouseId:    &id,
	}
}

func GetRandomHouseDataWithSpecifiedId(f faker.Faker, i uint16, limit uint16) House {
	return House{
		ID:         i,
		FloorCount: f.UInt16Between(1, 10),
		Price:      f.RandomFloat(2, 50000, 100000),
	}
}

func GetRandomCarDataWithSpecifiedId(f faker.Faker, i uint16, limit uint16) Car {
	id := f.UInt16Between(1, limit)
	return Car{
		ID:         i,
		Mark:       f.Car().Maker(),
		Model:      f.Car().Model(),
		Price:      f.RandomFloat(2, 50000, 100000),
		EngineType: RandomEngineType(f),
		PersonId:   &id,
	}
}

func GetRandomParkingPlaceDataWithSpecifiedId(f faker.Faker, i uint16, limit uint16) ParkingPlace {
	id := f.UInt16Between(1, limit)
	sw := f.UInt16Between(0, 2)
	return ParkingPlace{
		ID:          i,
		PlacesCount: f.UInt16Between(1, 4),
		IsWarm:      sw == 0,
		IsCovered:   sw%2 == 0,
		HouseId:     &id,
	}
}

func RandomBirthDayInLimits(f faker.Faker) customtypes.Date {
	return customtypes.Date(
		civil.Date{
			Year:  f.IntBetween(1950, 2000),
			Month: time.Month(f.IntBetween(0, 12)),
			Day:   f.IntBetween(0, 28),
		},
	)
}

func RandomEngineType(f faker.Faker) string {
	switch f.IntBetween(1, 5) {
	case 1:
		return "Diesel"
	case 2:
		return "CNG"
	case 3:
		return "Hydrogenic"
	case 4:
		return "Electric"
	case 5:
		return "PHEV"
	default:
		return "Gasoline"
	}
}
