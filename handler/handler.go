package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yana-golang/errs"

	"github.com/gin-gonic/gin"
)

// handler util
func HandleReturnError(c *gin.Context, err error) {
	switch err := err.(type) {
	case errs.AppError:
		{
			jsonByte, _ := json.Marshal(err.ErrorMsg)
			jsonString := string(jsonByte)
			fmt.Printf("## Handler Request ## StatusCode:%v, message: %v, error: %v\n", err.Code, err.Message, jsonString)
			c.JSON(err.Code, err)
		}
	case error:
		{
			// default
			fmt.Printf("## Handler Request ## StatusCode:500, message: %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, err)
		}
	}
}

func HandlerReturnData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
