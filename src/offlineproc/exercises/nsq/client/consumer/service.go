package consumer

import (
	"log"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

const DefaultNSQDLookupHost = "127.0.0.1:4161"

// Service is the consumer service that will process nsq tasks
type Service struct {
	host   string
	config *nsq.Config

	mu        sync.RWMutex
	topics    []string
	consumers []*nsq.Consumer

	path string
}

// New returns an instance of the Service
func New(host, path string, config *nsq.Config) *Service {
	return &Service{
		host:   host,
		config: config,
		path:   path,
	}
}

func (s *Service) Open() error {
	// Add my consumers
	if err := s.addConsumer(pingTopic, monitorChannel, newPingHandler(s)); err != nil {
		return err
	}
	return nil
}

func (s *Service) addConsumer(topic, channel string, handler nsq.HandlerFunc) error {
	log.Printf("adding consumer topic:%q channel:%q", topic, channel)

	q, _ := nsq.NewConsumer(topic, channel, s.config)
	q.AddHandler(handler)
	if err := q.ConnectToNSQLookupd(s.host); err != nil {
		return err
	}
	s.mu.Lock()
	s.consumers = append(s.consumers, q)
	s.mu.Unlock()
	return nil

}

func (s *Service) Close() error {
	var wg sync.WaitGroup

	s.mu.RLock()
	wg.Add(len(s.consumers))
	for _, consumer := range s.consumers {
		go func(c *nsq.Consumer) {
			c.Stop()
			<-c.StopChan
			wg.Done()
		}(consumer)
	}
	wg.Wait()
	s.mu.RUnlock()
	return nil
}
