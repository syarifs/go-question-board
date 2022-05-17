package controller_test

import (
	"bytes"
	"encoding/json"
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

func TestGetEvaluateQuest(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewEvaluateRepository(gdb)
	service := service.NewEvaluateService(mockRepo)
	controller := controller.NewEvaluateController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewEvaluateRoutes(api, controller)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Evaluate", 1)
		rowTeacherSubject := sqlmock.NewRows([]string{"id", "user_id", "subject_id", "class"}).
			AddRow(1, 2, 1, "A")
		rowJoinStudent := sqlmock.NewRows([]string{"id", "user_id", "subject_id"}).
			AddRow(2, 3, 1)
		rowTeacher := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(2, "Teacher", "teacher@web.io", utils.HashPassword("password"), 1, 1, 1)
		rowStudent := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(3, "Student", "student@web.io", utils.HashPassword("password"), 1, 1, 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", utils.HashPassword("password"), 1, 1, 1)
		rowQuestion := sqlmock.NewRows([]string{"id", "questionnaire_id", "question"}).
			AddRow(1, 1, "Quest 1")
		rowAnswer := sqlmock.NewRows([]string{"id", "question_id", "answer"}).
			AddRow(1, 1, "Answer 1")
		rowUserComplete := sqlmock.NewRows([]string{"questionnaire_id", "user_id"})
		rowTags := sqlmock.NewRows([]string{"questionnaire_id", "tag_id"})

		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id = ? ORDER BY `subjects`.`id` LIMIT 1").
			WithArgs(1).
			WillReturnRows(rowSubject)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`subject_id` = ?").
			WithArgs(1).
			WillReturnRows(rowJoinStudent)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").
			WithArgs(3).
			WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`subject_id` = ? AND class = ?").
			WithArgs(1, "A").
			WillReturnRows(rowTeacherSubject)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").
			WithArgs(2).
			WillReturnRows(rowTeacher)
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE type = 'Evaluate' ORDER BY `questionnaire`.`id` LIMIT 1").
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
		mock.ExpectQuery("SELECT * FROM `answer_option` WHERE `answer_option`.`question_id` = ?").
			WithArgs(1).
			WillReturnRows(rowAnswer)
		mock.ExpectQuery("SELECT * FROM `questionnaire_tags` WHERE `questionnaire_tags`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowTags)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.QueryParams().Add("class", "A")
		ctx.QueryParams().Add("subject_id", "1")
		ctx.SetPath("/api/student/evaluate")
		
		assert.NoError(t, controller.GetQuest(ctx))
		assert.Equal(t, 200, rec.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id = ? ORDER BY `subjects`.`id` LIMIT 1").
			WithArgs(0).
			WillReturnError(gorm.ErrRecordNotFound)
		
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE type = 'Evaluate' ORDER BY `questionnaire`.`id` LIMIT 1").
			WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/student/evaluate")
		
		assert.NoError(t, controller.GetQuest(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}

func TestGetEvaluateResponse(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewEvaluateRepository(gdb)
	service := service.NewEvaluateService(mockRepo)
	controller := controller.NewEvaluateController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewEvaluateRoutes(api, controller)

	user := response.UserList{
		ID: 1,
		Name: "Teacher",
	}
	body, _ := json.Marshal(user)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		rowEvaluate := sqlmock.NewRows([]string{"id"}).AddRow(1)
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Evaluate", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", utils.HashPassword("password"), 1, 1, 1)
		rowQuestion := sqlmock.NewRows([]string{"id", "questionnaire_id", "question"}).
			AddRow(1, 1, "Quest 1")
		rowAnswer := sqlmock.NewRows([]string{"id", "question_id", "answer"}).
			AddRow(1, 1, "Answer 1")
		rowUserAnswer := sqlmock.NewRows([]string{"id", "answer", "question_id", "user_id", "evaluate_teacher_id"})
		rowUserComplete := sqlmock.NewRows([]string{"questionnaire_id", "user_id"})
		rowTags := sqlmock.NewRows([]string{"questionnaire_id", "tag_id"})

		mock.ExpectQuery("SELECT `id` FROM `evaluate_teachers` WHERE teacher_id = ? AND subject_id = ? AND class = ?").
			WithArgs(1, 1, "A").
			WillReturnRows(rowEvaluate)
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE type = 'Evaluate' ORDER BY `questionnaire`.`id` LIMIT 1").
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
		mock.ExpectQuery("SELECT * FROM `answer_option` WHERE `answer_option`.`question_id` = ?").
			WithArgs(1).
			WillReturnRows(rowAnswer)
		mock.ExpectQuery("SELECT * FROM `user_answers` WHERE `user_answers`.`question_id` = ? AND evaluate_teacher_id IN (?)").
			WithArgs(1, 1).
			WillReturnRows(rowUserAnswer)
		mock.ExpectQuery("SELECT * FROM `questionnaire_tags` WHERE `questionnaire_tags`.`questionnaire_id` = ?").
			WithArgs(1).
			WillReturnRows(rowTags)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.QueryParams().Add("class", "A")
		ctx.QueryParams().Add("subject_id", "1")
		ctx.SetPath("/api/teacher/evaluate")
		
		assert.NoError(t, controller.ViewEvaluateResponse(ctx))
		assert.Equal(t, 200, rec.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		mock.ExpectQuery("SELECT `id` FROM `evaluate_teachers` WHERE teacher_id = ? AND subject_id = ? AND class = ?").
			WithArgs(1, 0, "").
			WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `questionnaire` WHERE type = 'Evaluate' ORDER BY `questionnaire`.`id` LIMIT 1").
			WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/teacher/evaluate")
		
		assert.NoError(t, controller.ViewEvaluateResponse(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}
