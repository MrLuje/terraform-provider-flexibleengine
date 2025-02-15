package flexibleengine

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDcsProductV1DataSource_basic(t *testing.T) {
	dataSourceName := "data.flexibleengine_dcs_product_v1.product1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcsProductV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcsProductV1DataSourceID(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "spec_code", "dcs.single_node"),
				),
			},
		},
	})
}

func testAccCheckDcsProductV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Dcs product data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Dcs product data source ID not set")
		}

		return nil
	}
}

var testAccDcsProductV1DataSource_basic = fmt.Sprintf(`
data "flexibleengine_dcs_product_v1" "product1" {
  spec_code = "dcs.single_node"
}
`)
