package rest_response

import (
	"encoding/json"
	"errors"
	"net/http"
)

// in case of failure. in this case not found http error
// {
// 	"status" : 404,
// 	"error" : true,
// 	"message" : "item not found",
// 	"data" : {}
// }

// case of success
// {
// 	"status" : 200/201,
// 	"error" : false,
// 	"message" : "celebrating message!",
// 	"data" : {
//		"key" : "value"
//		...
//	}
//}

type RestResp interface {
	Message() string
	Status() int
	Error() bool
	Data() interface{}
}

type restResp struct {
	RespMessage string      `json:"message"`
	RespStatus  int         `json:"status"`
	RespError   bool        `json:"error"`
	RespData    interface{} `json:"data"`
}

func (e restResp) Error() bool {
	return e.RespError
}

func (e restResp) Message() string {
	return e.RespMessage
}

func (e restResp) Status() int {
	return e.RespStatus
}

func (e restResp) Data() interface{} {
	return e.RespData
}

func NewRestResponse(message string, status int, err bool, data interface{}) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  status,
		RespError:   err,
		RespData:    data,
	}
}

func NewRestRespFromBytes(bytes []byte) (RestResp, error) {
	var apiErr restResp
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string, data interface{}) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  http.StatusBadRequest,
		RespError:   true,
		RespData:    data,
	}
}

func NewNotFoundError(message string, data interface{}) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  http.StatusNotFound,
		RespError:   true,
		RespData:    data,
	}
}

func NewUnauthorizedError(message string, data interface{}) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  http.StatusUnauthorized,
		RespError:   true,
		RespData:    data,
	}
}

func NewInternalServerError(message string, data interface{}) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  http.StatusInternalServerError,
		RespError:   true,
		RespData:    data,
	}
}

func NewSuccessResponse(message string, data interface{}, status int) RestResp {
	return restResp{
		RespMessage: message,
		RespStatus:  status,
		RespError:   false,
		RespData:    data,
	}
}
