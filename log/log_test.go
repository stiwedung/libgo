package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	SetLevel("debug")
	Debug("Debug")
	Debugf("Debugf")
	Info("Info")
	Infof("Infof")
	Warn("Warn")
	Warnf("Warnf")
	Error("Error")
	Errorf("Errorf")
	Fatal("Fatal")
	//Fatalf("Fatalf")
	Close()
}
