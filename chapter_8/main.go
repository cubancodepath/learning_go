package main

import (
	"fmt"
	"strings"
)

type Number interface {
	int | float64
}

func Double[T Number](value T) T {
	return value * 2
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyType int

func (m MyType) String() string {
	return fmt.Sprintf("%d", m)
}

func Print[T Printable](val T) {
	fmt.Println(val.String())
}

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
}

func (l *LinkedList[T]) Add(val T) {
	newNode := &Element[T]{value: val}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList[T]) String() string {
	var final strings.Builder
	final.WriteString("[")
	for current := l.head; current != nil; current = current.next {
		fmt.Fprintf(&final, "%v", current.value)
		if current.next != nil {
			final.WriteString(" -> ")
		}
	}
	final.WriteString("]")
	return final.String()
}

func (l *LinkedList[T]) Insert(val T, pos int) {
	newNode := &Element[T]{value: val}
	if pos <= 0 || l.head == nil {
		newNode.next = l.head
		l.head = newNode
		return
	}

	current := l.head
	for i := 0; i < pos-1 && current.next != nil; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
}

func (l *LinkedList[T]) Index(val T) int {
	current := l.head
	for i := 0; current != nil; i, current = i+1, current.next {
		if current.value == val {
			return i
		}
	}
	return -1
}

func main() {
	result := Double(10)
	fmt.Println(result)

	Print(MyType(2))

	// Ejemplo de uso de LinkedList
	var list LinkedList[int]
	list.Add(10)
	list.Add(20)
	list.Add(30)

	fmt.Printf("Índice de 20: %d\n", list.Index(20))

	list.Insert(15, 1) // Insertar el número 15 en la posición 1
	fmt.Printf("Índice de 15 (recién insertado): %d\n", list.Index(15))
	fmt.Printf("Nuevo índice de 20 (se desplazó): %d\n", list.Index(20))

	fmt.Printf("Índice de 99 (no existe): %d\n", list.Index(99))
	fmt.Println(list.String())
}
