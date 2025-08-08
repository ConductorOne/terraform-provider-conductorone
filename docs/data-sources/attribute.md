# conductorone_attribute

The `conductorone_attribute` data source allows you to read existing attribute values in ConductorOne.

## Example Usage

```hcl
data "conductorone_attribute" "example" {
  id = "existing-attribute-id"
}

output "attribute_value" {
  value = data.conductorone_attribute.example.value
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) The ID of the AttributeValue to read.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `attribute_type_id` - The ID of the AttributeType that this AttributeValue belongs to.
* `value` - The value of the attribute.
* `created_at` - The timestamp when the attribute value was created (RFC3339 format).
* `updated_at` - The timestamp when the attribute value was last updated (RFC3339 format). 