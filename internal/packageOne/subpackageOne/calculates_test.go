package subpackageone

import "testing"

func TestMax(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	actual := Max(input)
	expected := 5

	if actual != expected {
		t.FailNow()
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxEmptySlice(t *testing.T) {
	input := []int{}
	actual := Max(input)
	expected := -1

	if actual != expected {
		t.Fail()
		t.Logf("Expected %d, got %d", expected, actual)
	}
}

type TestCase struct {
	input    []int
	expected int
}

func TestMaxTD(t *testing.T) {
	cases := []TestCase{
		TestCase{
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		TestCase{
			input:    []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		TestCase{
			input:    []int{0},
			expected: 0,
		},
	}

	for _, c := range cases {
		actual := Max(c.input)
		expected := c.expected

		if actual != expected {
			t.Logf("Expected %d, got %d", expected, actual)
			// t.Fail()
			t.FailNow()
		}
	}
}

func TestMax2(t *testing.T) {
	cases := []TestCase{
		TestCase{
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		TestCase{
			input:    []int{-1, -2, -3, -4, -5},
			expected: -1,
		},
		TestCase{
			input:    []int{0},
			expected: 0,
		},
	}

	for _, c := range cases {
		actual := Max(c.input)
		expected := c.expected

		if actual != expected {
			t.Fatalf("Expected %d, got %d", expected, actual)
		}
	}
}

func TestMaxEmptySlice2(t *testing.T) {
	input := []int{}
	actual := Max(input)
	expected := -1

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual) // == t.Fail + t.Logf
	}
}

/*
* Error()  = Fail()    + Log()
* Errorf() = Fail()    + Logf()
* Fatal()  = FailNow() + Log()
* Fatalf() = FailNow() + Logf()
 */
