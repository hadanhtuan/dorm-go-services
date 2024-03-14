package enum


type UserRoleValue string
type UserRoleEnt struct {
	User  UserRoleValue
	Host  UserRoleValue
	Admin UserRoleValue
}
var UserRole = &UserRoleEnt{
	User:  "USER",
	Host:  "HOST",
	Admin: "ADMIN",
}


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
