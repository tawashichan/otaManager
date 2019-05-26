package dynamo

import (
	"errors"
	"testing"
)

func TestIsConditionalUpdateFailed(t *testing.T) {
	err := errors.New(`ConditionalCheckFailedException: The conditional request failed
	status code: 400, request id: TL0RPCQTG33LN2C27Q41CV8R7BVV4KQNSO5AEMVJF66Q9ASUAAJG
	`)
	if !IsDynamoDBConditionalUpdateFailed(err) {
		t.Error("")
	}
}

func TestIsConditionalUpdateFailed2(t *testing.T) {
	err := errors.New(`internal server error`)
	if IsDynamoDBConditionalUpdateFailed(err) {
		t.Error("")
	}
}
