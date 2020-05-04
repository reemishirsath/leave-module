package models

//Role Model
type Role struct {
	ID       int    `gorm:"id" json:"id"`
	RoleName string `gorm:"role_name" json:"roleName"`
}
