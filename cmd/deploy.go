package cmd

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
	"time"
)

var (
	deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deploy a docker image",
		Long: `
Description:
  EXAMPLE
`,
		Example: `  runtime deploy`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)

			defer cancel()

			client := godo.NewFromToken(API_KEY)

			spec := &godo.AppSpec{
				Name: "test",
				Services: []*godo.AppServiceSpec{
					{
						Name: "test",
						Image: &godo.ImageSourceSpec{
							RegistryType: "DOCR",
							Registry:     "fabric-registry",
							Repository:   "test",
							Tag:          "",
						},
					},
				},
				StaticSites: nil,
				Workers:     nil,
				Jobs:        nil,
				Functions:   nil,
				Databases:   nil,
				Domains:     nil,
				Region:      "",
				Envs:        nil,
				Alerts:      nil,
				Ingress:     nil,
			}
			appsCreateRequest := &godo.AppCreateRequest{
				Spec: spec,
			}
			createCreateResponse, r, err := client.Apps.Create(ctx, appsCreateRequest)
			if err != nil {
				return
			}
			fmt.Println(createCreateResponse)
			fmt.Println(r)
		},
	}
)

func init() {
	rootCmd.AddCommand(deployCmd)
}
