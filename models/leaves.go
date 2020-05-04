package models

import "time"

//Leave Model
type Leave struct {
	ID          int       `gorm:"id;PRIMARY_KEY;AUTO_INCREMENT" json:"id,omitempty"`
	UserID      int       `gorm:"user_id" json:"userID,omitempty"`
	Status      string    `gorm:"status" json:"status,omitempty"`
	LeaveReason string    `gorm:"leave_reason" json:"leaveReason,omitempty"`
	Feedback    string    `gorm:"feedback" json:"feedback,omitempty"`
	FromDate    time.Time `gorm:"from_date" json:"fromDate,omitempty"`
	ToDate      time.Time `gorm:"to_date" json:"toDate,omitempty"`
	LeaveDays   int       `gorm:"leave_days" json:"leaveDays,omitempty"`
	CreatedAt   time.Time `gorm:"created_at" json:"-"`
}

//GetLeavesResponse Model
type GetLeavesResponse struct {
	User   User    `json:"user"`
	Leaves []Leave `json:"leaves"`
}

//GetAllLeavesResponse Model
type GetAllLeavesResponse struct {
	Leaves []Leave `json:"leaves"`
}
