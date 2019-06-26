package roomGroupModel

import (
	"reservationManager/domain/model/common"
	"reservationManager/domain/model/reservationModel"
	"reservationManager/domain/model/roomModel"
	"sort"
)

// 複数のRoomをもち、Roomの空き情報を管理する

type RoomGroup struct {
	Id                  common.RoomGroupId
	Availability        RoomAvailability
	AvailabilityVersion RoomAvailabilityVersion
	Rooms               roomModel.Rooms
	Reservations        *reservationModel.Reservations
}

func (rg RoomGroup) CanStay(span common.DateSpan) error {
	return nil
}

func (rg RoomGroup) Reassign() error {
	if rg.Reservations.Len() <= 2 {
		return nil
	}
	sort.Sort(rg.Reservations)

	start := rg.Reservations.First().StaySpan.StartDate
	end := rg.Reservations.Last().StaySpan.EndDate

	days, err := common.NewDateSpan(start, end)
	if err != nil {
		return err
	}

	dateList := days.GetDateList()
	var dateMap = map[string]int{}
	for _, date := range dateList {
		dateMap[date.String()] = 0
	}

	return nil
}
