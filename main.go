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

	startDate := common.NewDate(2019, 1, 1)

	for i := 0; i < 20; i++ {
		num := i
		go func() {
			shouldExec := true
			for shouldExec {
				if err := client.UpdateAvailability(id, event.UpdateRoomGroupAvailabilities{
					&event.UpdateRoomGroupAvailability{
						Date:   startDate,
						Change: event.ChangeInAvailableRoomNum(num),
					},
				}); err != nil {
					if dynamo.IsDynamoDBConditionalUpdateFailed(err) {
						fmt.Println(err.Error())
					} else {
						shouldExec = false
					}
				} else {
					shouldExec = false
				}
			}
		}()
	}
	time.Sleep(100 * time.Second)
}
