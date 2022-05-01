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

func (srv TagService) CreateTag(tag models.TagModel) (res models.TagModel, err error) {
	err  = srv.repo.CreateTag(tag)
	res = tag
	return
}

func (srv TagService) ReadTag() (res []models.TagModel, err error) {
	var tag *[]models.TagModel
	tag, err  = srv.repo.ReadTag()
	if err == nil {
		for _, um := range *tag {
			res = append(res, um)
		}
	}
	return
}

func (srv TagService) UpdateTag(id int, tag models.TagModel) (res models.TagModel, err error) {
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
