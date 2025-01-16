package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/spf13/cobra"
)

func NewMsgSoftwareUpgrade(from sdk.AccAddress, pln types.Plan) (*types.MsgSoftwareUpgrade, error) {
	m := &types.MsgSoftwareUpgrade{
		Authority: from.String(),
		Plan:      pln,
	}
	return m, nil
}
func NewCmdManualUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manual-software-upgrade name height info",
		Args:  cobra.ExactArgs(3),
		Short: "Create a software upgrade",
		Long: "Create a software upgrade.\n" +
			"Please specify a unique name and height for the upgrade to take effect.\n" +
			"You may include info to reference a binary download link, in a format compatible with: https://github.com/cosmos/cosmos-sdk/tree/main/cosmovisor",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			height := args[1]
			info := args[2]
			heightInt, err := strconv.ParseInt(height, 10, 64)
			if err != nil {
				return err
			}
			pln := types.Plan{
				Name:   name,
				Height: heightInt,
				Info:   info,
			}
			from := clientCtx.GetFromAddress()
			msg, err := NewMsgSoftwareUpgrade(from, pln)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
func NewMsgCancelUpgrade(from sdk.AccAddress) (*types.MsgCancelUpgrade, error) {
	m := &types.MsgCancelUpgrade{
		Authority: from.String(),
	}
	return m, nil
}
func NewCmdManualCancelUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "manual-cancel-upgrade",
		Args:  cobra.ExactArgs(0),
		Short: "Cancel the current software upgrade",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			msg, err := NewMsgCancelUpgrade(from)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
