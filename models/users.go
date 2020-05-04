package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//User Model
type User struct {
	ID          int        `gorm:"id" json:"id"`
	Name        string     `gorm:"name" json:"name"`
	Email       string     `gorm:"email" json:"email"`
	Password    string     `gorm:"password" json:"-"`
	RoleID      int        `gorm:"role_id" json:"roleID"`
	TotalLeaves int        `gorm:"total_leaves" json:"totalLeaves"`
	CreatedAt   time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"deleted_at" json:"-"`
}

//UserLoginRequest Model
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required" email:"true"`
	Password string `json:"password" validate:"required"`
}

//UserLoginResponse Model
type UserLoginResponse struct {
	Token string `json:"token"`
}

// JWTClaims describes Claims in token.
type JWTClaims struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
	jwt.StandardClaims
}
