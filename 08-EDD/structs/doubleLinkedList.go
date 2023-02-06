package structs

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *DoubleLinkedList) isEmpty() bool {
	return l.Size == 0
}

func (l *DoubleLinkedList) First() *Node {
	return l.Head
}

func (l *DoubleLinkedList) Last() *Node {
	return l.Tail
}

func (l *DoubleLinkedList) AddFirst(value interface{}) {
	newNode := &Node{value, nil, nil} // Create a new node
	if l.isEmpty() {                  // Verify if the list is empty
		l.Head = newNode // If it is empty, the new node is the head and the tail
		l.Tail = newNode
	} else { // If it is not empty, the new node is the head and the old head is the next node
		newNode.Next = l.Head // The old head is the next node  <- [newNode] -> [oldHead]
		l.Head.Prev = newNode // The new node is the previous node <- [newNode] <- [oldHead]
		l.Head = newNode      // The new node is the head
	}
	l.Size++
}

func (l *DoubleLinkedList) AddLast(value interface{}) {
	newNode := &Node{value, nil, nil} // Create a new node
	if l.isEmpty() {                  // Verify if the list is empty
		l.Head = newNode // If it is empty, the new node is the head and the tail
		l.Tail = newNode
	} else { // If it is not empty, the new node is the tail and the old tail is the previous node
		newNode.Prev = l.Tail // The old tail is the previous node <- [oldTail] <- [newNode]
		l.Tail.Next = newNode // The new node is the next node <- [oldTail] -> [newNode]
		l.Tail = newNode      // The new node is the tail
	}
	l.Size++
}

func (l *DoubleLinkedList) RemoveFirst() {
	if l.isEmpty() { // Verify if the list is empty
		return
	}
	if l.Head == l.Tail { // Verify if the list has only one element
		l.Head = nil // If it has only one element, the head and the tail are nil
		l.Tail = nil
	} else { // If it has more than one element, the head is the next node
		l.Head = l.Head.Next // The next node is the head
		l.Head.Prev = nil    // The previous node of the new head is nil
	}
	l.Size--
}

func (l *DoubleLinkedList) RemoveLast() {
	if l.isEmpty() { // Verify if the list is empty
		return
	}
	if l.Head == l.Tail { // Verify if the list has only one element
		l.Head = nil // If it has only one element, the head and the tail are nil
		l.Tail = nil
	} else { // If it has more than one element, the tail is the previous node
		l.Tail = l.Tail.Prev // The previous node is the tail
		l.Tail.Next = nil    // The next node of the new tail is nil
	}
	l.Size--
}

func (l *DoubleLinkedList) Remove(node *Node) {
	if l.isEmpty() { // Verify if the list is empty
		return
	}
	if node == l.Head { // Verify if the node is the head
		l.RemoveFirst() // If it is the head, remove the first node
		return
	}
	if node == l.Tail { // Verify if the node is the tail
		l.RemoveLast() // If it is the tail, remove the last node
		return
	}
	node.Prev.Next = node.Next // The previous node of the node to remove points to the next node of the node to remove
	node.Next.Prev = node.Prev // The next node of the node to remove points to the previous node of the node to remove
	l.Size--
}

func (l *DoubleLinkedList) RemoveByIndex(index int) {
	if l.isEmpty() { // Verify if the list is empty
		return
	}
	if index == 0 { // Verify if the index is the first
		l.RemoveFirst() // If it is the first, remove the first node
		return
	}
	if index == l.Size-1 { // Verify if the index is the last
		l.RemoveLast() // If it is the last, remove the last node
		return
	}
	node := l.Head
	for i := 0; i < index; i++ { // Iterate until the index
		node = node.Next
	}
	node.Prev.Next = node.Next // The previous node of the node to remove points to the next node of the node to remove
	node.Next.Prev = node.Prev // The next node of the node to remove points to the previous node of the node to remove
	l.Size--
}

func (l *DoubleLinkedList) Print() {
	if l.isEmpty() { // Verify if the list is empty
		return
	}
	node := l.Head
	for node != nil { // Iterate until the last node
		print(node, " ")
		node = node.Next
	}
	println()
}
