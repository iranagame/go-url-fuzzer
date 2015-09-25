package broadcaster

import (
	"log"

	"net/url"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/trustmaster/goflow"
)

// ResultBroadcaster receives found relative URLs and broadcasts to active output components.
type ResultBroadcaster struct {
	flow.Component

	FoundEntry <-chan messages.FoundEntry
	baseURL    url.URL
}

// NewResultBroadcaster creates new instance of result broadcaster.
func NewResultBroadcaster(configuration *configuration.Configuration) *ResultBroadcaster {
	return &ResultBroadcaster{
		baseURL: configuration.BaseURL(),
	}
}

// OnFoundEntry performs broadcasting.
func (r *ResultBroadcaster) OnFoundEntry(foundEntry messages.FoundEntry) {
	log.Println(foundEntry.String(r.baseURL))
}