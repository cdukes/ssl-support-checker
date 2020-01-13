package checker

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SupportsSSL(url *url.URL) bool {
	host := strings.ToLower(url.Host)

	h := sha1.New()
	h.Write([]byte(host))

	s := hex.EncodeToString(h.Sum(nil))
	prefix := s[0:4]

	client := &http.Client{
		Timeout: time.Second,
	}

	u := "https://duckduckgo.com/smarter_encryption.js?pv1=" + prefix

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	results := []string{}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return false
	}

	for _, r := range results {
		if r != s {
			continue
		}

		return true
	}

	return false
}

func MaybeUpgradeURL(s string) string {
	if !strings.HasPrefix(s, "http://") {
		return s
	}

	parts, err := url.Parse(s)
	if err != nil {
		return s
	}

	if !SupportsSSL(parts) {
		return s
	}

	parts.Scheme = "https"

	return parts.String()
}
