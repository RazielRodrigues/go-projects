package main

import "testing"

// ============================================================================
// BENCHMARKS - Comparing Stack vs Heap Allocation Performance
// ============================================================================

// Simple Value Operations
func BenchmarkStackSimpleValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunStackSimpleValue()
	}
}

func BenchmarkHeapSimpleValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunHeapSimpleValue()
	}
}

func BenchmarkStackPassByPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunStackPassByPointer()
	}
}

// Struct Operations
func BenchmarkStackStructAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunStackStructAllocation()
	}
}

func BenchmarkHeapStructAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunHeapStructAllocation()
	}
}

func BenchmarkStackStructModification(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunStackStructModification()
	}
}

func BenchmarkHeapStructModification(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunHeapStructModification()
	}
}

// Slice Operations
func BenchmarkStackSliceAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunStackSliceAllocation()
	}
}

func BenchmarkHeapSliceAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunHeapSliceAllocation()
	}
}

// ============================================================================
// HUGESTRUCT SPECIFIC BENCHMARKS - Demonstrating large struct allocations
// ============================================================================

func BenchmarkHugeStructStackAllocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = HugeStruct{
			Data:   [1024]byte{1, 2, 3},
			Values: [100]int{1, 2, 3, 4, 5},
			Flags:  [32]bool{true, false, true},
			Text:   "Direct stack allocation",
		}
	}
}

func BenchmarkHugeStructHeapAllocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = &HugeStruct{
			Data:   [1024]byte{1, 2, 3},
			Values: [100]int{1, 2, 3, 4, 5},
			Flags:  [32]bool{true, false, true},
			Text:   "Direct heap allocation",
		}
	}
}

func BenchmarkHugeStructStackCopy(b *testing.B) {
	b.ReportAllocs()
	original := HugeStruct{
		Data:   [1024]byte{1, 2, 3},
		Values: [100]int{1, 2, 3, 4, 5},
		Flags:  [32]bool{true, false, true},
		Text:   "Original struct",
	}

	for i := 0; i < b.N; i++ {
		copy := original // Copy by value (stack)
		copy.Values[0] = i
		_ = copy
	}
}

func BenchmarkHugeStructHeapCopy(b *testing.B) {
	b.ReportAllocs()
	original := HugeStruct{
		Data:   [1024]byte{1, 2, 3},
		Values: [100]int{1, 2, 3, 4, 5},
		Flags:  [32]bool{true, false, true},
		Text:   "Original struct",
	}

	for i := 0; i < b.N; i++ {
		copy := &original // Create pointer (potential heap escape)
		copy.Values[0] = i
		_ = copy
	}
}

func BenchmarkHugeStructPassByValue(b *testing.B) {
	b.ReportAllocs()
	s := HugeStruct{
		Data:   [1024]byte{1, 2, 3},
		Values: [100]int{1, 2, 3, 4, 5},
		Flags:  [32]bool{true, false, true},
		Text:   "Pass by value test",
	}

	for i := 0; i < b.N; i++ {
		_ = processHugeStructByValue(s)
	}
}

func BenchmarkHugeStructPassByPointer(b *testing.B) {
	b.ReportAllocs()
	s := HugeStruct{
		Data:   [1024]byte{1, 2, 3},
		Values: [100]int{1, 2, 3, 4, 5},
		Flags:  [32]bool{true, false, true},
		Text:   "Pass by pointer test",
	}

	for i := 0; i < b.N; i++ {
		_ = processHugeStructByPointer(&s)
	}
}

// Helper functions for the benchmarks
//
//go:noinline
func processHugeStructByValue(s HugeStruct) int {
	return int(s.Data[0]) + s.Values[0]
}

//go:noinline
func processHugeStructByPointer(s *HugeStruct) int {
	return int(s.Data[0]) + s.Values[0]
}

const creations = 20_000_000

func TestStackStructAllocation(t *testing.T) {
	for range creations {
		RunStackStructAllocation()
	}
}

func TestHeapStructAllocation(t *testing.T) {
	for range creations {
		RunHeapStructAllocation()
	}
}
