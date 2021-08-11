// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVolumeProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_volume_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_volume_profiles" "test1" {
      }`)
}
