// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

// LinkedHashMapExample to demonstrate basic usage of LinkedHashMapExample
func main() {
	m := linkedhashmap.New() // empty (keys are of type int)
	m.Put(5, "b")
	m.Put(1, "b")
	m.Put(2, "x")
	m.Put(3, "a")
	m.Put(4, "b")

	// 按插入顺序打印
	m.Each(func(key interface{}, value interface{}) {
		fmt.Println(key)
		fmt.Println(value)
	})

	_, _ = m.Get(2)          // b, true
	_, _ = m.Get(6)          // nil, false
	_ = m.Values()           // (insertion-order)
	_ = m.Keys()             // (insertion-order)
	m.Remove(1)
	m.Clear()                // empty
	m.Empty()                // true
	m.Size()                 // 0
}
