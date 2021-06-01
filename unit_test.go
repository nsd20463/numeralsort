package numeralsort

import (
	"strconv"
	"testing"
)

func TestLess(t *testing.T) {
	// test cases pairs where cases[2*i] < cases[2*i+1]
	cases := []string{
		"a", "b",
		"1", "a",
		"", "a",
		"", "1",
		"0", "1",
		"0", "00",
		"1", "2",
		"1", "01",
		"1", "001",
		"001", "010",
		"1", "10",
		"2", "10",
		"file0.txt", "file1.txt",
		"file1.txt", "file2.txt",
		"file1.txt", "file10.txt",
		"file2.txt", "file10.txt",
		"file03.txt", "file10.txt",
		"file123.txt", "file124.txt",
		"file123.txt", "file143.txt",
		"file123.txt", "file423.txt",
		"file123.txt", "file4123.txt",
		"file123.txt", "file0123.txt",
		"a1b2c", "a1b2z",
		"a1b2c", "a1b3c",
		"a1b2c", "a1z2c",
		"a1b2c", "a2b2c",
		"a1b2c", "z1b2c",
		"a1b2c", "a11b2c",
		"a1b22c", "a11b2c",
		"a1b2c", "a01b2c",
		"a1b2c", "a1b22c",
		"a1b2c", "a1b02c",
		"1b2", "1b3",
		"1b2", "1z2",
		"1b2", "2b2",
		"v1.2.0", "v1.10.0",
		"1.10", "10.1",
	}

	for i := 0; i < len(cases); i += 2 {
		x := cases[i]
		y := cases[i+1]
		t.Logf("Less(%q, %q) = %v", x, y, Less(x, y))
		if !Less(x, y) {
			t.Errorf("Less(%q, %q) != true", x, y)
		}
		// and the reverse should not be true
		t.Logf("Less(%q, %q) = %v", y, x, Less(y, x))
		if Less(y, x) {
			t.Errorf("Less(%q, %q) != false", y, x)
		}
	}
}

func TestEqual(t *testing.T) {
	// test cases for equality (Less(x,x) should be false)
	cases := []string{
		"",
		"a",
		"0",
		"00",
		"1",
		"10",
		"010",
		"123",
	}

	for _, x := range cases {
		if Less(x, x) {
			t.Errorf("Less(%q, %q) != false", x, x)
		}
	}
}

func TestSort(t *testing.T) {
	result := []string{"100", "10", "1000", "1", "0"}
	Strings(result)
	t.Logf("%q", result)
	prev := -1
	for _, s := range result {
		n, err := strconv.Atoi(s)
		if err != nil {
			t.Errorf("unexpected error %s", err)
			break
		}
		if prev >= n {
			t.Errorf("out of order. %d came before %d", prev, n)
		}
		prev = n
	}
}
