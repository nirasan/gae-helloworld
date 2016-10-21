package hello

import (
    "net/http"
    "github.com/nirasan/gae-helloworld"
)

func init() {
    http.Handle("/", gae_helloworld.CreateHandler())
}
