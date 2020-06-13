package log

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	DisableLevel
)

type levelMeta struct {
	name         string
	aliasName    []string
	rawText      string
	colorfulText string
}

var levels = [...]levelMeta{
	{
		name:         "debug",
		rawText:      "[DBUG]",
		colorfulText: CyanText("[DBUG]"),
	},
	{
		name:         "info",
		rawText:      "[INFO]",
		colorfulText: BlueText("[INFO]"),
	},
	{
		name:         "warn",
		aliasName:    []string{"warning"},
		rawText:      "[WARN]",
		colorfulText: YellowText("[WARN]"),
	},
	{
		name:         "error",
		rawText:      "[ERRO]",
		colorfulText: RedText("[ERRO]"),
	},
	{
		name:         "fatal",
		rawText:      "[FTAL]",
		colorfulText: MagentaText("[FTAL]"),
	},
	{
		name:      "disable",
		aliasName: []string{"disabled"},
	},
}

func ParseLevel(name string) Level {
	for i, lvlMeta := range levels {
		if lvlMeta.name == name {
			return Level(i)
		}
		for _, alias := range lvlMeta.aliasName {
			if alias == name {
				return Level(i)
			}
		}
	}
	return DisableLevel
}

func (lvl Level) rawText() string {
	return levels[lvl].rawText
}

func (lvl Level) colorfulText() string {
	return levels[lvl].colorfulText
}
