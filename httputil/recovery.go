package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/sirupsen/logrus"
	"gitlab.com/mr.14/xiv-go-core/exception"
)

// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
type Recovery struct {
	Logger     *logrus.Logger
	PrintStack bool
	StackAll   bool
	StackSize  int
}

// NewRecovery returns a new instance of Recovery
func NewRecovery(logger *logrus.Logger) *Recovery {
	return &Recovery{
		Logger:    logger,
		StackAll:  false,
		StackSize: 1024 * 8,
	}
}

func (rec *Recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			stack := make([]byte, rec.StackSize)
			stack = stack[:runtime.Stack(stack, rec.StackAll)]

			status, code, fieldErrors := rec.handleError(err)

			if status >= 400 && status < 500 {
				rec.Logger.Warnf("%s\n%s", err, stack)
			}

			if status >= 500 {
				rec.Logger.Errorf("%s\n%s", err, stack)
			}

			if rw.Header().Get("Content-Type") == "" {
				rw.Header().Set("Content-Type", "application/json; charset=utf-8")
			}

			rw.WriteHeader(status)
			json.NewEncoder(rw).Encode(struct {
				Code   string                  `json:"code"`
				Fields []*exception.FieldError `json:"fields,omitempty"`
			}{code, fieldErrors})
		}
	}()

	next(rw, r)
}

func (rec *Recovery) handleError(err interface{}) (int, string, []*exception.FieldError) {
	switch e := err.(type) {
	case *exception.HTTPError:
		return e.Status, e.Code, e.Fields
	default:
		return http.StatusInternalServerError, fmt.Errorf("%v", e).Error(), nil
	}
}
