package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func sendErrorFormat(ctx echo.Context, code int, message string) error {
	petErr := generated.ErrorResponseApi{
		Code:    code,
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

func (s *Server) PostEstate(ctx echo.Context) error {
	var resp generated.UuidResponse
	var params generated.PostEstateJSONBody
	ctx.Bind(&params)
	if (params.Length < 1 && params.Length > 50000) || (params.Width < 1 && params.Width > 50000) || params.Length == 0 || params.Width == 0 {
		return sendErrorFormat(ctx, http.StatusBadRequest, "Internal Error")
	}
	input := repository.InsertEstateInput(params)

	id, err := s.Repository.InsertEstate(ctx.Request().Context(), input)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusInternalServerError, "Internal Error")
	}
	resp.Id = &id
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstateIdTree(ctx echo.Context, id string) error {
	var resp generated.UuidResponse
	var params generated.PostEstateIdTreeJSONBody

	estate, err := s.Repository.GetEstateById(ctx.Request().Context(), id)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusNotFound, "Estate Not Found")
	}

	ctx.Bind(&params)
	if params.X < 1 || params.X > estate.Length || params.Y < 1 || params.Y > estate.Width {
		return sendErrorFormat(ctx, http.StatusBadRequest, "Bad Client Request")
	}

	input := repository.InsertEstateObjectInput(params)

	idObject, err := s.Repository.InsertEstateObject(ctx.Request().Context(), id, input)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusBadRequest, "Bad Client Request")
	}
	resp.Id = &idObject
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateIdStats(ctx echo.Context, id string) error {
	estate, err := s.Repository.GetEstateById(ctx.Request().Context(), id)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusNotFound, "Estate Not Found")
	}

	stats, err := s.Repository.GetEstateStats(ctx.Request().Context(), estate.Id)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusInternalServerError, "Something Went Wrong")
	}

	resp := generated.StatsResponse{Count: &stats.Count, Median: &stats.Median, Max: &stats.Max, Min: &stats.Min}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id string, param generated.GetEstateIdDronePlanParams) error {
	estate, err := s.Repository.GetEstateById(ctx.Request().Context(), id)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusNotFound, "Estate Not Found")
	}

	plan, err := s.Repository.GetDronePlanByEstateId(ctx.Request().Context(), estate.Id)
	if err != nil {
		return sendErrorFormat(ctx, http.StatusInternalServerError, "Something Went Wrong")
	}

	resp := generated.DronePlanResponse{Distance: &plan}
	return ctx.JSON(http.StatusOK, resp)
}
