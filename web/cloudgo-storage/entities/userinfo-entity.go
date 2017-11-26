package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
    UID        int   `orm:"id,auto-inc"` //语义标签
    UserName   string
    DepartName string
    Created   *time.Time   `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
    if len(u.UserName) == 0 {
        panic("UserName shold not null!")
    }
    if u.Created == nil {
        t := time.Now()
        u.Created = &t
    }
    return &u
}