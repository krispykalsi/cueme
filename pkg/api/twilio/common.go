package twilio

import (
	"cueme/pkg/api/failure"
	"log"
	"net/http"
	"strconv"
)

func executeRequest(c *http.Client, r *http.Request) error {
	res, err := c.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	if strconv.Itoa(res.StatusCode)[0] == '2' {
		return nil
	} else {
		return failure.FromTwilio(res.Status)
	}
}
