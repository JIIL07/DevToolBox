import json
import sys

from typing import Any, Dict, List


def to_camel_case(snake_str: str) -> str:

    components = snake_str.split('_')
    return components[0] + ''.join(x.title() for x in components[1:])


def to_pascal_case(snake_str: str) -> str:

    components = snake_str.split('_')
    return ''.join(x.title() for x in components)


def get_typescript_type(value: Any) -> str:

    if value is None:
        return "any"
    elif isinstance(value, bool):
        return "boolean"
    elif isinstance(value, int):
        return "number"
    elif isinstance(value, float):
        return "number"
    elif isinstance(value, str):
        return "string"
    elif isinstance(value, list):
        if not value:
            return "any[]"
        element_type = get_typescript_type(value[0])
        return f"{element_type}[]"
    elif isinstance(value, dict):
        return "object"
    else:
        return "any"


def generate_interface_name(key: str) -> str:

    return to_pascal_case(key)


def generate_interface(data: Dict[str, Any], interface_name: str = "GeneratedInterface") -> str:

    lines = [f"interface {interface_name} {{"]
    
    for key, value in data.items():
        field_name = to_camel_case(key)
        field_type = get_typescript_type(value)
        
        if isinstance(value, dict):
            nested_interface_name = generate_interface_name(key)
            lines.append(f"  {field_name}: {nested_interface_name};")
        elif isinstance(value, list) and value and isinstance(value[0], dict):
            nested_interface_name = generate_interface_name(key)
            lines.append(f"  {field_name}: {nested_interface_name}[];")
        else:
            lines.append(f"  {field_name}: {field_type};")
    
    lines.append("}")
    return "\n".join(lines)


def generate_nested_interfaces(data: Dict[str, Any], base_name: str = "GeneratedInterface") -> List[str]:

    interfaces = []
    
    for key, value in data.items():
        if isinstance(value, dict):
            interface_name = generate_interface_name(key)
            interface_code = generate_interface(value, interface_name)
            interfaces.append(interface_code)
            
            nested_interfaces = generate_nested_interfaces(value, interface_name)
            interfaces.extend(nested_interfaces)
            
        elif isinstance(value, list) and value and isinstance(value[0], dict):
            interface_name = generate_interface_name(key)
            interface_code = generate_interface(value[0], interface_name)
            interfaces.append(interface_code)
            
            nested_interfaces = generate_nested_interfaces(value[0], interface_name)
            interfaces.extend(nested_interfaces)
    
    return interfaces


def generate(input_json: str) -> str:

    try:
        data = json.loads(input_json)
        
        if not isinstance(data, dict):
            return "Error: Input must be a JSON object"
        
        main_interface = generate_interface(data)
        nested_interfaces = generate_nested_interfaces(data)
        
        all_interfaces = [main_interface] + nested_interfaces
        return "\n\n".join(all_interfaces)
        
    except json.JSONDecodeError as e:
        return f"Error: Invalid JSON - {str(e)}"
    except Exception as e:
        return f"Error: {str(e)}"


if __name__ == "__main__":
    if len(sys.argv) > 1:
        input_data = sys.argv[1]
    else:
        input_data = sys.stdin.read()
    
    result = generate(input_data)
    print(result)
