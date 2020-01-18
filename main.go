package main

import(
  "fmt"
  "time"
  "encoding/json"
  "io/ioutil"
)

type AudioAnalysis struct {
	Meta     Meta       `json:"meta"`
	Track    Track      `json:"track"`
	Bars     []Bars     `json:"bars"`
	Beats    []Beats    `json:"beats"`
	Tatums   []Tatums   `json:"tatums"`
	Sections []Sections `json:"sections"`
	Segments []Segments `json:"segments"`
}
type Meta struct {
	AnalyzerVersion string  `json:"analyzer_version"`
	Platform        string  `json:"platform"`
	DetailedStatus  string  `json:"detailed_status"`
	StatusCode      int     `json:"status_code"`
	Timestamp       int     `json:"timestamp"`
	AnalysisTime    float64 `json:"analysis_time"`
	InputProcess    string  `json:"input_process"`
}
type Track struct {
	NumSamples              int     `json:"num_samples"`
	Duration                float64 `json:"duration"`
	SampleMd5               string  `json:"sample_md5"`
	OffsetSeconds           int     `json:"offset_seconds"`
	WindowSeconds           int     `json:"window_seconds"`
	AnalysisSampleRate      int     `json:"analysis_sample_rate"`
	AnalysisChannels        int     `json:"analysis_channels"`
	EndOfFadeIn             float64 `json:"end_of_fade_in"`
	StartOfFadeOut          float64 `json:"start_of_fade_out"`
	Loudness                float64 `json:"loudness"`
	Tempo                   float64 `json:"tempo"`
	TempoConfidence         float64 `json:"tempo_confidence"`
	TimeSignature           int     `json:"time_signature"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	Key                     int     `json:"key"`
	KeyConfidence           float64 `json:"key_confidence"`
	Mode                    int     `json:"mode"`
	ModeConfidence          float64 `json:"mode_confidence"`
	Codestring              string  `json:"codestring"`
	CodeVersion             float64 `json:"code_version"`
	Echoprintstring         string  `json:"echoprintstring"`
	EchoprintVersion        float64 `json:"echoprint_version"`
	Synchstring             string  `json:"synchstring"`
	SynchVersion            float64 `json:"synch_version"`
	Rhythmstring            string  `json:"rhythmstring"`
	RhythmVersion           float64 `json:"rhythm_version"`
}
type Bars struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}
type Beats struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}
type Tatums struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}
type Sections struct {
	Start                   float64 `json:"start"`
	Duration                float64 `json:"duration"`
	Confidence              float64 `json:"confidence"`
	Loudness                float64 `json:"loudness"`
	Tempo                   float64 `json:"tempo"`
	TempoConfidence         float64 `json:"tempo_confidence"`
	Key                     int     `json:"key"`
	KeyConfidence           float64 `json:"key_confidence"`
	Mode                    int     `json:"mode"`
	ModeConfidence          float64 `json:"mode_confidence"`
	TimeSignature           int     `json:"time_signature"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
}
type Segments struct {
	Start           float64   `json:"start"`
	Duration        float64   `json:"duration"`
	Confidence      float64   `json:"confidence"`
	LoudnessStart   float64   `json:"loudness_start"`
	LoudnessMaxTime float64   `json:"loudness_max_time"`
	LoudnessMax     float64   `json:"loudness_max"`
	Pitches         []float64 `json:"pitches"`
	Timbre          []float64 `json:"timbre"`
	LoudnessEnd     float64   `json:"loudness_end,omitempty"`
}
type CurrentPosition struct {
  Bars      Bars
  Beats     Beats
  Tatums    Tatums
  Sections  Sections
  Segments  Segments
}

func (analysis AudioAnalysis) FindCurrentPosition(position float64) CurrentPosition {
  var current CurrentPosition
  previousBar := analysis.Bars[0]
  for i, n := range analysis.Bars {
    if position < n.Start {
      current.Bars = previousBar
      break
    }
    previousBar = analysis.Bars[i]
  }
  previousBeat := analysis.Beats[0]
  for i, n := range analysis.Beats {
    if position < n.Start {
      current.Beats = previousBeat
      break
    }
    previousBeat = analysis.Beats[i]
  }
  previousTatum := analysis.Tatums[0]
  for i, n := range analysis.Tatums {
    if position < n.Start {
      current.Tatums = previousTatum
      break
    }
    previousTatum = analysis.Tatums[i]
  }
  previousSection := analysis.Sections[0]
  for i, n := range analysis.Sections {
    if position < n.Start {
      current.Sections = previousSection
      break
    }
    previousSection = analysis.Sections[i]
  }
  previousSegment := analysis.Segments[0]
  for i, n := range analysis.Segments {
    if position < n.Start {
      current.Segments = previousSegment
      break
    }
    previousSegment = analysis.Segments[i]
  }
  return current
}

var audioAnalysis AudioAnalysis

func main() {
  file, err := ioutil.ReadFile("sample.json")
  if err != nil {
    panic("--- No sample json found ---")
  }
  initializeAudioAnalysis([]byte(file))
  startTime := time.Now()
  for {
    duration := time.Now().Sub(startTime).Seconds()
    audioAnalysis.FindCurrentPosition(time.Now().Sub(startTime).Seconds())
    fmt.Println(duration)
  }
}

func initializeAudioAnalysis(jsonData []byte) {
  err := json.Unmarshal(jsonData, &audioAnalysis)
  if err != nil {
    panic("--- Invalid audio analysis ---")
  }
}
