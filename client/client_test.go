package client

import (
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	for i := 0; i < 10; i++ {
		go InitClient()
	}

	time.Sleep(50 * time.Second)
}
