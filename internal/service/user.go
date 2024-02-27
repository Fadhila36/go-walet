package service

import (
	"context"
	"encoding/json"
	"go-walet/domain"
	"go-walet/dto"
	"go-walet/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

// Aunthenticate implements domain.UserService.
func (u UserService) Aunthenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.UserRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	token := util.GenerateRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)
	return dto.AuthRes{
		Token: token,
	}, nil
}

// ValidateToken implements domain.UserService.
func (u UserService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, &user)
	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository) domain.UserService {
	return &UserService{
		UserRepository:  userRepository,
		cacheRepository: cacheRepository,
	}

}
