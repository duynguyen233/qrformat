/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/duynguyen233/qrformat/cmd/format"
	"github.com/spf13/cobra"
)

var FormatData string

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format the data in the QR code",
	Long:  `Format the QR code into specific format.`,
	Run: func(cmd *cobra.Command, args []string) {
		if FormatData != "" {
			response, err := format.FormatQR(FormatData)
			if err != nil {
				fmt.Println("Invalid Format of QR: ", err)
				return
			}
			fmt.Print(response)
		} else {
			fmt.Println("⚠️ Warning: You are building without specifying the -d (--data) flag.")
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.Flags().StringVarP(&FormatData, "data", "d", "", "Data to format (required)")
}
