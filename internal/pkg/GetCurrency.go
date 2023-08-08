package pkg

import (
	"context"
	"io"
	"net/http"
	"time"
)

var (
	apiMaxTime = 300
)

func GetCurrency(endpoint string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(apiMaxTime))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(res), nil

}
