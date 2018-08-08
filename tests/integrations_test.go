// +build !appengine

package tests

import (
	"os/exec"
	"testing"
)

// Starts an app, sends a request, and tests that the resulting bugsnag
// error report has the correct values.

func TestNegroniRequestPanic(t *testing.T) {
	defer pkill("negroni")
	startTestServer()
	body := startPanickingApp(t,
		"./fixtures/negroni.go", "http://localhost:9078", "default")
	assertSeverityReasonEqual(t, body, "error", "unhandledErrorMiddleware", true)
}

func TestNegroniRequestPanicCallbackAltered(t *testing.T) {
	defer pkill("negroni")
	startTestServer()
	body := startPanickingApp(t,
		"./fixtures/negroni.go", "http://localhost:9078", "beforenotify")
	assertSeverityReasonEqual(t, body, "info", "userCallbackSetSeverity", true)
}

func TestGinRequestPanic(t *testing.T) {
	defer pkill("gin")
	startTestServer()
	body := startPanickingApp(t, "./fixtures/gin.go", "http://localhost:9079", "default")
	assertSeverityReasonEqual(t, body, "error", "unhandledErrorMiddleware", true)
}

func TestGinRequestPanicCallbackAltered(t *testing.T) {
	defer pkill("gin")
	startTestServer()
	body := startPanickingApp(t, "./fixtures/gin.go", "http://localhost:9079", "beforenotify")
	assertSeverityReasonEqual(t, body, "info", "userCallbackSetSeverity", true)
}

func TestMartiniRequestPanic(t *testing.T) {
	defer pkill("martini")
	startTestServer()
	body := startPanickingApp(t, "./fixtures/martini.go", "http://localhost:3000", "default")
	assertSeverityReasonEqual(t, body, "error", "unhandledErrorMiddleware", true)
}

func TestMartiniRequestPanicCallbackAltered(t *testing.T) {
	defer pkill("martini")
	startTestServer()
	body := startPanickingApp(t, "./fixtures/martini.go", "http://localhost:3000", "beforenotify")
	assertSeverityReasonEqual(t, body, "info", "userCallbackSetSeverity", true)
}

func TestRevelRequestPanic(t *testing.T) {
	defer pkill("revel")
	startTestServer()
	body := startRevelApp(t, "default")
	assertSeverityReasonEqual(t, body, "error", "unhandledErrorMiddleware", true)
}

func TestRevelRequestPanicCallbackAltered(t *testing.T) {
	defer pkill("revel")
	startTestServer()
	body := startRevelApp(t, "beforenotify")
	assertSeverityReasonEqual(t, body, "info", "userCallbackSetSeverity", true)
}

func pkill(process string) {
	cmd := exec.Command("pkill", "-x", process)
	cmd.Start()
	cmd.Wait()
}
