package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"migrationCobra/database"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate down command",
	Long:  "migrate up handles migrations from V2 to V1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate down command")
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
			fmt.Printf("cerating migration error: %v \n", err)
		}

		if err := m.Down(); err != nil {
			fmt.Printf("migrate down error: %v \n", err)
		}

		fmt.Println("Migrate down done!")
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
