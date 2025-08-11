/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/spf13/cobra"
)

// registerTimeCmd represents the registerTime command
var registerTimeDate string
var registerTimeMinutes int32
var registerTimeCategory string
var registerTimeDescription string

var registerTimeCmd = &cobra.Command{
	Use:   "registerTime",
	Short: "The command to register time spent",
	Long: `This command will allow you to register time spent.
	It will require the date of the activity, the time spent in minutes and what it was spent on.
	Later this entry is modifiable and deletable.`,
	Run: func(cmd *cobra.Command, args []string) {

		time := database.AddTimeRegistrationParams{
			DateActivity:  utils.TimeParse(registerTimeDate),
			LengthMinutes: registerTimeMinutes,
			Description:   registerTimeDescription,
			Catagory:      registerTimeCategory,
		}
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Println("Error retrieving current user from session")
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Println("Error during retrieval of user permissions from database")
			return
		}
		if permissions.AccessTimeregistration != true {
			fmt.Println("Current user is not allowed in the time registartion database")
			return
		}
		entry, err := queries.AddTimeRegistration(context.Background(), time)
		if err != nil {
			fmt.Printf("error during inserting data into the database: %s \n", err)
			return
		}
		layout := "02-01-2006"
		fmt.Printf("Databse entry created!\n")
		fmt.Printf("Entry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %d \n",
			entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.LengthMinutes)

	},
}

func init() {
	rootCmd.AddCommand(registerTimeCmd)

	registerTimeCmd.Flags().StringVarP(&registerTimeDate, "date", "d", "", "Flag to specify the date worked on a project. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	err := registerTimeCmd.MarkFlagRequired("date")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTimeCmd.Flags().Int32VarP(&registerTimeMinutes, "time", "t", 0, "Flag to specify the amount of time worked on a project in minutes.")
	err = registerTimeCmd.MarkFlagRequired("time")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTimeCmd.Flags().StringVarP(&registerTimeCategory, "category", "c", "", "Flag to specify the category/project name of the project.")
	err = registerTimeCmd.MarkFlagRequired("category")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTimeCmd.Flags().StringVarP(&registerTimeDescription, "description", "e", "", "Flag to specify the description of the work performed.")
	err = registerTimeCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerTimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
