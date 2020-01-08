package stmp_test

import (
	"github.com/acrazing/stmp-go/stmp"
	"log"
	"time"
)

func ExampleClient_DialKCP() {
	sc := stmp.NewClient(nil)
	sc.HandleConnected(func(header stmp.Header, message string) {
		log.Printf("stmp connected: %q.", message)
		time.Sleep(time.Second)
		// the connection will auto reconnect by default
		sc.Close(stmp.StatusNetworkError, "test retry")
	})
	sc.HandleDisconnected(func(reason stmp.StatusError, willRetry bool, retryCount int, retryWait time.Duration) {
		log.Printf("stmp disconnected, reason: %q, will retry: %t the %d time in %d seconds.", reason, willRetry, retryCount, retryWait)
	})
	sc.DialTCP("127.0.0.1:9991")
}
