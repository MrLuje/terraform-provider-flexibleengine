---
subcategory: "Virtual Private Cloud (VPC)"
---

# flexibleengine_networking_secgroup_v2

Manages a Security Group resource within FlexibleEngine.

## Example Usage

```hcl
resource "flexibleengine_networking_secgroup_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "My neutron security group"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to obtain the V2 networking client.
    A networking client is needed to create a port. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group.

* `name` - (Required) A unique name for the security group.

* `description` - (Optional) A unique name for the security group.

* `delete_default_rules` - (Optional) Whether or not to delete the default
    egress security rules. This is `false` by default. See the below note
    for more information.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

## Default Security Group Rules

In most cases, FlexibleEngine will create some egress security group rules for each
new security group. These security group rules will not be managed by
Terraform, so if you prefer to have *all* aspects of your infrastructure
managed by Terraform, set `delete_default_rules` to `true` and then create
separate security group rules such as the following:

```hcl
resource "flexibleengine_networking_secgroup_rule_v2" "secgroup_rule_v4" {
  direction = "egress"
  ethertype = "IPv4"
  security_group_id = flexibleengine_networking_secgroup_v2.secgroup.id
}

resource "flexibleengine_networking_secgroup_rule_v2" "secgroup_rule_v6" {
  direction = "egress"
  ethertype = "IPv6"
  security_group_id = flexibleengine_networking_secgroup_v2.secgroup.id
}
```

Please note that this behavior may differ depending on the configuration of
the FlexibleEngine cloud. The above illustrates the current default Neutron
behavior. Some FlexibleEngine clouds might provide additional rules and some might
not provide any rules at all (in which case the `delete_default_rules` setting
is moot).

## Import

Security Groups can be imported using the `id`, e.g.

```
$ terraform import flexibleengine_networking_secgroup_v2.secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
```
