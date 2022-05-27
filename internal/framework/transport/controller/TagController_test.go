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

func TestGetTag(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewTagRepository(gdb)
	service := service.NewTagService(mockRepo)
	controller := controller.NewTagController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewTagRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		tagRow := sqlmock.NewRows([]string{"id", "name", "value"}).
			AddRow(1, "Year", "2019")

		mock.ExpectQuery("SELECT * FROM `tags`").WillReturnRows(tagRow)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/tag")
		
		assert.NoError(t, controller.ReadTag(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectQuery("SELECT * FROM `tags`").WillReturnError(gorm.ErrRecordNotFound)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/tag")

		assert.NoError(t, controller.ReadTag(ctx))
		assert.Equal(t, 404, rec.Code)
		t.Log(rec.Code)
	})
}

func TestCreateTag(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewTagRepository(gdb)
	service := service.NewTagService(mockRepo)
	controller := controller.NewTagController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewTagRoutes(api, controller)
	tag := models.Tag{
		Name: "Year",
		Value: "2019",
	}
	body, _ := json.Marshal(tag)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `tags` (`name`,`value`) VALUES (?,?)").
			WithArgs(tag.Name, tag.Value).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/tag")
		
		assert.NoError(t, controller.CreateTag(ctx))
		assert.Equal(t, 201, rec.Code)
	})
}

func TestUpdateTag(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewTagRepository(gdb)
	service := service.NewTagService(mockRepo)
	controller := controller.NewTagController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewTagRoutes(api, controller)
	tag := models.Tag{
		Name: "Year",
		Value: "2019",
	}
	body, _ := json.Marshal(tag)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `tags` SET `name`=?,`value`=? WHERE `id` = ?").
			WithArgs(tag.Name, tag.Value, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/tag/:id/update")
		
		assert.NoError(t, controller.UpdateTag(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}

func TestDeleteTag(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewTagRepository(gdb)
	service := service.NewTagService(mockRepo)
	controller := controller.NewTagController(service)

	e := echo.New()
	api := e.Group("/api")

	routes.NewTagRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `tags` WHERE `tags`.`id` = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath("/api/admin/tag/:id/delete")
		
		assert.NoError(t, controller.DeleteTag(ctx))
		assert.Equal(t, 200, rec.Code)
	})
}
