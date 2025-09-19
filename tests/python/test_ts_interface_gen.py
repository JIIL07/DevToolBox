#!/usr/bin/env python3
import json
import sys
import os

sys.path.append(os.path.join(os.path.dirname(__file__), '../../plugins/official'))

from ts_interface_gen import generate, to_camel_case, to_pascal_case, get_typescript_type


def test_to_camel_case():
    assert to_camel_case("user_name") == "userName"
    assert to_camel_case("user_id_field") == "userIdField"
    assert to_camel_case("name") == "name"
    assert to_camel_case("") == ""


def test_to_pascal_case():
    assert to_pascal_case("user_name") == "UserName"
    assert to_pascal_case("user_id_field") == "UserIdField"
    assert to_pascal_case("name") == "Name"
    assert to_pascal_case("") == ""


def test_get_typescript_type():
    assert get_typescript_type("string") == "string"
    assert get_typescript_type(123) == "number"
    assert get_typescript_type(123.45) == "number"
    assert get_typescript_type(True) == "boolean"
    assert get_typescript_type(None) == "any"
    assert get_typescript_type([]) == "any[]"
    assert get_typescript_type(["a", "b"]) == "string[]"
    assert get_typescript_type([1, 2]) == "number[]"
    assert get_typescript_type({}) == "object"


def test_generate_simple_interface():
    input_json = '{"name": "string", "age": "number"}'
    result = generate(input_json)
    
    assert "interface GeneratedInterface" in result
    assert "name: string;" in result
    assert "age: string;" in result


def test_generate_nested_interface():
    input_json = '{"user": {"name": "John", "age": 30}, "active": true}'
    result = generate(input_json)
    
    assert "interface GeneratedInterface" in result
    assert "interface User" in result
    assert "user: User;" in result
    assert "name: string;" in result
    assert "age: number;" in result
    assert "active: boolean;" in result


def test_generate_array_interface():
    input_json = '{"users": [{"name": "John", "age": 30}]}'
    result = generate(input_json)
    
    assert "interface GeneratedInterface" in result
    assert "interface Users" in result
    assert "users: Users[];" in result


def test_generate_invalid_json():
    result = generate('{"invalid": json}')
    assert result.startswith("Error:")


def test_generate_non_object():
    result = generate('"just a string"')
    assert "Error: Input must be a JSON object" in result


if __name__ == "__main__":
    test_to_camel_case()
    test_to_pascal_case()
    test_get_typescript_type()
    test_generate_simple_interface()
    test_generate_nested_interface()
    test_generate_array_interface()
    test_generate_invalid_json()
    test_generate_non_object()
    
    print("All tests passed!")
