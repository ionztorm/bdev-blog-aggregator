package utils

import (
	"fmt"
	"io"
)

func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		// Log the error or handle it as needed
		fmt.Printf("Warning: error closing resource: %v\n", err)
	}
}
