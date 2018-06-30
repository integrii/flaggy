package flaggy

import "testing"

func TestDoubleParse(t *testing.T) {
	ResetParser()

	err := DefaultParser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	err = DefaultParser.Parse()
	if err == nil {
		t.Fatal(err)
	}
}
