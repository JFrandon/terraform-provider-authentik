---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "authentik_stage_authenticator_sms Resource - terraform-provider-authentik"
subcategory: ""
description: |-
  
---

# authentik_stage_authenticator_sms (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account_sid** (String, Sensitive)
- **auth** (String, Sensitive)
- **from_number** (String)
- **name** (String)

### Optional

- **auth_password** (String, Sensitive)
- **auth_type** (String) Defaults to `basic`.
- **configure_flow** (String)
- **id** (String) The ID of this resource.
- **sms_provider** (String) Defaults to `twilio`.

