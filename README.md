# MQTTViz
## MQTT Music Visualizer for Tasmota Devices

This is a very rough working prototype of a music visualizer utilising [Spotify's Audio Analysis API](https://developer.spotify.com/documentation/web-api/reference/tracks/get-audio-analysis/) to control lighting devices running [Tasmota](https://tasmota.github.io/docs/) firmware for ESP8266 devices.

[![MQTTViz](https://img.youtube.com/vi/50cjBKx-N0U/0.jpg)](https://www.youtube.com/watch?v=50cjBKx-N0U)

[![HomeAssistant](https://img.youtube.com/vi/iz8wpWk6qoE/0.jpg)](https://www.youtube.com/watch?v=iz8wpWk6qoE)

### Requirements
* GoLang 1.13.0
* An MQTT Broker (such as [Eclipse Mosquitto](https://mosquitto.org/))
* ESP8266 LED light bulbs (RGB or white) flashed with Tasmota - I found [this](https://github.com/ct-Open-Source/tuya-convert) very helpful
* [Spotify Developer Account](https://developer.spotify.com/dashboard/)

### Initial Setup
1. Create a Spotify App via the Spotify Developer [Dashboard](https://developer.spotify.com/dashboard/) with the following callback URL: `http://localhost:5555/callback`.

   Ensure you take note of your Client ID and Secret.

   If you are not running locally, replace callback URL `localhost` with the appropriate IP address.

2. For each of your bulbs, use the Tasmota web UI to ensure your MQTT broker is configured and full topic format is:

   ```
   %topic%/%prefix%/
   ```

3. Set a group topic for each of your white LED bulbs (i.e. `white_lights`) using the Tasmota console:

   ```
   GroupTopic white_lights
   ```

4. and one for each of your RGB LED bulbs (i.e. `rgb_lights`):

   ```
   GroupTopic rgb_lights
   ```

5. Also ensure fade is disabled as this breaks sync:

   ```
   Fade 0
   ```

### Required Environment Variables
* `MQTT_HOST` - The URI of your MQTT broker (i.e. `tcp://192.168.1.1:1883`)
* `WHITE_LIGHT_GROUP` - The group topic you set above for white lights
* `RGB_LIGHT_GROUP` - The group topic you set above for RGB lights
* `WHITE_LIGHTS` - A comma delimitered list of each of your white lights (MQTT topic)
* `RGB_LIGHTS` - A comma delimitered list of each of your RGB lights (MQTT topic)
* `CLIENT_ID` - Your Spotify Client ID
* `CLIENT_SECRET` - Your Spotify Client Secret
* `REDIRECT_URI` - Your callback URI

### Running Locally

```bash
$ go run mqttviz.go
```

Then navigate to `http://localhost:5555` to authenticate to Spotify (or appropriate IP address).

Start playing a track on Spotify and that's it!

#### Stopping Visualizer
You can stop the visualizer by making a POST request to `http://localhost:5555/stop`

#### Resuming the Visualizer
You can resume the visualizer by making a POST request to `http://localhost:5555/start`

#### Getting Visualizer Status
Make a GET request to `http://localhost:5555/status`

#### Running in Docker
`Dockerfile` has been supplied, please refer to Required Environment Variables above.
