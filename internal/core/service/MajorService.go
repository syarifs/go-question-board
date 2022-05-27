package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils/errors"
)

type MajorService struct {
	repo repository.MajorRepository
}

func NewMajorService(repo repository.MajorRepository) *MajorService {
	return &MajorService{repo: repo}
}

func (srv MajorService) CreateMajor(major models.Major) (err error) {
	err  = srv.repo.CreateMajor(major)

	err = errors.CheckError(nil, err)

	return
}

func (srv MajorService) ReadMajor() (res *[]models.Major, err error) {
	res, err  = srv.repo.ReadMajor()

	err = errors.CheckError(res, err)

	return
}

func (srv MajorService) UpdateMajor(id int, major models.Major) (err error) {
	major.ID = uint(id)
	err  = srv.repo.UpdateMajor(major)

	err = errors.CheckError(nil, err)

	return
}

func (srv MajorService) DeleteMajor(id int) (err error) {
	err  = srv.repo.DeleteMajor(id)

	err = errors.CheckError(nil, err)

	return
}
