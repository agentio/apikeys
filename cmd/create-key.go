package cmd

import (
	"errors"
	"fmt"

	"github.com/agentio/apikeys/genproto/apikeys/apiv2/apikeyspb"
	"github.com/agentio/apikeys/genproto/longrunningpb"
	"github.com/agentio/sidecar"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func createKeyCmd() *cobra.Command {
	var address string
	var parent string
	var service string
	var keyid string
	var displayName string
	cmd := &cobra.Command{
		Use:   "create-key",
		Short: "Create key",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sidecar.NewClient(address)
			if parent == "" {
				return errors.New("--parent must be specified")
			}
			if service == "" {
				return errors.New("--service must be specified")
			}
			response, err := sidecar.CallUnary[apikeyspb.CreateKeyRequest, longrunningpb.Operation](
				client,
				"/google.api.apikeys.v2.ApiKeys/CreateKey",
				sidecar.NewRequest(&apikeyspb.CreateKeyRequest{
					Parent: parent,
					Key: &apikeyspb.Key{
						DisplayName: displayName,
						Restrictions: &apikeyspb.Restrictions{
							ApiTargets: []*apikeyspb.ApiTarget{
								{
									Service: service,
								},
							},
						},
					},
					KeyId: keyid,
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
	cmd.Flags().StringVar(&parent, "parent", "", "parent (projects/PROJECTNAME)")
	cmd.Flags().StringVar(&service, "service", "", "service to be used with this key")
	cmd.Flags().StringVar(&keyid, "keyid", "", "key id")
	cmd.Flags().StringVar(&displayName, "display-name", "", "display name")
	cmd.Flags().StringVarP(&address, "address", "a", "localhost:4444", "service address")
	return cmd
}
