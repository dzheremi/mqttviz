package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetPlayerStatus - gets the current player status from Spotify
func GetPlayerStatus() {
	Pause = false
	lastTrackID := ""
	lastPlayStatus := false
	firstLoop := true
	for {
		if Pause {
			stopDetection()
			lastPlayStatus = false
		}
		for Pause {
			time.Sleep(3 * time.Second)
		}
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authCredentials.AccessToken))
		response, err := client.Do(request)
		lastPlayerUpdate = time.Now()
		if err != nil || response.StatusCode != 200 {
			if response.StatusCode == 401 {
				refreshAccessToken()
			}
			continue
		}
		defer response.Body.Close()
		buffer, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(buffer, &player)
		if lastTrackID != player.Item.ID {
			if lastPlayStatus {
				stopDetection()
			}
			for !getTrackAnalysis(player.Item.ID) {
			}
			lastTrackID = player.Item.ID
			if player.IsPlaying {
				startDetection()
				lastPlayStatus = player.IsPlaying
				continue
			}
		}
		if firstLoop {
			lastPlayStatus = player.IsPlaying
			firstLoop = false
		}
		if lastPlayStatus != player.IsPlaying {
			if player.IsPlaying {
				startDetection()
			} else {
				stopDetection()
			}
			lastPlayStatus = player.IsPlaying
		}
		time.Sleep(3 * time.Second)
	}
}

func getTrackAnalysis(trackID string) bool {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/audio-analysis/%s", trackID), nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authCredentials.AccessToken))
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		return false
	}
	defer response.Body.Close()
	buffer, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(buffer, &audioAnalysis)
	return true
}

func durationTracker() {
	for {
		select {
		case <-stop:
			return
		default:
			durationMs := player.ProgressMs + time.Now().Sub(lastPlayerUpdate).Milliseconds()
			durationS := float64(durationMs) / 1000.0
			currentPosition <- audioAnalysis.FindCurrentPosition(durationS)
			time.Sleep(1 * time.Millisecond)
		}
	}
}
