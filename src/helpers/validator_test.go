package helpers

import "testing"

func TestIsValidIdentifier(t *testing.T) {
	tests := []struct {
		id       string
		length   int
		expected bool
	}{
		{"0x224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c", 66, true},
		{"0x224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1", 66, false},
		{"0xG224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c", 66, false},
		{"224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c", 66, false},
		{"0x224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c", 65, false},
	}

	for _, test := range tests {
		result := IsValidIdentifier(test.id, test.length)
		if result != test.expected {
			t.Errorf("For ID %q and length %d, expected %t but got %t", test.id, test.length, test.expected, result)
		}
	}
}
