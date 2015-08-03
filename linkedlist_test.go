// Package algostudy provides algostudy
package algostudy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll.head != nil || ll.tail != nil || ll.length != 0 {
		t.Errorf("Creation of new linked list failed!")
	}
}

func TestInsert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	ll := NewLinkedList()

	for i := 0; i < rand.Intn(50)+1; i++ {
		msg := fmt.Sprintf("test%d", i)
		err := ll.Insert(msg, i)
		if err != nil {
			t.Errorf("Error inserting node in list: %v\n", err)
		}
		if ll.length != i+1 {
			t.Errorf("Length not incremented on isert: %d", ll.length)
		}
	}

	err := ll.Insert("center", ll.length/2)
	if err != nil {
		t.Errorf("Failed to insert into middle of list: %v", err)
	}

	err = ll.Insert("error", 1000)
	if err == nil {
		t.Errorf("Failed to catch overflow!")
	}

}
