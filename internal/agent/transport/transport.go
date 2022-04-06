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

	response, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		return fmt.Errorf("SendMetric: %s", err)
	}
	response.Body.Close()
	return nil
}
