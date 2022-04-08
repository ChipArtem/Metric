package transport

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChipArtem/Metric/internal/models"
)

type HTTPclient struct {
	client  *http.Client
	hostURL string
}

func NewHTTPClient(url string, client *http.Client) HTTPclient {
	return HTTPclient{
		client:  client,
		hostURL: url,
	}
}

func (c HTTPclient) SendMetric(m models.Metric) error {
	client := &http.Client{}
	url := c.hostURL + "/update/" + m.Type + "/" + m.Name + "/" + m.Value
	r, err := client.Post(url, "text/plain", nil)
	if err != nil {
		log.Printf("SendMetric error %v", http.StatusOK)
		return err
	}
	if r.StatusCode != http.StatusOK {
		log.Printf("SendMetric status %v", http.StatusOK)
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	r.Body.Close()
	return nil
}
