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

func (rs Reservations) Len() int {
	return len(rs)
}

func (rs Reservations) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs Reservations) Less(i, j int) bool {
	return rs[i].StaySpan.StartDate.IsEarlierEq(rs[j].StaySpan.StartDate)
}

func (rs Reservations) First() *Reservation {
	return rs[0]
}

func (rs Reservations) Last() *Reservation {
	return rs[rs.Len()-1]
}
