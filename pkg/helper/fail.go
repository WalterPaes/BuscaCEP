package helper

import "encoding/json"

type Error struct {
	Message string `json:"message"`
}

func (e *Error) ToByte() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return data
}