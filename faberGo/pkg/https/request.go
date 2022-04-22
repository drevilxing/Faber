package https

import (
	"errors"
	"net/http"
	"net/url"
)

func CheckNeedArgs(values url.Values, args ...string) error {
	message := ""
	for _, value := range args {
		if !values.Has(value) {
			message += value + ", "
		}
	}
	if "" != message {
		return errors.New(message + "arg(s) is(are) not defined.")
	}
	return nil
}

type GetRequest struct {
	Query url.Values `json:"query"`
	Err   error      `json:"err"`
}

func DealGetRequest(w http.ResponseWriter, r *http.Request, args ...string) GetRequest {
	SetDefaultHeaders(w)
	request := GetRequest{
		Query: r.URL.Query(),
		Err:   nil,
	}
	if "GET" != r.Method {
		if "OPTIONS" != r.Method {
			request.Err = errors.New("Request Must be 'Get'. ")
			StringResponse{
				Status:  ResponseFalse,
				Message: "Server Error. " + request.Err.Error(),
				Code:    http.StatusMethodNotAllowed,
			}.Send(w, r)
		} else {
			request.Err = errors.New("OPTIONS")
			SendResponseOK(w, r)
		}
		return request
	}
	request.Err = CheckNeedArgs(request.Query, args...)
	if nil != request.Err {
		StringResponse{
			Status:  ResponseFalse,
			Message: "Server Error. " + request.Err.Error(),
			Code:    http.StatusExpectationFailed,
		}.Send(w, r)
	}
	return request
}

// POST

type PostRequest struct {
	Err error `json:"err"`
}

func DealPostRequest(w http.ResponseWriter, r *http.Request, args ...string) PostRequest {
	SetDefaultHeaders(w)
	request := PostRequest{
		Err: r.ParseForm(),
	}
	if "POST" != r.Method {
		if "OPTIONS" != r.Method {
			request.Err = errors.New("Request Must be 'Post'. ")
			StringResponse{
				Status:  ResponseFalse,
				Message: "Server Error. " + request.Err.Error(),
				Code:    http.StatusMethodNotAllowed,
			}.Send(w, r)
		} else {
			request.Err = errors.New("OPTIONS")
			SendResponseOK(w, r)
		}
		return request
	}
	request.Err = CheckNeedArgs(r.Form, args...)
	if nil != request.Err {
		StringResponse{
			Status:  ResponseFalse,
			Message: "Server Error. " + request.Err.Error(),
			Code:    http.StatusExpectationFailed,
		}.Send(w, r)
	}
	return request
}

func DealPostError() {}
