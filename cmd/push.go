package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"time"
)

var (
	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "push a docker image",
		Long: `
Description:
  EXAMPLE
`,
		Example: `  runtime push`,
		Run: func(cmd *cobra.Command, args []string) {
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = imagePush(cli)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

func imagePush(dockerClient *client.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	var authConfig = types.AuthConfig{
		Username:      API_KEY,
		Password:      API_KEY,
		ServerAddress: REGISTRY,
	}

	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	tag := REGISTRY + "test"
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	rd, err := dockerClient.ImagePush(ctx, tag, opts)
	if err != nil {
		return err
	}

	defer rd.Close()

	//err = print(rd)
	if err != nil {
		return err
	}

	return nil
}
