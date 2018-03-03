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
func (tx *DynamoTX) PutItem(input *dynamodb.PutItemInput) {
	inputJSON, err := json.Marshal(input)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.PutItem(input)
	tx.logTransaction(err, "UpdateItem", string(inputJSON))
}

// PutItems puts items in dynamodb
func (tx *DynamoTX) PutItems(inputs []*dynamodb.PutItemInput) {
	for _, input := range inputs {
		tx.PutItem(input)
	}
}

// UpdateItem updates an item in dynamodb
func (tx *DynamoTX) UpdateItem(input *dynamodb.UpdateItemInput) {
	inputJSON, err := json.Marshal(input)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.UpdateItem(input)
	tx.logTransaction(err, "UpdateItem", string(inputJSON))
}

// UpdateItems updates items in dynamodb
func (tx *DynamoTX) UpdateItems(inputs []*dynamodb.UpdateItemInput) {
	for _, input := range inputs {
		tx.UpdateItem(input)
	}
}

// DeleteItem deletes an item in dynamodb
func (tx *DynamoTX) DeleteItem(input *dynamodb.DeleteItemInput) {
	inputJSON, err := json.Marshal(input)
	errorutil.PanicIfError(err)

	_, err = tx.DynamoDB.DeleteItem(input)
	tx.logTransaction(err, "DeleteItem", string(inputJSON))
}

// DeleteItems deletes items in dynamodb
func (tx *DynamoTX) DeleteItems(inputs []*dynamodb.DeleteItemInput) {
	for _, input := range inputs {
		tx.DeleteItem(input)
	}
}

func (tx *DynamoTX) logTransaction(err error, inputType string, inputJSON string) {
	if err == nil {
		logForm, err := json.Marshal(&logForm{
			TxID:  tx.TxID,
			Code:  "SUCCESS",
			Type:  inputType,
			Input: inputJSON,
		})
		log.Println(string(logForm))
		errorutil.PanicIfError(err)
	} else {
		logForm, err := json.Marshal(&logForm{
			TxID: tx.TxID,
			Code: "ERROR",
			Type: inputType,
		})
		log.Println(string(logForm))
		errorutil.PanicIfError(err)
	}
}
