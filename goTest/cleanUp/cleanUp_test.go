package cleanUp

import (
	"fmt"
	"os"
	"testing"
)

// createFile is helper function that is called from various tests
func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// do the test, no needs to worry 'bout resource cleanup
	fmt.Println(fName)
}
