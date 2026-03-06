package main

import (
	"fmt"
)

type HugeStruct struct {
	Data   [1024]byte // 1KB of data
	Values [100]int   // Array of integers
	Flags  [32]bool   // Array of booleans
	Text   string     // String field
}

func main() {
	fmt.Println("Stack vs Heap Memory Allocation Benchmarks")
	fmt.Println("Run 'go test -bench=.' to see performance differences")
	fmt.Println()
	fmt.Println("Stack allocations: Fast, limited size, automatic cleanup")
	fmt.Println("Heap allocations: Slower, unlimited size, garbage collected")
}

// ============================================================================
// SIMPLE VALUE OPERATIONS - Demonstrating basic stack vs heap behavior
// ============================================================================

// StackSimpleValue demonstrates allocation on the stack
// The integer is allocated on the stack and returned by value
//
//go:noinline
func StackSimpleValue() int {
	value := 42
	result := value * 2
	return result // returned by value, stays on stack
}

// HeapSimpleValue demonstrates escape to heap
// Returning a pointer forces the value to escape to the heap
//
//go:noinline
func HeapSimpleValue() *int {
	value := 42
	result := value * 2
	return &result // escapes to heap because we return a pointer
}

// StackPassByPointer demonstrates passing pointers without heap allocation
// The pointer is passed down the stack, but the original value stays on stack
//
//go:noinline
func StackPassByPointer(value *int) int {
	result := *value * 2
	return result // returned by value
}

// ============================================================================
// STRUCT OPERATIONS - Demonstrating larger data structure allocations
// ============================================================================

// StackStructAllocation creates a struct on the stack
// Large structs on stack are fast but consume stack space
//
//go:noinline
func StackStructAllocation() HugeStruct {
	return HugeStruct{
		Data:   [1024]byte{1, 2, 3},     // Initialize with some data
		Values: [100]int{1, 2, 3, 4, 5}, // Initialize first few values
		Flags:  [32]bool{true, false, true},
		Text:   "Stack allocated struct",
	}
}

// HeapStructAllocation creates a struct on the heap
// Returning a pointer forces heap allocation and involves GC overhead
//
//go:noinline
func HeapStructAllocation() *HugeStruct {
	return &HugeStruct{
		Data:   [1024]byte{1, 2, 3},     // Initialize with some data
		Values: [100]int{1, 2, 3, 4, 5}, // Initialize first few values
		Flags:  [32]bool{true, false, true},
		Text:   "Heap allocated struct",
	}
}

// StackStructModification modifies a struct without heap allocation
// Takes a struct by value, modifies it, and returns by value
//
//go:noinline
func StackStructModification(s HugeStruct) HugeStruct {
	s.Values[0] = s.Values[0] * 2
	s.Text = "Modified on Stack"
	s.Flags[0] = !s.Flags[0]
	return s
}

// HeapStructModification forces heap allocation through pointer return
// Similar operation but with heap allocation
//
//go:noinline
func HeapStructModification(s HugeStruct) *HugeStruct {
	s.Values[0] = s.Values[0] * 2
	s.Text = "Modified on Heap"
	s.Flags[0] = !s.Flags[0]
	return &s // forces heap allocation
}

// ============================================================================
// SLICE OPERATIONS - Demonstrating dynamic memory allocation patterns
// ============================================================================

// StackSliceAllocation creates a small slice that may stay on stack
// Small, fixed-size slices might be stack-allocated
//
//go:noinline
func StackSliceAllocation() [10]int {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i * i
	}
	return arr
}

// HeapSliceAllocation creates a slice on the heap
// Dynamic slices typically allocate on the heap
//
//go:noinline
func HeapSliceAllocation(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i * i
	}
	return slice
}

// ============================================================================
// WRAPPER FUNCTIONS FOR BENCHMARKING
// ============================================================================

func RunStackSimpleValue() {
	_ = StackSimpleValue()
}

func RunHeapSimpleValue() {
	_ = HeapSimpleValue()
}

func RunStackPassByPointer() {
	value := 42
	_ = StackPassByPointer(&value)
}

func RunStackStructAllocation() {
	_ = StackStructAllocation()
}

func RunHeapStructAllocation() {
	_ = HeapStructAllocation()
}

func RunStackStructModification() {
	s := HugeStruct{
		Values: [100]int{10, 20, 30},
		Text:   "Original",
		Flags:  [32]bool{true},
	}
	_ = StackStructModification(s)
}

func RunHeapStructModification() {
	s := HugeStruct{
		Values: [100]int{10, 20, 30},
		Text:   "Original",
		Flags:  [32]bool{true},
	}
	_ = HeapStructModification(s)
}

func RunStackSliceAllocation() {
	_ = StackSliceAllocation()
}

func RunHeapSliceAllocation() {
	_ = HeapSliceAllocation(100)
}
