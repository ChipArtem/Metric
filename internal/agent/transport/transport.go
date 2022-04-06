package transport

import (
	"fmt"
	"net/http"

	"github.com/ChipArtem/Metric/internal/models"
	"github.com/go-resty/resty/v2"
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

	r, err := resty.New().R().SetHeader("Content-Type", "text/plain").SetBody([]byte("")).Post(url)
	// response, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		panic("SendMetric: err:_" + err.Error() + "_")
		return err
	}

	if r.StatusCode() != http.StatusOK {
		panic("SendMetric status: _" + url + "_")
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	return nil
}
