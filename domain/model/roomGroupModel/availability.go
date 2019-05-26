package roomGroupModel

import (
	"reservationManager/domain/model/common"
	"strconv"
)

type ReservedCount uint

type BlockedCount uint

// 楽観ロックのために使用
type RoomAvailabilityVersion uint

type DateAvailability struct {
	Date          common.Date
	ReservedCount ReservedCount
	BlockedCount  BlockedCount
}

type DateAvailabilities []*DateAvailability

func (r RoomAvailabilityVersion) String() string {
	return strconv.FormatUint(uint64(r), 10)
}
