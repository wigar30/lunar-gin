package model

import "github.com/gin-gonic/gin"


type ApiResponse struct {
	Data   interface{}   `json:"data"`
	Status string        `json:"status,omitempty"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"details,omitempty"`
}
func (e ErrorResponse) Error() string {
	return e.Message
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

func OnSuccess(c *gin.Context, data interface{}) error {
	resp := ApiResponse{
		Status: "OK",
		Data: data,
	}

	c.JSON(200, resp)
	return nil
}

func OnError(c *gin.Context, err *ErrorResponse) error {

	c.JSON(err.Code, err)
	return nil
}
