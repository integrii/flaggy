package flaggy

import "testing"

func TestDoubleParse(t *testing.T) {
	ResetParser()

	err := mainParser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	err = mainParser.Parse()
	if err == nil {
		t.Fatal(err)
	}
}
