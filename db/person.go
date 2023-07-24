package db

import (
	"tc_go/customtypes"
	u "tc_go/utils"
)

type Person struct {
	ID         uint16           `json:"id" sql:"primary key"`
	FirstName  string           `json:"first_name" gorm:"not null;type:varchar(31);"`
	SecondName string           `json:"last_name" gorm:"not null;type:varchar(31);"`
	BirthDay   customtypes.Date `json:"birth_day" gorm:"not null;type:date;"`
	Sex        string           `json:"sex" sql:"type:varchar(6) not null check((sex)::text ~ similar_to_escape('(male|female)'::text))"`
	Money      float64          `json:"money" gorm:"not null;type:numeric(19,2);"`
	HouseId    *uint16          `json:"house_id" gorm:"<-"`
}

func (person *Person) Validate() (map[string]interface{}, bool) {
	if person.FirstName == "" {
		return u.Message(false, "First name should be on the payload"), false
	}

	if person.SecondName == "" {
		return u.Message(false, "Second name should be on the payload"), false
	}

	//if person.BirthDay.After(civil.DateOf(time.Now())) {
	//	return u.Message(false, "Birthday must be in 20th or 21th century and in 11-11-2000 format"), false
	//}

	if person.Sex != "MALE" && person.Sex != "FEMALE" {
		return u.Message(false, "Sex should be \"male\" or \"female\""), false
	}

	if person.Money < 0 {
		return u.Message(false, "Money should be positive float or 0"), false
	}

	return u.Message(true, "success"), true
}

func (person *Person) Create() map[string]interface{} {
	if resp, ok := person.Validate(); !ok {
		return resp
	}
	GetDB().Create(person)
	resp := u.Message(true, "success")
	resp["data"] = person
	return resp
}

func GetPerson(id int) *Person {
	person := &Person{}
	err := GetDB().Table("people").Where("id = ?", id).First(person).Error
	if err != nil {
		return nil
	}
	return person
}

func GetPersons() *[]Person {
	person := &[]Person{}
	err := GetDB().Table("people").Find(person).Error
	if err != nil {
		return nil
	}
	return person
}

func GetPersonsByHouseId(id uint16) *[]Person {
	person := &[]Person{}
	err := GetDB().Table("people").Where("house_id = ?", id).Find(person).Error
	if err != nil {
		return nil
	}
	return person
}
