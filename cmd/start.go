package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/toposender/internal/config"
	"github.com/stackvista/toposender/internal/domain"
	"github.com/stackvista/toposender/internal/stackstate"
)

type StartArgs struct {
	TopologyFile string
}

func StartCommand(cfg *config.Config) *cobra.Command {
	args := &StartArgs{}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start a new snapshot",
		Long:  "Start a new snapshot",
		RunE:  SendTopology(cfg, args),
	}

	cmd.Flags().StringVarP(&args.TopologyFile, "topology-file", "f", "", "Topology file to send")

	return cmd
}

func SendTopology(cfg *config.Config, startArgs *StartArgs) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		t, err := domain.ReadTopology(startArgs.TopologyFile)
		if err != nil {
			return err
		}

		tm := domain.ToTopologyMessage(cfg, t)

		if err := stackstate.NewStackStateClient(cfg.ApiKey, cfg.StackStateAddress, cfg.StackStatePort, cfg.StackStatePrefix).SendTopology(tm); err != nil {
			return err
		}

		return nil
	}
}
