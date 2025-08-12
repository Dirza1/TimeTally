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
		if registerTransactionDate == "" {
			fmt.Printf("\n -d or --date flag was not set. Please set a correct date\n")
			return
		}
		if registerTransactionAmmount == 0 {
			fmt.Printf("\n Either the -a or --ammount flag was not set, or 0 minutes was filled in. Either ensure the flag is set, or register a minimum of 1 minute\n")
			return
		}
		if registerTransactionType == "" {
			fmt.Printf("\n -t or --type flag was not set. Please set a corect type flag\n")
			return
		}
		if registerTransactionDescription == "" {
			fmt.Printf("\n-e or --description flag is not set. Ensure a description is given to the transaction\n")
			return
		}
		if registerTransactionCatagory == "" {
			fmt.Printf("\n-c or --category flag not set. Ensure category is set for the transaction\n")
			return
		}

		queries := utils.DatabaseConnection()
		currentUser, err := utils.LoadSession()
		if err != nil {
			fmt.Printf("\nError retrieving current user from session. Err: \n%s\n", err)
			return
		}
		permissions, err := queries.GetUserPermissions(context.Background(), currentUser.UserName)
		if err != nil {
			fmt.Printf("\nError during retrieval of user permissions from database. Err: \n%s\n", err)
			return
		}
		if permissions.AccessFinance != true {
			fmt.Println("Current user is not allowed in the financial database")
			return
		}
		date := utils.TimeParse(registerTransactionDate)

		transaction := database.AddTransactionParams{
			DateTransaction: date,
			AmmountCent:     registerTransactionAmmount,
			Type:            registerTransactionType,
			Description:     registerTransactionDescription,
			Catagory:        registerTransactionCatagory,
		}

		transactions, err := queries.AddTransaction(context.Background(), transaction)
		if err != nil {
			fmt.Printf("\nerror during inserting data into the database:\n%s \n", err)
			return
		}
		layout := "02-01-2006"
		fmt.Println("Transaction added!")
		fmt.Printf("\nEntry ID: %s. Transaction date: %s. Category: %s, Description: %s, Total ammount(Cent): %d, Transactio type: %s \n",
			transactions.ID, transactions.DateTransaction.Format(layout), transactions.Catagory, transactions.Description, transactions.AmmountCent, transactions.Type)
		utils.UpdateSession()
	},
}

func init() {
	rootCmd.AddCommand(registerTransactionCmd)

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDate, "date", "d", "", "Flag denote the date of the transaction. Use full date notateion. e.g. 22-11-2025 for 22 november 2025")

	registerTransactionCmd.Flags().Int32VarP(&registerTransactionAmmount, "amount", "a", 0, "Flag denote the amount spent or gaained in the transaction. Input the ammount in cents e.g. 1 euro is 100")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionType, "type", "t", "", "Flag denote the type of the transaction. Use either spent or gained")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionDescription, "description", "e", "", "Flag denote the description of the transaction.")

	registerTransactionCmd.Flags().StringVarP(&registerTransactionCatagory, "catagory", "c", "", "Flag denote the catagory of the transaction. Use a project for the name.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerTransactionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerTransactionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
