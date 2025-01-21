package api

import (
	"fmt"
	"io"
	"net/http"
)

var typesPath = "/types"
var oidcrpPath = "/oidc_rp"
var authenticatorsPath = "/config/authenticators"
var resourcesPath = "/config/resources"
var spPath = "/config/samlsp"
var reloadPath = "/state"

func PostState(client *http.Client, w http.ResponseWriter) error {
	req, err := http.NewRequest("POST", baseApiUrl+reloadPath, nil)
	if err != nil {
		return fmt.Errorf("error creating POST request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
	return nil
}

func GetTypes(client *http.Client, w http.ResponseWriter) error {
	resp, err := client.Get(baseApiUrl + typesPath)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return nil
}

func GetOidcRp(client *http.Client, w http.ResponseWriter) error {
	resp, err := client.Get(baseApiUrl + oidcrpPath)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return nil
}

func GetAuthenticators(client *http.Client, w http.ResponseWriter) error {
	resp, err := client.Get(baseApiUrl + authenticatorsPath)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return nil
}

func GetMetaData(client *http.Client, w http.ResponseWriter) error {
	resp, err := client.Get(baseApiUrl + resourcesPath)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return nil
}

func GetSp(client *http.Client, w http.ResponseWriter) error {
	resp, err := client.Get(baseApiUrl + spPath)
	if err != nil {
		return fmt.Errorf("error making API call: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
	return nil
}
