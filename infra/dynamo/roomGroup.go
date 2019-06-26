package dynamo

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"reservationManager/domain/model/common"
	"reservationManager/domain/model/event"
	"reservationManager/domain/model/roomGroupModel"
)

type roomGroupRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func NewRoomGroupRepository(tableName string) roomGroupModel.IRoomGroupRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	client := dynamodb.New(sess)
	return roomGroupRepository{
		client:    client,
		tableName: tableName,
	}
}

func (repo roomGroupRepository) Save(roomGroup *roomGroupModel.RoomGroup) error {
	item, err := dynamodbattribute.MarshalMap(roomGroup)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(repo.tableName),
	}

	_, e := repo.client.PutItem(input)
	if e != nil {
		return e
	}
	return nil
}

func (repo roomGroupRepository) Get(id common.RoomGroupId) (*roomGroupModel.RoomGroup, error) {
	roomGroup := &roomGroupModel.RoomGroup{}
	result, err := repo.client.GetItem(&dynamodb.GetItemInput{
		TableName: &repo.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(string(id)),
			},
		},
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, roomGroup); err != nil {
		return nil, err
	}
	return roomGroup, nil
}

func (repo roomGroupRepository) UpdateAvailability(groupId common.RoomGroupId, update event.UpdateRoomGroupAvailabilities) error {
	roomGroup, err := repo.Get(groupId)
	if err != nil {
		return err
	}
	/*availability := append(roomGroup.Availability, &roomGroupModel.DateAvailability{
		Date:          update[0].Date,
		ReservedCount: roomGroupModel.ReservedCount(update[0].Change),
	})*/

	count := roomGroupModel.ReservedCount(int(roomGroup.Availability[update[0].Date.String()].ReservedCount) + int(update[0].Change))

	fmt.Println(count)

	roomGroup.Availability[update[0].Date.String()].ReservedCount = roomGroupModel.ReservedCount(func() int {
		if count < 0 {
			return 0
		} else {
			return int(count)
		}
	}())
	availabilityInput, err := dynamodbattribute.Marshal(roomGroup.Availability)
	if err != nil {
		return err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: &repo.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(string(roomGroup.Id)),
			},
		},
		UpdateExpression:    aws.String("Set Availability = :availability, AvailabilityVersion = :nextVersion"),
		ConditionExpression: aws.String("AvailabilityVersion = :currentVersion"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":currentVersion": {
				N: aws.String(roomGroup.AvailabilityVersion.String()),
			},
			":nextVersion": {
				N: aws.String(roomGroup.AvailabilityVersion.NextVersion().String()),
			},
			":availability": availabilityInput,
		},
	}
	_, e := repo.client.UpdateItem(updateInput)
	if e != nil {
		if IsDynamoDBConditionalUpdateFailed(e) {
			return conditionalUpdateFailed{
				Err: e,
			}
		} else {
			return e
		}
	}
	return nil
}
