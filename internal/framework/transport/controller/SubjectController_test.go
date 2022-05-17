package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/repository"
	"go-question-board/internal/framework/routes"
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/utils/mocktesting"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetSubject(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewSubjectRepository(gdb)
	service := service.NewSubjectService(mockRepo)
	controller := controller.NewSubjectController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewSubjectRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")

		mock.ExpectQuery("SELECT * FROM `subjects`").WillReturnRows(rowSubject)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/subject")
		
		assert.NoError(t, controller.ReadSubject(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `subjects`").WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/subject")

		assert.NoError(t, controller.ReadSubject(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}

func TestGetSubjectByID(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewSubjectRepository(gdb)
	service := service.NewSubjectService(mockRepo)
	controller := controller.NewSubjectController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewSubjectRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")
		rowTeacher := sqlmock.NewRows([]string{"teacher_id", "subject_id"})
		rowStudent := sqlmock.NewRows([]string{"teacher_id", "subject_id"})

		mock.ExpectQuery("SELECT * FROM `subjects` WHERE `subjects`.`id` = ? ORDER BY `subjects`.`id` LIMIT 1").WithArgs(1).WillReturnRows(rowSubject)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`subject_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`subject_id` = ?").WithArgs(1).WillReturnRows(rowTeacher)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/subject/:id")
		
		assert.NoError(t, controller.ReadSubjectByID(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `subjects` WHERE `subjects`.`id` = ? ORDER BY `subjects`.`id` LIMIT 1").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`subject_id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`subject_id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/subject/:id")

		assert.NoError(t, controller.ReadSubjectByID(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}

func TestGetTeacherSubject(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}
	
	user := request.Teacher{
		ID: 1,
		Name: "Admin",
	}

	mockRepo := repository.NewSubjectRepository(gdb)
	service := service.NewSubjectService(mockRepo)
	controller := controller.NewSubjectController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewSubjectRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))

		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")
		rowTeacher := sqlmock.NewRows([]string{"subject_id"}).AddRow(1)
		rowStudent := sqlmock.NewRows([]string{"user_id", "subject_id"}).AddRow(2, 1)
		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(2, "Student A", "student@web.io", "this is password", 1, 1, 1)
		rowMajor := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "INF", "Informatics")
		rowUserTags := sqlmock.NewRows([]string{"tag_id", "user_id"}).AddRow(1, 2)
		rowTags := sqlmock.NewRows([]string{"id", "name", "value"}).AddRow(1, "Class", "A")

		mock.ExpectQuery("SELECT subject_id FROM `teacher_subjects` WHERE user_id = ?").WithArgs(1).WillReturnRows(rowTeacher)
		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id IN (?)").WithArgs(1).WillReturnRows(rowSubject)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`subject_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").WithArgs(2).WillReturnRows(rowUsers)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnRows(rowMajor)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(2).WillReturnRows(rowUserTags)
		mock.ExpectQuery("SELECT * FROM `tags` WHERE `tags`.`id` = ?").WithArgs(1).WillReturnRows(rowTags)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/teacher/subject")
		
		assert.NoError(t, controller.ReadTeacherSubject(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT subject_id FROM `teacher_subjects` WHERE user_id = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id IN (?)").WithArgs(nil).WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/teacher/subject")

		assert.NoError(t, controller.ReadTeacherSubject(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}

func TestGetStudentSubject(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}
	
	user := request.Teacher{
		ID: 1,
		Name: "Admin",
	}

	mockRepo := repository.NewSubjectRepository(gdb)
	service := service.NewSubjectService(mockRepo)
	controller := controller.NewSubjectController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewSubjectRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rowSubject := sqlmock.NewRows([]string{"id", "code", "name"}).AddRow(1, "BSPGM", "Basic Programming")
		rowSubjectTeacher := sqlmock.NewRows([]string{"id", "user_id", "subject_id", "class"}).AddRow(1, 2, 1, "A")
		rowUserTags := sqlmock.NewRows([]string{"value"}).AddRow("A")
		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(2, "Student A", "student@web.io", "this is password", 1, 1, 1)

		mock.ExpectQuery("SELECT `value` FROM `tags` WHERE id IN (SELECT tag_id FROM `user_tags` WHERE user_id = ?) AND name = 'Class'").
			WithArgs(1).WillReturnRows(rowUserTags)
		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id IN (SELECT subject_id FROM `student_subject` WHERE user_id = ?)").WithArgs(1).WillReturnRows(rowSubject)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`subject_id` = ? AND class = ?").
			WithArgs(1, "A").WillReturnRows(rowSubjectTeacher)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ?").WithArgs(2).WillReturnRows(rowUsers)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/student/subject")
		
		assert.NoError(t, controller.ReadStudentSubject(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT `value` FROM `tags` WHERE id IN (SELECT tag_id FROM `user_tags` WHERE user_id = ?) AND name = 'Class'").
			WithArgs(0).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `subjects` WHERE id IN (SELECT subject_id FROM `student_subject` WHERE user_id = ?)").
			WithArgs(0).WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/student/subject")

		assert.NoError(t, controller.ReadStudentSubject(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}
