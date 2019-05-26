package event

import "reservationManager/domain/model/common"

type ChangeInAvailableRoomNum int

type UpdateRoomGroupAvailability struct {
	Date   common.Date
	Change ChangeInAvailableRoomNum
}

type UpdateRoomGroupAvailabilities []*UpdateRoomGroupAvailability
