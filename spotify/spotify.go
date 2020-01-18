package spotify

import (
	"time"
)

// ClientCredentials - exported client credentials
var ClientCredentials Client

// Pause - used to stop the application polling
var Pause bool

var authCredentials Auth
var player Player
var audioAnalysis AudioAnalysis
var lastPlayerUpdate time.Time
var currentPosition chan CurrentPosition
var stop chan bool
var detecting bool
