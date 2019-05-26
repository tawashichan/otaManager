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
	startDate := common.NewDate(2019, 1, 1)
	roomGroup := &roomGroupModel.RoomGroup{
		Id: id,
		Availability: roomGroupModel.DateAvailabilities{
			&roomGroupModel.DateAvailability{
				Date:          startDate,
				ReservedCount: 0,
			},
		},
	}
	if e := client.Save(roomGroup); e != nil {
		panic(e)
	}

	for i := 0; i < 50; i++ {
		num := i
		go func() {
			shouldExec := true
			for shouldExec {
				if err := client.UpdateAvailability(id, event.UpdateRoomGroupAvailabilities{
					&event.UpdateRoomGroupAvailability{
						Date: startDate,
						Change: event.ChangeInAvailableRoomNum(func() int {
							if num%2 == 0 {
								return 1
							} else {
								return -1
							}
						}()),
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
