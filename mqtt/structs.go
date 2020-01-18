package mqtt

import (
	"math/rand"
)

// Configuration - settings struct
type Configuration struct {
	Server          string
	WhiteLightGroup string
	RGBLightGroup   string
	WhiteLights     []string
	RGBLights       []string
}

// RandomWhiteLight - returns a random white light
func (configuration Configuration) RandomWhiteLight() string {
	return configuration.WhiteLights[rand.Intn(len(configuration.WhiteLights))]
}
