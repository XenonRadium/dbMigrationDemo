package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"migrationCobra/database"
)

var migrateForcecmd = &cobra.Command{
	Use:   "force",
	Short: "migrate force command",
	Long:  "migrate force forces the database to a defined previous version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate force command")
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
			fmt.Printf("creating migration error: %v \n", err)
		}

		version, err := cmd.Flags().GetInt("version")
		if err != nil {
			fmt.Println("Please specify a valid version")
		}

		if err := m.Force(version); err != nil {
			fmt.Printf("migrate down error: %v \n", err)
		}

		fmt.Println("Migrate force done!")
	},
}

func init() {
	migrateForcecmd.Flags().IntP("version", "v", 0, "States ")
	migrateCmd.AddCommand(migrateForcecmd)
}
