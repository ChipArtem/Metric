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
	// response, err := c.client.Post(url, "text/plain", nil)
	r, err := resty.New().R().
		SetBody([]byte("")).
		SetHeader("Content-Type", "text/plain").Post(url)
	if err != nil {
		panic("url_" + url + "__ __" + err.Error() + "_")
		return err
	}
	if r.StatusCode() != http.StatusOK {
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	// response.Body.Close()
	return nil
}
