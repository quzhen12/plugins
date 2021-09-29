package common

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ReadAsLine("file.go", func(line string, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(line)
	})
}
