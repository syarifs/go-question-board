package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/repository"
	"go-question-board/internal/framework/routes"
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/mocktesting"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var Token string

func TestLogin(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	loginBody := request.LoginRequest{
		Email: "admin@web.io",
		Password: "admin",
	}

	password, _ := utils.HashPassword(loginBody.Password)

	mockAuth := repository.NewAuthRepository(gdb, nil)
	authService := service.NewAuthService(mockAuth)
	authController := controller.NewAuthController(authService)

	e := echo.New()
	api := e.Group("/api")

	routes.NewAuthRoutes(api, authController)
	
	t.Run("Success", func(t *testing.T) {
		body, _ := json.Marshal(loginBody)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Administrator", loginBody.Email, password, 1, 1, 1)

		rowRole := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Administrator")
		rowMajor := sqlmock.NewRows([]string{"id", "code", "name"})
		rowStudent := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTeacher := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTags := sqlmock.NewRows([]string{"tag_id", "user_id"})

		mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").WithArgs(loginBody.Email).WillReturnRows(rowUsers)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnRows(rowRole)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnRows(rowMajor)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`user_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTags)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTeacher)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/login")
		
		assert.NoError(t, authController.Login(ctx))

		var result response.MessageDataJWT
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
				log.Fatalln(err)
		}

		Token = result.JWT.AccessToken

		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test Wrong Password", func(t *testing.T) {

		rowUsers := sqlmock.NewRows([]string{"id", "name", "email", "password", "level_id", "major_id", "status"}).
			AddRow(1, "Administrator", loginBody.Email, password, 1, 1, 1)

		rowRole := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Administrator")
		rowMajor := sqlmock.NewRows([]string{"id", "code", "name"})
		rowStudent := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTeacher := sqlmock.NewRows([]string{"user_id", "subject_id"})
		rowTags := sqlmock.NewRows([]string{"tag_id", "user_id"})

		mock.ExpectQuery("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1").WithArgs(loginBody.Email).WillReturnRows(rowUsers)
		mock.ExpectQuery("SELECT * FROM `roles` WHERE `roles`.`id` = ?").WithArgs(1).WillReturnRows(rowRole)
		mock.ExpectQuery("SELECT * FROM `majors` WHERE `majors`.`id` = ?").WithArgs(1).WillReturnRows(rowMajor)
		mock.ExpectQuery("SELECT * FROM `student_subject` WHERE `student_subject`.`user_id` = ?").WithArgs(1).WillReturnRows(rowStudent)
		mock.ExpectQuery("SELECT * FROM `user_tags` WHERE `user_tags`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTags)
		mock.ExpectQuery("SELECT * FROM `teacher_subjects` WHERE `teacher_subjects`.`user_id` = ?").WithArgs(1).WillReturnRows(rowTeacher)

		loginBody.Password = "password"
		body, _ := json.Marshal(loginBody)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/login")

		assert.NoError(t, authController.Login(ctx))
		assert.Equal(t, 417, rec.Code)
	})
}
