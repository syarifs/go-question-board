package controller_test

import (
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
)

func TestGetTag(t *testing.T) {
	gdb, mock, err := mocktesting.InitGORMSQLMock()
	if err != nil {
		panic(err)
	}

	mockRepo := repository.NewTagRepository(gdb)
	service := service.NewTagService(mockRepo)
	controller := controller.NewTagController(service)

	tagRow := sqlmock.NewRows([]string{"id", "name", "value"}).
		AddRow(1, "Year", "2019")

	mock.ExpectQuery("SELECT * FROM `tags`").WillReturnRows(tagRow)

	e := echo.New()
	api := e.Group("/api")

	routes.NewTagRoutes(api, controller)
	
	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/tag")
		
		assert.NoError(t, controller.ReadTag(ctx))
		assert.Equal(t, 200, rec.Code)
	})
	
	t.Run("Test No Data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/admin/tag")

		assert.NoError(t, controller.ReadTag(ctx))
		assert.Equal(t, 417, rec.Code)
		t.Log(rec.Code)
	})
}
