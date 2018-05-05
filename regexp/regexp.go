// Package regexp parses simple regular expressions into a tree structure.
package regexp

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var Trace = os.Getenv("TRACE") == "1"

// Node represents a regular expression tree node.
type Node struct {
	Type  rune
	Left  *Node
	Right *Node
}

// String prints the regular expression tree rooted at this node.
func (n *Node) String() string {
	if n == nil {
		return ""
	}
	switch {
	case n.Type >= 'a' && n.Type <= 'z':
		return string(n.Type)
	case n.Type == '*':
		return "*{" + n.Left.String() + "}"
	case n.Type == '|':
		return "|{" + n.Left.String() + "," + n.Right.String() + "}"
	case n.Type == '+':
		return n.Left.String() + n.Right.String()
	default:
		return fmt.Sprintf("unknown type %s", string(n.Type))
	}
}

// Parse reads RuneScanner s and returns a regular expression tree.
// Parentheses have highest precedence, followed by closure, concatenation,
// and alternation, in that order.
//
// The grammar is:
//     alternation   <-- concatenation
//     alternation   <-- concatenation | alternation
//     concatenation <-- closure
//     concatenation <-- closure concatenation
//     closure       <-- item
//     closure       <-- item *
//     item          <-- letter
//     item          <-- ( alternation )
//     letter        <-- a..z
func Parse(s io.RuneScanner) (*Node, error) {
	return alternation(s)
}

func alternation(s io.RuneScanner) (n *Node, e error) {
	if Trace {
		defer func() {
			fmt.Println("  TRACE: alternation =", n)
		}()
	}

	left, err := concatenation(s)
	if err != nil {
		return nil, err
	}

	r, _, err := s.ReadRune()
	if err == io.EOF {
		return left, nil
	}
	if err != nil {
		return nil, err
	}
	if r != '|' {
		if err := s.UnreadRune(); err != nil {
			return nil, err
		}
		return left, nil
	}

	right, err := alternation(s)
	if err != nil {
		return nil, err
	}

	return &Node{Type: '|', Left: left, Right: right}, nil
}

func concatenation(s io.RuneScanner) (n *Node, e error) {
	if Trace {
		defer func() {
			fmt.Println("  TRACE: concatenation =", n)
		}()
	}

	left, err := closure(s)
	if err != nil {
		return nil, err
	}

	r, _, err := s.ReadRune()
	if err == io.EOF {
		return left, nil
	}
	if err != nil {
		return nil, err
	}
	if err = s.UnreadRune(); err != nil {
		return nil, err
	}

	if r == '(' || r >= 'a' && r <= 'z' {
		right, err := concatenation(s)
		if err != nil {
			return nil, err
		}
		return &Node{Type: '+', Left: left, Right: right}, nil
	} else {
		return left, nil
	}
}

func closure(s io.RuneScanner) (n *Node, e error) {
	if Trace {
		defer func() {
			fmt.Println("  TRACE: closure =", n)
		}()
	}

	left, err := item(s)
	if err != nil {
		return nil, err
	}

	r, _, err := s.ReadRune()
	if err == io.EOF {
		return left, nil
	}
	if err != nil {
		return nil, err
	}

	if r != '*' {
		if err := s.UnreadRune(); err != nil {
			return nil, err
		}
		return left, nil
	}

	return &Node{Type: '*', Left: left}, nil
}

func item(s io.RuneScanner) (n *Node, e error) {
	if Trace {
		defer func() {
			fmt.Println("  TRACE: item =", n)
		}()
	}

	left, _, err := s.ReadRune()
	if err == io.EOF {
		return nil, errors.New("unexpected EOF")
	}
	if err != nil {
		return nil, err
	}
	if left == '(' {
		inner, err := alternation(s)
		if err != nil {
			return nil, err
		}
		right, _, err := s.ReadRune()
		if err == io.EOF {
			return nil, fmt.Errorf("expected ')' got EOF")
		}
		if err != nil {
			return nil, err
		}
		if right != ')' {
			return nil, fmt.Errorf("expected ')' got %q", right)
		}
		return inner, nil
	}

	if left < 'a' || left > 'z' {
		return nil, fmt.Errorf("unexpected character %q", left)
	}
	return &Node{Type: left}, nil
}
