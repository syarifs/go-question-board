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

func (srv TagService) CreateTag(tag models.Tag) (res models.Tag, err error) {
	err  = srv.repo.CreateTag(tag)
	res = tag
	return
}

func (srv TagService) ReadTag() (res []models.Tag, err error) {
	var tag *[]models.Tag
	tag, err  = srv.repo.ReadTag()
	if err == nil {
		for _, um := range *tag {
			res = append(res, um)
		}
	}
	return
}

func (srv TagService) UpdateTag(id int, tag models.Tag) (res models.Tag, err error) {
	err  = srv.repo.UpdateTag(id, tag)
	if err == nil {
		res = tag
	}
	return
}

func (srv TagService) DeleteTag(id int) (err error) {
	err  = srv.repo.DeleteTag(id)
	return
}
