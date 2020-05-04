package repositories

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	log "github.com/reemishirsath/leave-module/logs"
	"github.com/reemishirsath/leave-module/models"
)

//UserRepository implimets all methods in UserRepository
type UserRepository interface {
	Login(context.Context, models.UserLoginRequest) (models.User, error)
	ProcessLeave(context.Context, int, models.Leave) (models.Leave, error)
	GetLeavesByUser(context.Context, int) (models.GetLeavesResponse, error)
	GetLeavesByStatus(context.Context, int, string) (models.GetLeavesResponse, error)
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

//ProcessLeave **
func (userRepositoryImpl UserRepositoryImpl) ProcessLeave(ctx context.Context, id int, leave models.Leave) (models.Leave, error) {
	dbConn := userRepositoryImpl.dbConn

	txn := dbConn.Begin()
	leave1 := leave
	err := txn.Table("leaves").Where("id=?", id).Find(&leave).Error
	if err != nil {
		txn.Rollback()
		return leave, err
	}
	if leave.Status == "APPROVED" {
		err = txn.Model(&models.User{}).Where("id=? AND !total_leaves < ?", leave.UserID, leave.LeaveDays).Update("total_leaves", gorm.Expr("total_leaves-?", leave.LeaveDays)).Error
		if err != nil {
			log.Logger(ctx).Error(err)
			txn.Rollback()
			return leave, err
		}
	}
	err = txn.Table("leaves").Where("id=?", id).Update(&leave1).Error
	if err != nil {
		txn.Rollback()
		return leave, err
	}
	txn.Commit()
	leave1.FromDate = leave.FromDate
	leave1.ToDate = leave.ToDate
	return leave1, nil
}

//GetLeavesByUser **
func (userRepositoryImpl UserRepositoryImpl) GetLeavesByUser(ctx context.Context, id int) (resp models.GetLeavesResponse, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("users").Where("id=?", id).First(&resp.User).Error
	if err != nil {
		return resp, err
	}
	err = dbConn.Table("leaves").Where("user_id=?", id).Find(&resp.Leaves).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

//GetLeavesByStatus **
func (userRepositoryImpl UserRepositoryImpl) GetLeavesByStatus(ctx context.Context, id int, status string) (resp models.GetLeavesResponse, err error) {
	dbConn := userRepositoryImpl.dbConn
	err = dbConn.Table("users").Where("id=?", id).First(&resp.User).Error
	if err != nil {
		return resp, err
	}
	err = dbConn.Table("leaves").Where("user_id=? AND status = ?", id, status).Find(&resp.Leaves).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}
