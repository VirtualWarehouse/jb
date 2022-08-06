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
	text := strings.Join(args, " ")

	resp, err := c.PostChatCommand(config.Cfg.TouchChannel, text, "shrug")
	if err != nil {
		return err
	}
	if !resp.OK {
		return errors.New(resp.Error)
	}

	fmt.Printf("Successfully touch: \"%s %s\"\n", text, `¯\_(ツ)_/¯`)

	return nil
}
