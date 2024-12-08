package service

import (
	"context"

	"github.com/tapeds/fp-pbkk-golang/dto"
	"github.com/tapeds/fp-pbkk-golang/entity"
	"github.com/tapeds/fp-pbkk-golang/repository"
)

type (
	AdminService interface {
		GetAllPenerbanganWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.PenerbanganPaginationResponse, error)
		CreateBandara(ctx context.Context, req dto.BandaraCreateRequest) (dto.BandaraResponse, error)
		CreateMaskapai(ctx context.Context, req dto.MaskapaiCreateRequest) (dto.MaskapaiResponse, error)
	}

	adminService struct {
		adminRepo  repository.AdminRepository
		jwtService JWTService
	}
)

func NewAdminService(adminRepo repository.AdminRepository, jwtService JWTService) AdminService {
	return &adminService{
		adminRepo:  adminRepo,
		jwtService: jwtService,
	}
}

func (as *adminService) GetAllPenerbanganWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.PenerbanganPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.PenerbanganPaginationResponse{}, err
	}

	var datas []dto.PenerbanganResponse
	for _, penerbangan := range dataWithPaginate.Penerbangans {
		data := dto.PenerbanganResponse{
			ID:              penerbangan.ID,
			NoPenerbangan:   penerbangan.NoPenerbangan,
			JadwalBerangkat: penerbangan.JadwalBerangkat,
			JadwalDatang:    penerbangan.JadwalDatang,
			Harga:           penerbangan.Harga,
			Kapasitas:       penerbangan.Kapasitas,
		}

		datas = append(datas, data)
	}

	return dto.PenerbanganPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (as *adminService) CreateBandara(ctx context.Context, req dto.BandaraCreateRequest) (dto.BandaraResponse, error) {
	_, flag, _ := as.adminRepo.CheckBandaraCode(ctx, nil, req.Kode)
	if flag {
		return dto.BandaraResponse{}, dto.ErrBandaraAlreadyExists
	}

	bandara := entity.Bandara{
		Name: req.Name,
		Kota: req.Kota,
		Kode: req.Kode,
	}

	bandaraReg, err := as.adminRepo.CreateBandara(ctx, nil, bandara)

	if err != nil {
		return dto.BandaraResponse{}, dto.ErrCreateBandara
	}

	return dto.BandaraResponse{
		ID:   bandaraReg.ID,
		Name: bandara.Name,
		Kode: bandara.Kode,
		Kota: bandara.Kota,
	}, nil
}

func (as *adminService) CreateMaskapai(ctx context.Context, req dto.MaskapaiCreateRequest) (dto.MaskapaiResponse, error) {
	maskapai := entity.Maskapai{
		Name:  req.Name,
		Image: req.Image,
	}

	maskapaiReg, err := as.adminRepo.CreateMaskapai(ctx, nil, maskapai)

	if err != nil {
		return dto.MaskapaiResponse{}, dto.ErrCreateMaskapai
	}

	return dto.MaskapaiResponse{
		ID:    maskapaiReg.ID,
		Name:  maskapaiReg.Name,
		Image: maskapaiReg.Image,
	}, nil
}
