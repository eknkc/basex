package basex

import "testing"

func Test_AlphabetTooShort(t *testing.T) {
	for _, alphabet := range []string{"", "0", "α"} {
		_, err := NewEncoding(alphabet)
		if err == nil {
			t.Fatalf("NewEncoding(%q) returned nil error", alphabet)
		}
		expect(err.Error(), "Alphabet must contain at least two characters.", t)
	}
}

func Test_MinimalAlphabet(t *testing.T) {
	enc, err := NewEncoding("01")
	if err != nil {
		t.Fatalf("NewEncoding(\"01\"): %v", err)
	}

	if got := enc.Encode([]byte{0}); got != "0" {
		t.Fatalf("Encode([]byte{0}) = %q, want %q", got, "0")
	}
	if got := enc.Encode([]byte{1}); got != "1" {
		t.Fatalf("Encode([]byte{1}) = %q, want %q", got, "1")
	}
	if got := enc.Encode([]byte{0, 1}); got != "01" {
		t.Fatalf("Encode([]byte{0, 1}) = %q, want %q", got, "01")
	}

	dec, err := enc.Decode("01")
	if err != nil {
		t.Fatalf("Decode(\"01\"): %v", err)
	}
	if len(dec) != 2 || dec[0] != 0 || dec[1] != 1 {
		t.Fatalf("Decode(\"01\") = %v, want [0 1]", dec)
	}
}
