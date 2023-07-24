package customtypes

import (
	"database/sql"
	"database/sql/driver"
	"github.com/golang-sql/civil"
	"time"
)

type Date civil.Date

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date(civil.DateOf(nullTime.Time))
	return
}

func (date Date) Value() (driver.Value, error) {
	return time.Date(date.Year, date.Month, date.Day, 0, 0, 0, 0, time.UTC), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}

func (date Date) GobEncode() ([]byte, error) {
	return civil.Date(date).In(time.UTC).GobEncode()
}

func (date *Date) GobDecode(b []byte) error {
	var timeVal time.Time
	err := timeVal.GobDecode(b)
	if err != nil {
		return err
	}
	dateVal := Date(civil.DateOf(timeVal))
	*date = dateVal
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	marshalled := make([]byte, 0)
	text, err := civil.Date(date).MarshalText()
	marshalled = append(marshalled, byte('"'))
	marshalled = append(marshalled, text...)
	marshalled = append(marshalled, byte('"'))
	return marshalled, err
}

func (date *Date) UnmarshalJSON(b []byte) error {
	c := civil.Date{}
	err := c.UnmarshalText(b[1 : len(b)-1])
	*date = Date(c)
	return err
}
