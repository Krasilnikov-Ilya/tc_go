package db

import (
	u "tc_go/utils"
)

type Car struct {
	ID         uint16  `json:"id" sql:"primary key"`
	Mark       string  `json:"mark" gorm:"not null;type:varchar(31);"`
	Model      string  `json:"model" gorm:"not null;type:varchar(31);"`
	Price      float64 `json:"money" gorm:"not null;type:numeric(19,2);"`
	EngineType string  `json:"engine_type" sql:"type:varchar(10) not null check((engine_type)::text ~ similar_to_escape('(Diesel|CNG|Hydrogenic|Electric|PHEV|Gasoline)'::text))"`
	PersonId   *uint16 `json:"person_id" gorm:"<-"`
}

func (car *Car) Validate() (map[string]interface{}, bool) {

	if car.Mark == "" {
		return u.Message(false, "Mark should be on the payload"), false
	}

	if car.Model == "" {
		return u.Message(false, "Model should be on the payload"), false
	}

	if car.Price < 0 {
		return u.Message(false, "Price should be positive float or 0"), false
	}

	if car.EngineType != "Diesel" && car.EngineType != "CNG" &&
		car.EngineType != "Hydrogenic" && car.EngineType != "Electric" &&
		car.EngineType != "PHEV" && car.EngineType != "Gasoline" {
		return u.Message(false,
			"EngineType can be Diesel, CNG, Hydrogenic, Electric, PHEV or Gasoline"), false
	}

	return u.Message(true, "success"), true
}

func (car *Car) Create() map[string]interface{} {
	if resp, ok := car.Validate(); !ok {
		return resp
	}
	GetDB().Create(car)
	resp := u.Message(true, "success")
	resp["data"] = car
	return resp
}

func GetCar(id int) *Car {
	car := &Car{}
	err := GetDB().Table("cars").Where("id = ?", id).First(car).Error
	if err != nil {
		return nil
	}
	return car
}

func GetCars() *[]Car {
	car := &[]Car{}
	err := GetDB().Table("cars").Find(car).Error
	if err != nil {
		return nil
	}
	return car
}

func GetPersonsCarsById(id int) *[]Car {
	cars := &[]Car{}
	err := GetDB().Table("cars").Where("person_id = ?", id).Find(cars).Error
	if err != nil {
		return nil
	}
	return cars
}
