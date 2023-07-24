package db

type House struct {
	ID         uint16  `json:"id" sql:"primary key"`
	FloorCount uint16  `json:"floor_count" sql:"type:int4 not null check(((floor_count >= 1) AND (floor_count <= 10)))"`
	Price      float64 `json:"price" gorm:"not null;type:numeric(19,2);"`
}

func GetHouse(id uint16) *House {
	house := &House{}
	err := GetDB().Table("houses").Where("id = ?", id).First(house).Error
	if err != nil {
		return nil
	}
	return house
}

func GetHouses() *[]House {
	house := &[]House{}
	err := GetDB().Table("houses").Find(house).Error
	if err != nil {
		return nil
	}
	return house
}
