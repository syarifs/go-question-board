package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/repository"
	"go-question-board/internal/framework/routes"
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/mocktesting"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)


func TestGetMyQuest(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)

	user := response.User{
		ID: 1,
		Name: "Teacher",
	}
	body, _ := json.Marshal(user)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		password, _ := utils.HashPassword("password")
		
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Questionnaire", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", password, 1, 1, 1)
		rowQuestion := sqlmock.NewRows([]string{"id", "questionnaire_id", "question"}).
			AddRow(1, 1, "Quest 1")
		rowAnswer := sqlmock.NewRows([]string{"id", "question_id", "answer"}).
			AddRow(1, 1, "Answer 1")
		rowUserComplete := sqlmock.NewRows([]string{"questionnaire_id", "user_id"})
		rowTags := sqlmock.NewRows([]string{"questionnaire_id", "tag_id"})

		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE created_by = ?").
			WithArgs(1).
			WillReturnRows(rowQuest)
		mock.ExpectQuery("SELECT * FROM `quest_user_complete` WHERE `quest_user_complete`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowUserComplete)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").
			WithArgs(1).
			WillReturnRows(rowCreator)
		mock.ExpectQuery("SELECT * FROM `question` WHERE `question`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowQuestion)
		mock.ExpectQuery("SELECT * FROM `questionnaire_tags` WHERE `questionnaire_tags`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowTags)
		mock.ExpectQuery("SELECT * FROM `answer_option` WHERE `answer_option`.`question_id` = ?").
			WithArgs(1).
			WillReturnRows(rowAnswer)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/quest")
		
		assert.NoError(t, controller.MyQuest(ctx))
		assert.Equal(t, 200, rec.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE created_by = ?").
			WithArgs(0).
			WillReturnError(gorm.ErrRecordNotFound)
		
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE type = 'Questionnaire' ORDER BY `questionnaire`.`id` LIMIT 1").
			WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/quest")
		
		assert.NoError(t, controller.MyQuest(ctx))
		assert.Equal(t, 404, rec.Code)
	})
}

func TestGetQuestForMe(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)

	user := response.UserDetails{
		ID: 1,
		Name: "Teacher",
		Tags: []models.Tag{
			{
				ID: 1,
				Name: "Year",
				Value: "2019",
			},
		},
	}

	body, _ := json.Marshal(user)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		password, _ := utils.HashPassword("password")
		
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Questionnaire", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", password, 1, 1, 1)
		rowCreatorRole := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Admin")
		rowQuestion := sqlmock.NewRows([]string{"id", "questionnaire_id", "question"}).
			AddRow(1, 1, "Quest 1")
		rowUserComplete := sqlmock.NewRows([]string{"questionnaire_id", "user_id"})
		rowTags := sqlmock.NewRows([]string{"questionnaire_id", "tag_id"})

		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE created_by != ? AND type != 'Evaluate' AND id IN (SELECT questionnaire_id FROM `questionnaire_tags` WHERE tag_id IN (?))").
			WithArgs(1, 1).
			WillReturnRows(rowQuest)
		mock.ExpectQuery("SELECT * FROM `quest_user_complete` WHERE `quest_user_complete`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowUserComplete)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").
			WithArgs(1).
			WillReturnRows(rowCreator)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").
			WithArgs(1).
			WillReturnRows(rowCreatorRole)
		mock.ExpectQuery("SELECT * FROM `question` WHERE `question`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowQuestion)
		mock.ExpectQuery("SELECT * FROM `questionnaire_tags` WHERE `questionnaire_tags`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowTags)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/quest/available")
		
		assert.NoError(t, controller.QuestForMe(ctx))
		assert.Equal(t, 200, rec.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE created_by != ? AND type != 'Evaluate' AND id IN (SELECT questionnaire_id FROM `questionnaire_tags` WHERE tag_id IN (?))").
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/quest/available")
		
		assert.NoError(t, controller.QuestForMe(ctx))
		assert.Equal(t, 404, rec.Code)
	})
}

func TestCreateQuestionnaire(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)

	ans := models.AnswerOption{ 
		Answer: "A",
	}

	quest := models.Question {
		Question:        "Test Quest",
		AnswerOption:    []models.AnswerOption{ans},
	}

	data := models.Questionnaire{
		ID:          0,
		Title:       "Test",
		Description: "Test",
		Tags:        nil,
		Type:        "Evaluate",
		Question:    []models.Question{quest},
		CreatedBy:   1,
		Completor:   nil,
	}
	body, _ := json.Marshal(data)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		passing := sqlmock.AnyArg()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `questionnaire` (`title`,`description`,`type`,`created_by`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?)").
			WithArgs(data.Title, data.Description, data.Type, data.CreatedBy, passing, passing, passing).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `question` (`questionnaire_id`,`question`) VALUES (?,?) ON DUPLICATE KEY UPDATE `questionnaire_id`=VALUES(`questionnaire_id`)").
			WithArgs(1, quest.Question).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `answer_option` (`question_id`,`answer`) VALUES (?,?) ON DUPLICATE KEY UPDATE `question_id`=VALUES(`question_id`)").
			WithArgs(1, ans.Answer).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/quest")

		assert.NoError(t, controller.CreateQuest(ctx))
		assert.Equal(t, 201, rec.Code)
		})
}

