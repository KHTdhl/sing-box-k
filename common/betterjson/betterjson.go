package betterjson

import (
	"github.com/hjson/hjson-go/v4"
	"github.com/sagernet/sing/common/json"
	"gopkg.in/yaml.v3"
)

func convertYamlToJSON(content []byte) ([]byte, error) {
	var mapRaw map[string]any
	if err := yaml.Unmarshal(content, &mapRaw); err != nil {
		return nil, err
	}
	mapClear := make(map[string]any)
	for _, key := range []string{"$schema", "log", "dns", "ntp", "inbounds", "outbounds", "route", "outbound_providers", "experimental"} {
		if value, ok := mapRaw[key]; ok {
			mapClear[key] = value
		}
	}

	return json.Marshal(&mapClear)
}

func convertHjsonToJSON(content []byte) ([]byte, error) {
	var mapRaw map[string]any
	if err := hjson.Unmarshal(content, &mapRaw); err != nil {
		return nil, err
	}
	return json.Marshal(&mapRaw)
}

func PreConvert(content []byte) ([]byte, error) {
	if parsedContent, err := convertHjsonToJSON(content); err == nil {
		return parsedContent, nil
	}
	return convertYamlToJSON(content)
}
