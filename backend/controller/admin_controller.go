package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
	"github.com/tapeds/fp-pbkk-golang/utils"
)

type (
	AdminController interface {
		GetPenerbangan(ctx *gin.Context)
		AddBandara(ctx *gin.Context)
		AddMaskapai(ctx *gin.Context)
	}

	adminController struct {
		adminService service.AdminService
	}
)

func NewAdminController(as service.AdminService) AdminController {
	return &adminController{
		adminService: as,
	}
}

func (ac *adminController) GetPenerbangan(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.GetAllPenerbanganWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) AddBandara(ctx *gin.Context) {
	var req dto.BandaraCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.CreateBandara(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_BANDARA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_CREATE_BANDARA,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) AddMaskapai(ctx *gin.Context) {
	var req dto.MaskapaiCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.CreateMaskapai(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_MASKAPAI, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_CREATE_MASKAPAI,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}
