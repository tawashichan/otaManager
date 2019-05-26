package dynamo

import "regexp"

type conditionalUpdateFailed struct {
	Err error
}

var conditionalUpdateFailedRegExp = regexp.MustCompile("ConditionalCheckFailedException")

func (c conditionalUpdateFailed) Error() string {
	return c.Err.Error()
}

func IsDynamoDBConditionalUpdateFailed(err error) bool {
	return conditionalUpdateFailedRegExp.MatchString(err.Error())
}
