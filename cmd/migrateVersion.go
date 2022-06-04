package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"migrationCobra/database"
	"strconv"
)

var migrateVersioncmd = &cobra.Command{
	Use:   "version",
	Short: "migrate version command",
	Long:  "shows the database version according to the schema migration table",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Showing current database version")
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

		version, dirtyCheck, err := m.Version()

		if err != nil {
			fmt.Printf("migrate version error: %v \n", err)
		}

		if dirtyCheck != false {
			fmt.Println("database is dirty")
		}

		fmt.Println("database version : " + strconv.FormatUint(uint64(version), 10))

		fmt.Println("Migrate version done!")
	},
}

func init() {
	migrateCmd.AddCommand(migrateVersioncmd)
}
