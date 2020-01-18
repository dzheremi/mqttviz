package spotify

// Client - spotify client credentials
type Client struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Code         string
}

// Auth - spotify user credentials
type Auth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// CurrentPosition - container for current position of player
type CurrentPosition struct {
	Bar     Bar
	Beat    Beat
	Tatum   Tatum
	Section Section
	Segment Segment
}

// Player - spotify player details
type Player struct {
	Device struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"device"`
	ShuffleState bool        `json:"shuffle_state"`
	RepeatState  string      `json:"repeat_state"`
	Timestamp    int64       `json:"timestamp"`
	Context      interface{} `json:"context"`
	ProgressMs   int64       `json:"progress_ms"`
	Item         struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Pausing      bool `json:"pausing"`
			SkippingPrev bool `json:"skipping_prev"`
		} `json:"disallows"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}

// AudioAnalysis - spotify audio analysis header
type AudioAnalysis struct {
	Meta     Meta      `json:"meta"`
	Track    Track     `json:"track"`
	Bars     []Bar     `json:"bars"`
	Beats    []Beat    `json:"beats"`
	Tatums   []Tatum   `json:"tatums"`
	Sections []Section `json:"sections"`
	Segments []Segment `json:"segments"`
}

// FindCurrentPosition - returns current position given seconds into a track
func (analysis AudioAnalysis) FindCurrentPosition(position float64) CurrentPosition {
	var current CurrentPosition
	if len(analysis.Bars) > 0 {
		previousBar := analysis.Bars[0]
		for i, n := range analysis.Bars {
			if position < n.Start {
				current.Bar = previousBar
				break
			}
			previousBar = analysis.Bars[i]
		}
	}
	if len(analysis.Beats) > 0 {
		previousBeat := analysis.Beats[0]
		for i, n := range analysis.Beats {
			if position < n.Start {
				current.Beat = previousBeat
				break
			}
			previousBeat = analysis.Beats[i]
		}
	}
	if len(analysis.Tatums) > 0 {
		previousTatum := analysis.Tatums[0]
		for i, n := range analysis.Tatums {
			if position < n.Start {
				current.Tatum = previousTatum
				break
			}
			previousTatum = analysis.Tatums[i]
		}
	}
	if len(analysis.Sections) > 0 {
		previousSection := analysis.Sections[0]
		for i, n := range analysis.Sections {
			if position < n.Start {
				current.Section = previousSection
				break
			}
			previousSection = analysis.Sections[i]
		}
	}
	if len(analysis.Segments) > 0 {
		previousSegment := analysis.Segments[0]
		for i, n := range analysis.Segments {
			if position < n.Start {
				current.Segment = previousSegment
				break
			}
			previousSegment = analysis.Segments[i]
		}
	}
	return current
}

// Meta - spotify track metadata
type Meta struct {
	AnalyzerVersion string  `json:"analyzer_version"`
	Platform        string  `json:"platform"`
	DetailedStatus  string  `json:"detailed_status"`
	StatusCode      int     `json:"status_code"`
	Timestamp       int     `json:"timestamp"`
	AnalysisTime    float64 `json:"analysis_time"`
	InputProcess    string  `json:"input_process"`
}

// Track - spotify track details
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

// Bar - bar of a song
type Bar struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}

// Beat - beat of a song
type Beat struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}

// Tatum - subunit of a beat
type Tatum struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}

// Section - section of a song (when time signature changes)
type Section struct {
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

// Segment - segment of a song
type Segment struct {
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
