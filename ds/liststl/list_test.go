// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package liststl

import (
	"container/list"
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	l := New[string]()
	l.PushBack("n1")
	l.PushBack("n2")
	node := l.Front()
	for node != nil {
		fmt.Printf("node.Value: %v\n", node.Value)
		node = node.Next()
	}
}
func TestFunctionList(t *testing.T) {
	l := list.New()
	l.PushBack("n1")
	l.PushBack("n2")
	node := l.Front()
	for node != nil {
		fmt.Printf("node.Value: %v\n", node.Value)
		node = node.Next()
	}
}
