package main

import (
	"fmt"
	"reservationManager/domain/model/common"
	"reservationManager/domain/model/event"
	"reservationManager/domain/model/roomGroupModel"
	"reservationManager/infra/dynamo"
	"time"
)

func main() {
	id := common.RoomGroupId("hoge")
	client := dynamo.NewRoomGroupRespository("otaManagerRoomGroup")
	roomGroup := &roomGroupModel.RoomGroup{Id: id}
	if e := client.Save(roomGroup); e != nil {
		panic(e)
	}
	for i := 0; i < 100; i++ {
		go func() {
			if err := client.UpdateAvailability(id, event.UpdateRoomGroupAvailabilities{}); err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
