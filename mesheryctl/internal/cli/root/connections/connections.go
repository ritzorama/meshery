package connections

import (
	"fmt"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/config"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	availableSubcommands = []*cobra.Command{listConnectionsCmd, deleteConnectionCmd, viewConnectionCmd, createConnectionCmd}

	pageNumberFlag int
	outFormatFlag  string
	saveFlag       bool
)

var ConnectionsCmd = &cobra.Command{
	Use:   "connection",
	Short: "Manage connections",
	Long:  `Manage Meshery connections to various platforms and services`,
	Example: `
// List all the connections
mesheryctl connection list

// View a specific connection
mesheryctl connection view [connection-name]

// Delete a connection
mesheryctl connection delete [connection-id]

// Create a new Kubernetes connection
mesheryctl connection create --type gke
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if err := cmd.Usage(); err != nil {
				return nil
			}
			return errors.New("please provide a subcommand")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		if ok := utils.IsValidSubcommand(availableSubcommands, args[0]); !ok {
			return errors.New(utils.ConnectionSubError(fmt.Sprintf("'%s' is an invalid subcommand. Please provide required options from [list|view|delete|create]. Use 'mesheryctl connection --help' to display usage guide.\n", args[0]), "connection"))
		}
		_, err := config.GetMesheryCtl(viper.GetViper())
		if err != nil {
			return utils.ErrLoadConfig(err)
		}
		err = cmd.Usage()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	listConnectionsCmd.Flags().BoolP("count", "c", false, "Display the count of connections")
	listConnectionsCmd.Flags().IntVarP(&pageNumberFlag, "page", "p", 1, "Page number")
	deleteConnectionCmd.Flags().StringP("id", "i", "", "ID of the connection to be deleted")
	ConnectionsCmd.AddCommand(availableSubcommands...)
}
