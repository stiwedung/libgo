package log

import "testing"

func TestColor(t *testing.T) {
	t.Log(RedText("Red"))
	t.Log(GreenText("Green"))
	t.Log(YellowText("Yellow"))
	t.Log(BlueText("Blue"))
	t.Log(MagentaText("Magenta"))
	t.Log(CyanText("Cyan"))
}
