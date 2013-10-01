package lab4

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestUnsafeStack(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	unsafeStack := new(UnsafeStack)
	fmt.Println("Unsafe Stack Test")
	testConcurrentStackAccess(unsafeStack)
}

func TestSafeStack(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	safeStack := new(SafeStack)
	fmt.Println("Safe Stack Test")
	testConcurrentStackAccess(safeStack)
}

func TestCspStack(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cspStack := NewCspStack()
	fmt.Println("CSP Stack Test")
	testConcurrentStackAccess(cspStack)
}

func TestSliceStack(t *testing.T) {
	sliceStack := NewSliceStack()
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Slice Stack Test")
	testConcurrentStackAccess(sliceStack)
}

func TestOpsUnsafeStack(t *testing.T) {
	fmt.Println("Test operations UnsafeStack")
	unsafeStack := new(UnsafeStack)
	testStackOperations(unsafeStack, t)
}

func TestOpsSafeStack(t *testing.T) {
	fmt.Println("Test operations SafeStack")
	safeStack := new(SafeStack)
	testStackOperations(safeStack, t)
}

func TestOpsCspStack(t *testing.T) {
	fmt.Println("Test operations CspStack")
	cspStack := NewCspStack()
	testStackOperations(cspStack, t)
}

func TestOpsSliceStack(t *testing.T) {
	fmt.Println("Test operations SliceStack")
	sliceStack := NewSliceStack()
	testStackOperations(sliceStack, t)
}

func TestOpsAllStacks(t *testing.T) {
	fmt.Println("Test operations all stacks")
	TestOpsUnsafeStack(t)
	TestOpsSafeStack(t)
	TestOpsCspStack(t)
	TestOpsSliceStack(t)
}

func BenchmarkSafeStack(b *testing.B) {
	safeStack := new(SafeStack)
	for i := 0; i < b.N; i++ {
		benchStackOperations(safeStack)
	}
}

func BenchmarkSliceStack(b *testing.B) {
	sliceStack := NewSliceStack()
	for i := 0; i < b.N; i++ {
		benchStackOperations(sliceStack)
	}
}

func BenchmarkCspStack(b *testing.B) {
	sliceStack := NewCspStack()
	for i := 0; i < b.N; i++ {
		benchStackOperations(sliceStack)
	}
}

const (
	NumberOfGoroutines = 4
	NumberOfOperations = 10
)

const (
	Len = iota
	Push
	Pop
)

func testConcurrentStackAccess(stack Stack) {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	wg.Add(NumberOfGoroutines)

	for i := 0; i < NumberOfGoroutines; i++ {
		go func(i int) {
			for j := 0; j < NumberOfOperations; j++ {
				op := rand.Intn(3)
				switch op {
				case Len:
					// fmt.Printf("#%d-%d: Len() was %d\n", i, j, stack.Len())

					stack.Len()
				case Push:
					// data := "Data" + strconv.Itoa(i) + strconv.Itoa(j)
					// fmt.Printf("#%d-%d: Push() with value %v\n", i, j, data)

					stack.Push("Data" + strconv.Itoa(i) + strconv.Itoa(j))
				case Pop:
					_ = stack.Pop()

					// fmt.Printf("#%d-%d: Pop() gave value %v\n", i, j, value)
				}
			}

			defer wg.Done()
		}(i)
	}

	wg.Wait()
}

func testStackOperations(stack Stack, t *testing.T) {
	if stack.Len() != 0 {
		t.Errorf("Initial Len() not 0")
	}

	stack.Push("Item1")
	if stack.Len() != 1 {
		t.Errorf("Len() not 1")
	}

	item1 := stack.Pop()
	if stack.Len() != 0 {
		t.Errorf("Len() not 0")
	}
	if item1 != "Item1" {
		t.Errorf("item1 not Item1, was %v")
	}

	stack.Push("Item2")
	stack.Push(3)
	stack.Push(4.0001)
	if stack.Len() != 3 {
		t.Errorf("Len() not 3")
	}

	item4 := stack.Pop()
	if stack.Len() != 2 {
		t.Errorf("Len() not 2")
	}
	if item4 != 4.0001 {
		t.Errorf("item4 not 4.0001, was %v", item4)
	}

	item3 := stack.Pop()
	if stack.Len() != 1 {
		t.Errorf("Len() not 1")
	}
	if item3 != 3 {
		t.Errorf("item3 not 3, was %v", item3)
	}

	item2 := stack.Pop()
	if stack.Len() != 0 {
		t.Errorf("Len() not 0")
	}
	if item2 != "Item2" {
		t.Errorf("item2 not Item2, was %v", item2)
	}

	item5 := stack.Pop()
	if item5 != nil {
		t.Errorf("item5 not nil, was %v", item5)
	}
	if stack.Len() != 0 {
		t.Errorf("Len() not 0")
	}

	// StackSlice realloc check
	for i := 0; i < 50; i++ {
		stack.Push(i)
	}
	if stack.Len() != 50 {
		t.Errorf("Len() not 50")
	}
	for j := 49; j >= 0; j-- {
		if x := stack.Pop(); x != j {
			t.Errorf("Pop: Want %d, got %d", j, x)
		}
	}
}

func benchStackOperations(stack Stack) {
	const nrOfOps = 10000

	for i := 0; i < nrOfOps; i++ {
		stack.Push(i)
	}

	for j := 0; j < nrOfOps; j++ {
		stack.Pop()
	}
}
