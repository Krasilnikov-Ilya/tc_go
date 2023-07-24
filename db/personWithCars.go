package db

type PersonWithCars struct {
	Person
	Cars []Car `json:"cars"`
}

func GetPersonWithCars(id int) *PersonWithCars {
	personWithCars := &PersonWithCars{}
	personWithCars.Person = *GetPerson(id)
	personWithCars.Cars = *GetPersonsCarsById(id)
	return personWithCars
}
