package log

import "fmt"

// Color
const (
	Black = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Decorate
const (
	HightLight = 1 << iota
	Underline
	Bright
	Reverse

	Base = 0
)

const colorBase = "\033[%dm%s\033[0m"

var decorateBase = []string{colorBase}

func init() {
	for i := 1; i <= HightLight|Underline|Bright|Reverse; i++ {
		base := "\033[%s"
		if i&HightLight != 0 {
			base = fmt.Sprintf(base, "1;%s")
		}
		if i&Underline != 0 {
			base = fmt.Sprintf(base, "4;%s")
		}
		if i&Bright != 0 {
			base = fmt.Sprintf(base, "5;%s")
		}
		if i&Reverse != 0 {
			base = fmt.Sprintf(base, "7;%s")
		}
		base = fmt.Sprintf(base, "%dm%s\033[0m")
		decorateBase = append(decorateBase, base)
	}
}

func RichText(text string, color int, decorate int) string {
	base := colorBase
	if decorate >= 0 && decorate < len(decorateBase) {
		base = decorateBase[decorate]
	}
	return fmt.Sprintf(base, color, text)
}

func RedText(text string) string {
	return RichText(text, Red, Base)
}

func GreenText(text string) string {
	return RichText(text, Green, Base)
}

func YellowText(text string) string {
	return RichText(text, Yellow, Base)
}

func BlueText(text string) string {
	return RichText(text, Blue, Base)
}

func MagentaText(text string) string {
	return RichText(text, Magenta, Base)
}

func CyanText(text string) string {
	return RichText(text, Cyan, Base)
}
