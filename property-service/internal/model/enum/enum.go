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
	Apartment PropertyTypeValue
	Home      PropertyTypeValue
	Hotel     PropertyTypeValue
}

var PropertyType = &PropertyTypeEnt{
	Apartment: "APARTMENT",
	Home:      "HOME",
	Hotel:     "HOTEL",
}
