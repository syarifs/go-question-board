package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/repository"
)

type MajorService struct {
	repo repository.MajorRepository
}

func NewMajorService(repo repository.MajorRepository) *MajorService {
	return &MajorService{repo: repo}
}

func (srv MajorService) CreateMajor(major models.Major) (err error) {
	err  = srv.repo.CreateMajor(major)
	return
}

func (srv MajorService) ReadMajor() (res []models.Major, err error) {
	var major *[]models.Major
	major, err  = srv.repo.ReadMajor()
	if err == nil {
		for _, um := range *major {
			res = append(res, um)
		}
	}
	return
}

func (srv MajorService) UpdateMajor(id int, major models.Major) (err error) {
	major.ID = uint(id)
	err  = srv.repo.UpdateMajor(major)
	return
}

func (srv MajorService) DeleteMajor(id int) (err error) {
	err  = srv.repo.DeleteMajor(id)
	return
}
