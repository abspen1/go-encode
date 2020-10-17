package output

import "testing"

func TestPrintEncodedLine(t *testing.T) {
	last := make([]string, 11)
	last[0] = "i"
	last[1] = "p"
	last[2] = "s"
	last[3] = "s"
	last[4] = "M"
	last[5] = "p"
	last[6] = "i"
	last[7] = "s"
	last[8] = "s"
	last[9] = "i"
	last[10] = "i"
	got := PrintEncodedLine(last, 11)
	if got != "1 i 1 p 2 s 1 M 1 p 1 i 2 s 2 i" {
		t.Errorf("PrintEncodedLine = %s; want %s", got, "1 i 1 p 2 s 1 M 1 p 1 i 2 s 2 i")
	}
}
