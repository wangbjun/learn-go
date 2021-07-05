package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	itemCount := 0
	go func() {
		for {
			itemCount++
			item := <-out
			log.Printf("Got item #%d: %v\n", itemCount, item)
		}
	}()

	return out
}
