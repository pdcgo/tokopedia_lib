package api

import "net/http"

type BaseApi struct{}

type BaseResponse struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (a *BaseApi) BaseResponseSuccess() (int, BaseResponse) {
	return http.StatusOK, BaseResponse{
		Code:   http.StatusOK,
		Detail: "success",
	}
}

func (a *BaseApi) BaseResponseBadRequest(err error) (int, BaseResponse) {
	return http.StatusBadRequest, BaseResponse{
		Code:   http.StatusBadRequest,
		Detail: err.Error(),
	}
}

func (a *BaseApi) BaseResponseInternalServerError(err error) (int, BaseResponse) {
	return http.StatusInternalServerError, BaseResponse{
		Code:   http.StatusInternalServerError,
		Detail: err.Error(),
	}
}
