// Package mylsn implements a bpf lsm.
package main

import (
	"os"

	"aduu.dev/utils/aduuapp"
	"github.com/spf13/cobra"
)

// RootCMD creates a root command of a program.
func RootCMD() *cobra.Command {
	viperSetup, v := aduuapp.SetupViper("mylsm", &aduuapp.SetupViperConfig{})
	_ = v

	cmd := &cobra.Command{
		Use:          "mylsm",
		Short:        "runs its own lsm",
		SilenceUsage: true,
	}

	viperSetup.SetupFlags(cmd, &aduuapp.SetupFlagsConfig{})

	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		if viperSetup.IsWriteConfigSet() {
			return viperSetup.WriteConfig()
		}

		return run()
	}
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		return nil
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.AddCommand()
	return cmd
}

func main() {
	errorExitCode := 1

	if err := RootCMD().Execute(); err != nil {
		// if !helper.ContainsDoNotPrintHelp(err) {
		//    fmt.Println(err)
		// }

		os.Exit(errorExitCode)
	}
}
