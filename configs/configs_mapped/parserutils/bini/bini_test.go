package bini

import (
	"os"
	"strings"
	"testing"
)

func TestBini(t *testing.T) {
	data := Dump("mbases.ini")
	os.WriteFile("output.txt", []byte(strings.Join(data, "\n")), 0644)
}
