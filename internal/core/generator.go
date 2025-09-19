package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type CodeGenerator interface {
	Generate(input string) (string, error)
	GetName() string
	GetDescription() string
}

type GoStructGenerator struct {
	name        string
	description string
}

func NewGoStructGenerator() *GoStructGenerator {
	return &GoStructGenerator{
		name:        "go-struct",
		description: "Генерирует Go структуры с JSON тегами из JSON схемы",
	}
}

func (g *GoStructGenerator) GetName() string {
	return g.name
}

func (g *GoStructGenerator) GetDescription() string {
	return g.description
}

func (g *GoStructGenerator) Generate(input string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	structName := "GeneratedStruct"
	code, err := g.generateStruct(structName, data)
	if err != nil {
		return "", fmt.Errorf("ошибка генерации структуры: %w", err)
	}

	return code, nil
}

func (g *GoStructGenerator) generateStruct(structName string, data interface{}) (string, error) {
	var builder strings.Builder
	var nestedStructs []string

	builder.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	switch v := data.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := v[key]
			fieldName := g.ToPascalCase(key)
			fieldType := g.GetGoType(value)
			jsonTag := fmt.Sprintf("`json:\"%s\"`", key)

			if nestedMap, ok := value.(map[string]interface{}); ok {
				nestedStructName := g.ToPascalCase(key)
				builder.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, nestedStructName, jsonTag))

				nestedStruct, err := g.generateStruct(nestedStructName, nestedMap)
				if err != nil {
					return "", err
				}
				nestedStructs = append(nestedStructs, nestedStruct)
			} else if slice, ok := value.([]interface{}); ok && len(slice) > 0 {
				elementType := g.GetGoType(slice[0])
				if nestedMap, ok := slice[0].(map[string]interface{}); ok {
					nestedStructName := g.ToPascalCase(key)
					builder.WriteString(fmt.Sprintf("\t%s []%s %s\n", fieldName, nestedStructName, jsonTag))

					nestedStruct, err := g.generateStruct(nestedStructName, nestedMap)
					if err != nil {
						return "", err
					}
					nestedStructs = append(nestedStructs, nestedStruct)
				} else {
					builder.WriteString(fmt.Sprintf("\t%s []%s %s\n", fieldName, elementType, jsonTag))
				}
			} else {
				builder.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, fieldType, jsonTag))
			}
		}
	default:
		return "", fmt.Errorf("неподдерживаемый тип данных: %T", data)
	}

	builder.WriteString("}")

	if len(nestedStructs) > 0 {
		builder.WriteString("\n\n")
		for _, nested := range nestedStructs {
			builder.WriteString(nested)
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}

func (g *GoStructGenerator) ToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	var parts []string
	var current strings.Builder
	
	for i, r := range s {
		if r == '_' || r == '-' || r == ' ' {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		} else if i > 0 && r >= 'A' && r <= 'Z' {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			current.WriteRune(r)
		} else {
			current.WriteRune(r)
		}
	}
	
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		result.WriteString(strings.ToUpper(part[:1]) + strings.ToLower(part[1:]))
	}

	return result.String()
}

func (g *GoStructGenerator) GetGoType(value interface{}) string {
	if value == nil {
		return "interface{}"
	}

	switch v := value.(type) {
	case bool:
		return "bool"
	case float64:
		if v == float64(int64(v)) {
			return "int"
		}
		return "float64"
	case string:
		return "string"
	case []interface{}:
		if len(v) == 0 {
			return "[]interface{}"
		}
		elementType := g.GetGoType(v[0])
		return fmt.Sprintf("[]%s", elementType)
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return reflect.TypeOf(value).String()
	}
}

type GeneratorRegistry struct {
	generators map[string]CodeGenerator
}

func NewGeneratorRegistry() *GeneratorRegistry {
	registry := &GeneratorRegistry{
		generators: make(map[string]CodeGenerator),
	}
	
	registry.Register(NewGoStructGenerator())
	
	return registry
}

func (r *GeneratorRegistry) Register(generator CodeGenerator) {
	r.generators[generator.GetName()] = generator
}

func (r *GeneratorRegistry) Get(name string) (CodeGenerator, bool) {
	generator, exists := r.generators[name]
	return generator, exists
}

func (r *GeneratorRegistry) List() []CodeGenerator {
	generators := make([]CodeGenerator, 0, len(r.generators))
	for _, generator := range r.generators {
		generators = append(generators, generator)
	}
	return generators
}

func (r *GeneratorRegistry) GetNames() []string {
	names := make([]string, 0, len(r.generators))
	for name := range r.generators {
		names = append(names, name)
	}
	return names
}