package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/service"
	"github.com/tapeds/fp-pbkk-golang/utils"
)

type (
	AdminController interface {
		GetPenerbangan(ctx *gin.Context)
		GetBandara(ctx *gin.Context)
		GetMaskapai(ctx *gin.Context)
		AddBandara(ctx *gin.Context)
		AddMaskapai(ctx *gin.Context)
		AddPenerbangan(ctx *gin.Context)
		EditPenerbangan(ctx *gin.Context)
		EditMaskapai(ctx *gin.Context)
		EditBandara(ctx *gin.Context)
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
	userRole := ctx.MustGet("role").(string)

	fmt.Println(userRole)

	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.GetAllPenerbanganWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_PENERBANGAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_PENERBANGAN,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) GetBandara(ctx *gin.Context) {
	result, err := ac.adminService.GetAllBandara(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_BANDARA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_BANDARA,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) GetMaskapai(ctx *gin.Context) {
	result, err := ac.adminService.GetAllMaskapai(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_MASKAPAI, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_MASKAPAI,
		Data:    result,
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

func (ac *adminController) AddPenerbangan(ctx *gin.Context) {
	var req dto.PenerbanganCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.CreatePenerbangan(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PENERBANGAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_CREATE_PENERBANGAN,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) EditPenerbangan(ctx *gin.Context) {
	var req dto.PenerbanganEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.EditPenerbangan(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_EDIT_PENERBANGAN, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_EDIT_PENERBANGAN,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) EditMaskapai(ctx *gin.Context) {
	var req dto.MaskapaiEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.EditMaskapai(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_EDIT_MASKAPAI, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_EDIT_MASKAPAI,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (ac *adminController) EditBandara(ctx *gin.Context) {
	var req dto.BandaraEditRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := ac.adminService.EditBandara(ctx.Request.Context(), req)

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_EDIT_BANDARA, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_EDIT_BANDARA,
		Data:    result,
	}

	ctx.JSON(http.StatusOK, resp)
}
