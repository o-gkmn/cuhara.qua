package app

import (
	"gorm.io/gorm"

	authqry "cuhara.qua.go/internal/auth/application/query"
	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/infra/config"
	"cuhara.qua.go/internal/infra/db"
	userread "cuhara.qua.go/internal/readmodel/user"
	rolecmd "cuhara.qua.go/internal/roles/application/command"
	roleinfra "cuhara.qua.go/internal/roles/infrastructure"
	tennantcmd "cuhara.qua.go/internal/tennants/application/command"
	tenantinfra "cuhara.qua.go/internal/tennants/infrastructure"
	usercmd "cuhara.qua.go/internal/users/application/command"
	userinfra "cuhara.qua.go/internal/users/infrastructure"
)

type Deps struct {
	WriteDB *gorm.DB
	ReadDB  *gorm.DB
}

func BuildCommandBus(deps Deps) *cqrs.CommandBus {
	cb := cqrs.NewCommandBus()

	userRepo := userinfra.NewPGUserRepository(deps.WriteDB)
	createUserHandler := usercmd.NewCreateUserHandler(userRepo)
	cqrs.Register(cb, usercmd.CreateUserCommandType, createUserHandler)

	roleRepo := roleinfra.NewPGRoleRepository(deps.WriteDB)
	createRoleHandler := rolecmd.NewCreateRoleHandler(roleRepo)
	cqrs.Register(cb, rolecmd.CreateRoleCommandType, createRoleHandler)

	tennantRepo := tenantinfra.NewPGTennantRepository(deps.WriteDB)
	createTennantHandler := tennantcmd.NewCreateTennantHandler(tennantRepo)
	cqrs.Register(cb, tennantcmd.CreateTennantCommandType, createTennantHandler)

	return cb
}

func BuildQueryBus(deps Deps, cfg *config.Config) *cqrs.QueryBus {
	qb := cqrs.NewQueryBus()

	userRepo := userread.NewUserReadRepository(deps.ReadDB)
	loginHandler := authqry.NewLoginHandler(userRepo, cfg)
	cqrs.RegisterQuery(qb, authqry.LoginQueryType, loginHandler)

	return qb
}

func BuildDeps(writeDBURL, readDBUrl string) (Deps, error) {
	writeDB, err := db.NewGormWriteDB(writeDBURL)
	if err != nil {
		return Deps{}, err
	}

	readDB, err := db.NewGormReadDB(readDBUrl)
	if err != nil {
		return Deps{}, err
	}

	return Deps{
		WriteDB: writeDB,
		ReadDB:  readDB,
	}, nil
}
