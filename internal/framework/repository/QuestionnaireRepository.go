package repository

import (
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

func (repo questionnaireRepository) MyQuest(user_id int) (quests *[]m.Questionnaire, err error) {
	err = repo.db.Debug().Preload(clause.Associations).Where("created_by = ?", user_id).Find(&quests).Error
	return
}

func (repo questionnaireRepository) QuestForMe(id int, tag []int) (quests *[]m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).
		Preload("Creator.Level").
		Where("created_by != ?", id).
		Where("type != 'Evaluate'").
		Where("id IN (?)", repo.db.Table("questionnaire_tags").
			Select("questionnaire_id").Where("tag_id IN ?", tag)).
		Find(&quests).Error
	return
}

func (repo questionnaireRepository) UpdateQuest(quest m.Questionnaire) (err error) {
	err = repo.db.Updates(&quest).Error
	if err == nil {
		err = repo.db.Model(&quest).Association("Tags").Replace(&quest.Tags)
		err = repo.db.Model(&quest).Association("Question").Replace(&quest.Question)
	}
	return
}

func (repo questionnaireRepository) DeleteQuest(id int) (err error) {
	err = repo.db.Delete(&m.Questionnaire{}, id).Error
	return
}

func (repo questionnaireRepository) Answer(quest m.Questionnaire, ans []m.UserAnswer) (err error) {
	err = repo.db.Create(&ans).Error
	if err == nil {
		err = repo.db.Model(&quest).Association("Completor").Append(&quest.Completor)
	}
	return
}

func (repo questionnaireRepository) ViewQuestResponse(id int) (quests *m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).
		Preload("Question.AnswerOption").
		Preload("Question.UserResponse").
		Preload("Question.UserResponse.User").
		Preload("Question.UserResponse.User.Level").
		First(&quests, id).Error
	return
}

func (repo questionnaireRepository) ViewQuestByID(id int) (quest *m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).Preload("Question.AnswerOption").First(&quest, id).Error
	return
}
