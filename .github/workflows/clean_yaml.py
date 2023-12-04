import yaml
import sys

def modify_yaml(obj, parents):
    if isinstance(obj, dict):
        for key, value in list(obj.items()):
            if key == 'updateMask':
                obj['updateMask']['x-speakeasy-terraform-ignore'] = True
            else:
                modify_yaml(value, parents + [key])
    elif isinstance(obj, list):
        for item in obj:
            modify_yaml(item, parents)

if len(sys.argv) != 2:
    print("Usage: python clean_yaml.py <filename>")
    sys.exit(1)

filename = sys.argv[1]

with open(filename, 'r') as file:
    data = yaml.safe_load(file)

modify_yaml(data, [])

with open(filename, 'w') as file:
    yaml.dump(data, file, default_flow_style=False, sort_keys=False)
