package logging

import "testing"

func TestWriteToFile(t *testing.T) {
	for i := 0; i < 3; i++ {
		WriteToFile()
	}
}
