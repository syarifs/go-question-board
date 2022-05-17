package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
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

	user := response.UserList{
		ID: 1,
		Name: "Teacher",
	}
	body, _ := json.Marshal(user)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Questionnaire", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", utils.HashPassword("password"), 1, 1, 1)
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
		assert.Equal(t, 417, rec.Code)
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
		
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Questionnaire", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", utils.HashPassword("password"), 1, 1, 1)
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
		assert.Equal(t, 417, rec.Code)
	})
}
