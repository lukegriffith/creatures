package LinkedList

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	key  T
}

func (n *Node[T]) Value() T {
	return n.key
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (L *List[T]) Insert(key T) {
	list := &Node[T]{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List[T]) Reverse() {
	curr := l.head
	var prev *Node[T]
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}
