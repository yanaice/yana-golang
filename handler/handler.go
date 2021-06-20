package handler

import (
	"encoding/json"
	"fmt"
	"yana-golang/errs"
)

// handler util
func HandleError(err error) { // change to handleError
	switch err := err.(type) {
	case errs.AppError:
		{
			jsonByte, _ := json.Marshal(err.ErrorMsg)
			jsonString := string(jsonByte)
			fmt.Printf("StatusCode:%v, message: %v, error: %v\n", err.Code, err.Message, jsonString)
		}
	case error:
		{
			// default
			fmt.Printf("StatusCode:500, message: %v\n", err.Error())
		}
	}
}
