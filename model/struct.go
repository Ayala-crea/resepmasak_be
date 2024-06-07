package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Users struct {
	IdUser     uint   `gorm:"primaryKey;column:id_user" json:"id_user"`
	IdRole     int    `gorm:"column:id_role" json:"id_role"`
	Nama       string `gorm:"column:nama" json:"nama"`
	Username   string `gorm:"column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	Email      string `gorm:"column:email" json:"email"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

type Roles struct {
	IdRole int    `gorm:"primaryKey;column:id_role" json:"id_role"`
	Nama   string `gorm:"column:nama" json:"nama"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole int  `json:"id_role"`
}
