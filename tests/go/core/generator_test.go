package core

import (
	"strings"
	"testing"

	"github.com/yourname/devtoolbox/internal/core"
)

func TestGoStructGenerator_Generate(t *testing.T) {
	generator := core.NewGoStructGenerator()

	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:  "простая структура",
			input: `{"name":"string","age":"int"}`,
			expected: `type GeneratedStruct struct {
	Age string ` + "`json:\"age\"`" + `
	Name string ` + "`json:\"name\"`" + `
}`,
			hasError: false,
		},
		{
			name:  "структура с числами",
			input: `{"id":123,"price":99.99,"active":true}`,
			expected: `type GeneratedStruct struct {
	Active bool ` + "`json:\"active\"`" + `
	Id int ` + "`json:\"id\"`" + `
	Price float64 ` + "`json:\"price\"`" + `
}`,
			hasError: false,
		},
		{
			name:  "вложенная структура",
			input: `{"user":{"name":"John","email":"john@example.com"},"posts":["post1","post2"]}`,
			expected: `type GeneratedStruct struct {
	Posts []string ` + "`json:\"posts\"`" + `
	User User ` + "`json:\"user\"`" + `
}

type User struct {
	Email string ` + "`json:\"email\"`" + `
	Name string ` + "`json:\"name\"`" + `
}`,
			hasError: false,
		},
		{
			name:  "массив объектов",
			input: `{"users":[{"name":"John","age":30},{"name":"Jane","age":25}]}`,
			expected: `type GeneratedStruct struct {
	Users []Users ` + "`json:\"users\"`" + `
}

type Users struct {
	Age int ` + "`json:\"age\"`" + `
	Name string ` + "`json:\"name\"`" + `
}`,
			hasError: false,
		},
		{
			name:     "невалидный JSON",
			input:    `{"name":"string","age":}`,
			expected: "",
			hasError: true,
		},
		{
			name:  "пустой объект",
			input: `{}`,
			expected: `type GeneratedStruct struct {
}`,
			hasError: false,
		},
		{
			name:  "поля с подчеркиваниями",
			input: `{"user_name":"string","user_email":"string"}`,
			expected: `type GeneratedStruct struct {
	UserEmail string ` + "`json:\"user_email\"`" + `
	UserName string ` + "`json:\"user_name\"`" + `
}`,
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := generator.Generate(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("ожидалась ошибка, но получили результат: %s", result)
				}
				return
			}

			if err != nil {
				t.Errorf("неожиданная ошибка: %v", err)
				return
			}

			result = strings.TrimSpace(result)
			expected := strings.TrimSpace(tt.expected)

			if result != expected {
				t.Errorf("результат не соответствует ожидаемому:\nПолучено:\n%s\n\nОжидалось:\n%s", result, expected)
			}
		})
	}
}

func TestGoStructGenerator_GetName(t *testing.T) {
	generator := core.NewGoStructGenerator()
	expected := "go-struct"
	
	if generator.GetName() != expected {
		t.Errorf("ожидалось имя %s, получили %s", expected, generator.GetName())
	}
}

func TestGoStructGenerator_GetDescription(t *testing.T) {
	generator := core.NewGoStructGenerator()
	expected := "Генерирует Go структуры с JSON тегами из JSON схемы"
	
	if generator.GetDescription() != expected {
		t.Errorf("ожидалось описание %s, получили %s", expected, generator.GetDescription())
	}
}

func TestGoStructGenerator_ToPascalCase(t *testing.T) {
	generator := core.NewGoStructGenerator()

	tests := []struct {
		input    string
		expected string
	}{
		{"user_name", "UserName"},
		{"user-name", "UserName"},
		{"user name", "UserName"},
		{"userName", "UserName"},
		{"user", "User"},
		{"", ""},
		{"_user_", "User"},
		{"user_id_field", "UserIdField"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := generator.ToPascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("для входа %s ожидалось %s, получили %s", tt.input, tt.expected, result)
			}
		})
	}
}

func TestGoStructGenerator_GetGoType(t *testing.T) {
	generator := core.NewGoStructGenerator()

	tests := []struct {
		input    interface{}
		expected string
	}{
		{true, "bool"},
		{false, "bool"},
		{123, "int"},
		{123.45, "float64"},
		{"string", "string"},
		{[]interface{}{"a", "b"}, "[]string"},
		{[]interface{}{1, 2}, "[]int"},
		{[]interface{}{}, "[]interface{}"},
		{map[string]interface{}{}, "map[string]interface{}"},
		{nil, "interface{}"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := generator.GetGoType(tt.input)
			if result != tt.expected {
				t.Errorf("для входа %v ожидался тип %s, получили %s", tt.input, tt.expected, result)
			}
		})
	}
}

func TestGeneratorRegistry(t *testing.T) {
	registry := core.NewGeneratorRegistry()

	generator, exists := registry.Get("go-struct")
	if !exists {
		t.Error("генератор go-struct не найден в реестре")
	}

	if generator.GetName() != "go-struct" {
		t.Errorf("ожидалось имя go-struct, получили %s", generator.GetName())
	}

	generators := registry.List()
	if len(generators) == 0 {
		t.Error("реестр не содержит генераторов")
	}

	names := registry.GetNames()
	if len(names) == 0 {
		t.Error("реестр не содержит имен генераторов")
	}

	found := false
	for _, name := range names {
		if name == "go-struct" {
			found = true
			break
		}
	}
	if !found {
		t.Error("go-struct не найден в списке имен генераторов")
	}
}

func TestGeneratorRegistry_Register(t *testing.T) {
	registry := core.NewGeneratorRegistry()

	testGenerator := &testGenerator{
		name:        "test-generator",
		description: "Тестовый генератор",
	}

	registry.Register(testGenerator)

	generator, exists := registry.Get("test-generator")
	if !exists {
		t.Error("тестовый генератор не найден в реестре")
	}

	if generator.GetName() != "test-generator" {
		t.Errorf("ожидалось имя test-generator, получили %s", generator.GetName())
	}
}

type testGenerator struct {
	name        string
	description string
}

func (t *testGenerator) Generate(input string) (string, error) {
	return "test output", nil
}

func (t *testGenerator) GetName() string {
	return t.name
}

func (t *testGenerator) GetDescription() string {
	return t.description
}