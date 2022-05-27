package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/entity/models"
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

func TestGetUser(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewUserRepository(gdb)
	service := service.NewUserService(mockRepo)
	controller := controller.NewUserController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewUserRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		password, _ := utils.HashPassword("password")

		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Administrator", "admin@web.io", password, 1, 1, 1)

		rowRole := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Administrator")
		rowMajor := sqlmock.NewRows([]string{"id", "code", "name"})
		rowStudent := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTeacher := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTags := sqlmock.NewRows([]string{"tag_id", "user_id"})

		mock.ExpectQuery("SELECT * FROM `users`").WillReturnRows(rowUsers)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnRows(rowRole)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnRows(rowMajor)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`user_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTags)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTeacher)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/user")
		
		assert.NoError(t, controller.ReadUser(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `users`").WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`user_id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`user_id` = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/user")

		assert.NoError(t, controller.ReadUser(ctx))
		assert.Equal(t, 404, rec.Code)
	})
}

func TestGetUserByID(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewUserRepository(gdb)
	service := service.NewUserService(mockRepo)
	controller := controller.NewUserController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewUserRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		
		password, _ := utils.HashPassword("password")

		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Administrator", "admin@web.io", password, 1, 1, 1)

		rowRole := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Administrator")
		rowMajor := sqlmock.NewRows([]string{"id", "code", "name"})
		rowStudent := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTeacher := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowUserTags := sqlmock.NewRows([]string{"user_id", "tag_id"})
		rowTags := sqlmock.NewRows([]string{"value"})

		mock.ExpectQuery("SELECT `value` FROM `tags` WHERE id IN (SELECT tag_id FROM `user_tags` WHERE user_id = ?) AND name = 'Class'").WithArgs(1).WillReturnRows(rowTags)
		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1").WithArgs(1).WillReturnRows(rowUsers)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnRows(rowRole)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnRows(rowMajor)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`user_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(1).WillReturnRows(rowUserTags)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTeacher)
		mock.ExpectQuery("SELECT `value` FROM `tags` WHERE id IN (SELECT tag_id FROM `user_tags` WHERE user_id = ?) AND name = 'Class'").WithArgs(1).WillReturnRows(rowTags)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/user/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		
		assert.NoError(t, controller.ReadUserByID(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/user/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		assert.NoError(t, controller.ReadUserByID(ctx))
		assert.Equal(t, 404, rec.Code)
	})
}

func TestCreateUser(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewUserRepository(gdb)
	service := service.NewUserService(mockRepo)
	controller := controller.NewUserController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewUserRoutes(api, controller)

	major := 1
	user := models.User{
		ID:             0,
		Email:          "user@web.io",
		Name:           "Unit",
		Password:       "testing",
		Status:         1,
		LevelID:        3,
		Tags:           nil,
		MajorID:        &major,
	}
	body, _ := json.Marshal(user)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		pass := sqlmock.AnyArg()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `users` (`email`,`name`,`password`,`status`,`level_id`,`major_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?)").
			WithArgs(user.Email, user.Name, pass, user.Status, user.LevelID, user.MajorID, pass, pass, pass).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/user")
		
		assert.NoError(t, controller.CreateUser(ctx))
		assert.Equal(t, 201, rec.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewUserRepository(gdb)
	service := service.NewUserService(mockRepo)
	controller := controller.NewUserController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewUserRoutes(api, controller)

	major := 1
	user := models.User{
		ID:             1,
		Email:          "user@web.io",
		Name:           "Unit",
		Password:       "testing",
		Status:         1,
		LevelID:        3,
		Tags:           nil,
		MajorID:        &major,
	}
	body, _ := json.Marshal(user)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		pass := sqlmock.AnyArg()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users` SET `email`=?,`name`=?,`password`=?,`status`=?,`level_id`=?,`major_id`=?,`updated_at`=? WHERE `id` = ?").
			WithArgs(user.Email, user.Name, pass, user.Status, user.LevelID, user.MajorID, pass, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users` SET `updated_at`=? WHERE `id` = ?").
			WithArgs(pass, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `user_tags` WHERE `user_tags`.`user_id` = ?").
			WithArgs(user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users` SET `updated_at`=? WHERE `id` = ?").
			WithArgs(pass, user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `student_subject` WHERE `student_subject`.`user_id` = ?").
			WithArgs(user.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/user/:id/update")
		
		assert.NoError(t, controller.UpdateUser(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewUserRepository(gdb)
	service := service.NewUserService(mockRepo)
	controller := controller.NewUserController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewUserRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `users` WHERE `users`.`id` = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/user/:id/delete")
		
		assert.NoError(t, controller.DeleteUser(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}