func TestUpdateQuestionnaire(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)


	ans := models.AnswerOption{ 
		Answer: "A",
		QuestionID: 1,
	}

	quest := models.Question {
		QuestionnaireID: 1,
		Question:        "Test Quest",
		AnswerOption:    []models.AnswerOption{ans},
	}

	data := models.Questionnaire{
		ID:          1,
		Title:       "Test",
		Description: "Test",
		Tags:        nil,
		Type:        "Evaluate",
		Question:    []models.Question{quest},
		CreatedBy:   1,
		Completor:   nil,
	}

	body, _ := json.Marshal(data)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		passing := sqlmock.AnyArg()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `questionnaire` SET `title`=?,`description`=?,`type`=?,`created_by`=?,`updated_at`=? WHERE `id` = ?").
			WithArgs(data.Title, data.Description, data.Type, data.CreatedBy, passing, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `question` (`questionnaire_id`,`question`) VALUES (?,?) ON DUPLICATE KEY UPDATE `questionnaire_id`=VALUES(`questionnaire_id`)").
			WithArgs(quest.QuestionnaireID, quest.Question).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `answer_option` (`question_id`,`answer`) VALUES (?,?) ON DUPLICATE KEY UPDATE `question_id`=VALUES(`question_id`)").
			WithArgs(ans.QuestionID, ans.Answer).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `questionnaire` SET `updated_at`=? WHERE `id` = ?").
			WithArgs(passing, data.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `questionnaire_tags` WHERE `questionnaire_tags`.`questionnaire_id` = ?").
			WithArgs(data.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `question` WHERE `question`.`questionnaire_id` = ?").
			WithArgs(data.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `question` (`questionnaire_id`,`question`) VALUES (?,?)").
			WithArgs(1, quest.Question).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `answer_option` (`question_id`,`answer`) VALUES (?,?) ON DUPLICATE KEY UPDATE `question_id`=VALUES(`question_id`)").
			WithArgs(ans.QuestionID, ans.Answer).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/quest/:id/update")

		assert.NoError(t, controller.UpdateQuest(ctx))
		assert.Equal(t, 200, rec.Code)
		})
}

func TestDeleteQuestionnaire(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `questionnaire` WHERE `questionnaire`.`id` = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/quest/:id/delete")

		assert.NoError(t, controller.DeleteQuest(ctx))
		assert.Equal(t, 200, rec.Code)
		})
}

func TestAnswerQuestionnaire(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewQuestionnaireRepository(gdb)
	service := service.NewQuestionnaireService(mockRepo)
	controller := controller.NewQuestionnaireController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewQuestionnaireRoutes(api, controller)

	quest := models.Questionnaire{
		ID: 1,
		Title: "Test Quest",
		Description: "Test Quest",
		Tags: []models.Tag{
			{ID: 1, Name: "Year", Value: "2019"},
		},
		Question: []models.Question{
			{Question: "Test Quest 1"},
		},
		CreatedBy: 1,
		Completor: []models.User{{ID: 1},},
	}
	userAnswer := []request.UserAnswer{
		{
			QuestionID: 1,
			Answer: "A",
		},
	}
	answer := request.Answer{
		Questionnaire: quest,
		Answer: userAnswer,
		User: models.User{ID: 1},
	}
	
	body, _ := json.Marshal(answer)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		passing := sqlmock.AnyArg()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `user_answers` (`answer`,`question_id`,`user_id`,`evaluate_teacher_id`) VALUES (?,?,?,?)").
			WithArgs(userAnswer[0].Answer, userAnswer[0].QuestionID, answer.User.ID, nil).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `questionnaire` SET `updated_at`=? WHERE `id` = ?").
			WithArgs(passing, quest.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `quest_user_complete` (`questionnaire_id`,`user_id`) VALUES (?,?),(?,?) ON DUPLICATE KEY UPDATE `questionnaire_id`=`questionnaire_id`").
			WithArgs(answer.Questionnaire.ID, answer.User.ID, answer.Questionnaire.ID, answer.User.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/quest/answer")

		assert.NoError(t, controller.QuestAnswer(ctx))
		assert.Equal(t, 200, rec.Code)
		})
}
