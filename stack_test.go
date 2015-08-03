// Package algostudy provides algostudy
package algostudy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.Head != nil || stack.Size != 0 {
		t.Errorf("Failed to initiaize new stack!")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()

	stack.Push("Test")
	if stack.Head == nil {
		t.Errorf("Failed to add element to stack!")
	}
	if stack.Head.Data != "Test" {
		t.Errorf("Incorrect element added to stack!")
	}
	if stack.Size != 1 {
		t.Errorf("Failed to increment size of stack!")
	}
	stack.Push("Test2")
	if stack.Head.Data != "Test2" {
		t.Errorf("Incorrect element added to stack!")
	}
	if stack.Head.Next == nil {
		t.Errorf("Failed to update pointer to next element!")
	}
	if stack.Size != 2 {
		t.Errorf("Failed to increment size of stack!")
	}
}

func TestPop(t *testing.T) {
	var nums []int
	var test string

	stack := NewStack()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(50); i++ {
		nums = append(nums, rand.Intn(50))
	}
	for _, num := range nums {
		test = fmt.Sprintf("Test%d", num)
		stack.Push(test)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		data, err := stack.Pop()
		if err != nil {
			t.Errorf("Underflowed stack!")
		}
		test = fmt.Sprintf("Test%d", nums[i])
		if data != test {
			t.Errorf("Failed to retrive data in correct order!")
		}
	}
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Failed to error on stack underflow!")
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := NewStack()
		stack.Push("test")
		_, _ = stack.Pop()
	}
}
