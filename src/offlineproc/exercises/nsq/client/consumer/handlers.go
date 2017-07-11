package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	nsq "github.com/bitly/go-nsq"
)

const (
	pingTopic      = "ping"
	monitorChannel = "monitor"
)

func newPingHandler(s *Service) nsq.HandlerFunc {
	return nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("processing message id %s for channel: %s, topic: %s", message.ID, monitorChannel, pingTopic)

		// Get the website we want to ping
		website := string(message.Body)

		filename := fmt.Sprintf("%d-%s.json", message.Timestamp, message.ID)
		filename = path.Join(s.path, filename)
		// Check to see if we processed this already
		if _, err := os.Stat(filename); err == nil {
			// we already processed this, nothing to do
			log.Println("previously processed, exiting")
		}

		log.Printf("attempting to ping website %s", website)
		// start a timer for this request
		begin := time.Now()

		// Retrieve the site
		resp, err := http.Get(website)
		if err != nil {
			log.Printf("failed to retrieved site: website: %q", website)
			return recordPingError(filename, website, err)
		}
		duration := time.Since(begin)

		// NOTE: if this fails to record, we could have a disk full, etc.
		// but we want the message to replay so we want the error to bubble up to
		// nsq and leave the message in the que until it can successfully complete
		log.Printf("successfull retrieved site: website: %q, statusCode: %d, duration: %s", website, resp.StatusCode, duration.String())
		return recordPingSuccess(filename, website, resp.StatusCode, duration)
	})
}

func recordPingSuccess(filename, website string, statusCode int, duration time.Duration) error {
	data := struct {
		Website    string `json:"website"`
		StatusCode int    `json:"statusCode"`
		Duration   string `json:"duration"`
	}{
		Website:    website,
		StatusCode: statusCode,
		Duration:   duration.String(),
	}
	return writePingFile(filename, &data)
}

func recordPingError(filename, website string, err error) error {
	data := struct {
		Website string `json:"website"`
		Err     string `json:"error"`
	}{
		Website: website,
		Err:     err.Error(),
	}
	return writePingFile(filename, &data)
}

func writePingFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}

	return nil
}
