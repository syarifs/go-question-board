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
		assert.Equal(t, 417, rec.Code)
	})
}
