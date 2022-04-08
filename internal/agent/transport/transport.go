package transport

import (
	"fmt"
	"net/http"
	"os"

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
	url := c.hostURL + "/update/" + m.Type + "/" + m.Name + "/" + m.Value
	fmt.Fprintf(os.Stderr, "transport url ", url, "\n")
	r, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "transport error ", err.Error(), "\n")
		return err
	}
	if r.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "transport status code: ", r.StatusCode, "\n")
		return fmt.Errorf("SendMetric status %v", http.StatusOK)
	}
	r.Body.Close()
	return nil
}
