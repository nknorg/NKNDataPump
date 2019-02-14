package apiServerResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const API_VERSION  = "NKNDataPump.v1"

type ResponseData struct {
	Data interface{}
	Timestamp string
	ApiVersion string
}

type Response struct {
	ctx   *gin.Context
}

func New(ctx *gin.Context) *Response {
	return &Response{
		ctx:   ctx,
	}
}

func (r *Response) BadRequest(data interface{}) {
	r.ctx.JSON(http.StatusBadRequest, data)
}

func (r *Response) InternalServerError(data interface{}) {
	r.ctx.JSON(http.StatusInternalServerError, data)
}

func (r *Response) Success(data interface{}) {
	responseData := &ResponseData{
		data,
		time.Now().Format("2006-01-02 15:04:05"),
		API_VERSION,
	}
	r.ctx.JSON(http.StatusOK, responseData)
}
