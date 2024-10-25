#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        schema['properties'][ntype].pop('nullable')

# Fix "site.asns" and "interface.tagged_vlans"
data["components"]["schemas"]["Site"]["properties"]["asns"]["items"].pop("type")
data["components"]["schemas"]["Interface"]["properties"]["tagged_vlans"]["items"].pop("type")
data["components"]["schemas"]["Site"]["properties"]["asns"]["items"]["$ref"] = "#/components/schemas/NestedASN"
data["components"]["schemas"]["Interface"]["properties"]["tagged_vlans"]["items"]["$ref"] = "#/components/schemas/NestedVLAN"
data["components"]["schemas"]["NestedASN"] = {
    "type": "object",
    "description": """Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
        dictionary of attributes which can be used to uniquely identify the related object. This class should be
        subclassed to return a full representation of the related object on read.""",
    "properties": {
        "id": {"type": "integer", "readOnly": True},
        "url": {"type": "string", "format": "uri", "readOnly": True},
        "display": {"type": "string", "readOnly": True},
        "asn": {
            "type": "integer",
            "maximum": 4294967295,
            "minimum": 1,
            "format": "int64",
            "description": "16- or 32-bit autonomous system number",
        },
    },
    "required": ["asn", "display", "id", "url"],
}

# Fix required aggragted counters absent from server responses
# Upstream Netbox issue: https://github.com/netbox-community/netbox/issues/14953

MISSING_TAG_FIELDS = ("tagged_items",)
MISSING_MANUFACTURER_FIELDS = ("devicetype_count", "inventoryitem_count", "platform_count")
MISSING_DEVICE_ROLE_FIELDS = ("device_count", "virtualmachine_count")
MISSING_DEVICE_TYPE_FIELDS = (
    "console_port_template_count",
    "console_server_port_template_count",
    "device_bay_template_count",
    "device_count",
    "front_port_template_count",
    "interface_template_count",
    "inventory_item_template_count",
    "module_bay_template_count",
    "power_outlet_template_count",
    "power_port_template_count",
    "rear_port_template_count",
)
MISSING_SITE_FIELDS = (
    "circuit_count",
    "device_count",
    "prefix_count",
    "rack_count",
    "virtualmachine_count",
    "vlan_count",
)

data["components"]["schemas"]["Tag"]["required"] = [e for e in data["components"]["schemas"]["Tag"]["required"] if e not in MISSING_TAG_FIELDS]
data["components"]["schemas"]["Manufacturer"]["required"] = [e for e in data["components"]["schemas"]["Manufacturer"]["required"] if e not in MISSING_MANUFACTURER_FIELDS]
data["components"]["schemas"]["DeviceRole"]["required"] = [e for e in data["components"]["schemas"]["DeviceRole"]["required"] if e not in MISSING_DEVICE_ROLE_FIELDS]
data["components"]["schemas"]["DeviceType"]["required"] = [e for e in data["components"]["schemas"]["DeviceType"]["required"] if e not in MISSING_DEVICE_TYPE_FIELDS]
data["components"]["schemas"]["Site"]["required"] = [e for e in data["components"]["schemas"]["Site"]["required"] if e not in MISSING_SITE_FIELDS]

# Fix broken vdisk queries

MISSING_VDISK_FIELDS = ("display",)

data["components"]["schemas"]["VirtualDisk"]["required"] = [e for e in data["components"]["schemas"]["VirtualDisk"]["required"] if e not in MISSING_VDISK_FIELDS]

# End fix for missing aggragted counters

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
