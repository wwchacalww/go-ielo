/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wwchacalww/go-psyc/adapters/cli"

	"github.com/spf13/cobra"
)

var action string
var userID string
var userName string
var userEmail string
var userPassword string
var userRole string
var userStatus bool

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&userRepository, action, userID, userName, userEmail, userPassword, userRole)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "list", "List users")
	cliCmd.Flags().StringVarP(&userID, "id", "i", "", "Get user ID")
	cliCmd.Flags().StringVarP(&userName, "username", "n", "", "Get user name")
	cliCmd.Flags().StringVarP(&userEmail, "email", "e", "", "Get user email")
	cliCmd.Flags().StringVarP(&userPassword, "password", "p", "", "Get user password")
	cliCmd.Flags().StringVarP(&userRole, "role", "r", "", "Get user role")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
