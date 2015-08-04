// Package algostudy provides algostudy
package algostudy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func testSetupList(t *testing.T) LinkedList {
	rand.Seed(time.Now().UnixNano())

	ll := NewLinkedList()

	for i := 0; i < rand.Intn(50)+10; i++ {
		msg := fmt.Sprintf("test%d", i)
		err := ll.Insert(msg, i)
		if err != nil {
			t.Errorf("Error inserting node in list: %v\n", err)
		}
		if ll.length != i+1 {
			t.Errorf("Length not incremented on isert: %d", ll.length)
		}
	}
	return ll

}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll.head != nil || ll.tail != nil || ll.length != 0 {
		t.Errorf("Creation of new linked list failed!")
	}
}

func TestInsert(t *testing.T) {
	ll := testSetupList(t)

	err := ll.Insert("center", ll.length/2)
	if err != nil {
		t.Errorf("Failed to insert into middle of list: %v", err)
	}

	err = ll.Insert("error", 1000)
	if err == nil {
		t.Errorf("Failed to catch overflow!")
	}
	err = ll.Insert("error", -1)
	if err == nil {
		t.Errorf("Failed to error on negative position!")
	}

	err = ll.Insert("HEAD", 0)
	if err != nil {
		t.Errorf("Head insert failed!")
	}
}

func TestRemoveNodeByValue(t *testing.T) {
	ll := testSetupList(t)

	index := ll.length / 2

	value, err := ll.Get(index)
	if err != nil {
		t.Errorf("Failed to get test value!")
	}
	err = ll.RemoveNodeByValue(value)
	if err != nil {
		t.Errorf("Error on removing node! %v", err)
	}
	test, _ := ll.Get(index)
	if test == value {
		t.Errorf("remove failed! Value at position %d still equals original value!", index)
	}

	value = ll.GetHead()
	err = ll.RemoveNodeByValue(value)
	if err != nil {
		t.Errorf("Removal of head by value failed! %v", err)
	}
	test = ll.GetHead()
	if test == value {
		t.Errorf("Removal of head failed! Old value returned!")
	}

	err = ll.RemoveNodeByValue("bananna")
	if err == nil {
		t.Errorf("Failed to error on attempted removal of non-existant value!")
	}
}

func TestRemoveHead(t *testing.T) {
	ll := testSetupList(t)

	head := ll.GetHead()
	ll.RemoveHead()
	test := ll.GetHead()
	if test == head {
		t.Errorf("Removal of head failed!")
	}
}

func TestRemoveTail(t *testing.T) {
	ll := testSetupList(t)

	tail := ll.GetTail()
	ll.RemoveTail()
	test := ll.GetTail()
	if test == tail {
		t.Errorf("Removal of tail failed!")
	}
}

func TestRemoveByPosition(t *testing.T) {
	ll := testSetupList(t)

	err := ll.RemoveByPosition(-1)
	if err == nil {
		t.Errorf("failed to error on negative position")
	}

	val := ll.GetHead()
	err = ll.RemoveByPosition(0)
	if err != nil {
		t.Errorf("removal of tail by position failed! %v", err)
	}
	test := ll.GetHead()
	if test == val {
		t.Errorf("removal of head node failed! Old value remains!")
	}

	val = ll.GetTail()
	err = ll.RemoveByPosition(ll.length - 1)
	if err != nil {
		t.Errorf("removal of tail by position failed! %v", err)
	}
	test = ll.GetTail()
	if test == val {
		t.Errorf("removal of tail node failed! Old value remains!")
	}

	idx := ll.length / 2
	val, _ = ll.Get(idx)
	err = ll.RemoveByPosition(idx)
	if err != nil {
		t.Errorf("Failed to remove node from middle of list! %v", err)
	}
	test, _ = ll.Get(idx)
	if test == val {
		t.Errorf("removal of node from middle of list failed! Old value remains!")
	}

	err = ll.RemoveByPosition(ll.length + 5)
	if err == nil {
		t.Errorf("failed to error when end of list was hit!")
	}

	ll = NewLinkedList()
	err = ll.RemoveByPosition(0)
	if err == nil {
		t.Errorf("failed to error on removal of nodes from empty list!")
	}
}

func TestGet(t *testing.T) {
	ll := testSetupList(t)

	data, err := ll.Get(0)
	if err != nil {
		t.Errorf("Failed to get data at head")
	}
	if data != "test0" {
		t.Errorf("Wrong value returned from Get(0)!: %v", data)
	}

	data, err = ll.Get(ll.length / 2)
	if err != nil {
		t.Errorf("Failed to get data at tail! %v", err)
	}
	test := fmt.Sprintf("test%d", ll.length/2)
	if data != test {
		t.Errorf("Wrong value returned from tail!: %v", data)
	}

	data = ll.GetHead()
	if data != "test0" {
		t.Errorf("Got wrong value at Head of Linked List. Returned %s should have been test0", data)
	}

	test = fmt.Sprintf("test%d", ll.length-1)
	data = ll.GetTail()
	if data != test {
		t.Errorf("Got wrong value at Tail of Linked List. Returned %s should have been %s", data, test)
	}

	_, err = ll.Get(-1)
	if err == nil {
		t.Errorf("Failed to error on negative position")
	}

	_, err = ll.Get(ll.length)
	if err == nil {
		t.Errorf("Failed to error on list overflow")
	}

}

func BenchmarkInsertHead(b *testing.B) {
	ll := NewLinkedList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ll.Insert("test", 0)
	}
	b.StopTimer()
	var prev, cur *Node
	cur = ll.head
	for cur != nil {
		prev = cur
		cur = cur.Next
		prev.Next = nil
	}
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

func BenchmarkInsertTail(b *testing.B) {
	ll := NewLinkedList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Insert("test", ll.length)
	}
	b.StopTimer()
	var prev, cur *Node
	cur = ll.head
	for cur != nil {
		prev = cur
		cur = cur.Next
		prev.Next = nil
	}
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

func BenchmarkInsertRandom(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	ll := NewLinkedList()

	for i := 0; i < rand.Intn(50)+10; i++ {
		err := ll.Insert("test", i)
		if err != nil {
			fmt.Println(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		idx := rand.Intn(ll.length)
		b.StartTimer()
		_ = ll.Insert("test", idx)
	}
	b.StopTimer()
	var prev, cur *Node
	cur = ll.head
	for cur != nil {
		prev = cur
		cur = cur.Next
		prev.Next = nil
	}
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

func ExampleLinkedList() {
	ll := NewLinkedList()

	// Insert
	err := ll.Insert("test", 0)
	if err != nil {
		fmt.Errorf("ERROR: Could not insert item in list: %v", err)
	}
	ll.InsertHead("test2")
	ll.InsertTail("test3")
	err = ll.Insert("test4", 1)
	if err != nil {
		fmt.Errorf("ERROR: Could not insert item in list: %v", err)
	}

	//Remove By Value
	err = ll.RemoveNodeByValue("Test4")
	if err != nil {
		fmt.Errorf("ERROR: Failed to remove item from list: %v", err)
	}

	//Remove By Position
	err = ll.RemoveByPosition(1)
	if err != nil {
		fmt.Errorf("ERROR: Failed to remove item from list: %v", err)
	}

	//Get item
	data, _ := ll.Get(1)
	fmt.Printf("%s", data)
	//Get Tail
	data = ll.GetTail()
	fmt.Printf(" %s", data)
	//Get Head
	data = ll.GetHead()
	fmt.Printf(" %s", data)

	// Output: test test3 test2

}
