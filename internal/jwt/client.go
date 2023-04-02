package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Client struct {
	logger     logger.StructuredLogger
	duration   time.Duration
	signingKey string
}

func NewClient(logger logger.StructuredLogger, duration time.Duration, signingKey string) *Client {
	return &Client{
		logger:     logger,
		duration:   duration,
		signingKey: signingKey,
	}
}

func (c *Client) GenerateJWTToken(userID int64) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(c.duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(c.signingKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign token: %s", err.Error())
	}

	return &signedToken, nil
}

func (c *Client) ValidateJWTToken(token string) (*int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(c.signingKey), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token: %s", err.Error())
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		exp, ok := claims["exp"].(float64) //nolint:varnamelen
		if !ok {
			return nil, errors.New("failed to parse exp claim")
		}

		exp64 := int64(exp)

		if exp64 < time.Now().Unix() {
			return nil, errors.New("token expired")
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return nil, errors.New("failed to parse user_id claim")
		}

		userID64 := int64(userID)

		return &userID64, nil
	}

	return nil, errors.New("failed to validate token")
}
