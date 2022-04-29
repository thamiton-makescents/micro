package cmd

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime/pkg/appinit"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.ExactArgs(1),
		Short: "create a new runtime or application project",
		Long: `
Description:
  EXAMPLE
`,
		Example: `  runtime create [name]`,
		Run: func(cmd *cobra.Command, args []string) {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("could not detect current directory (this is a very unusual error)")
			}

			dir, err := filepath.Abs(pwd)
			if err != nil {
				fmt.Printf("could not determine absolute path of current directory")
			}

			name := args[0]

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			fs := afero.NewOsFs()
			init := appinit.NewAppInitializer(dir, fs, cmd.OutOrStdout())
			_, err = init.Run(name, "python")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
}
