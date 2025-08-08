# conductorone_attribute

The `conductorone_attribute` resource allows you to manage attribute values in ConductorOne.

## Example Usage

```hcl
resource "conductorone_attribute" "example" {
  attribute_type_id = "your-attribute-type-id"
  value             = "example-value"
}
```

## Argument Reference

The following arguments are supported:

* `attribute_type_id` - (Required) The ID of the AttributeType that this AttributeValue belongs to.
* `value` - (Required) The value of the attribute. This field requires replacement if changed.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the AttributeValue.
* `created_at` - The timestamp when the attribute value was created (RFC3339 format).
* `updated_at` - The timestamp when the attribute value was last updated (RFC3339 format).

## Import

Attribute values can be imported using their ID:

```bash
terraform import conductorone_attribute.example <attribute-id>
```

## Notes

* Attribute values cannot be updated once created. If you need to change the value, you must delete and recreate the resource.
* The `value` field requires replacement if changed, meaning Terraform will delete and recreate the resource. 