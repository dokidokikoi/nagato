package postgres

import (
	"fmt"
	"nagato/apiservice/internal/config"
	"nagato/apiservice/internal/model"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store struct {
	db *gorm.DB
}

func (d *Store) Blanks() *blanks {
	return newBlanks(d)
}

func (d *Store) BlankMatters() *blankMatters {
	return newBlankMatters(d)
}

func (d *Store) Bridges() *bridges {
	return newBridges(d)
}

func (d *Store) Installs() *installs {
	return newInstalls(d)
}

func (d *Store) Matters() *matters {
	return newMatters(d)
}

func (d *Store) Shares() *shares {
	return newShares(d)
}

func (d *Store) SmallFileCaches() *smallFileCaches {
	return newSmallFileCaches(d)
}

func (d *Store) Tags() *tags {
	return newTags(d)
}

func (d *Store) Users() *users {
	return newUsers(d)
}

func (d *Store) Transaction() *transaction {
	return newTransaction(d)
}

func GetPGFactory() (*Store, error) {
	db, err := gorm.Open(postgres.Open(config.Config().PGConfig.Dns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}
	// 设定数据库最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设定数据库最长连接时长
	sqlDB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	// 设定数据库最大空闲数
	sqlDB.SetMaxIdleConns(100)

	// 自动化迁移
	if err := migrateDatabase(db); err != nil {
		fmt.Println(err)
	}

	return &Store{db: db}, nil
}

// cleanDatabase tear downs the database tables.
func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&model.User{}); err != nil {
		return errors.Wrap(err, "drop user model failed")
	}
	if err := db.Migrator().DropTable(&model.Blank{}); err != nil {
		return errors.Wrap(err, "drop blank model failed")
	}
	if err := db.Migrator().DropTable(&model.Tag{}); err != nil {
		return errors.Wrap(err, "drop tag model failed")
	}
	if err := db.Migrator().DropTable(&model.Matter{}); err != nil {
		return errors.Wrap(err, "drop matter model failed")
	}
	if err := db.Migrator().DropTable(&model.Share{}); err != nil {
		return errors.Wrap(err, "drop share model failed")
	}
	if err := db.Migrator().DropTable(&model.Bridge{}); err != nil {
		return errors.Wrap(err, "drop bridge model failed")
	}
	if err := db.Migrator().DropTable(&model.Install{}); err != nil {
		return errors.Wrap(err, "drop install model failed")
	}
	// if err := db.Migrator().DropTable(&model.BlankMatter{}); err != nil {
	// 	return errors.Wrap(err, "drop blank_matter model failed")
	// }
	if err := db.Migrator().DropTable(&model.SmallFileCache{}); err != nil {
		return errors.Wrap(err, "drop small_file_cache model failed")
	}

	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}
	if err := db.AutoMigrate(&model.Blank{}); err != nil {
		return errors.Wrap(err, "migrate blank model failed")
	}
	if err := db.AutoMigrate(&model.Tag{}); err != nil {
		return errors.Wrap(err, "migrate tag model failed")
	}
	if err := db.AutoMigrate(&model.Matter{}); err != nil {
		return errors.Wrap(err, "migrate matter model failed")
	}
	if err := db.AutoMigrate(&model.Share{}); err != nil {
		return errors.Wrap(err, "migrate share model failed")
	}
	if err := db.AutoMigrate(&model.Bridge{}); err != nil {
		return errors.Wrap(err, "migrate bridge model failed")
	}
	if err := db.AutoMigrate(&model.Install{}); err != nil {
		return errors.Wrap(err, "migrate install model failed")
	}
	// if err := db.AutoMigrate(&model.BlankMatter{}); err != nil {
	// 	return errors.Wrap(err, "migrate blank_matter model failed")
	// }
	if err := db.AutoMigrate(&model.SmallFileCache{}); err != nil {
		return errors.Wrap(err, "migrate small_file_cache model failed")
	}

	return nil
}

// resetDatabase resets the database tables.
func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}
