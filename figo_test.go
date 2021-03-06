package figo

import "testing"

func TestNew(t *testing.T) {
	_ = New()
}

func Test_CanPush(t *testing.T) {
	q := New()
	q.Push("Nada")
}

func Test_CanPop(t *testing.T) {
	q := New()
	_ = q.Pop()
}

func Test_CanIsEmpty(t *testing.T) {
	q := New()
	_ = q.IsEmpty()
}

func Test_CanLen(t *testing.T) {
	q := New()
	_ = q.Len()
}

func Test_CanAddOneItem(t *testing.T) {
	q := New()
	if q.Len() != 0 {
		t.Error("Didn't get the right size")
	}

	q.Push("Nada")
	if q.Len() != 1 {
		t.Error("Should have increased in size")
	}

	val := q.Pop()
	if val != "Nada" {
		t.Error("Did not get the correct value")
	}

	if q.Len() != 0 {
		t.Error("Should have decreased in size")
	}
}

func Test_CanAddMultipleItems(t *testing.T) {
	q := New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	if q.IsEmpty() {
		t.Error("The queue should not be empty")
	}

	for i := 0; !q.IsEmpty(); i++ {
		v := q.Pop()
		if v != i {
			t.Error("Did not get the correct value")
		}
	}
}

func Test_CanIterateEmptyList(t *testing.T) {
	q := New()

	i := 0
	for elem := q.Front(); q.Next(elem) != nil; elem = q.Next(elem) {
	}

	if i != 0 {
		t.Error("Should have been empty")
		return
	}
}

func Test_CanIterateItems(t *testing.T) {
	q := New()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	i := 0
	for elem := q.Front(); q.Next(elem) != nil; elem = q.Next(elem) {
		if elem.Value.(int) != i {
			t.Error("Did not match the correct order")
			return
		}

		i++
	}

	if i != 9 {
		t.Error("Did not iterate the whole list")
		return
	}
}

func TestAsync(t *testing.T) {
	_ = NewAsync()
}

func TestAsync_CanPush(t *testing.T) {
	q := NewAsync()
	q.Push("Nada")
}

func TestAsync_CanPop(t *testing.T) {
	q := NewAsync()
	_ = q.Pop()
}

func TestAsync_CanIsEmpty(t *testing.T) {
	q := NewAsync()
	_ = q.IsEmpty()
}

func TestAsync_CanLen(t *testing.T) {
	q := NewAsync()
	_ = q.Len()
}

func TestAsync_CanAddOneItem(t *testing.T) {
	q := NewAsync()
	if q.Len() != 0 {
		t.Error("Didn't get the right size")
	}

	q.Push("Nada")
	if q.Len() != 1 {
		t.Error("Should have increased in size")
	}

	val := q.Pop()
	if val != "Nada" {
		t.Error("Did not get the correct value")
	}

	if q.Len() != 0 {
		t.Error("Should have decreased in size")
	}
}

func TestAsync_CanAddMultipleItems(t *testing.T) {
	q := NewAsync()

	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	if q.IsEmpty() {
		t.Error("The queue should not be empty")
	}

	for i := 0; !q.IsEmpty(); i++ {
		v := q.Pop()
		if v != i {
			t.Error("Did not get the correct value")
		}
	}
}
