package cmd

import (
	"github.com/spf13/cobra"
)

const API_KEY = "6d2b0b8e3f7db82dc942cf8ed6b6e7339318f3d7d47a4701846e6892e7d1c697"
const REGISTRY = "registry.digitalocean.com/fabric-registry/"

var (
	rootCmd = &cobra.Command{
		Use:   "runtime",
		Short: "Example short description",
		Long:  ``,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//client := godo.NewFromToken(API_KEY)

}
