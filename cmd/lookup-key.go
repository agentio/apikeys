package cmd

import (
	"fmt"

	"cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func lookupKeyCmd() *cobra.Command {
	var address string
	cmd := &cobra.Command{
		Use:   "lookup-key",
		Short: "Lookup key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sidecar.NewClient(address)
			response, err := sidecar.CallUnary[apikeyspb.LookupKeyRequest, apikeyspb.LookupKeyResponse](
				client,
				"/google.api.apikeys.v2.ApiKeys/LookupKey",
				sidecar.NewRequest(&apikeyspb.LookupKeyRequest{
					KeyString: args[0],
				}))
			if err != nil {
				return err
			}
			b, err := protojson.MarshalOptions{Indent: "  "}.Marshal(response.Msg)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(b))
			return nil
		},
	}
	cmd.Flags().StringVarP(&address, "address", "a", "localhost:4444", "service address")
	return cmd
}
