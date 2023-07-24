package db

type HouseWithParkingPlaces struct {
	House
	ParkingPlaces []ParkingPlace `json:"parking_places"`
	Lodgers       []Person       `json:"lodgers"`
}

func GetHouseWithParkingPlacesAndPersons(id int) *HouseWithParkingPlaces {
	uid := uint16(id)
	houseWithParkingPlaces := &HouseWithParkingPlaces{}
	houseWithParkingPlaces.House = *GetHouse(uid)
	houseWithParkingPlaces.ParkingPlaces = *GetParkingPlacesByHouseId(uid)
	houseWithParkingPlaces.Lodgers = *GetPersonsByHouseId(uid)
	return houseWithParkingPlaces
}

func GetHousesWithParkingPlacesAndPersons() []*HouseWithParkingPlaces {
	h := *GetHouses()
	hwp := make([]*HouseWithParkingPlaces, len(h))
	for i, house := range h {
		hwp[i] = GetHouseWithParkingPlacesAndPersons(int(house.ID))
	}
	return hwp
}
