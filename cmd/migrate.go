/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

// MigrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Updates the database to the latest version.",
	Long: `This command uses a tool called goose to update your database.
	The database is updates based on what version the DB is on and what the latest version is.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := goose.SetDialect("postgres")
		if err != nil {
			fmt.Printf("\nError seeting goose dialect. Err:\n%s\n", err)
			return
		}
		fmt.Println("Running Database Migration")
		err = goose.Up(utils.Database(), "sql/schema")
		if err != nil {
			fmt.Printf("\nError during database migration. Err:\n%s\n", err)
			return
		}
		fmt.Println("migration completed.")
	},
}

func init() {
	rootCmd.AddCommand(MigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
