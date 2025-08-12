/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Dirza1/Time-and-expence-registration/internal/database"
	"github.com/Dirza1/Time-and-expence-registration/internal/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// updateTimeCmd represents the updateTime command
var updateTimeCmdDate string
var updateTimeCmdMinutes int32
var updateTimeCmdCategory string
var updateTimeCmdDescription string
var updateTimeID string

var updateTimeCmd = &cobra.Command{
	Use:   "updateTime",
	Short: "The command to update time spent",
	Long: `This command will allow you to update time spent.
	It will require the date of the activity, the time spent in minutes and what it was spent on.
	Later this entry is modifiable and deletable.`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateTimeCmdDate == "" {
			fmt.Printf("\n -d or --date flag was not set. Please set a correct date\n")
			return
		}
		if updateTimeCmdMinutes == 0 {
			fmt.Printf("\n Either the -t or --time flag was not set, or 0 minutes was filled in. Either ensure the flag is set, or register a minimum of 1 cent\n")
			return
		}
		if updateTimeCmdDescription == "" {
			fmt.Printf("\n-e or --description flag is not set. Ensure a description is given to the time registration\n")
			return
		}
		if updateTimeCmdCategory == "" {
			fmt.Printf("\n-c or --category flag not set. Ensure category is set for the time registration\n")
			return
		}
		if updateTimeID == "" {
			fmt.Printf("\n-i or --id flag not set. Please supply a ID to update\n")
			return
		}

		layout := "02-01-2006"
		ID, err := uuid.Parse(updateTimeID)
		if err != nil {
			fmt.Printf("\nerror during parsing of the ID: %s \n", err)
			return
		}
		time := database.UpdateTimeParams{
			DateActivity:  utils.TimeParse(updateTimeCmdDate),
			LengthMinutes: updateTimeCmdMinutes,
			Description:   updateTimeCmdDescription,
			Catagory:      updateTimeCmdCategory,
			ID:            ID,
		}
		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err: \n%s\n", err)
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database. Err: \n%s\n", err)
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}
		entry, err := queries.UpdateTime(context.Background(), time)
		if err != nil {
			fmt.Printf("\nerror during updating of the entry: %s \n", err)
			return
		}
		fmt.Println("Time updated to: ")
		fmt.Printf("\nEntry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %d \n",
			entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.LengthMinutes)
	},
}

func init() {
	rootCmd.AddCommand(updateTimeCmd)

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdDate, "date", "d", "", "Flag to specify the date worked on a project. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")

	updateTimeCmd.Flags().Int32VarP(&updateTimeCmdMinutes, "time", "t", 0, "Flag to specify the amount of time worked on a project in minutes.")

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdCategory, "category", "c", "", "Flag to specify the category/project name of the project.")

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdDescription, "description", "e", "", "Flag to specify the description of the work performed.")

	updateTimeCmd.Flags().StringVarP(&updateTimeID, "id", "i", "", "Flag to specify the ID of the work performed.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
