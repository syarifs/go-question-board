package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/repository"
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
