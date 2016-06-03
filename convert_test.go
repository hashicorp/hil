package hil

import (
	"reflect"
	"testing"

	"github.com/hashicorp/hil/ast"
)

func TestInterfaceToVariable(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected ast.Variable
	}{
		{
			name:  "string",
			input: "Hello world",
			expected: ast.Variable{
				Type:  ast.TypeString,
				Value: "Hello world",
			},
		},
		{
			name:  "empty list",
			input: []interface{}{},
			expected: ast.Variable{
				Type:  ast.TypeList,
				Value: []ast.Variable{},
			},
		},
		{
			name:  "int",
			input: 1,
			expected: ast.Variable{
				Type:  ast.TypeString,
				Value: "1",
			},
		},
		{
			name:  "list of strings",
			input: []string{"Hello", "World"},
			expected: ast.Variable{
				Type: ast.TypeList,
				Value: []ast.Variable{
					ast.Variable{
						Type:  ast.TypeString,
						Value: "Hello",
					},
					ast.Variable{
						Type:  ast.TypeString,
						Value: "World",
					},
				},
			},
		},
		{
			name:  "list of lists of strings",
			input: [][]string{[]string{"Hello", "World"}, []string{"Goodbye", "World"}},
			expected: ast.Variable{
				Type: ast.TypeList,
				Value: []ast.Variable{
					ast.Variable{
						Type: ast.TypeList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TypeString,
								Value: "Hello",
							},
							ast.Variable{
								Type:  ast.TypeString,
								Value: "World",
							},
						},
					},
					ast.Variable{
						Type: ast.TypeList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TypeString,
								Value: "Goodbye",
							},
							ast.Variable{
								Type:  ast.TypeString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:  "map of string->string",
			input: map[string]string{"Hello": "World", "Foo": "Bar"},
			expected: ast.Variable{
				Type: ast.TypeMap,
				Value: map[string]ast.Variable{
					"Hello": ast.Variable{
						Type:  ast.TypeString,
						Value: "World",
					},
					"Foo": ast.Variable{
						Type:  ast.TypeString,
						Value: "Bar",
					},
				},
			},
		},
		{
			name: "map of lists of strings",
			input: map[string][]string{
				"Hello":   []string{"Hello", "World"},
				"Goodbye": []string{"Goodbye", "World"},
			},
			expected: ast.Variable{
				Type: ast.TypeMap,
				Value: map[string]ast.Variable{
					"Hello": ast.Variable{
						Type: ast.TypeList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TypeString,
								Value: "Hello",
							},
							ast.Variable{
								Type:  ast.TypeString,
								Value: "World",
							},
						},
					},
					"Goodbye": ast.Variable{
						Type: ast.TypeList,
						Value: []ast.Variable{
							ast.Variable{
								Type:  ast.TypeString,
								Value: "Goodbye",
							},
							ast.Variable{
								Type:  ast.TypeString,
								Value: "World",
							},
						},
					},
				},
			},
		},
		{
			name:  "empty map",
			input: map[string]string{},
			expected: ast.Variable{
				Type:  ast.TypeMap,
				Value: map[string]ast.Variable{},
			},
		},
		{
			name: "three-element map",
			input: map[string]string{
				"us-west-1": "ami-123456",
				"us-west-2": "ami-456789",
				"eu-west-1": "ami-012345",
			},
			expected: ast.Variable{
				Type: ast.TypeMap,
				Value: map[string]ast.Variable{
					"us-west-1": ast.Variable{
						Type:  ast.TypeString,
						Value: "ami-123456",
					},
					"us-west-2": ast.Variable{
						Type:  ast.TypeString,
						Value: "ami-456789",
					},
					"eu-west-1": ast.Variable{
						Type:  ast.TypeString,
						Value: "ami-012345",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		output, err := InterfaceToVariable(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("%s:\nExpected: %s\n     Got: %s\n", tc.name, tc.expected, output)
		}
	}
}
