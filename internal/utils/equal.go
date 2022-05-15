package utils

import (
	"go-question-board/internal/core/models"
	"reflect"
)

func TagEqual(user_tag, quest_tag []models.Tag) bool {

	if len(user_tag) > len(quest_tag) {
		checkEqual(user_tag, quest_tag)
	}

	if len(quest_tag) > len(user_tag) {
		checkEqual(quest_tag, user_tag)
	}

	if len(quest_tag) == len(user_tag) {
		if !reflect.DeepEqual(quest_tag, user_tag) {
			return false
		}
	}

	return true
}

func checkEqual(slice1, slice2 []models.Tag) bool {
	for _, t := range slice1 {
		for _, t2 := range slice2 {
			if t != t2 {
				return false
			}
		}
	}
	return true
}
