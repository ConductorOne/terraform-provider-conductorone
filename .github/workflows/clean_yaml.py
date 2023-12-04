import yaml

with open('openapi.yaml', 'r') as file:
    data = yaml.safe_load(file)

# Function to recursively search and modify the data
def modify_yaml(obj):
    if isinstance(obj, dict):
        for key, value in obj.items():
            if key == 'updateMask':
                obj['x-speakeasy-terraform-ignore'] = True
            else:
                modify_yaml(value)
    elif isinstance(obj, list):
        for item in obj:
            modify_yaml(item)

modify_yaml(data)

with open('openapi.yaml', 'w') as file:
    yaml.dump(data, file, default_flow_style=False)