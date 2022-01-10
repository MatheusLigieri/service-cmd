package cmd

import (
	"fmt"

	"github.com/matheusLigieri/service-cmd/app"
	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use: "example",
	RunE: func(cmd *cobra.Command, args []string) error {

		config, err := app.InitConfig()
		if err != nil {
			return err
		}

		fmt.Println(config.Version)
		fmt.Println("Exemplo de comando, aqui podemos simular um 'adapter' que cuida de toda a regra recebendo diversas funções de primeiro nivel")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(exampleCmd)
}
