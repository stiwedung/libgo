package log

import "testing"

func TestLevel(t *testing.T) {
	for lvl := DisableLevel; lvl <= FatalLevel; lvl++ {
		t.Log(lvl.rawText(), lvl.colorfulText())
	}
}
