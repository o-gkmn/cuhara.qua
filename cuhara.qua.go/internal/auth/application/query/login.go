package query

import (
	"context"
	"fmt"
	"time"

	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/infra/config"
	"cuhara.qua.go/internal/readmodel/user"
	"github.com/golang-jwt/jwt/v5"
)

const LoginQueryType = "auth.login"

type LoginRequest struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginQuery struct {
	cqrs.BaseQuery[LoginResponse]

	Email string `json:"email"`
}

func NewLoginQuery(r LoginRequest) *LoginQuery {
	query := cqrs.NewBaseQuery[LoginResponse](LoginQueryType)
	return &LoginQuery{BaseQuery: query, Email: r.Email}
}

type LoginHandler struct {
	userReadRepository *user.UserReadRepository
	cfg                *config.Config
}

func NewLoginHandler(userReadRepository *user.UserReadRepository, cfg *config.Config) *LoginHandler {
	return &LoginHandler{userReadRepository: userReadRepository, cfg: cfg}
}

func (h *LoginHandler) Handle(ctx context.Context, q cqrs.Query[LoginResponse]) (LoginResponse, error) {
	loginQuery, ok := q.(*LoginQuery)
	if !ok {
		return LoginResponse{}, fmt.Errorf("invalid query type: expected *LoginQuery")
	}

	foundUser, err := h.userReadRepository.GetByEmail(loginQuery.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	token, err := generateToken(foundUser, h.cfg)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{Token: token}, nil

}

func generateToken(foundUser *user.UserDTO, configuration *config.Config) (string, error) {
	if foundUser == nil || configuration == nil || configuration.JWTSecret == "" {
		return "", fmt.Errorf("invalid input for token generation")
	}

	now := time.Now().UTC()
	exp := now.Add(time.Duration(configuration.JWTTTLMinute) * time.Minute)

	claims := jwt.MapClaims{
		"sub":       fmt.Sprintf("%d", foundUser.ID),
		"email":     foundUser.Email,
		"roleId":    foundUser.Role.ID,
		"tennantId": foundUser.Tennant.ID,
		"iss":       configuration.JWTIssuer,
		"iat":       now.Unix(),
		"exp":       exp.Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString([]byte(configuration.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt: %w", err)
	}
	return signed, nil
}
