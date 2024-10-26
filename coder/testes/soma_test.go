package soma

import "testing"

func TestSoma(t *testing.T) {
	/* 	t.Parallel()
	   	t.Skip()
	   	t.Fail() */
	dataset := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, data := range dataset {
		result := Soma(data.a, data.b)
		if result != data.expected {
			t.Errorf("Soma(%d, %d) = %d, want %d", data.a, data.b, result, data.expected)
		}
	}

}
