package db

type ParkingPlace struct {
	ID          uint16  `json:"id" sql:"primary key"`
	PlacesCount uint16  `json:"places_count" sql:"type:int4 not null check((places_count between 1 and 4))"`
	IsCovered   bool    `json:"is_covered" sql:"type:bool not null"`
	IsWarm      bool    `json:"is_warm" sql:"type:bool not null check((((is_warm IS NOT TRUE) OR (is_covered IS NOT FALSE))))"`
	HouseId     *uint16 `json:"house_id" gorm:"<-;not null"`
}

func GetParkingPlacesByHouseId(id uint16) *[]ParkingPlace {
	pp := &[]ParkingPlace{}
	err := GetDB().Table("parking_places").Where("house_id = ?", id).Find(pp).Error
	if err != nil {
		return nil
	}
	return pp
}
