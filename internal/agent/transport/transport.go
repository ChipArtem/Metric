package transport

import (
	"fmt"
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
		hostURL: url}
}

func (c HTTPclient) SendMetric(m models.Metric) error {
	url := fmt.Sprintf("%s/update/%s/%s/%s", c.hostURL, m.Type, m.Name, m.Value)
	fmt.Println("url ", url)
	response, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		panic(fmt.Sprintf("SendMetric: %s", err))
		return err
	}

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("SendMetric: %v", err))
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	response.Body.Close()
	return nil
}
