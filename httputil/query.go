package httputil

import (
	"net/url"

	"gitlab.com/mr.14/xiv-go-core/converter"
)

// Query defines the API query credentials
type Query struct {
	Key string      `json:"key"`
	Val interface{} `json:"val"`
	Op  string      `json:"op"`
}

// NewQuery creates a query instance
func NewQuery(key string, val interface{}, op string) *Query {
	return &Query{
		Key: key,
		Val: val,
		Op:  op,
	}
}

// GetParams extracts queries from URL
func GetParams(url *url.URL) []*Query {
	var params []*Query
	if q := url.Query().Get("q"); q != "" {
		converter.StringToModel(&params, q)
	}
	return params
}
