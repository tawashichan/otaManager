package roomGroupModel

import (
	"reservationManager/domain/model/common"
	"reservationManager/domain/model/event"
)

type IRoomGroupRepository interface {
	Save(roomGroup *RoomGroup) error
	UpdateAvailability(groupId common.RoomGroupId, update event.UpdateRoomGroupAvailabilities) error
}
