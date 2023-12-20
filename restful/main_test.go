package restful

import (
	"testing"
	"time"
)

func TestSetup(t *testing.T) {

	timeout := time.After(3 * time.Second)
	_panic := make(chan bool)
	go func() {
		r := App(true)
		// Listen and Server in 0.0.0.0:8080
		err := r.Run(":8080")
		t.Error(err)
		_panic <- true
	}()
	select {
	case <-timeout:
		println("success")
	case <-_panic:
	}

}
