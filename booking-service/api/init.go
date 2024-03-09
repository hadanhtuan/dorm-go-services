package apiBooking

import (
	protoBooking "booking-service/proto/booking"
)

type BookingController struct {
	protoBooking.UnimplementedBookingServiceServer
}
