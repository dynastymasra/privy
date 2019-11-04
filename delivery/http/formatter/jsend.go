package formatter

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

//JSend used to format JSON with JSend rules
type JSend struct {
	Status  string      `json:"status" binding:"required"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// FailResponse is used to return response with JSON format if failure
func FailResponse(msg string) JSend {
	return JSend{Status: "failed", Message: msg}
}

// SuccessResponse used to return response with JSON format success
func SuccessResponse() JSend {
	return JSend{Status: "success"}
}

// ObjectResponse used to return response JSON format if have data value
func ObjectResponse(data interface{}) JSend {
	return JSend{Status: "success", Data: data}
}

// Stringify used to stringify json object
func (j JSend) Stringify() string {
	toJSON, err := json.Marshal(j)
	if err != nil {
		logrus.WithError(err).Errorln("Unable to stringify JSON")
		return ""
	}
	return string(toJSON)
}
