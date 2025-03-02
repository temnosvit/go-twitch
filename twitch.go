package twitch

import (
	"github.com/temnosvit/go-twitch/api"
	"github.com/temnosvit/go-twitch/irc"
	"github.com/temnosvit/go-twitch/pubsub"
)

// API provides tools for developing integrations with Twitch.
func API(clientID string, opts ...api.ClientOption) *api.Client {
	return api.New(clientID, opts...)
}

// IRC is the Twitch interface for chat functionality.
func IRC() *irc.Client {
	return irc.New()
}

// PubSub enables you to subscribe to a topic for updates.
func PubSub() *pubsub.Client {
	return pubsub.New()
}
