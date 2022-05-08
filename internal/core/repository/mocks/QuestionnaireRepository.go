// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "go-question-board/internal/core/models"

	mock "github.com/stretchr/testify/mock"
)

// QuestionnaireRepository is an autogenerated mock type for the QuestionnaireRepository type
type QuestionnaireRepository struct {
	mock.Mock
}

// AvailableQuest provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) AvailableQuest(_a0 []uint) (*[]models.Questionnaire, error) {
	ret := _m.Called(_a0)

	var r0 *[]models.Questionnaire
	if rf, ok := ret.Get(0).(func([]uint) *[]models.Questionnaire); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Questionnaire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQuest provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) CreateQuest(_a0 models.Questionnaire) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Questionnaire) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteQuest provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) DeleteQuest(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMyQuest provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) ListMyQuest(_a0 int) (*[]models.Questionnaire, error) {
	ret := _m.Called(_a0)

	var r0 *[]models.Questionnaire
	if rf, ok := ret.Get(0).(func(int) *[]models.Questionnaire); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Questionnaire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuest provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) UpdateQuest(_a0 models.Questionnaire) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Questionnaire) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ViewQuestByID provides a mock function with given fields: _a0
func (_m *QuestionnaireRepository) ViewQuestByID(_a0 int) (*models.Questionnaire, error) {
	ret := _m.Called(_a0)

	var r0 *models.Questionnaire
	if rf, ok := ret.Get(0).(func(int) *models.Questionnaire); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Questionnaire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
