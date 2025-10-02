package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/modules/auth"
	"cuhara.qua.go/internal/modules/claim"
	"cuhara.qua.go/internal/modules/role"
	tenant "cuhara.qua.go/internal/modules/tennant"
	"cuhara.qua.go/internal/modules/topic"
	"cuhara.qua.go/internal/modules/user"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Router struct {
	Routes        []*echo.Route
	Root          *echo.Group
	APIV1Auth     *echo.Group
	APIV1Users    *echo.Group
	APIV1Roles    *echo.Group
	APIV1Tennants *echo.Group
	APIV1Topics   *echo.Group
	APIV1Claims   *echo.Group
	APIV1SubTopics *echo.Group
}

type Server struct {
	Config  config.Server
	DB      *sql.DB
	Echo    *echo.Echo
	Router  *Router
	Auth    AuthService
	User    UserService
	Role    RoleService
	Tennant TennantService
	Topic   TopicService
	Claim   ClaimService
}

type AuthService interface {
	Login(context.Context, dto.LoginRequest) (dto.LoginResponse, error)
	Register(context.Context, dto.RegisterRequest) (dto.LoginResponse, error)
}

type UserService interface {
	GetUsers(context.Context) ([]dto.UserDTO, error)
	Update(context.Context, dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
	Delete(context.Context, dto.DeleteUserRequest) (dto.DeleteUserResponse, error)
}

type RoleService interface {
	Create(context.Context, dto.CreateRoleRequest) (dto.CreateRoleResponse, error)
	Update(context.Context, dto.UpdateRoleRequest) (dto.UpdateRoleResponse, error)
	Delete(context.Context, dto.DeleteRoleRequest) (dto.DeleteRoleResponse, error)
	GetRoles(context.Context) ([]dto.RoleDTO, error)
}

type TennantService interface {
	Create(context.Context, dto.CreateTenantRequest) (dto.CreateTenantResponse, error)
	Update(context.Context, dto.UpdateTenantRequest) (dto.UpdateTenantResponse, error)
	Delete(context.Context, dto.DeleteTenantRequest) (dto.DeleteTenantResponse, error)
	GetAll(context.Context) ([]dto.TenantDTO, error)
}

type TopicService interface {
	GetAll(context.Context) ([]dto.TopicDTO, error)
	Create(context.Context, dto.CreateTopicRequest) (dto.CreateTopicResponse, error)
	Update(context.Context, dto.UpdateTopicRequest) (dto.UpdateTopicResponse, error)
	Delete(context.Context, dto.DeleteTopicRequest) (dto.DeleteTopicResponse, error)
	GetSubTopics(context.Context, dto.GetSubTopicsRequest) ([]dto.SubTopicDTO, error)
	CreateSubTopic(context.Context, dto.CreateSubTopicRequest) (dto.CreateSubTopicResponse, error)
	UpdateSubTopic(context.Context, dto.UpdateSubTopicRequest) (dto.UpdateSubTopicResponse, error)
	DeleteSubTopic(context.Context, dto.DeleteSubTopicRequest) (dto.DeleteSubTopicResponse, error)
}

type ClaimService interface {
	GetAll(context.Context) ([]dto.ClaimDTO, error)
	Create(context.Context, dto.CreateClaimRequest) (dto.CreateClaimResponse, error)
	Update(context.Context, dto.UpdateClaimRequest) (dto.UpdateClaimResponse, error)
	Delete(context.Context, dto.DeleteClaimRequest) (dto.DeleteClaimResponse, error)
}


func NewServer(config config.Server) *Server {
	s := &Server{
		Config:  config,
		DB:      nil,
		Echo:    nil,
		Router:  nil,
		Auth:    nil,
		User:    nil,
		Role:    nil,
		Tennant: nil,
		Topic:   nil,
		Claim:   nil,
	}

	return s
}

func (s *Server) Ready() bool {
	return s.DB != nil &&
		s.Echo != nil &&
		s.Router != nil &&
		s.Auth != nil &&
		s.User != nil &&
		s.Role != nil &&
		s.Tennant != nil &&
		s.Topic != nil &&
		s.Claim != nil
}

func (s *Server) InitCmd() *Server {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err := s.InitDB(ctx); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	cancel()

	if err := s.InitAuthService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize auth service")
	}

	if err := s.InitUserService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize user service")
	}

	if err := s.InitRoleService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize role service")
	}

	if err := s.InitTennantService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize tennant service")
	}

	if err := s.InitTopicService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize topic service")
	}

	if err := s.InitClaimService(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize claim service")
	}

	return s
}

func (s *Server) InitAuthService() error {
	s.Auth = auth.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitUserService() error {
	s.User = user.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitRoleService() error {
	s.Role = role.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitTennantService() error {
	s.Tennant = tenant.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitTopicService() error {
	s.Topic = topic.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitClaimService() error {
	s.Claim = claim.NewService(s.Config, s.DB)

	return nil
}

func (s *Server) InitDB(ctx context.Context) error {
	connStr := s.Config.Database.ConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if s.Config.Database.MaxIdleConns > 0 {
		db.SetMaxIdleConns(s.Config.Database.MaxIdleConns)
	}
	if s.Config.Database.MaxOpenConns > 0 {
		db.SetMaxOpenConns(s.Config.Database.MaxOpenConns)
	}
	if s.Config.Database.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(s.Config.Database.ConnMaxLifetime)
	}

	s.DB = db

	return nil
}

func (s *Server) Start() error {
	if !s.Ready() {
		return errors.New("server is not ready")
	}

	return s.Echo.Start(s.Config.Echo.ListenAddress)
}

func (s *Server) Shutdown(ctx context.Context) []error {
	log.Warn().Msg("Shutting down server")

	var errs []error

	if s.DB != nil {
		log.Debug().Msg("Closing database connection")

		if err := s.DB.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if s.Echo != nil {
		log.Debug().Msg("Shutting down echo server")

		if err := s.Echo.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed to shutdown echo server")
			errs = append(errs, err)
		}
	}

	return errs
}
