package db

import (
	"log"

	"gorm.io/gorm"

	roledomain "cuhara.qua.go/internal/roles/domain"
	tennantdomain "cuhara.qua.go/internal/tennants/domain"
	userdomain "cuhara.qua.go/internal/users/domain"
)

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&roledomain.Role{}); err != nil {
		return err
	}

	seedRoles(db)

	if err := db.AutoMigrate(&tennantdomain.Tennant{}); err != nil {
		return err
	}

	seedTennants(db)

	if err := db.AutoMigrate(&userdomain.User{}); err != nil {
		return err
	}

	log.Println("INFO: Database migration completed successfully")
	return nil
}

func seedRoles(db *gorm.DB) error {
	var count int64
	db.Model(&roledomain.Role{}).Count(&count)

	if count == 0 {
		roles := []roledomain.Role{
			{Name: "ADMIN"},
			{Name: "USER"},
		}
		return db.Create(&roles).Error
	}
	return nil
}

func seedTennants(db *gorm.DB) error {
	var count int64
	db.Model(&tennantdomain.Tennant{}).Count(&count)

	if count == 0 {
		tennants := []tennantdomain.Tennant{
			{Name: "KOÇ HOLDİNG AŞ"},
			{Name: "SABANCI HOLDİNG AŞ"},
		}
		return db.Create(tennants).Error
	}
	return nil
}
