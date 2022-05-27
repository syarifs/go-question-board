package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv UserService) CreateUser(user models.User) (err error) {

	if user.Password != "" {
		user.Password, err = utils.HashPassword(user.Password)
	}

	err = srv.repo.CreateUser(user)

	err = errors.CheckError(nil, err)
	return
}

func (srv UserService) ReadUser() (res *[]response.User, err error) {
	var user *[]models.User
	user, err  = srv.repo.ReadUser()

	if err = errors.CheckError(user, err); err == nil {
		res, _ = utils.TypeConverter[[]response.User](&user)
	}

	return
}

func (srv UserService) ReadUserByID(id int) (res *response.UserDetails, err error) {
	var users *models.User
	users, err  = srv.repo.ReadUserByID(id)

	if err = errors.CheckError(users, err); err == nil {
		res, _ = utils.TypeConverter[response.UserDetails](&users)
	}

	return
}

func (srv UserService) UpdateUser(id int, user models.User) (err error) {
	user.ID = uint(id)

	if user.Password != "" {
		user.Password, err = utils.HashPassword(user.Password)
	}

	if ! utils.IsEmpty(user.TeacherSubject) {
		for i := range user.TeacherSubject {
			user.TeacherSubject[i].SubjectID = user.TeacherSubject[i].Subject.ID
			user.TeacherSubject[i].UserID = &user.ID
		}
	}
	err  = srv.repo.UpdateUser(user)
	err = errors.CheckError(nil, err)
	return
}

func (srv UserService) DeleteUser(id int) (err error) {
	err  = srv.repo.DeleteUser(id)
	err = errors.CheckError(nil, err)
	return
}
