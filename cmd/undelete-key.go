package cmd

import (
	"fmt"

	"github.com/agentio/apikeys/genproto/apikeys/apiv2/apikeyspb"
	"github.com/agentio/apikeys/genproto/longrunningpb"
	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func undeleteKeyCmd() *cobra.Command {
	var address string
	cmd := &cobra.Command{
		Use:   "undelete-key",
		Short: "Undelete key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			operation, err := sidecar.CallUnary[apikeyspb.UndeleteKeyRequest, longrunningpb.Operation](
				sidecar.NewClient(address),
				"/google.api.apikeys.v2.ApiKeys/UndeleteKey",
				sidecar.NewRequest(&apikeyspb.UndeleteKeyRequest{
					Name: args[0],
				}))
			if err != nil {
				return err
			}
			b, err := protojson.MarshalOptions{Indent: "  "}.Marshal(operation.Msg)
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
