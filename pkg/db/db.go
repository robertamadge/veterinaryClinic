package db

import (
	"log"

	"github.com/robertamadge/veterinayClinic/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Owner{})
	db.Migrator().CreateConstraint(&models.Owner{}, "Pets")
	db.Migrator().CreateConstraint(&models.Owner{}, "fk_owners_pets")
	db.AutoMigrate(&models.Pet{})
	db.Migrator().CreateConstraint(&models.Pet{}, "Owners")
	db.Migrator().CreateConstraint(&models.Pet{}, "fk_owners_pets")
	db.AutoMigrate(&models.Appointment{})
	db.Migrator().CreateConstraint(&models.Appointment{}, "Pets")
	db.Migrator().CreateConstraint(&models.Appointment{}, "fk_appointments_pets")
	db.Migrator().CreateConstraint(&models.Appointment{}, "Doctors")
	db.Migrator().CreateConstraint(&models.Appointment{}, "fk_appointments_doctors")
	db.AutoMigrate(&models.Doctor{})

	return db
}
