package sessions

import (
	"os"
	"runtime"
	"time"
)

// notifierPayload defines the .notifier subobject of the payload
type notifierPayload struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// appPayload defines the .app subobject of the payload
type appPayload struct {
	Type         string `json:"type"`
	ReleaseStage string `json:"releaseStage"`
	Version      string `json:"version"`
}

// devicePayload defines the .device subobject of the payload
type devicePayload struct {
	OsName   string `json:"osName"`
	Hostname string `json:"hostname"`
}

// sessionCountsPayload defines the .sessionCounts subobject of the payload
type sessionCountsPayload struct {
	StartedAt       string `json:"startedAt"`
	SessionsStarted int    `json:"sessionsStarted"`
}

// sessionPayload defines the top level payload object
type sessionPayload struct {
	Notifier      notifierPayload        `json:"notifier"`
	App           appPayload             `json:"app"`
	Device        devicePayload          `json:"device"`
	SessionCounts []sessionCountsPayload `json:"sessionCounts"`
}

// makeSessionPayload creates a sessionPayload based off of the given sessions and config
func makeSessionPayload(sessions []*Session, config *SessionTrackingConfiguration) *sessionPayload {
	releaseStage := config.ReleaseStage
	if releaseStage == "" {
		releaseStage = "production"
	}
	hostname := config.Hostname
	if hostname == "" {
		hostname, _ = os.Hostname() //Ignore the hostname if this call errors
	}

	return &sessionPayload{
		Notifier: notifierPayload{
			Name:    "Bugsnag Go",
			URL:     "https://github.com/bugsnag/bugsnag-go",
			Version: config.Version,
		},
		App: appPayload{
			Type:         config.AppType,
			Version:      config.AppVersion,
			ReleaseStage: releaseStage,
		},
		Device: devicePayload{
			OsName:   runtime.GOOS,
			Hostname: hostname,
		},
		SessionCounts: []sessionCountsPayload{
			{
				StartedAt:       sessions[0].StartedAt.UTC().Format(time.RFC3339),
				SessionsStarted: len(sessions),
			},
		},
	}
}
