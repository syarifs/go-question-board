package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils/errors"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (srv TagService) CreateTag(tag models.Tag) (err error) {
	err  = srv.repo.CreateTag(tag)
	err = errors.CheckError(nil, err)
	return
}

func (srv TagService) ReadTag() (res *[]models.Tag, err error) {
	res, err  = srv.repo.ReadTag()

	err = errors.CheckError(res, err)

	return
}

func (srv TagService) UpdateTag(id int, tag models.Tag) (err error) {
	tag.ID = uint(id)
	err  = srv.repo.UpdateTag(tag)

	err = errors.CheckError(nil, err)

	return
}

func (srv TagService) DeleteTag(id int) (err error) {
	err  = srv.repo.DeleteTag(id)

	err = errors.CheckError(nil, err)

	return
}
