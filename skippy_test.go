package skippy

import (
	"fmt"
	"testing"
)

/*
(2) -------> b ----------------> nil
(1) -------> b ------> d -> e -> nil
(0) --> a -> b -> c -> d -> e -> nil
*/

// var a = &node{key: "a", value: []byte("a"), successors: []*node{b}}
// var b = &node{key: "b", value: []byte("b"), successors: []*node{c, d, nil}}
// var c = &node{key: "c", value: []byte("c"), successors: []*node{d}}
// var d = &node{key: "d", value: []byte("d"), successors: []*node{e, e}}
// var e = &node{key: "e", value: []byte("e"), successors: make([]*node, 2)}

// s.headers[2] = b
// s.headers[1] = b
// s.headers[0] = a

/*
(2) --------------> c --------------> nil
(1) --> a --------> c --> d --------> nil
(0) --> a --> b --> c --> d --> e --> nil
*/

var a = &node{key: "a", value: []byte("a"), successors: []*node{b, c}}
var b = &node{key: "b", value: []byte("b"), successors: []*node{c}}
var c = &node{key: "c", value: []byte("c"), successors: []*node{d, d, nil}}
var d = &node{key: "d", value: []byte("d"), successors: []*node{e, nil}}
var e = &node{key: "e", value: []byte("e"), successors: make([]*node, 1)}

func TestString(t *testing.T) {
	s := New()

	s.headers[2] = c
	s.headers[1] = a
	s.headers[0] = a

	/*
		(2) --> nil
		(1) --> nil
		(0) --> nil
	*/
	// s.headers[2] = nil
	// s.headers[1] = nil
	// s.headers[0] = nil

	fmt.Println(s.String())
}
