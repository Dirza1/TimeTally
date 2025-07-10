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
		layout := "02-01-2006"
		ID, err := uuid.Parse(updateTimeID)
		if err != nil {
			fmt.Printf("error during parsing of the ID: %s \n", err)
			return
		}
		time := database.UpdateTimeParams{
			DateActivity:  utils.TimeParse(updateTimeCmdDate),
			LengthMinutes: updateTimeCmdMinutes,
			Description:   updateTimeCmdDescription,
			Catagory:      updateTimeCmdCategory,
			ID:            ID,
		}
		query := utils.DatabaseConnection()
		entry, err := query.UpdateTime(context.Background(), time)
		if err != nil {
			fmt.Printf("error during updating of the entry: %s \n", err)
			return
		}
		fmt.Println("Time updated to: ")
		fmt.Printf("Entry ID: %s. Activity date: %s. Category: %s, Description: %s, Time spent(Hours): %d \n",
			entry.ID, entry.DateActivity.Format(layout), entry.Catagory, entry.Description, entry.LengthMinutes)
	},
}

func init() {
	rootCmd.AddCommand(updateTimeCmd)

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdDate, "date", "d", "", "Flag to specify the date worked on a project. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	updateTimeCmd.MarkFlagRequired("date")

	updateTimeCmd.Flags().Int32VarP(&updateTimeCmdMinutes, "time", "t", 0, "Flag to specify the amount of time worked on a project in minutes.")
	updateTimeCmd.MarkFlagRequired("time")

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdCategory, "category", "c", "", "Flag to specify the category/project name of the project.")
	updateTimeCmd.MarkFlagRequired("category")

	updateTimeCmd.Flags().StringVarP(&updateTimeCmdDescription, "description", "e", "", "Flag to specify the description of the work performed.")
	updateTimeCmd.MarkFlagRequired("description")

	updateTimeCmd.Flags().StringVarP(&updateTimeID, "id", "i", "", "Flag to specify the ID of the work performed.")
	updateTimeCmd.MarkFlagRequired("id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTimeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
