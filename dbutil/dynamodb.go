package dbutil

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func toJSON(value events.DynamoDBAttributeValue) string {
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
			result += delim + toJSON(val)
			delim = ","
		}

		return result + "]"
	case events.DataTypeMap:
		result := "{"
		delim := ""

		for key, val := range value.Map() {
			result += delim + "\"" + key + "\":" + toJSON(val)
			delim = ","
		}

		return result + "}"
	default:
		json, _ := value.MarshalJSON()
		return string(json)
	}
}
