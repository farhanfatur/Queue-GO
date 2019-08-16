package queue

import (
	"fmt"
	"testing"
)

var testValues = []interface{}{
	"lorem",
	"ipsum",
	1,
	2,
	3,
	"jack",
	"jill",
	"felix",
	"donking",
}

// TestPush validate evict old item policy
func TestEvictPolicy(t *testing.T) {
	q := New()

	for i, v := range testValues {
		q.Push(v)

		t.Log("current: ", q.Keys())

		// validate
		// item existence
		if !q.Contains(v) {
			t.Errorf("policy: newly inserted %v must be exists", v)
		}

		if i < 10 && q.Len() != (i+1) {
			t.Errorf("expected length %d but actual: %d", i+1, q.Len())
		} else if i >= 10 && q.Len() != 10 {
			t.Errorf("expexted length: %d but actual: %d", i+1, q.Len())
		}
	}
}

// TestPop validate pop item policy
func TestPop(t *testing.T) {
	q := New()

	for _, v := range testValues {
		q.Push(v)
	}

	for q.Len() > 0 {
		t.Log("current: ", q.Keys())

		v := q.Pop()

		// validate
		expect := testValues[len(testValues)-(q.Len()+1)]
		if v != expect {
			fmt.Printf("expected %v but recevied %v", expect, v)
		}
	}

}
