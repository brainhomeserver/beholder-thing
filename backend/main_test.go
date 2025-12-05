package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cucumber/godog"
)

type apiFeature struct {
	response *http.Response
	body     []byte
}

var testServer *httptest.Server

func (a *apiFeature) resetResponse() {
	if a.response != nil {
		a.response.Body.Close()
	}
	a.response = nil
	a.body = nil
}

func (a *apiFeature) iSendAGETRequestTo(path string) error {
	if testServer == nil {
		return errors.New("test server is not running")
	}

	resp, err := http.Get(testServer.URL + path)
	if err != nil {
		return err
	}

	a.response = resp
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	a.body = data
	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(expected int) error {
	if a.response == nil {
		return errors.New("no response received")
	}

	if a.response.StatusCode != expected {
		return fmt.Errorf("expected status code %d, got %d", expected, a.response.StatusCode)
	}
	return nil
}

func (a *apiFeature) theHeaderShouldBe(key, value string) error {
	if a.response == nil {
		return errors.New("no response received")
	}

	if got := a.response.Header.Get(key); got != value {
		return fmt.Errorf("expected header %s to be %q, got %q", key, value, got)
	}
	return nil
}

func (a *apiFeature) theResponseBodyShouldBe(body *godog.DocString) error {
	expected := strings.TrimSpace(body.Content)
	actual := strings.TrimSpace(string(a.body))

	if expected != actual {
		return fmt.Errorf("expected body %q, got %q", expected, actual)
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		testServer = httptest.NewServer(newMux())
	})

	ctx.AfterSuite(func() {
		if testServer != nil {
			testServer.Close()
		}
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	feature := &apiFeature{}

	ctx.Before(func(c context.Context, _ *godog.Scenario) (context.Context, error) {
		feature.resetResponse()
		return c, nil
	})

	ctx.Step(`^I send a GET request to "([^"]*)"$`, feature.iSendAGETRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, feature.theResponseCodeShouldBe)
	ctx.Step(`^the header "([^"]*)" should be "([^"]*)"$`, feature.theHeaderShouldBe)
	ctx.Step(`^the response body should be:$`, feature.theResponseBodyShouldBe)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "backend",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("godog feature tests failed")
	}
}
