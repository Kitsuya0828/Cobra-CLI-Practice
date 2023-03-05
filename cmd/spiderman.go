/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	"example.com/m/aa"
	"github.com/spf13/cobra"
)

// spidermanCmd represents the spiderman command
var spidermanCmd = &cobra.Command{
	Use:   "spiderman",
	Short: "Print the ascii art of Spiderman",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := aa.Aa.Open("spiderman.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	
		// ファイルを読み込んで出力
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)

		fmt.Print(buf.String())
	},
}

func init() {
	rootCmd.AddCommand(spidermanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spidermanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spidermanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
