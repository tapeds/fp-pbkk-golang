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
		GetAllBandara(ctx context.Context) ([]dto.BandaraResponse, error)
		GetAllMaskapai(ctx context.Context) ([]dto.MaskapaiResponse, error)
		CreateBandara(ctx context.Context, req dto.BandaraCreateRequest) (dto.BandaraResponse, error)
		CreateMaskapai(ctx context.Context, req dto.MaskapaiCreateRequest) (dto.MaskapaiResponse, error)
		CreatePenerbangan(ctx context.Context, req dto.PenerbanganCreateRequest) (dto.PenerbanganResponse, error)
		EditPenerbangan(ctx context.Context, req dto.PenerbanganEditRequest) (dto.PenerbanganResponse, error)
		EditMaskapai(ctx context.Context, req dto.MaskapaiEditRequest) (dto.MaskapaiResponse, error)
		EditBandara(ctx context.Context, req dto.BandaraEditRequest) (dto.BandaraResponse, error)
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
		var bandaras []dto.BandaraArahResponse
		for _, bp := range penerbangan.BandaraPenerbangan {
			bandara := dto.BandaraArahResponse{
				ID:   bp.Bandara.ID,
				Name: bp.Bandara.Name,
				Kota: bp.Bandara.Kota,
				Kode: bp.Bandara.Kode,
				Arah: bp.Arah,
			}
			bandaras = append(bandaras, bandara)
		}

		data := dto.PenerbanganResponse{
			ID:              penerbangan.ID,
			NoPenerbangan:   penerbangan.NoPenerbangan,
			JadwalBerangkat: penerbangan.JadwalBerangkat,
			JadwalDatang:    penerbangan.JadwalDatang,
			Harga:           penerbangan.Harga,
			Kapasitas:       penerbangan.Kapasitas,
			Maskapai: dto.MaskapaiResponse{
				ID:    penerbangan.Maskapai.ID,
				Name:  penerbangan.Maskapai.Name,
				Image: penerbangan.Maskapai.Image,
			},
			Bandaras: bandaras,
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

func (as *adminService) GetAllBandara(ctx context.Context) ([]dto.BandaraResponse, error) {
	bandaraData, err := as.adminRepo.GetAllBandara(ctx, nil)

	if err != nil {
		return []dto.BandaraResponse{}, nil
	}

	var bandaras []dto.BandaraResponse
	for _, bandara := range bandaraData {
		bandara := dto.BandaraResponse{
			ID:   bandara.ID,
			Name: bandara.Name,
			Kota: bandara.Kota,
			Kode: bandara.Kode,
		}
		bandaras = append(bandaras, bandara)
	}

	return bandaras, nil
}

func (as *adminService) GetAllMaskapai(ctx context.Context) ([]dto.MaskapaiResponse, error) {
	maskapaiData, err := as.adminRepo.GetAllMaskapai(ctx, nil)

	if err != nil {
		return []dto.MaskapaiResponse{}, nil
	}

	var maskapais []dto.MaskapaiResponse
	for _, maskapai := range maskapaiData {
		maskapai := dto.MaskapaiResponse{
			ID:    maskapai.ID,
			Name:  maskapai.Name,
			Image: maskapai.Image,
		}
		maskapais = append(maskapais, maskapai)
	}

	return maskapais, nil
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

func (as *adminService) CreatePenerbangan(ctx context.Context, req dto.PenerbanganCreateRequest) (dto.PenerbanganResponse, error) {
	if req.Harga <= 0 {
		return dto.PenerbanganResponse{}, dto.ErrPriceBelowZero
	}

	if req.Kapasitas <= 0 {
		return dto.PenerbanganResponse{}, dto.ErrCapacityBelowZero
	}

	if !req.JadwalBerangkat.Before(req.JadwalDatang) {
		return dto.PenerbanganResponse{}, dto.ErrScheduleUnmatch
	}

	if req.BandaraBerangkatID == req.BandaraDatangID {
		return dto.PenerbanganResponse{}, dto.ErrMatchingAirport
	}

	_, flag, _ := as.adminRepo.CheckPenerbanganNumber(ctx, nil, req.NoPenerbangan)
	if flag {
		return dto.PenerbanganResponse{}, dto.ErrPenerbanganAlreadyExists
	}

	var bandaraPenerbangan []entity.BandaraPenerbangan

	bandaraBerangkat := entity.BandaraPenerbangan{
		BandaraID: req.BandaraBerangkatID,
		Arah:      entity.ArahBerangkat,
	}

	bandaraPenerbangan = append(bandaraPenerbangan, bandaraBerangkat)

	bandaraDatang := entity.BandaraPenerbangan{
		BandaraID: req.BandaraDatangID,
		Arah:      entity.ArahDatang,
	}

	bandaraPenerbangan = append(bandaraPenerbangan, bandaraDatang)

	penerbangan := entity.Penerbangan{
		NoPenerbangan:      req.NoPenerbangan,
		JadwalBerangkat:    req.JadwalBerangkat,
		JadwalDatang:       req.JadwalDatang,
		Harga:              req.Harga,
		Kapasitas:          req.Kapasitas,
		BandaraPenerbangan: bandaraPenerbangan,
		MaskapaiID:         req.MaskapaiID,
	}

	penerbanganReg, err := as.adminRepo.CreatePenerbangan(ctx, nil, penerbangan)

	if err != nil {
		return dto.PenerbanganResponse{}, dto.ErrCreateMaskapai
	}

	var bandaras []dto.BandaraArahResponse
	for _, bp := range penerbanganReg.BandaraPenerbangan {
		bandara := dto.BandaraArahResponse{
			ID:   bp.Bandara.ID,
			Name: bp.Bandara.Name,
			Kota: bp.Bandara.Kota,
			Kode: bp.Bandara.Kode,
			Arah: bp.Arah,
		}
		bandaras = append(bandaras, bandara)
	}

	return dto.PenerbanganResponse{
		ID:              penerbanganReg.ID,
		NoPenerbangan:   penerbanganReg.NoPenerbangan,
		JadwalBerangkat: penerbanganReg.JadwalDatang,
		JadwalDatang:    penerbanganReg.JadwalDatang,
		Harga:           penerbanganReg.Harga,
		Kapasitas:       penerbanganReg.Kapasitas,
		Maskapai: dto.MaskapaiResponse{
			ID:    penerbanganReg.Maskapai.ID,
			Name:  penerbanganReg.Maskapai.Name,
			Image: penerbanganReg.Maskapai.Image,
		},
		Bandaras: bandaras,
	}, nil
}

func (as *adminService) EditPenerbangan(ctx context.Context, req dto.PenerbanganEditRequest) (dto.PenerbanganResponse, error) {
	existingPenerbangan, err := as.adminRepo.GetPenerbanganByID(ctx, nil, req.ID)
	if err != nil {
		return dto.PenerbanganResponse{}, dto.ErrPenerbanganNotFound
	}

	if req.Harga <= 0 {
		return dto.PenerbanganResponse{}, dto.ErrPriceBelowZero
	}

	if req.Kapasitas <= 0 {
		return dto.PenerbanganResponse{}, dto.ErrCapacityBelowZero
	}

	if !req.JadwalBerangkat.Before(req.JadwalDatang) {
		return dto.PenerbanganResponse{}, dto.ErrScheduleUnmatch
	}

	if req.BandaraBerangkatID == req.BandaraDatangID {
		return dto.PenerbanganResponse{}, dto.ErrMatchingAirport
	}

	if existingPenerbangan.NoPenerbangan != req.NoPenerbangan {
		_, flag, _ := as.adminRepo.CheckPenerbanganNumber(ctx, nil, req.NoPenerbangan)
		if flag {
			return dto.PenerbanganResponse{}, dto.ErrPenerbanganAlreadyExists
		}
	}

	var bandaraPenerbangan []entity.BandaraPenerbangan

	bandaraBerangkat := entity.BandaraPenerbangan{
		BandaraID: req.BandaraBerangkatID,
		Arah:      entity.ArahBerangkat,
	}

	bandaraPenerbangan = append(bandaraPenerbangan, bandaraBerangkat)

	bandaraDatang := entity.BandaraPenerbangan{
		BandaraID: req.BandaraDatangID,
		Arah:      entity.ArahDatang,
	}

	bandaraPenerbangan = append(bandaraPenerbangan, bandaraDatang)

	penerbangan := entity.Penerbangan{
		NoPenerbangan:      req.NoPenerbangan,
		JadwalBerangkat:    req.JadwalBerangkat,
		JadwalDatang:       req.JadwalDatang,
		Harga:              req.Harga,
		Kapasitas:          req.Kapasitas,
		BandaraPenerbangan: bandaraPenerbangan,
		MaskapaiID:         req.MaskapaiID,
	}

	updatedPenerbangan, err := as.adminRepo.UpdatePenerbangan(ctx, nil, existingPenerbangan.ID, penerbangan)
	if err != nil {
		return dto.PenerbanganResponse{}, dto.ErrEditPenerbangan
	}

	var bandaras []dto.BandaraArahResponse
	for _, bp := range updatedPenerbangan.BandaraPenerbangan {
		bandara := dto.BandaraArahResponse{
			ID:   bp.Bandara.ID,
			Name: bp.Bandara.Name,
			Kota: bp.Bandara.Kota,
			Kode: bp.Bandara.Kode,
			Arah: bp.Arah,
		}
		bandaras = append(bandaras, bandara)
	}

	return dto.PenerbanganResponse{
		ID:              updatedPenerbangan.ID,
		NoPenerbangan:   updatedPenerbangan.NoPenerbangan,
		JadwalBerangkat: updatedPenerbangan.JadwalBerangkat,
		JadwalDatang:    updatedPenerbangan.JadwalDatang,
		Harga:           updatedPenerbangan.Harga,
		Kapasitas:       updatedPenerbangan.Kapasitas,
		Maskapai: dto.MaskapaiResponse{
			ID:    updatedPenerbangan.Maskapai.ID,
			Name:  updatedPenerbangan.Maskapai.Name,
			Image: updatedPenerbangan.Maskapai.Image,
		},
		Bandaras: bandaras,
	}, nil
}

func (as *adminService) EditMaskapai(ctx context.Context, req dto.MaskapaiEditRequest) (dto.MaskapaiResponse, error) {
	existingMaskapai, err := as.adminRepo.GetMaskapaiByID(ctx, nil, req.ID)
	if err != nil {
		return dto.MaskapaiResponse{}, dto.ErrMaskapaiNotFound
	}

	maskapai := entity.Maskapai{
		Name:  req.Name,
		Image: req.Image,
	}

	updatedMaskapai, err := as.adminRepo.UpdateMaskapai(ctx, nil, existingMaskapai.ID, maskapai)
	if err != nil {
		return dto.MaskapaiResponse{}, dto.ErrCreateMaskapai
	}

	return dto.MaskapaiResponse{
		ID:    updatedMaskapai.ID,
		Name:  updatedMaskapai.Name,
		Image: updatedMaskapai.Image,
	}, nil
}

func (as *adminService) EditBandara(ctx context.Context, req dto.BandaraEditRequest) (dto.BandaraResponse, error) {
	existingBandara, err := as.adminRepo.GetBandaraByID(ctx, nil, req.ID)
	if err != nil {
		return dto.BandaraResponse{}, dto.ErrMaskapaiNotFound
	}

	maskapai := entity.Bandara{
		Name: req.Name,
		Kode: req.Kode,
		Kota: req.Kota,
	}

	updatedBandara, err := as.adminRepo.UpdateBandara(ctx, nil, existingBandara.ID, maskapai)
	if err != nil {
		return dto.BandaraResponse{}, dto.ErrCreateBandara
	}

	return dto.BandaraResponse{
		ID:   updatedBandara.ID,
		Name: updatedBandara.Name,
		Kode: updatedBandara.Kode,
		Kota: updatedBandara.Kota,
	}, nil
}
