package cli

import (
	"fmt"
	"os"

	"cosmossdk.io/core/address"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func NewSetMetadataCmd(ac address.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-metadata [metadata-file | metadata-json]",
		Short: "Set metadata for a bank denomination",
		Long: `Set metadata for a bank denomination. Metadata can be provided either as a JSON file or as a JSON string.
Example:
  # From a file
  $ simd tx bank set-metadata metadata.json
metadata.json:
{
  "description": "The native staking token of an arbitrary cosmos sdk chain.",
  "denom_units": [{ "denom": "stake", "exponent": 0, "aliases": ["stake"] }],
  "base": "stake",
  "display": "stake",
  "name": "stake",
  "symbol": "stake"
}
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var metadataBytes []byte
			input := args[0]

			if _, err := os.Stat(input); err == nil {
				metadataBytes, err = os.ReadFile(input)
				if err != nil {
					return fmt.Errorf("failed to read metadata file: %w", err)
				}
			} else {
				metadataBytes = []byte(input)
			}

			var metadata types.Metadata
			if err := clientCtx.Codec.UnmarshalJSON(metadataBytes, &metadata); err != nil {
				return fmt.Errorf("failed to parse metadata JSON: %w", err)
			}

			msg := types.NewMsgSetMetadata(clientCtx.GetFromAddress(), metadata)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
