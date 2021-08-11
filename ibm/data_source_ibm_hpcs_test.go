// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMHPCSDatasourceBasic(t *testing.T) {
	instanceName := "test-hpcs"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMHPCSDatasourceConfig(instanceName),
				Destroy: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_hpcs.hpcs", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_hpcs.hpcs", "service", "hs-crypto"),
				),
			},
		},
	})
}

func testAccCheckIBMHPCSDatasourceConfig(instanceName string) string {
	return fmt.Sprintf(`
	data "ibm_hpcs" "hpcs" {
		name              = "%s"
	}
	`, instanceName)

}
