package utils

import "go-question-board/internal/core/models"

func GetTagByName(name string, tags []models.Tag) (value string) {
	for _, v := range tags {
		if v.Name == name {
			value = v.Value
			return
		}
	}
	return
}
