// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceDiskDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_disk.test1"
	var instance string
	insResName := "ibm_is_instance.testacc_instance"
	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists(insResName, instance),
					resource.TestCheckResourceAttr(
						insResName, "name", name),
					resource.TestCheckResourceAttr(
						insResName, "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						insResName, "disks.#", "1"),
					resource.TestCheckResourceAttrSet(
						insResName, "disks.0.name"),
					resource.TestCheckResourceAttrSet(
						insResName, "disks.0.size"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMISInstanceDiskDataSourceConfig(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "size"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceDiskDataSourceConfig(vpcname, subnetname, sshname, publicKey, volname, name string) string {
	return testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volname, name) + fmt.Sprintf(`
	data "ibm_is_instance" "ins" {
		name = "%s"
		private_key = file("test-fixtures/.ssh/id_rsa")
  		passphrase  = ""
	  }
      data "ibm_is_instance_disks" "test1" {
		instance = data.ibm_is_instance.ins.id
		
      }
	  data "ibm_is_instance_disk" "test1" {
	    instance = data.ibm_is_instance.ins.id
	    disk = data.ibm_is_instance_disks.test1.disks.0.id
	}`, name)
}
