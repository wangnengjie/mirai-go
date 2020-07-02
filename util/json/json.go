package json

import (
	jsoniter "github.com/json-iterator/go"
)

type RawMessage jsoniter.RawMessage

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal = Json.Marshal
	Unmarshal = Json.Unmarshal
	MarshalToString = Json.MarshalToString
	NewDecoder = Json.NewDecoder
	NewEncoder = Json.NewEncoder
	Get = Json.Get
)