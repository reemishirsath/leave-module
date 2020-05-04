package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/reemishirsath/leave-module/models"
)

//UserRepository implimets all methods in UserRepository
type UserRepository interface {
	Login(context.Context, models.UserLoginRequest) (models.User, error)
	ApplyLeave(context.Context, models.Leave) (models.Leave, error)
	GetLeaves(context.Context, int) (models.GetAllLeavesResponse, error)
	GetLeavesByStatus(context.Context, int, string) (models.GetAllLeavesResponse, error)
}

//UserRepositoryImpl **
type UserRepositoryImpl struct {
	dbConn *gorm.DB
}

//NewUserRepositoryImpl inject dependancies of DataStore
func NewUserRepositoryImpl(dbConn *gorm.DB) UserRepository {
	return &UserRepositoryImpl{dbConn: dbConn}
}

//Login returns jwt token
func (userRepositoryImpl UserRepositoryImpl) Login(ctx context.Context, user models.UserLoginRequest) (userData models.User, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Where("email=?", user.Email).First(&userData).Error
	if err != nil {
		return userData, errors.New("Invalid login details")
	}
	return userData, nil

}

//ApplyLeave create users entry in database
func (userRepositoryImpl UserRepositoryImpl) ApplyLeave(ctx context.Context, leave models.Leave) (models.Leave, error) {
	dbConn := userRepositoryImpl.dbConn

	if err := dbConn.Create(&leave).Error; err != nil {
		return leave, err
	}
	return leave, nil
}

//GetLeaves **
func (userRepositoryImpl UserRepositoryImpl) GetLeaves(ctx context.Context, id int) (resp models.GetAllLeavesResponse, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("leaves").Where("user_id=?", id).Find(&resp.Leaves).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

//GetLeavesByStatus **
func (userRepositoryImpl UserRepositoryImpl) GetLeavesByStatus(ctx context.Context, id int, status string) (resp models.GetAllLeavesResponse, err error) {
	fmt.Println("stauts", status)
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("leaves").Where("user_id=? AND status = ?", id, status).Find(&resp.Leaves).Error
	if err != nil {
		return resp, err
	}
	fmt.Println(resp)
	return resp, err
}
