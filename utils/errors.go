package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const validationError = "ValidationError"
const serverError = "ServerError"

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func HandleServiceCodes(ctx *gin.Context, baseRes BaseResponse, err error) {
	var smerror Error
	if err != nil {
		smerror.Code = baseRes.StatusCode
		smerror.Type = serverError
		smerror.Message = err.Error()
		baseRes.Error = smerror
	}

	ctx.JSON(baseRes.StatusCode, baseRes)
	ctx.Abort()
}

func InternalServer(ctx *gin.Context, err string) {
	var res BaseResponse
	var smerror Error
	errorCode := http.StatusInternalServerError

	smerror.Code = errorCode
	smerror.Type = serverError
	smerror.Message = err

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}

func ValidationError(ctx *gin.Context, err error) {
	var res BaseResponse
	var smerror Error
	errorCode := http.StatusUnprocessableEntity

	smerror.Code = errorCode
	smerror.Type = validationError
	smerror.Message = err.Error()

	res.Error = smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
