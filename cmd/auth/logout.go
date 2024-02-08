package auth

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/noop"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logout",
		//TODO more fleshed-out descriptions
		Short: "Logout with Ping",
		Long:  "Logout with Ping",
		Run: func(cmd *cobra.Command, args []string) {
			l := logger.Get()
			// Just use the no-op connector for now by default
			authConnectors := []connector.Authenticatable{
				noop.Connector(),
			}
			err := authConnectors[0].Logout()
			if err != nil {
				l.Fatal().Err(err).Msg("Logout failed")
			}
		},
	}

	return cmd
}
