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

// registerTransactionCmd represents the registerTransaction command
var registerTransactionDate string
var registerTransactionAmmount int32
var registerTransactionType string
var registerTransactionDescription string
var registerTransactionCatagory string

var registerTransactionCmd = &cobra.Command{
	Use:   "registerTransaction",
	Short: "The command to register a traansaction ",
	Long: `This command will allow you to register transactions.
	It will require the date of the activity, theamount in cents and what it was spent on.
	Later this entry is modifiable and deletable.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := utils.DatabaseConnection()
		date := utils.TimeParse(registerTransactionDate)

		transaction := database.AddTransactionParams{
			DateTransaction: date,
			AmmountCent:     registerTransactionAmmount,
			Type:            registerTransactionType,
			Description:     registerTransactionDescription,
			Catagory:        registerTransactionCatagory,
		}

		transactions, err := query.AddTransaction(context.Background(), transaction)
		if err != nil {
			fmt.Printf("error during inserting data into the database: %s \n", err)
			return
		}

		fmt.Println("Transaction added!")
		fmt.Println(transactions)
	},
}

func init() {
	rootCmd.AddCommand(registerTransactionCmd)

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDate, "date", "d", "", "Flag denote the date of the transaction. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	err := registerTransactionCmd.MarkFlagRequired("date")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTransactionCmd.Flags().Int32VarP(&registerTransactionAmmount, "amount", "a", 0, "Flag denote the amount spent or gaained in the transaction. Input the ammount in cents e.g. 1 euro is 100")
	err = registerTransactionCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTransactionCmd.Flags().StringVarP(&registerTransactionType, "type", "t", "", "Flag denote the type of the transaction. Use either spent or gained")
	err = registerTransactionCmd.MarkFlagRequired("type")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDescription, "description", "e", "", "Flag denote the description of the transaction.")
	err = registerTransactionCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	registerTransactionCmd.Flags().StringVarP(&registerTransactionCatagory, "catagory", "c", "", "Flag denote the catagory of the transaction. Use a project for the name.")
	err = registerTransactionCmd.MarkFlagRequired("catagory")
	if err != nil {
		fmt.Printf("required flag not set")
		return
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
