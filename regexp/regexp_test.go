package regexp

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"a*", "*{a}"},
		{"a", "a"},
		{"(a*)*", "*{*{a}}"},
		{"", "unexpected EOF"},
		{"a|b", "|{a,b}"},
		{"ab", "ab"},
		{"abcd", "abcd"},
		{"(((z)))", "z"},
		{"abc*d", "ab*{c}d"},
		{"abc|def", "|{abc,def}"},
		{"ab*c*|d*ef*", "|{a*{b}*{c},*{d}e*{f}}"},
		{"a(b|c)*", "a*{|{b,c}}"},
	}

	for _, c := range cases {
		if Trace {
			fmt.Println("TRACE: testing", c.input)
		}
		r := strings.NewReader(c.input)
		result, err := Parse(r)
		if err != nil {
			if err.Error() == c.expected {
				continue
			}
			t.Errorf("%s on %q, want %q", err, c.input, c.expected)
			continue
		}
		if result.String() != c.expected {
			t.Errorf("Parse(%q) = %q, want %q", c.input, result, c.expected)
		}
	}
}
