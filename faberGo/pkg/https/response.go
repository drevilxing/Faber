package https

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const OKString = "Request Executed Successfully."
const ResponseFalse = 1
const ResponseTrue = 0

type Dealer struct {
	Header   func(w http.ResponseWriter) `json:"header"`
	Handlers []Handler                   `json:"handlers"`
}

type Handler interface {
	Dealer(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)
}

func (dealer Dealer) Deal(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	if nil != dealer.Header {
		dealer.Header(w)
	}
	wc, rc := w, r
	var err error
	for i := range dealer.Handlers {
		if nil == dealer.Handlers[i].Dealer {
			continue
		}
		wc, rc, err = dealer.Handlers[i].Dealer(wc, rc)
		if nil != err {
			return w, r, err
		}
	}
	return wc, rc, nil
}

type StringResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func makeResponseOK() StringResponse {
	return StringResponse{
		Status:  ResponseTrue,
		Message: OKString,
		Code:    http.StatusOK,
	}
}

func (res StringResponse) Send(w http.ResponseWriter, r *http.Request) {
	result, err := json.Marshal(res)
	if nil != err {
		fmt.Println(err)
		return
	}
	SetContentJsonHeader(w)
	w.WriteHeader(res.Code)
	_, err = w.Write(result)
	if nil != err {
		fmt.Println(err)
	}
}

func SendStringResponse(w http.ResponseWriter, r *http.Request, response StringResponse) {
	result, err := json.Marshal(response)
	if nil != err {
		fmt.Println(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(response.Code)
	_, err = w.Write(result)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func SendResponseOK(w http.ResponseWriter, r *http.Request) {
	makeResponseOK().Send(w, r)
}

func SendResponseInternalError(w http.ResponseWriter, r *http.Request, err error, code int) {
	StringResponse{
		Status:  code,
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}.Send(w, r)
}

type JsonResponse struct {
	Status  int             `json:"status"`
	Message json.RawMessage `json:"message"`
	Code    int             `json:"code"`
}

func SendJsonResponse(w http.ResponseWriter, r *http.Request, response JsonResponse) {
	result, err := json.Marshal(response)
	if nil != err {
		fmt.Println(err.Error())
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(response.Code)
	_, err = w.Write(result)
	if nil != err {
		fmt.Println(err.Error())
	}
}

func DealJsonMarshal(w http.ResponseWriter, r *http.Request, v interface{}) []byte {
	data, err := json.Marshal(v)
	if nil != err {
		SendStringResponse(w, r, StringResponse{
			Status:  ResponseFalse,
			Message: "Server Error. Json: " + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return nil
	}
	return data
}

func WriteJsonMarshal(w http.ResponseWriter, r *http.Request, v interface{}) {
	message := DealJsonMarshal(w, r, v)
	if nil != message {
		SendJsonResponse(w, r, JsonResponse{
			Status:  ResponseTrue,
			Message: message,
			Code:    http.StatusOK,
		})
	}
}
