package roomGroupModel

import (
	"reservationManager/domain/model/common"
	"reservationManager/domain/model/reservationModel"
	"reservationManager/domain/model/roomModel"
)

// 複数のRoomをもち、Roomの空き情報を管理する

type RoomGroup struct {
	Rooms        roomModel.Rooms
	Reservations *reservationModel.Reservations
}

func (g RoomGroup) CanStay(span common.DateSpan) error {

	return nil
}
