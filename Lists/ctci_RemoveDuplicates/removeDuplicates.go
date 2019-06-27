package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(4)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(4)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(3)
	l.PushBack(5)
	l.PushBack(1)

	r := removeListDuplicates3(l)

	for e := r.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func removeListDuplicates(l *list.List) *list.List {
	r := list.New()
	m := make(map[interface{}]bool)

	for e := l.Front(); e != nil; e = e.Next() {
		m[e.Value] = true
	}

	for k, _ := range m {
		r.PushBack(k)
	}

	return r
}

func removeListDuplicates2(l *list.List) *list.List {
	m := make(map[interface{}]bool)

	for e := l.Front(); e != nil; {
		if m[e.Value] {
			t := e.Next()
			l.Remove(e)
			m[e.Value] = true
			e = t
		} else {
			m[e.Value] = true
			e = e.Next()
		}
	}

	return l
}

func removeListDuplicates3(l *list.List) *list.List {
	for e := l.Front(); e != nil; {
		for e1 := e.Next(); e1 != nil; {
			if e.Value == e1.Value {
				t := e1.Next()
				l.Remove(e1)
				e1 = t
			} else {
				e1 = e1.Next()
			}
		}
		e = e.Next()
	}

	return l
}
