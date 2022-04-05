package utils

const (
	StatusWaitPayment      int = 0
	StatusWaitPaymentProof int = 1
	StatusPaymentVerifying int = 2
	StatusPaymentVerified  int = 3
	StatusPaymentDeclined  int = 4
	StatusTicketProcess    int = 5
	StatusTicketConfirm    int = 6
	StatusBookingExpired   int = 7
	StatusBookingRejected  int = 8

	StatusAirlineBooking  int = 0
	StatusAirlineIssued   int = 1
	StatusAirlineCanceled int = 2

	StatusRobotQueue   int = 0
	StatusRobotRequest int = 1
	StatusRobotProcess int = 2
	StatusRobotSuccess int = 3
	StatusRobotError   int = 4
)

//GetStatusInfo func
func GetStatusInfo(x int) string {
	s := [9]string{
		"Waiting for payment",
		"Waiting for payment proof",
		"Payment is being verified",
		"Payment has been verified",
		"Payment declined",
		"Ticket on process",
		"Ticket confirm",
		"Booking Expired",
		"Booking Rejected",
	}

	if x > (len(s)-1) || x < 0 {
		return ""
	}

	return s[x]
}

//GetStatusAirlineInfo func
func GetStatusAirlineInfo(x int) string {
	s := [3]string{
		"Booking",
		"Issued",
		"Cancel",
	}

	if x > (len(s)-1) || x < 0 {
		return ""
	}

	return s[x]
}

//GetStatusRobotInfo func
func GetStatusRobotInfo(x int) string {
	status := [5]string{
		"In queue",
		"On Request",
		"On Process",
		"Success",
		"Error",
	}

	if x > (len(status)-1) || x < 0 {
		return ""
	}

	return status[x]
}
