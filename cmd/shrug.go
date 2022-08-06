package cmd

import (
	"fmt"
	"strings"

	"github.com/VirtualWarehouse/jb/internal/client"
	"github.com/VirtualWarehouse/jb/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// touchCmd represents the touch command
var shrugCmd = &cobra.Command{
	Use:   "shrug",
	Short: "shrug",
	Long:  `Arguments are sent as a message`,
	RunE:  shrugRun,
}

func init() {
	rootCmd.AddCommand(shrugCmd)
}

func shrugRun(_ *cobra.Command, args []string) error {
	c := client.New(config.Cfg)

	resp, err := c.PostChatCommand(config.Cfg.TouchChannel, strings.Join(args, " "), "shrug")
	if err != nil {
		return err
	}
	if !resp.OK {
		return errors.New(resp.Error)
	}

	fmt.Println("Successfully touch")

	return nil
}
