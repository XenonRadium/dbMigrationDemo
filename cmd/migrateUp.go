package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"migrationCobra/database"
)

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up command installs v1",
	Long:  "migrate up handles migrations for v1 of database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate up command")
		db := database.Open()
		dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Printf("error creating instance: %v \n", err)
		}

		fileSource, err := (&file.File{}).Open("file://migrate")
		if err != nil {
			fmt.Printf("error creating instance: %v \n", err)
		}

		m, err := migrate.NewWithInstance("file", fileSource, "mysql", dbDriver)
		if err != nil {
			fmt.Printf("crating migration errpr: %v \n", err)
		}

		if err := m.Up(); err != nil {
			fmt.Printf("migrate up errpr: %v \n", err)
		}

		fmt.Println("Migrate up done!")
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
