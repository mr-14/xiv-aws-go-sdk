package dbutil

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mr-14/xiv-aws-go-sdk/errorutil"
	uuid "github.com/satori/go.uuid"
)

// DynamoTX defines a dynamodb transaction
type DynamoTX struct {
	DynamoDB *dynamodb.DynamoDB
	TxID     string
}

type logForm struct {
	TxID  string
	Code  string
	Type  string
	Input string
}

// NewDynamoTX creates a dynamodb transaction
func NewDynamoTX(dynamodb *dynamodb.DynamoDB) *DynamoTX {
	txID, err := uuid.NewV4()
	errorutil.PanicIfError(err)
	return &DynamoTX{
		DynamoDB: dynamodb,
		TxID:     txID.String(),
	}
}

// PutItem puts an item in dynamodb
func (tx *DynamoTX) PutItem(input *dynamodb.PutItemInput, rollbackInput *dynamodb.DeleteItemInput) {
	rollbackJSON, err := json.Marshal(rollbackInput)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.PutItem(input)
	tx.logError(err, "UpdateItem", string(rollbackJSON))
}

// UpdateItem updates an item in dynamodb
func (tx *DynamoTX) UpdateItem(input *dynamodb.UpdateItemInput, rollbackInput *dynamodb.UpdateItemInput) {
	rollbackJSON, err := json.Marshal(rollbackInput)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.UpdateItem(input)
	tx.logError(err, "UpdateItem", string(rollbackJSON))
}

// DeleteItem puts an item in dynamodb
func (tx *DynamoTX) DeleteItem(input *dynamodb.DeleteItemInput, rollbackInput *dynamodb.PutItemInput) {
	rollbackJSON, err := json.Marshal(rollbackInput)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.DeleteItem(input)
	tx.logError(err, "DeleteItem", string(rollbackJSON))
}

func (tx *DynamoTX) logError(err error, inputType string, rollbackInput string) {
	if err == nil {
		logForm, err := json.Marshal(&logForm{
			TxID:  tx.TxID,
			Code:  "SUCCESS",
			Type:  inputType,
			Input: string(rollbackInput),
		})
		errorutil.PanicIfError(err)
		log.Println(string(logForm))
	} else {
		logForm, err := json.Marshal(&logForm{
			TxID: tx.TxID,
			Code: "ERROR",
			Type: inputType,
		})
		errorutil.PanicIfError(err)
		log.Println(string(logForm))
		panic(err)
	}
}
