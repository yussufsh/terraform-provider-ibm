// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"log"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
Datasource to get the list of dhcp servers in a power instance
*/

const PIDhcpServers = "servers"

func dataSourceIBMPIDhcps() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIDhcpServersRead,
		Schema: map[string]*schema.Schema{
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			// Computed Attributes
			PIDhcpServers: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PIDhcpId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the DHCP Server",
						},
						PIDhcpStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the DHCP Server",
						},
						PIDhcpNetwork: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DHCP Server private network",
						},
						PIDhcpLeases: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of DHCP Server PVM Instance leases",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PIDhcpInstanceIp: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP of the PVM Instance",
									},
									PIDhcpInstanceMac: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The MAC Address of the PVM Instance",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIDhcpServersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIDhcpClient(sess, cloudInstanceID)
	dhcpServers, err := client.GetAllWithContext(ctx, cloudInstanceID)
	if err != nil {
		log.Printf("[DEBUG] get all DHCP failed %v", err)
		return diag.Errorf("failed to perform get all DHCP operation with error %v", err)
	}

	result := make([]map[string]interface{}, 0, len(dhcpServers))
	for _, dhcpServer := range dhcpServers {
		server := map[string]interface{}{
			PIDhcpId:     *dhcpServer.ID,
			PIDhcpStatus: *dhcpServer.Status,
		}

		dhcpNetwork := dhcpServer.Network
		if dhcpNetwork != nil {
			server[PIDhcpNetwork] = *dhcpNetwork.ID
		}
		dhcpLeases := dhcpServer.Leases
		if dhcpLeases != nil {
			leaseList := make([]map[string]string, len(dhcpLeases))
			for i, lease := range dhcpLeases {
				leaseList[i] = map[string]string{
					PIDhcpInstanceIp:  *lease.InstanceIP,
					PIDhcpInstanceMac: *lease.InstanceMacAddress,
				}
			}
			server[PIDhcpLeases] = leaseList
		}
		result = append(result, server)
	}

	var genID, _ = uuid.GenerateUUID()
	d.SetId(genID)
	d.Set(PIDhcpServers, result)

	return nil
}
