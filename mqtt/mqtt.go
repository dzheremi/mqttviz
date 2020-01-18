package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Settings - exported MQTT settings
var Settings Configuration

// Client - exported MQTT Client
var Client MQTT.Client

// Setup - establish MQTT connection
func Setup() {
	opts := MQTT.NewClientOptions().AddBroker(Settings.Server)
	opts.SetClientID("mqttviz")
	Client = MQTT.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic("--- Could not connect to MQTT broker ---")
	}
}
