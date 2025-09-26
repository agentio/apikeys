package cmd

import (
	"fmt"

	"github.com/agentio/apikeys/genproto/longrunningpb"
	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func getOperationCmd() *cobra.Command {
	var address string
	cmd := &cobra.Command{
		Use:   "get-operation OPERATION",
		Short: "Get operation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sidecar.NewClient(address)
			response, err := sidecar.CallUnary[longrunningpb.GetOperationRequest, longrunningpb.Operation](
				client,
				"/google.longrunning.Operations/GetOperation",
				sidecar.NewRequest(&longrunningpb.GetOperationRequest{
					Name: args[0],
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
