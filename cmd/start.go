/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Neaj-Morshad-101/HTTP-API-server/apis"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var (
	//Port is flag to store the default port for http server.
	Port     int
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start the server on a default port",
		Long: `start the server on a default port ,
				but port can be specify using the port flag`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add startserver func here")
			apis.StartServer(Port)
		},
	}
)

func init() {
	startCmd.PersistentFlags().IntVarP(&Port, "port", "p", 5050, "default port for http server")
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
