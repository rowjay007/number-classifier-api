package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func FetchFunFact(n int) string {
	client := resty.New()
	resp, err := client.R().Get(fmt.Sprintf("http://numbersapi.com/%d/math", n))

	if err != nil {
		return "Fun fact not available"
	}

	return resp.String()
}
