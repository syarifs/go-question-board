package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
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

	request := request.User{
		ID:     1,
		Email:  "student@web.io",
		Name:   "Student",
		Level:  models.Level{},
		Tag:    []models.Tag{
			{
				ID: 1,
				Name: "Year",
				Value: "2019",
			},
			{
				ID: 2,
				Name: "Class",
				Value: "A",
			},
		},
		Status: 1,
	}

	t.Run("Success", func(t *testing.T) {
		body, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		password, _ := utils.HashPassword("password")
		
		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Evaluate", 1)
		rowTeacherSubject := sqlmock.NewRows([]string{"id", "user_id", "subject_id", "class"}).
			AddRow(1, 2, 1, "A")
		rowJoinStudent := sqlmock.NewRows([]string{"id", "user_id", "subject_id"}).
			AddRow(2, 3, 1)
		rowTeacher := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(2, "Teacher", "teacher@web.io", password, 1, 1, 1)
		rowStudent := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(3, "Student", "student@web.io", password, 1, 1, 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", password, 1, 1, 1)
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
		ctx.QueryParams().Add("subject_id", "1")
		ctx.SetPath("/api/student/evaluate")
		
		assert.NoError(t, controller.GetQuest(ctx))
		assert.Equal(t, 200, rec.Code)
	})

	t.Run("Fail", func(t *testing.T) {
		request.Tag = nil
		body, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
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
		assert.Equal(t, 404, rec.Code)
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

	request := request.User{
		ID:     1,
		Email:  "student@web.io",
		Name:   "Student",
		Level:  models.Level{},
		Tag:    []models.Tag{},
		Status: 1,
	}
	body, _ := json.Marshal(request)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		password, _ := utils.HashPassword("password")
		
		rowEvaluate := sqlmock.NewRows([]string{"id"}).AddRow(1)
		rowQuest := sqlmock.NewRows([]string{"id", "title", "description", "type", "created_by"}).
			AddRow(1, "Evaluasi Dosen", "Evaluasi setiap semester", "Evaluate", 1)
		rowCreator := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Admin", "admin@web.io", password, 1, 1, 1)
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
		assert.Equal(t, 404, rec.Code)
	})
}

func TestAnswerEvaluate(t *testing.T) {
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

	user := models.User{
		ID: 1,
		Name: "Syarif Ubaidillah",
	}

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
			ID: 1,
			QuestionID: 1,
			Answer: "A",
		},
	}
	evaluate := models.EvaluateTeacher{
		ID: 1,
		SubjectID: 1,
		TeacherID: 1,
		Class: "A",
	}
	answer := request.Answer{
		Questionnaire: quest,
		Answer: userAnswer,
		User: user,
	}
	
	body, _ := json.Marshal(answer)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `evaluate_teachers` (`subject_id`,`teacher_id`,`class`) VALUES (?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`").
			WithArgs(evaluate.SubjectID, evaluate.TeacherID, evaluate.Class).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `user_answers` (`answer`,`question_id`,`user_id`,`evaluate_teacher_id`,`id`) VALUES (?,?,?,?,?)").
			WithArgs(userAnswer[0].Answer, userAnswer[0].QuestionID, answer.User.ID, evaluate.ID, userAnswer[0].ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `evaluate_teachers` (`subject_id`,`teacher_id`,`class`,`id`) VALUES (?,?,?,?)").
			WithArgs(evaluate.SubjectID, evaluate.TeacherID, evaluate.Class, evaluate.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `evaluate_teachers` (`subject_id`,`teacher_id`,`class`,`id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `id`=`id`").
			WithArgs(evaluate.SubjectID, evaluate.TeacherID, evaluate.Class, evaluate.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE `user_answers` SET `evaluate_teacher_id`=? WHERE `id` = ?").
			WithArgs(evaluate.ID, userAnswer[0].ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.QueryParams().Add("teacher_id", "1")
		ctx.QueryParams().Add("subject_id", "1")
		ctx.QueryParams().Add("class", "A")
		ctx.SetPath("/api/student/evaluate")

		assert.NoError(t, controller.QuestAnswer(ctx))
		assert.Equal(t, 200, rec.Code)
		})
}
