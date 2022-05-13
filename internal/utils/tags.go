package utils

import (
	"encoding/json"
	"go-question-board/internal/core/models"
)

func GetTagByName(name string, tags []models.Tag) (value string) {
	for _, v := range tags {
		if v.Name == name {
			value = v.Value
			return
		}
	}
	return
}

func TypeConverter[R any](data any) (res *R, err error) {
	var b []byte
	b, err = json.Marshal(&data)
	if err != nil {
		res = nil
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	return
}
