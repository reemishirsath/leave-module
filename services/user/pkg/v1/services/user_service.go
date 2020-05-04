package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/reemishirsath/leave-module/logs"
	"golang.org/x/crypto/bcrypt"

	repository "github.com/reemishirsath/leave-module/services/user/pkg/v1/repositories"

	"github.com/reemishirsath/leave-module/models"
)

// UsersService describes the service.
type UsersService interface {
	Login(context.Context, models.UserLoginRequest) (models.UserLoginResponse, error)
	ApplyLeave(ctx context.Context, leave models.Leave) (interface{}, error)
	GetLeaves(context.Context, int) (models.GetAllLeavesResponse, error)
	GetLeavesByStatus(context.Context, int, string) (models.GetAllLeavesResponse, error)
}

//UsersServiceImpl **
type UsersServiceImpl struct {
	userRepo repository.UserRepository
}

//NewUserServiceImpl inject depedancies user repositiory
func NewUserServiceImpl(userRepo repository.UserRepository) UsersService {
	return &UsersServiceImpl{userRepo: userRepo}
}

//Login **
func (b *UsersServiceImpl) Login(ctx context.Context, user models.UserLoginRequest) (resp models.UserLoginResponse, err error) {
	log.Logger(ctx).Info("LoginUser")
	userData, err := b.userRepo.Login(ctx, user)
	if err != nil {
		return resp, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return resp, errors.New("Invalid login details")
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.JWTClaims{
		UserID: userData.ID,
		RoleID: userData.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenDetails := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenDetails.SignedString([]byte("95a31f74-a1cd-4321-8a9d-bdb0735e445a"))

	if err != nil {
		return resp, errors.New("Invalid login details")
	}

	resp.Token = token
	return resp, nil
}

//ApplyLeave **
func (b *UsersServiceImpl) ApplyLeave(ctx context.Context, leave models.Leave) (interface{}, error) {

	log.Logger(ctx).Info("ApplyLeave ", leave)
	leave.Status = "APPLIED"

	// calculate years, month, days and time betwen dates
	leaveDays := calcDays(leave.FromDate, leave.ToDate)

	// calculate total number of days
	duration := leave.ToDate.Sub(leave.FromDate)
	leaveDays = int(duration.Hours() / 24)
	leave.LeaveDays = leaveDays

	resp, err := b.userRepo.ApplyLeave(ctx, leave)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func calcDays(a, b time.Time) (day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	_, M2, d2 := b.Date()
	month := int(M2 - M1)
	day = int(d2 - d1)
	// Normalize negative values
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	return
}

//GetLeaves **
func (b *UsersServiceImpl) GetLeaves(ctx context.Context, id int) (resp models.GetAllLeavesResponse, err error) {
	log.Logger(ctx).Info(" GetLeaves")

	resp, err = b.userRepo.GetLeaves(ctx, id)

	if err != nil {
		return resp, err
	}

	return resp, err
}

//GetLeavesByStatus **
func (b *UsersServiceImpl) GetLeavesByStatus(ctx context.Context, id int, status string) (resp models.GetAllLeavesResponse, err error) {
	log.Logger(ctx).Info(" GetLeavesByStatus")

	resp, err = b.userRepo.GetLeavesByStatus(ctx, id, status)

	if err != nil {
		return resp, err
	}

	return resp, err
}
