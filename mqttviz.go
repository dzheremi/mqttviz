package main

import (
	"github.com/dzheremi/mqttviz/mqtt"
	"github.com/dzheremi/mqttviz/spotify"
	"github.com/dzheremi/mqttviz/webserver"
	"math/rand"
  "os"
  "strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
	mqtt.Settings = mqtt.Configuration{
		Server:          os.Getenv("MQTT_HOST"),
		WhiteLightGroup: os.Getenv("WHITE_LIGHT_GROUP"),
		RGBLightGroup:   os.Getenv("RGB_LIGHT_GROUP"),
		WhiteLights: strings.Split(os.Getenv("WHITE_LIGHTS"), ","),
		RGBLights: strings.Split(os.Getenv("RGB_LIGHTS"), ","),
	}
	spotify.ClientCredentials = spotify.Client{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURI:  os.Getenv("REDIRECT_URI"),
		Code:         "",
	}
}

func main() {
	mqtt.Setup()
	webserver.Start()
	spotify.GetAccessToken()
	go webserver.API()
	go spotify.GetPlayerStatus()
	select {}
}
