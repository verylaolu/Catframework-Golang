package Common

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code int
	Msg string
	Data map[string]interface{}
}

func Json_return(writer http.ResponseWriter,result Result)  {
	json.NewEncoder(writer).Encode(result)
}