package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"haha/internal/logger"
	"haha/internal/models"
	"strings"
	"time"
)

const (
	salt       = "akfdjlskjfweoi324d"
	signingKey = "kdnjsndjnd*jdnj212md"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo Authorization
	logg *logger.Logger
}

func NewAuthService(repo Authorization, logg *logger.Logger) *AuthService {
	return &AuthService{repo: repo, logg: logg}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func (s *AuthService) CreateUser(name, email, telegram, password, role string) (uint, error) {
	user := models.User{
		Name:        name,
		Email:       email,
		Telegram:    telegram,
		Password:    password,
		Role:        role,
		Status:      "",
		Description: "",
		Vacancies:   make([]models.Vacancy, 0),
		Resumes:     make([]models.Resume, 0),
	}

	if strings.EqualFold(user.Role, models.APPLICANT) {
		user.Status = models.ACTIVE
	}

	user.Password = generatePasswordHash(user.Password)

	_, err := s.repo.GetOne(user.Email, user.Password)
	if err == nil {
		s.logg.Error("email is not free")
		return 0, errors.New("email is not free")
	}

	return s.repo.Create(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, string, error) {
	user, err := s.repo.GetOne(email, generatePasswordHash(password))
	if err != nil {
		s.logg.Error(err)
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
		user.ID,
	})

	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		s.logg.Error(err)
		return "", "", err
	}

	return tokenStr, user.Role, err
}

func (s *AuthService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logg.Error("invalid signing method")
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		s.logg.Error(err)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		s.logg.Error("token claims are not of type *tokenClaims")
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *AuthService) GetUser(id uint) (models.User, error) {
	return s.repo.GetOneById(id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
