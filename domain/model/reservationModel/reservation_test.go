package reservationModel

import (
	"reservationManager/domain/model/common"
	"testing"
)

func TestReservation_CanStay(t *testing.T) {
	span := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	reservation := Reservation{
		StaySpan: span,
		Accepted: true,
	}
	if reservation.CanStay(span) {
		t.Error("")
	}
}

func TestReservation_CanStay2(t *testing.T) {
	span := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	rSpan := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 2),
		EndDate:   common.NewDate(2019, 1, 3),
	}
	reservation := Reservation{
		StaySpan: rSpan,
		Accepted: true,
	}
	if !reservation.CanStay(span) {
		t.Error("")
	}
}

func TestReservation_CanStay3(t *testing.T) {
	span := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	rSpan := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 2),
		EndDate:   common.NewDate(2019, 1, 3),
	}
	reservation := Reservation{
		StaySpan: rSpan,
		Accepted: false,
	}
	if !reservation.CanStay(span) {
		t.Error("")
	}
}

func TestReservation_CanStay4(t *testing.T) {
	span := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	rSpan := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	reservation := Reservation{
		StaySpan: rSpan,
		Accepted: false,
	}
	if !reservation.CanStay(span) {
		t.Error("")
	}
}

func TestReservations_CanStay(t *testing.T) {
	span := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	rSpan1 := common.DateSpan{
		StartDate: common.NewDate(2019, 1, 1),
		EndDate:   common.NewDate(2019, 1, 1),
	}
	reservation1 := Reservation{
		StaySpan: rSpan1,
		Accepted: false,
	}

	reservations := Reservations{
		&reservation1,
	}

	if !reservations.CanStay(span) {
		t.Error("")
	}
}
