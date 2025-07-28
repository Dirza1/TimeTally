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
		layout := "02-01-2006"
		query := utils.DatabaseConnection()

		ID, err := uuid.Parse(updateTransactionID)
		if err != nil {
			fmt.Printf("error during parsing of the ID: %s \n", err)
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

		transactions, err := query.UpdateTransaction(context.Background(), transaction)
		if err != nil {
			fmt.Printf("error during updating of the entry: %s \n", err)
			return
		}
		fmt.Println("Transaction updated to: ")
		fmt.Printf("Entry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Cent): %d \n",
			transactions.ID, transactions.DateTransaction.Format(layout), transactions.Catagory, transactions.Description, transactions.AmmountCent)
	},
}

func init() {
	rootCmd.AddCommand(updateTransactionCmd)

	updateTransactionCmd.Flags().StringVarP(&updateTransactionDate, "date", "d", "", "Flag denote the date of the transaction. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	err := updateTransactionCmd.MarkFlagRequired("date")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	updateTransactionCmd.Flags().Int32VarP(&updateTransactionAmmount, "amount", "a", 0, "Flag denote the amount spent or gaained in the transaction. Input the ammount in cents e.g. 1 euro is 100")
	err = updateTransactionCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	updateTransactionCmd.Flags().StringVarP(&updateTransactionType, "type", "t", "", "Flag denote the type of the transaction. Use either spent or gained")
	err = updateTransactionCmd.MarkFlagRequired("type")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	updateTransactionCmd.Flags().StringVarP(&updateTransactionDescription, "description", "e", "", "Flag denote the description of the transaction.")
	err = updateTransactionCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	updateTransactionCmd.Flags().StringVarP(&updateTransactionCatagory, "catagory", "c", "", "Flag denote the catagory of the transaction. Use a project for the name.")
	err = updateTransactionCmd.MarkFlagRequired("catagory")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}
	updateTransactionCmd.Flags().StringVarP(&updateTransactionID, "id", "i", "", "Flag denote the ID of the transaction.")
	err = updateTransactionCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
