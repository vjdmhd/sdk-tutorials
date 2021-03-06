package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/github-username/scavenge/x/scavenge/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	scavengeTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	scavengeTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding # 1
		GetCmdCommitSolution(cdc),
		GetCmdRevealSolution(cdc),
		GetCmdSetCommit(cdc),
		GetCmdDeleteCommit(cdc),
		GetCmdCreateScavenge(cdc),
		GetCmdSetScavenge(cdc),
		GetCmdDeleteScavenge(cdc),
	)...)

	return scavengeTxCmd
}
