package gorm

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainuser "github.com/mlevshin/my-finance-go-clean/internal/domain/user"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"gorm.io/gorm"
)

type userEntity struct {
	Base
	Name         string `gorm:"column:name;size:128;not null;"`
	Email        string `gorm:"column:email;size:128;not null;unique;"`
	PasswordHash []byte `gorm:"column:password_hash;not null;"`
}

func (*userEntity) TableName() string {
	return "users"
}

type userRw struct {
	db *gorm.DB
}

func NewUserRw(db *gorm.DB) (rw.UserRW, error) {
	err := db.AutoMigrate(&userEntity{})
	if err != nil {
		return nil, err
	}
	return &userRw{db: db}, nil
}

func (u *userRw) FindAll() ([]*domainuser.User, error) {
	var entities []userEntity
	if err := u.db.Find(&entities).Error; err != nil {
		return nil, err
	}

	users := make([]*domainuser.User, 0, len(entities))
	for _, entity := range entities {
		users = append(users, entity.mapToDomain())
	}
	return users, nil
}

func (u *userRw) FindById(id domain.Id) (*domainuser.User, error) {
	user := userEntity{}
	if err := u.db.Where("id = ?", uuid.UUID(id)).First(&user).Error; err != nil {
		return nil, err
	}
	return user.mapToDomain(), nil
}

func (u *userRw) Save(user domainuser.User) error {
	userToSave := mapToEntity(&user)
	err := u.db.Save(*userToSave).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRw) FindByEmail(email string) (*domainuser.User, error) {
	user := userEntity{}
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user.mapToDomain(), nil
}
