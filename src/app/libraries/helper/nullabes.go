package helper

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

func CastAsNullableInt(integerInterface interface{}) int {

	integer, err := integerInterface.(*sql.NullInt32)

	if err != true || integer.Valid == false {
		return -1
	}

	return int(integer.Int32)
}

func CastAsNullableString(stringValueInterface interface{}) string {

	stringValue, err := stringValueInterface.(*sql.NullString)

	if err != true || stringValue.Valid == false {
		return ""
	}

	return stringValue.String
}

func CastToNullableTime(timeDataInterface interface{}) NullTime {

	timeData, err := timeDataInterface.(*sql.NullString)

	if err != true || timeData.Valid == false {
		return NullTime{Valid: timeData.Valid}
	}

	newTime, error := time.Parse("2015-04-15 15:35:14", timeData.String)

	if error != nil {
		return NullTime{Valid: false}
	}

	return NullTime{Value: newTime, Valid: true}
}

type NullTime struct {
	Value time.Time
	Valid bool
}

func CastAsNullableUuid(timeDataInterface interface{}) NullUuid {

	timeData, err := timeDataInterface.(*sql.NullString)

	if err != true || timeData.Valid == false {
		return NullUuid{Valid: timeData.Valid}
	}

	newUuid, error := uuid.Parse(timeData.String)

	if error != nil {
		return NullUuid{Valid: false}
	}

	return NullUuid{Value: newUuid, Valid: true}
}

type NullUuid struct {
	Value uuid.UUID
	Valid bool
}
