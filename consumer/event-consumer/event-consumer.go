package event_consumer

import (
	"log"
	"sync"
	even "teleg_bot/events"
	"time"
)

type Consumer struct {
	fetcher   even.Fetcher
	processor even.Processor
	batchSize int
}

func New(fetcher even.Fetcher, processor even.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)
			continue
		}
	}
}

// Goroutine паралелльная обработка евентов
func (c *Consumer) handleEvents(events []even.Event) error {
	wg := sync.WaitGroup{}
	wg.Add(len(events))
	for _, event := range events {
		go func(event even.Event) {
			defer wg.Done()

			log.Printf("got new events: %s", event.Text)

			if err := c.processor.Process(event); err != nil {
				log.Printf("can't handle event: %s", err.Error())

			}
		}(event)

	}
	wg.Wait()
	return nil
}
