package utils

import (
	"MovieService/config"
	"encoding/json"
)

// StructToMap struct 转 map
func StructToMap(req interface{}) (map[string]interface{}, error) {
	mapData := make(map[string]interface{})
	reqData, marshalErr := json.Marshal(&req)
	if marshalErr != nil {
		return nil, config.New(config.InnerMarshalError)
	}
	unMarshalErr := json.Unmarshal(reqData, &mapData)
	if unMarshalErr != nil {
		return nil, config.New(config.InnerUnmarshalError)
	}

	return mapData, nil
}
