package cmd

import (
	"fmt"

	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"

	"github.com/agentio/apikeys/genproto/apikeys/apiv2/apikeyspb"
	"google.golang.org/protobuf/encoding/protojson"
)

func listKeysCmd() *cobra.Command {
	var address string
	cmd := &cobra.Command{
		Use:   "list-keys PROJECT",
		Short: "List keys",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sidecar.NewClient(address)
			nextPageToken := ""
			for {
				response, err := sidecar.CallUnary[apikeyspb.ListKeysRequest, apikeyspb.ListKeysResponse](
					client,
					"/google.api.apikeys.v2.ApiKeys/ListKeys",
					sidecar.NewRequest(&apikeyspb.ListKeysRequest{
						Parent:    "projects/" + args[0] + "/locations/global",
						PageSize:  100,
						PageToken: nextPageToken,
					}))
				if err != nil {
					return err
				}
				b, err := protojson.MarshalOptions{Indent: "  "}.Marshal(response.Msg)
				if err != nil {
					return err
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(b))
				nextPageToken = response.Msg.NextPageToken
				if nextPageToken == "" {
					break
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&address, "address", "a", "localhost:4444", "service address")
	return cmd
}
