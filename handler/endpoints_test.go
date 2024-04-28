package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var UUIDEstate = "0b689246-27b5-4fd3-8413-c719c8350ce6"
var UUIDEstateObject = "04d9ca2a-b026-45b1-ac55-8914cb0a861c"

func TestPostEstate(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(`{"width":10, "length":10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	opts := NewServerOptions{
		Repository: mockRepo,
	}
	h := NewServer(opts)

	gomock.InOrder(
		mockRepo.EXPECT().InsertEstate(c.Request().Context(), repository.InsertEstateInput{Width: 10, Length: 10}).Return(UUIDEstate, nil),
	)

	if assert.NoError(t, h.PostEstate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":\""+UUIDEstate+"\"}\n", rec.Body.String())
	}
}

func TestPostEstateIdTree(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/"+UUIDEstate+"/tree", strings.NewReader(`{"x":10, "y":10, "height":20}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	opts := NewServerOptions{
		Repository: mockRepo,
	}
	h := NewServer(opts)

	gomock.InOrder(
		mockRepo.EXPECT().GetEstateById(c.Request().Context(), UUIDEstate).Return(repository.GetEstateByIdOutput{Id: UUIDEstate, Length: 11, Width: 11}, nil),
		mockRepo.EXPECT().InsertEstateObject(c.Request().Context(), UUIDEstate, repository.InsertEstateObjectInput{X: 10, Y: 10, Height: 20}).Return(UUIDEstateObject, nil),
	)

	if assert.NoError(t, h.PostEstateIdTree(c, UUIDEstate)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":\""+UUIDEstateObject+"\"}\n", rec.Body.String())
	}
}

func TestGetEstateIdStats(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/"+UUIDEstate+"/stats", strings.NewReader(`{"x":10, "y":10, "height":20}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	opts := NewServerOptions{
		Repository: mockRepo,
	}
	h := NewServer(opts)

	gomock.InOrder(
		mockRepo.EXPECT().GetEstateById(c.Request().Context(), UUIDEstate).Return(repository.GetEstateByIdOutput{Id: UUIDEstate, Length: 11, Width: 11}, nil),
		mockRepo.EXPECT().GetEstateStats(c.Request().Context(), UUIDEstate).Return(repository.GetEstateStatsOutput{Count: 10, Max: 20, Median: 10, Min: 5}, nil),
	)

	if assert.NoError(t, h.GetEstateIdStats(c, UUIDEstate)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"count\":10,\"max\":20,\"median\":10,\"min\":5}\n", rec.Body.String())
	}
}

func TestGetEstateIdDronePlan(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/"+UUIDEstate+"/drone-plan", strings.NewReader(`{"max_distance":10, "y":10, "height":20}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	opts := NewServerOptions{
		Repository: mockRepo,
	}
	h := NewServer(opts)

	gomock.InOrder(
		mockRepo.EXPECT().GetEstateById(c.Request().Context(), UUIDEstate).Return(repository.GetEstateByIdOutput{Id: UUIDEstate, Length: 11, Width: 11}, nil),
		mockRepo.EXPECT().GetDronePlanByEstateId(c.Request().Context(), UUIDEstate).Return(120, nil),
	)

	if assert.NoError(t, h.GetEstateIdDronePlan(c, UUIDEstate, generated.GetEstateIdDronePlanParams{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"distance\":120}\n", rec.Body.String())
	}
}
