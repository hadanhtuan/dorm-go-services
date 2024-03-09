package enum

type BookingStatusValue string

type BookingStatusEnt struct {
	Pending   BookingStatusValue
	Confirmed BookingStatusValue
	Rejected  BookingStatusValue
}

var BookingStatus = &BookingStatusEnt{
	Pending:   "PENDING",
	Confirmed: "CONFIRMED",
	Rejected:  "REJECTED",
}

type PropertyTypeValue string

type PropertyTypeEnt struct {
	Room  PropertyTypeValue
	Home  PropertyTypeValue
	Hotel PropertyTypeValue
}

var PropertyType = &PropertyTypeEnt{
	Room:  "ROOM",
	Home:  "HOME",
	Hotel: "HOTEL",
}

func GetBookingStatusValue(value int64) BookingStatusValue {
	switch value {
	case 1:
		return BookingStatus.Pending
	case 2:
		return BookingStatus.Confirmed
	case 3:
		return BookingStatus.Rejected
	default:
		return BookingStatus.Pending
	}
}

func GetPropertyTypeValue(value int64) PropertyTypeValue {
	switch value {
	case 1:
		return PropertyType.Room
	case 2:
		return PropertyType.Home
	case 3:
		return PropertyType.Hotel
	default:
		return PropertyType.Room
	}
}
