package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"

	config "miata/init"
)

func main() {
	app := fx.New(
		config.Module,
		fx.Invoke(execute),
	)
	ctx := context.Background()
	app.Start(ctx)
	app.Stop(ctx)
}

func logError(log *zap.Logger, err error) {
	log.Sugar().Error(err.Error())
}

func execute(log *zap.Logger, db *gorm.DB) {
	instance, err := db.DB()
	if err != nil {
		log.Error("error", zap.String("error", err.Error()))
	}
	driver, err := mysql.WithInstance(instance, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	var rootCmd = &cobra.Command{
		Use: "migrate",
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "returns the currently active migration version.",
		Run: func(cmd *cobra.Command, args []string) {
			version, dirty, err := m.Version()
			if err != nil {
				logError(log, err)
				return
			}
			log.Info("version", zap.Uint("version", version), zap.Bool("dirty", dirty))
		},
	}
	var forceCmd = &cobra.Command{
		Use:   "force",
		Short: "sets a migration version.",
		Run: func(cmd *cobra.Command, args []string) {
			v, err := cmd.Flags().GetInt("version")
			if err != nil {
				logError(log, err)
				return
			}
			err = m.Force(v)
			if err != nil {
				logError(log, err)
				return
			}
		},
	}
	forceCmd.Flags().IntP("version", "v", -1, "version")
	forceCmd.MarkFlagRequired("version")
	var upCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate all the way up.",
		Run: func(cmd *cobra.Command, args []string) {
			err := m.Up()
			if err != nil {
				logError(log, err)
				return
			}
		},
	}
	var gotoCmd = &cobra.Command{
		Use:   "goto",
		Short: "migrates either up or down to the specified version.",
		Run: func(cmd *cobra.Command, args []string) {
			v, err := cmd.Flags().GetUint("version")
			if err != nil {
				logError(log, err)
				return
			}
			err = m.Migrate(v)
			if err != nil {
				logError(log, err)
				return
			}
		},
	}
	gotoCmd.Flags().UintP("version", "v", 1024, "version")
	gotoCmd.MarkFlagRequired("version")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(forceCmd)
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(gotoCmd)
	rootCmd.Execute()
}
