package tools

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GetExternalIPParams struct {
	URL string
}

var DefaultGetExternalIPParams = GetExternalIPParams{
	URL: "https://api.ipify.org",
}

func GetExternalIP(params *GetExternalIPParams) (string, error) {
	if params == nil {
		params = &DefaultGetExternalIPParams
	}

	resp, err := http.Get(params.URL)
	if err != nil {
		return "", fmt.Errorf("error fetching external IP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error fetching external IP: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return strings.TrimSpace(string(body)), nil
}

func MustGetExternalIP(params *GetExternalIPParams) string {
	ip, err := GetExternalIP(params)
	if err != nil {
		panic(err)
	}
	return ip
}
