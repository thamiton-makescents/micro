package cmd

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a docker image",
		Long: `
Description:
  EXAMPLE
`,
		Example: `  runtime delete`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)

			defer cancel()

			client := godo.NewFromToken(API_KEY)
			appsListRequest := &godo.ListOptions{
				Page:    0,
				PerPage: 0,
			}
			list, _, err := client.Apps.List(context.Background(), appsListRequest)
			if err != nil {
				return
			}
			for i := 0; i < len(list); i++ {
				if strings.Contains(list[i].Spec.Name, "test") {
					response, err := client.Apps.Delete(ctx, list[i].ID)
					if err != nil {
						return
					}
					fmt.Println(response)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
