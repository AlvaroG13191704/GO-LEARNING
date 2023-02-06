package structs

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}
