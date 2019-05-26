package roomGroupModel

import "reservationManager/domain/model/common"

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

type RoomGroupAvailability struct {
	Availability DateAvailabilities
	Version      RoomAvailabilityVersion
}
