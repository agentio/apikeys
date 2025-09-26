package cmd

import (
	"fmt"
	"os"

	"cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func updateKeyCmd() *cobra.Command {
	var address string
	cmd := &cobra.Command{
		Use:   "update-key FILE",
		Short: "Update key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sidecar.NewClient(address)
			b, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			var key apikeyspb.Key
			err = protojson.Unmarshal(b, &key)
			if err != nil {
				return err
			}
			response, err := sidecar.CallUnary[apikeyspb.UpdateKeyRequest, longrunningpb.Operation](
				client,
				"/google.api.apikeys.v2.ApiKeys/UpdateKey",
				sidecar.NewRequest(&apikeyspb.UpdateKeyRequest{
					Key: &key,
				}))
			if err != nil {
				return err
			}
			b2, err := protojson.MarshalOptions{Indent: "  "}.Marshal(response.Msg)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(b2))
			return nil
		},
	}
	cmd.Flags().StringVarP(&address, "address", "a", "localhost:4444", "service address")
	return cmd
}
