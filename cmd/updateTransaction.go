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

// updateTransactionCmd represents the updateTransaction command
var updateTransactionDate string
var updateTransactionAmmount int32
var updateTransactionType string
var updateTransactionDescription string
var updateTransactionCatagory string
var updateTransactionID string

var updateTransactionCmd = &cobra.Command{
	Use:   "updateTransaction",
	Short: "A brief description of your command",
	Long: `This command will allow you to update a transaction.
	It will require the date of the activity, theamount in cents and what it was spent on.
	Later this entry is modifiable and deletable.`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateTransactionDate == "" {
			fmt.Printf("\n -d or --date flag was not set. Please set a correct date\n")
			return
		}
		if updateTransactionAmmount == 0 {
			fmt.Printf("\n Either the -a or --ammount flag was not set, or 0 minutes was filled in. Either ensure the flag is set, or register a minimum of 1 minute\n")
			return
		}
		if updateTransactionType == "" {
			fmt.Printf("\n -t or --type flag was not set. Please set a corect type flag\n")
			return
		}
		if updateTransactionDescription == "" {
			fmt.Printf("\n-e or --description flag is not set. Ensure a description is given to the transaction\n")
			return
		}
		if updateTransactionCatagory == "" {
			fmt.Printf("\n-c or --category flag not set. Ensure category is set for the transaction\n")
			return
		}
		if updateTransactionID == "" {
			fmt.Printf("\n-i or --id flag not set. Please supply a ID to update\n")
			return
		}

		layout := "02-01-2006"

		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err: \n%s\n", err)
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database\n%s\n", err)
			return
		}
		if permissions.Administrator != true {
			fmt.Println("Current user is not an administrator")
			return
		}

		ID, err := uuid.Parse(updateTransactionID)
		if err != nil {
			fmt.Printf("\nerror during parsing of the ID: %s \n", err)
			return
		}
		transaction := database.UpdateTransactionParams{
			DateTransaction: utils.TimeParse(updateTransactionDate),
			AmmountCent:     updateTransactionAmmount,
			Type:            updateTransactionType,
			Description:     updateTransactionDescription,
			Catagory:        updateTransactionCatagory,
			ID:              ID,
		}

		transactions, err := queries.UpdateTransaction(context.Background(), transaction)
		if err != nil {
			fmt.Printf("\nerror during updating of the entry: %s \n", err)
			return
		}
		fmt.Println("Transaction updated to: ")
		fmt.Printf("\nEntry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Cent): %d \n",
			transactions.ID, transactions.DateTransaction.Format(layout), transactions.Catagory, transactions.Description, transactions.AmmountCent)
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(updateTransactionCmd)

	updateTransactionCmd.Flags().StringVarP(&updateTransactionDate, "date", "d", "", "Flag denote the date of the transaction. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")

	updateTransactionCmd.Flags().Int32VarP(&updateTransactionAmmount, "amount", "a", 0, "Flag denote the amount spent or gaained in the transaction. Input the ammount in cents e.g. 1 euro is 100")

	updateTransactionCmd.Flags().StringVarP(&updateTransactionType, "type", "t", "", "Flag denote the type of the transaction. Use either spent or gained")

	updateTransactionCmd.Flags().StringVarP(&updateTransactionDescription, "description", "e", "", "Flag denote the description of the transaction.")

	updateTransactionCmd.Flags().StringVarP(&updateTransactionCatagory, "catagory", "c", "", "Flag denote the catagory of the transaction. Use a project for the name.")

	updateTransactionCmd.Flags().StringVarP(&updateTransactionID, "id", "i", "", "Flag denote the ID of the transaction.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
