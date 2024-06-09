package user_domain

import (
	"context"

	"github.com/Upsiloner/UniTrend/domain"
)

const (
	CollectionUser = "users"
)

type User struct {
	domain.BaseModel
	ID        int    `gorm:"column:id; primary_key; not null" json:"-"`
	Name      string `gorm:"column:name;unique" json:"name"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
	Status    int    `gorm:"column:status" json:"status"`
	AvatarUrl string `gorm:"column:avatar_url" json:"avatar_url"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	GetUserByName(c context.Context, name string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	GetUserByUnionID(c context.Context, union_id string) (User, error)
	Update(c context.Context, user *User) error
	ChangeUserPwd(c context.Context, user *User, newPassword string) error
	UpdateUserAvatar(c context.Context, user *User, avatarUrl string) error
}
