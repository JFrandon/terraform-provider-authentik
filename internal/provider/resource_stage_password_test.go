package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceStagePassword(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceStagePassword(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("authentik_stage_password.name", "name", rName),
				),
			},
		},
	})
}

func testAccResourceStagePassword(name string) string {
	return fmt.Sprintf(`
resource "authentik_stage_password" "name" {
  name              = "%s"
  backends = ["django.contrib.auth.backends.ModelBackend"]
}
`, name)
}
