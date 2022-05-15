package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (srv TagService) CreateTag(tag models.Tag) (err error) {
	err  = srv.repo.CreateTag(tag)
	return
}

func (srv TagService) ReadTag() (res *[]models.Tag, err error) {
	res, err  = srv.repo.ReadTag()

	if utils.IsEmpty(res) {
		err = errors.New("Data Not Found")
	}

	return
}

func (srv TagService) UpdateTag(id int, tag models.Tag) (err error) {
	tag.ID = uint(id)
	err  = srv.repo.UpdateTag(tag)
	return
}

func (srv TagService) DeleteTag(id int) (err error) {
	err  = srv.repo.DeleteTag(id)
	return
}
