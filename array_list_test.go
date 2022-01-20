package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArrayListTest struct {
	al  *ArrayList
	al2 *ArrayList
}

func ArrayListTestSetup(tb testing.TB) (func(tb testing.TB), ArrayListTest) {
	at := ArrayListTest{}
	at.al = NewArrayList(0)
	at.al2 = NewArrayList(0)

	return func(tb testing.TB) {
		tb.Log("ArrayListTestSetup teardown")
	}, at
}

func TestNewArrayList_NewWithSlice(t *testing.T) {
	s := []interface{}{1, 2, 3, 4, 5}
	al := NewArrayListWithSlice(s)

	assert.Equal(t, []interface{}{1, 2, 3, 4, 5}, al.Iterator())

}

func TestArrayList_Add(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)

	assert.Equal(t, 4, at.al.GetSize())
}

func TestArrayList_Insert(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)

	at.al.Insert(2, 5)

	expected := []interface{}{1, 2, 5, 3, 4}

	for i := 0; i < 4; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_Remove(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)

	at.al.Remove(3)

	expected := []interface{}{1, 2, 4}

	for i := 0; i < 3; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_AddAll(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)

	at.al2.Add(5)
	at.al2.Add(6)
	at.al2.Add(7)
	at.al2.Add(8)

	at.al.AddAll(2, at.al2)

	expected := []interface{}{1, 2, 5, 6, 7, 8, 3, 4}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RemoveAll_whenInOrder(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al2.Add(3)
	at.al2.Add(4)
	at.al2.Add(5)
	at.al2.Add(6)

	at.al.RemoveAll(at.al2)

	expected := []interface{}{1, 2, 7, 8}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RemoveAll_whenDisorder(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al2.Add(2)
	at.al2.Add(4)
	at.al2.Add(6)
	at.al2.Add(8)

	at.al.RemoveAll(at.al2)

	expected := []interface{}{1, 3, 5, 7}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RemoveAll_whenDuplicate(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al.Add(2)

	at.al2.Add(2)
	at.al2.Add(4)
	at.al2.Add(6)
	at.al2.Add(8)

	at.al.RemoveAll(at.al2)

	expected := []interface{}{1, 3, 5, 7}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RemoveAll_whenNil(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(nil)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al2.Add(2)
	at.al2.Add(4)
	at.al2.Add(6)
	at.al2.Add(8)

	at.al.RemoveAll(at.al2)

	expected := []interface{}{1, 5, 7}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RetailAll_whenAllin(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al2.Add(2)
	at.al2.Add(4)
	at.al2.Add(6)
	at.al2.Add(8)

	at.al.RetainAll(at.al2)

	expected := []interface{}{2, 4, 6, 8}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}

func TestArrayList_RetailAll_whenNotAllin(t *testing.T) {
	teardownTest, at := ArrayListTestSetup(t)
	defer teardownTest(t)

	at.al.Add(1)
	at.al.Add(2)
	at.al.Add(3)
	at.al.Add(4)
	at.al.Add(5)
	at.al.Add(6)
	at.al.Add(7)
	at.al.Add(8)

	at.al2.Add(2)
	at.al2.Add(4)
	at.al2.Add(6)
	at.al2.Add(8)
	at.al2.Add(9)

	at.al.RetainAll(at.al2)

	expected := []interface{}{2, 4, 6, 8}
	for i := 0; i < at.al.size; i++ {
		assert.Equal(t, expected[i], at.al.Get(i))
	}
}
