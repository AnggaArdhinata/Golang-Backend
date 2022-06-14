package gorm

import (
	"backend-golang/src/database/gorm/models"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	RunE: dbMigrate,
	
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		
		{
			ID: "20220614001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{})
			},
			
			
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},

		{
			ID: "20220614002",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Vehicle{})
			},
			
			
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("vehicles")
			},
		},

		{
			ID: "20220614003",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Order{})
			},
			
			
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("orders")
			},
		},
	})

	if migUp {
		if err = m.Migrate(); err != nil {
			return err
		}
		log.Fatal("Migration did run successfully")
		return nil
	}

	if migDown {
		if err = m.RollbackLast(); err != nil {
			return err
		}
		log.Fatal("Rollback did run successfully")
		return nil
	}

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&models.User{},
			&models.Order{},
			&models.Vehicle{},
		)

		if err != nil {
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		return err

	}
	log.Fatal("Init schema successfully")
	return nil
}
