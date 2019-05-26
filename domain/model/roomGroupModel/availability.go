package roomGroupModel

import (
	"reservationManager/domain/model/common"
	"strconv"
)

type ReservedCount int

type BlockedCount int

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

func (r RoomAvailabilityVersion) NextVersion() RoomAvailabilityVersion {
	return r + 1
}
