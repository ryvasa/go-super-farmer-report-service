package database

import (
	"fmt"

	"github.com/ryvasa/go-super-farmer-report-service/internal/model/domain"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(env *env.Env) (*gorm.DB, error) {
	host := env.Database.Host
	port := env.Database.Port
	name := env.Database.Name
	user := env.Database.User
	password := env.Database.Password
	timezone := env.Database.Timezone

	if host == "" || port == "" || name == "" || user == "" || password == "" {
		return nil, fmt.Errorf("missing database environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, name, port, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Land{},
		&domain.Commodity{},
		&domain.LandCommodity{},
		&domain.Province{},
		&domain.City{},
		&domain.Price{},
		&domain.PriceHistory{},
		&domain.Supply{},
		&domain.SupplyHistory{},
		&domain.Demand{},
		&domain.DemandHistory{},
		&domain.Harvest{},
		&domain.Sale{},
	)

	// seeders.Seeders(db)

	return db, nil
}
