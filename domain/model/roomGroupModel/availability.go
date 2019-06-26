package roomGroupModel

import (
	"strconv"
)

type ReservedCount int

type BlockedCount int

// 楽観ロックのために使用
type RoomAvailabilityVersion uint

type DateAvailability struct {
	ReservedCount ReservedCount
	BlockedCount  BlockedCount
}

type DateAvailabilities []*DateAvailability

type RoomAvailability map[string]*DateAvailability

func (r RoomAvailabilityVersion) String() string {
	return strconv.FormatUint(uint64(r), 10)
}

func (r RoomAvailabilityVersion) NextVersion() RoomAvailabilityVersion {
	return r + 1
}
