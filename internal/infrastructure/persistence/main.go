package persistence

import (
	"fmt"

	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/rum-people-preseed/weather-harvester-svc/internal/config"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/entity"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/domain/repository"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/logger"
	"github.com/rum-people-preseed/weather-harvester-svc/internal/infrastructure/persistence/gorm_repo"
)

// DatabaseConnection is the database type
type DatabaseConnection struct {
	cfg  *config.DbConfig
	conn *gorm.DB
}

// Init initializes a database connection based on provided config
func (dc *DatabaseConnection) Init() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dc.cfg.Host, dc.cfg.User, dc.cfg.Password, dc.cfg.DBName, dc.cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   false,
		},
		FullSaveAssociations: false,
		Logger:               logger.New(zap.L()),
	})
	if err != nil {
		return err
	}
	dc.conn = db
	return nil
}

func (dc *DatabaseConnection) AutoMigrate() error {
	var entities []interface{}
	entities = append(entities, &entity.WeatherHistory{})
	return dc.conn.Migrator().AutoMigrate(entities...)
}

// ProvideDatabase provides a database type instance
func ProvideDatabase() interface{} {
	return func(config *config.DbConfig) (DBSetupOut, error) {
		db := &DatabaseConnection{cfg: config}
		err := db.Init()
		if err != nil {
			return DBSetupOut{}, err
		}

		if config.AutoMigrate == "true" {
			zap.S().Info("DB Auto-migrate enabled")
			if err := db.AutoMigrate(); err != nil {
				return DBSetupOut{}, err
			}
		}

		po := DBSetupOut{}
		po.setupRepository(db.conn)
		return po, nil
	}
}

type DBSetupOut struct {
	dig.Out

	// atomic repository to run transactional queries
	repository.AtomicRepository

	repository.WeatherHistoryRepository
}

func (po *DBSetupOut) setupRepository(db *gorm.DB) {
	po.AtomicRepository = gorm_repo.NewAtomicRepository(db)
	po.WeatherHistoryRepository = gorm_repo.NewWeatherHistoryRepository(db)
}
