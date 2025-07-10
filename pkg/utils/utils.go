package utils

import (
	"fmt"
	"io"
)

func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Printf("Warning: error closing resource: %v\n", err)
	}
}
