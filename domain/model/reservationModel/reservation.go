package reservationModel

import (
	"reservationManager/domain/model/common"
)

type Reservation struct {
	StaySpan common.DateSpan
	Accepted bool
}

type Reservations []*Reservation

func (r Reservation) CanStay(staySpan common.DateSpan) bool {
	return !(r.Accepted && r.StaySpan.IsOverlapping(staySpan))
}

func (rs Reservations) CanStay(staySpan common.DateSpan) bool {
	for _, r := range rs {
		if !r.CanStay(staySpan) {
			return false
		}
	}
	return true
}
