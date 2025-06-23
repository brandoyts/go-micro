package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateAccessToken(userUuid string) (string, error)
	GenerateRefreshToken(userUuid string) (string, error)
	VerifyToken(token string) (*jwt.RegisteredClaims, error)
}

type jwtService struct {
	secretKey            string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey:            secretKey,
		accessTokenDuration:  15 * time.Minute,
		refreshTokenDuration: 7 * 24 * time.Hour,
	}
}

type CustomClaims struct {
	UserUuid string `json:"uuid"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a short-lived JWT
func (j *jwtService) GenerateAccessToken(userUUID string) (string, error) {
	return j.generateToken(userUUID, j.accessTokenDuration)
}

// GenerateRefreshToken creates a long-lived JWT
func (j *jwtService) GenerateRefreshToken(userUUID string) (string, error) {
	return j.generateToken(userUUID, j.refreshTokenDuration)
}

// Internal helper to generate tokens
func (j *jwtService) generateToken(userUUID string, duration time.Duration) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		UserUuid: userUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userUUID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// VerifyToken parses and validates a token string
func (j *jwtService) VerifyToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &claims.RegisteredClaims, nil
}
