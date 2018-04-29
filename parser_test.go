package flaggy

import "testing"

func TestDoubleParse(t *testing.T) {
	err := mainParser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	err = mainParser.Parse()
	if err == nil {
		t.Fatal(err)
	}
}
