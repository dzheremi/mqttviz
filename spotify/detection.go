package spotify

import (
	"fmt"
	"github.com/dzheremi/mqttviz/mqtt"
	"math/rand"
	"time"
)

func startDetection() {
	if !detecting {
		currentPosition = make(chan CurrentPosition)
		stop = make(chan bool)
		go durationTracker()
		go sectionDetection()
		go beatDetection()
		go tatumDetection()
		detecting = true
	}
}

func stopDetection() {
	if detecting {
		close(stop)
		detecting = false
	}
}

func barDetection() {
	previousBar := <-currentPosition
	go barHook(previousBar)
	for {
		select {
		case <-stop:
			return
		default:
			thisBar := <-currentPosition
			if previousBar.Bar.Start != thisBar.Bar.Start {
				go barHook(thisBar)
				previousBar = thisBar
			}
		}
	}
}

func beatDetection() {
	previousBeat := <-currentPosition
	go beatHook(previousBeat)
	for {
		select {
		case <-stop:
			return
		default:
			thisBeat := <-currentPosition
			if previousBeat.Beat.Start != thisBeat.Beat.Start {
				go beatHook(thisBeat)
				previousBeat = thisBeat
			}
		}
	}
}

func tatumDetection() {
	previousTatum := <-currentPosition
	go tatumHook(previousTatum)
	for {
		select {
		case <-stop:
			return
		default:
			thisTatum := <-currentPosition
			if previousTatum.Tatum.Start != thisTatum.Tatum.Start {
				go tatumHook(thisTatum)
				previousTatum = thisTatum
			}
		}
	}
}

func sectionDetection() {
	previousSection := <-currentPosition
	go sectionHook(previousSection)
	for {
		select {
		case <-stop:
			return
		default:
			thisSection := <-currentPosition
			if previousSection.Section.Start != thisSection.Section.Start {
				go sectionHook(thisSection)
				previousSection = thisSection
			}
		}
	}
}

func segmentDetection() {
	previousSegment := <-currentPosition
	go segmentHook(previousSegment)
	for {
		select {
		case <-stop:
			return
		default:
			thisSegment := <-currentPosition
			if previousSegment.Segment.Start != thisSegment.Segment.Start {
				go segmentHook(thisSegment)
				previousSegment = thisSegment
			}
		}
	}
}

func barHook(currentPosition CurrentPosition) {
  // For future use
}

func beatHook(currentPosition CurrentPosition) {
	mqtt.Client.Publish(fmt.Sprintf("%s/cmnd/Dimmer", mqtt.Settings.RGBLightGroup), 0, false, "100")
	for _, n := range mqtt.Settings.RGBLights {
		mqtt.Client.Publish(fmt.Sprintf("%s/cmnd/Color", n), 0, false, fmt.Sprintf("%d", rand.Intn(10)+1))
	}
}

func tatumHook(currentPosition CurrentPosition) {
	mqtt.Client.Publish(fmt.Sprintf("%s/cmnd/Power", mqtt.Settings.WhiteLightGroup), 0, false, "Off")
	time.Sleep(time.Duration(currentPosition.Tatum.Duration/2) * time.Second)
	mqtt.Client.Publish(fmt.Sprintf("%s/cmnd/Dimmer", mqtt.Settings.RandomWhiteLight()), 0, false, "100")
}

func sectionHook(currentPosition CurrentPosition) {
  // For future use
}

func segmentHook(currentPosition CurrentPosition) {
  // For future use
}
