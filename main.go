package main

import (
	"fmt"
	"math/rand"
)

type List struct {
	head *Element
	count int
}

type Element struct {
	value int
	next *Element
}

func (l *List) Size() int {
	return l.count
}

func (l *List) IsEmpty() bool  {
	return l.count == 0
}

func (l *List) AddHead(val int)  {
	l.head = &Element{
		value: val,
		next: l.head,
	}
	l.count++
}

// high complexity of creation the list
// to optimize we can add tail to List
func (l *List) AddLast(val int)  {

	curr := l.head

	newElement := &Element{
		value: val,
		next:  nil,
	}

	// Empty list
	if curr == nil {
		l.head = newElement
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newElement
}

func (l *List) Print() string {
	temp := l.head

	res := ""
	for temp != nil {
		res = res + fmt.Sprintf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println(res)
	return res
}

// Delete first occurrence of value
func (l *List) Delete(val int) bool {
	curr := l.head

	// if list if empty
	if l.IsEmpty() {
		return false
	}

	// if deleted value is head
	if val == curr.value {
		l.head = curr.next
		l.count--
		return true
	}

	// iterate all values
	for curr.next != nil {
		if val == curr.next.value {
			curr.next = curr.next.next
			l.count--
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *List) FindElementByVal(val int) *Element  {

	curr := l.head

	if l.IsEmpty() {
		return nil
	}

	for curr.next != nil {
		if val == curr.value {
			return curr
		}
	}

	return nil
}

// Add loop to head
func (l *List) AddLoopToHead()  {
	curr := l.head

	loopElement := &Element{
		value: rand.Intn(l.Size()),
		next:  l.head,
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = loopElement
}

func (l *List) AddLoopToRandomEl()  {
	curr := l.head

	randEl := rand.Intn(l.Size())
	loopEl := &Element{
		value: randEl,
		next:  l.FindElementByVal(randEl),
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = loopEl
}

// IsLoop defines if there is a loop in linked list
func (l *List) IsLoop() bool  {

	if l.IsEmpty() {
		return false
	}

	// one element in list and next is not nil
	if l.Size() == 1 && l.head.next != nil {
		return true
	}

	oneStep := l.head
	twoStep := l.head

	for twoStep.next != nil && twoStep.next.next != nil {
		oneStep = oneStep.next
		twoStep = twoStep.next.next

		if oneStep == twoStep {
			return true
		}
	}

	return false
}

// hasLoop defines if there is a loop using map
func (l *List) HasLoop() bool {
	if l.IsEmpty() {
		return false
	}

	// one element in list and next is not nil
	if l.Size() == 1 && l.head.next != nil {
		return true
	}

	curr := l.head

	inList := map[interface{}]int{}
	for curr.next != nil {
		if val, ok := inList[curr]; ok {
			return true
		} else {
			inList[curr] = val
		}
		curr = curr.next
	}

	return false
}

func NewLinkedList(n int) *List {
	lst := &List{}

	for i := n; i > 0; i-- {
		lst.AddHead(i)
	}

	return lst
}


func main()  {

	lst := NewLinkedList(0)
	//lst.AddLoopToHead()
	//lst.AddLoopToRandomEl()

	lst.HasLoop()

	lst.Print()
}
