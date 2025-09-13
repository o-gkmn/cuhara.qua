package app

import (
	"gorm.io/gorm"

	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/infra/db"
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
