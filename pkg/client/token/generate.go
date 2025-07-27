package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/bandanascripts/phantompay/pkg/core"
	"github.com/bandanascripts/phantompay/pkg/service/redis"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	UserId string `json:"userid"`
	jwt.RegisteredClaims
}

func GenerateAccess(userId, signingKeyId string, signingKey *rsa.PrivateKey) (string, error) {

	var accessExpiration = time.Now().Add(15 * time.Minute)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &UserClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "localhost:8080",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(accessExpiration),
		},
	})

	accessToken.Header["KID"] = signingKeyId

	accessTokenStr, err := accessToken.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign access token : %w", err)
	}

	return accessTokenStr, nil
}

func GenerateRefresh(userId, signingKeyId string, signingKey *rsa.PrivateKey) (string, error) {

	var refreshExpiration = time.Now().Add(7 * 24 * time.Hour)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &UserClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "localhost:8080",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(refreshExpiration),
		},
	})

	refreshToken.Header["KID"] = signingKeyId

	refreshTokenStr, err := refreshToken.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token : %w", err)
	}

	return refreshTokenStr, nil
}

func GenerateTokens(ctx context.Context, userId string) (string, string, error) {

	activeKeyId, err := redis.GetFromRedis(ctx, "RSA:ACTIVEKEY")

	if err != nil {
		return "", "", err
	}

	privateKey, err := core.FetchAndParsePrivKey(ctx, "RSA:PRIVATEKEY:"+activeKeyId)

	if err != nil {
		return "", "", err
	}

	accessToken, err := GenerateAccess(userId, activeKeyId, privateKey)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefresh(userId, activeKeyId, privateKey)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
