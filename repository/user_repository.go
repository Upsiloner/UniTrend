package repository

import (
	"context"
	"errors"

	"github.com/Upsiloner/UniTrend/domain/user_domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) user_domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *user_domain.User) error {
	return ur.database.Create(user).Error
}

func (ur *userRepository) GetUserByName(c context.Context, name string) (user_domain.User, error) {
	var user user_domain.User
	err := ur.database.Where("name = ?", name).First(&user).Error
	return user, err
}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (user_domain.User, error) {
	var user user_domain.User
	err := ur.database.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *userRepository) GetUserByUnionID(c context.Context, union_id string) (user_domain.User, error) {
	var user user_domain.User
	err := ur.database.Where("union_id = ?", union_id).First(&user).Error
	return user, err
}

func (ur *userRepository) Update(c context.Context, user *user_domain.User) error {
	return ur.database.Save(user).Error
}

func (ur *userRepository) ChangeUserPwd(c context.Context, user *user_domain.User, newPassword string) error {
	if user.Password == newPassword {
		return errors.New("新密码不能与旧密码相同")
	}
	user.Password = newPassword
	return ur.Update(c, user)
}

func (ur *userRepository) UpdateUserAvatar(c context.Context, user *user_domain.User, avatarUrl string) error {
	user.AvatarUrl = avatarUrl
	return ur.Update(c, user)
}
