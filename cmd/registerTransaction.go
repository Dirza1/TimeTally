/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
			log.Fatal("error during registering of transaction: ", err)
		}

		fmt.Println("Transaction added!")
		fmt.Println(transactions)
	},
}

func init() {
	rootCmd.AddCommand(registerTransactionCmd)

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDate, "date", "d", "", "Flag denote the date of the transaction. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")
	registerTransactionCmd.MarkFlagRequired("date")

	registerTransactionCmd.Flags().Int32VarP(&registerTransactionAmmount, "amount", "a", 0, "Flag denote the amount spent or gaained in the transaction. Input the ammount in cents e.g. 1 euro is 100")
	registerTransactionCmd.MarkFlagRequired("amount")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionType, "type", "t", "", "Flag denote the type of the transaction. Use either spent or gained")
	registerTransactionCmd.MarkFlagRequired("type")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDescription, "description", "e", "", "Flag denote the description of the transaction.")
	registerTransactionCmd.MarkFlagRequired("description")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionCatagory, "catagory", "c", "", "Flag denote the catagory of the transaction. Use a project for the name.")
	registerTransactionCmd.MarkFlagRequired("catagory")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
