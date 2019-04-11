package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCmd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Simple",
			input:    "cmd arg1 arg2 arg3",
			expected: []string{"cmd", "arg1", "arg2", "arg3"},
		}, {
			name:     "Double Quoted String",
			input:    `cmd "arg 1" "arg two is a sentence" "arg three is another sentence"`,
			expected: []string{"cmd", "arg 1", "arg two is a sentence", "arg three is another sentence"},
		}, {
			name:     "Single Quoted Strings",
			input:    `cmd 'arg 1' 'arg two is a sentence' 'arg three is another sentence'`,
			expected: []string{"cmd", "arg 1", "arg two is a sentence", "arg three is another sentence"},
		}, {
			name:     "Escaped or Internal Quotes",
			input:    `cmd "'this arg has single quotes'" '"this arg has double quotes"' 'single quote: \', double quote \"`,
			expected: []string{"cmd", `'this arg has single quotes'`, `"this arg has double quotes"`, `single quote: ', double quote "`},
		}, {
			name:     "Adjacent Tokens",
			input:    `cmd "hello",'world' good"day"to"you'"`,
			expected: []string{`cmd`, `hello,world`, `gooddaytoyou'`},
		}, {
			name: "Internal Newlines",
			input: `cmd "hello			
"`,
			expected: []string{`cmd`, "hello\t\t\t\n"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {

			args, err := ParseCommand(strings.NewReader(test.input))

			if err != nil {
				t.Fatal(err)
			}

			if len(args) != len(test.expected) {
				t.Fatalf("len(%#v) != len(%#v)", args, test.expected)
			}

			if !reflect.DeepEqual(args, test.expected) {
				t.Fatalf("ParseCommand(%q) returned %#v; expected %#v", test.input, args, test.expected)
			}
		})
	}
}
