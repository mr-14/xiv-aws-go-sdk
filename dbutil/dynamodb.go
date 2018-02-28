package dbutil

import (
	"errors"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fumin/dtx"
)

// ToJSON converts DynamoDBAttributeValue to json string
func ToJSON(value events.DynamoDBAttributeValue) string {
	if value.IsNull() {
		return ""
	}

	switch value.DataType() {
	case events.DataTypeBoolean:
		return strconv.FormatBool(value.Boolean())
	case events.DataTypeString:
		return "\"" + value.String() + "\""
	case events.DataTypeStringSet:
		result := "["
		delim := ""
		for _, val := range value.StringSet() {
			result += delim + "\"" + val + "\""
			delim = ","
		}
		return result + "]"
	case events.DataTypeNumber:
		return value.Number()
	case events.DataTypeNumberSet:
		result := "["
		delim := ""
		for _, val := range value.NumberSet() {
			result += delim + val
			delim = ","
		}
		return result + "]"
	case events.DataTypeList:
		result := "["
		delim := ""

		for _, val := range value.List() {
			result += delim + ToJSON(val)
			delim = ","
		}

		return result + "]"
	case events.DataTypeMap:
		result := "{"
		delim := ""

		for key, val := range value.Map() {
			result += delim + "\"" + key + "\":" + ToJSON(val)
			delim = ","
		}

		return result + "}"
	default:
		json, _ := value.MarshalJSON()
		return string(json)
	}
}

// TxWrapper gets a transaction wrapper
func TxWrapper(transaction func(tx *dtx.Transaction)) func(tx *dtx.Transaction) error {
	return func(tx *dtx.Transaction) (txErr error) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error: %s", err)
				txErr = errors.New("error.dynamodb.transaction")
			}
		}()

		transaction(tx)
		return nil
	}
}
