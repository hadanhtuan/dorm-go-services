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
