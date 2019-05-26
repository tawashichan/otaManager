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

func NewRoomGroupRespository(tableName string) roomGroupModel.IRoomGroupRepository {
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

	output, err := repo.client.PutItem(input)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}

func (repo roomGroupRepository) UpdateAvailability(groupId common.RoomGroupId, update event.UpdateRoomGroupAvailabilities) error {
	roomGroup := &roomGroupModel.RoomGroup{}
	result, err := repo.client.GetItem(&dynamodb.GetItemInput{
		TableName: &repo.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(string(groupId)),
			},
		},
	})
	if err != nil {
		return err
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, roomGroup); err != nil {
		return err
	}

	fmt.Println(roomGroup)

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
				N: aws.String((roomGroup.AvailabilityVersion + 1).String()),
			},
			":availability": availabilityInput,
		},
	}
	updateResult, err := repo.client.UpdateItem(updateInput)
	fmt.Println(updateResult)
	if err != nil {
		return err
	}
	return nil
}
