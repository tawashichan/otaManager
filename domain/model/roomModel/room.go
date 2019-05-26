package roomModel

import "reservationManager/domain/model/common"

type Room struct {
	Id   common.RoomID
	Name string
}

type Rooms []*Room
