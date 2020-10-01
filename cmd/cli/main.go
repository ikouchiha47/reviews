package main

import (
	"fmt"
	"math/rand"
	"os"
	"reviews/internal/pkg/config"
	"time"

	"reviews/pkg/db"

	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := NewRootCommand().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "reviewsapp",
	}

	cmd.AddCommand(NewDBMigrate())
	cmd.AddCommand(NewDBRollback())

	return cmd
}

func NewDBMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dbmigrate",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return RunMigration(func(m db.Migrator) error { return m.Up() })
		},
	}
	return cmd
}

func NewDBRollback() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dbrollback",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return RunMigration(func(m db.Migrator) error { return m.Rollback() })
		},
	}
	return cmd
}

func RunMigration(f func(m db.Migrator) error) error {
	pgDB, err := db.InitDB(config.NewDBConfig())
	if err != nil {
		return err
	}
	defer pgDB.Close()

	m, err := db.NewPostgresMigrator(config.MigrationPath(), pgDB.DB)
	if err != nil {
		return err
	}

	return f(m)
}
