package controller_test

import (
	"bytes"
	"encoding/json"
	"go-question-board/internal/core/entity/models"
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

func TestGetMajor(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewMajorRepository(gdb)
	service := service.NewMajorService(mockRepo)
	controller := controller.NewMajorController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewMajorRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		majorRow := sqlmock.NewRows([]string{"id", "code", "name"}).
			AddRow(1, "INF", "Informatics")

		mock.ExpectQuery("SELECT * FROM `majors`").WillReturnRows(majorRow)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/major")
		
		assert.NoError(t, controller.ReadMajor(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `majors`").WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/major")

		assert.NoError(t, controller.ReadMajor(ctx))
		assert.Equal(t, 404, rec.Code)
	})
}

func TestCreateMajor(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewMajorRepository(gdb)
	service := service.NewMajorService(mockRepo)
	controller := controller.NewMajorController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewMajorRoutes(api, controller)
	major := models.Major{
		Code: "INF",
		Name: "Informatics",
	}
	body, _ := json.Marshal(major)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `majors` (`code`,`name`) VALUES (?,?)").
			WithArgs(major.Code, major.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/major")
		
		assert.NoError(t, controller.CreateMajor(ctx))
		assert.Equal(t, 201, rec.Code)
	})
}

func TestUpdateMajor(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewMajorRepository(gdb)
	service := service.NewMajorService(mockRepo)
	controller := controller.NewMajorController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewMajorRoutes(api, controller)
	major := models.Major{
		Code: "INF",
		Name: "Informatics",
	}
	body, _ := json.Marshal(major)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `majors` SET `code`=?,`name`=? WHERE `id` = ?").
			WithArgs(major.Code, major.Name, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/major/:id/update")
		
		assert.NoError(t, controller.UpdateMajor(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}

func TestDeleteMajor(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewMajorRepository(gdb)
	service := service.NewMajorService(mockRepo)
	controller := controller.NewMajorController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewMajorRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `majors` WHERE `majors`.`id` = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/major/:id/delete")
		
		assert.NoError(t, controller.DeleteMajor(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}
