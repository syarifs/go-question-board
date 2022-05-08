package repository

import (
	"fmt"
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type questionnaireRepository struct {
	db *gorm.DB
}

func NewQuestionnaireRepository(db *gorm.DB) *questionnaireRepository {
	return &questionnaireRepository{db: db}
}

func (repo questionnaireRepository)CreateQuest(quest m.Questionnaire) (err error) {
	err = repo.db.Create(&quest).Error
	return
}

func (repo questionnaireRepository) ListMyQuest(user_id int) (quests *[]m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).Find(&quests).Error
	return
}

func (repo questionnaireRepository) AvailableQuest(tag []uint) (quests *[]m.Questionnaire, err error) {
	err = repo.db.Debug().Preload("Tags").Where("id IN (?)",
		repo.db.Table("questionnaire_tags").Select("questionnaire_id").Where("tag_id in ?", tag)).Find(&quests).Error
	return
}

func (repo questionnaireRepository) UpdateQuest(quest m.Questionnaire) (err error) {
	err = repo.db.Debug().Updates(&quest).Error
	if err == nil {
		err = repo.db.Debug().Model(&quest).Association("Tags").Replace(&quest.Tags)
		fmt.Println(err)
		err = repo.db.Debug().Model(&quest).Association("Question").Replace(&quest.Question)
		fmt.Println(err)
	}
	return
}

func (repo questionnaireRepository) DeleteQuest(id int) (err error) {
	err = repo.db.Delete(&m.Questionnaire{}, id).Error
	return
}

func (repo questionnaireRepository) ViewQuestByID(id int) (quest *m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).Find(&quest, id).Error
	return
}
