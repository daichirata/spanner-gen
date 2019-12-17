package cmd

import (
	"fmt"

	"github.com/daichirata/spanner-gen/internal"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "spanner-gen SOURCE",
		Short: "",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must specify 2 arguments")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				out          string
				single       bool
				template     string
				suffix       string
				ignoreFields []string
				ignoreTables []string
			)
			cmd.Flags().StringVarP(&out, "out", "o", "", "output path or file name")
			cmd.Flags().StringVar(&suffix, "suffix", ".gen.go", "output file suffix")
			cmd.Flags().BoolVar(&single, "single", false, "toggle single file output")
			cmd.Flags().StringArrayVar(&ignoreFields, "ignore-fields", nil, "fields to exclude from the generated Go code types")
			cmd.Flags().StringArrayVar(&ignoreTables, "ignore-tables", nil, "tables to exclude from the generated Go code types")
			cmd.Flags().StringVar(&template, "template", "", "user supplied template path")

			source, err := internal.NewSource(args[0])
			if err != nil {
				return err
			}

			g := internal.NewGenerator(source, internal.GeneratorOption{
				Out:          out,
				Suffix:       suffix,
				Single:       single,
				Template:     template,
				IgnoreFields: ignoreFields,
				IgnoreTables: ignoreTables,
			})
			if err := g.Generate(); err != nil {
				return fmt.Errorf("error: %v", err)
			}
			return nil
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
