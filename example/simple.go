package main

import (
	"fmt"
	"strings"

	"github.com/hysios/broker/session"
	"github.com/hysios/log"
)

var _ session.Handler = (*Handler)(nil)

// Handler implements mqtt.Handler interface
type Handler struct {
}

// New creates new Event entity
func New() *Handler {
	return &Handler{}
}

// AuthConnect is called on device connection,
// prior forwarding to the MQTT broker
func (h *Handler) AuthConnect(c *session.Client) error {
	log.Info(fmt.Sprintf("AuthConnect() - clientID: %s, username: %s, password: %s, client_CN: %s", c.ID, c.Username, string(c.Password), c.Cert.Subject.CommonName))
	return nil
}

// AuthPublish is called on device publish,
// prior forwarding to the MQTT broker
func (h *Handler) AuthPublish(c *session.Client, topic *string, payload *[]byte) error {
	log.Info(fmt.Sprintf("AuthPublish() - clientID: %s, topic: %s, payload: %s", c.ID, *topic, string(*payload)))
	return nil
}

// AuthSubscribe is called on device publish,
// prior forwarding to the MQTT broker
func (h *Handler) AuthSubscribe(c *session.Client, topics *[]string) error {
	log.Info(fmt.Sprintf("AuthSubscribe() - clientID: %s, topics: %s", c.ID, strings.Join(*topics, ",")))
	return nil
}

// Connect - after client successfully connected
func (h *Handler) Connect(c *session.Client) {
	log.Info(fmt.Sprintf("Connect() - username: %s, clientID: %s", c.Username, c.ID))
}

// Publish - after client successfully published
func (h *Handler) Publish(c *session.Client, topic *string, payload *[]byte) {
	log.Info(fmt.Sprintf("Publish() - username: %s, clientID: %s, topic: %s, payload: %s", c.Username, c.ID, *topic, string(*payload)))
}

// Subscribe - after client successfully subscribed
func (h *Handler) Subscribe(c *session.Client, topics *[]string) {
	log.Info(fmt.Sprintf("Subscribe() - username: %s, clientID: %s, topics: %s", c.Username, c.ID, strings.Join(*topics, ",")))
}

// Unsubscribe - after client unsubscribed
func (h *Handler) Unsubscribe(c *session.Client, topics *[]string) {
	log.Info(fmt.Sprintf("Unsubscribe() - username: %s, clientID: %s, topics: %s", c.Username, c.ID, strings.Join(*topics, ",")))
}

// Disconnect on conection lost
func (h *Handler) Disconnect(c *session.Client) {
	log.Info(fmt.Sprintf("Disconnect() - client with username: %s and ID: %s disconenected", c.Username, c.ID))
}
