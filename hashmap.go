package main

import "fmt"

type HashMap struct {
	array [arraysize]*linkedlist
}

const arraysize = 7

type linkedlist struct {
	head *Node
}

type Node struct {
	key  string
	next *Node
}

func (h *HashMap) Insert(k string) {
	h.array[hash(k)].insert(k)
}

func (h *HashMap) Search(k string) {
	ans := h.array[hash(k)].search(k)
	fmt.Println(ans)
}

func (h *HashMap) Delete(k string) {
	h.array[hash(k)].delete(k)
}

func (l *linkedlist) insert(k string) {
	newNode := &Node{key: k}
	newNode.next = l.head
	l.head = newNode
}

func (l *linkedlist) search(k string) bool {
	curr := l.head
	for curr != nil {
		if curr.key == k {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *linkedlist) delete(k string) {
	if l.head.key == k {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.key == k {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func Init() *HashMap {
	hp := &HashMap{}
	for i := range hp.array {
		hp.array[i] = &linkedlist{}
	}

	return hp
}

func hash(key string) int {
	sum := 0
	for _, i := range key {
		sum += int(i)
	}

	return sum % arraysize
}

func main() {
	hm := Init()
	hm.Insert("Raj")
	hm.Insert("Gogo")
	hm.Insert("Dora")
	fmt.Println(hm.array[hash("Raj")].head)
	fmt.Println(hm.array[hash("Raj")].head.next)

}
