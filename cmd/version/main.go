package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var StartCmd *cobra.Command

func init() {
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "printf version info",
		Example: "bkAdmin version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
}

func run() error {
	//todo
	//fmt.Println(global.Version)
	fmt.Println("v1.0.1")
	return nil
}
