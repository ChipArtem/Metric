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
	url := c.hostURL + "/update/" + m.Type + "/" + m.Name + "/" + m.Value
	response, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		panic("SendMetric: err:_" + err.Error() + "_")
		return err
	}

	if response.StatusCode != http.StatusOK {
		panic("SendMetric status: _" + url + "_")
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	response.Body.Close()
	return nil
}
