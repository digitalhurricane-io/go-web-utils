package utils

import (
	"github.com/pkg/errors"
	"net/url"
)

// FormatUrl Ensure that url is correctly formatted. If the url does not contain a scheme ("https://")
// it will be added.
func FormatUrl(originalUrl string) (string, error) {
	u, err := url.Parse(originalUrl)
	if err != nil {
		return "", err
	}

	urlIsValid := u.Host != "" && u.Scheme != ""
	if urlIsValid {
		return originalUrl, nil
	}

	if u.Host == "" {
		return "", errors.New("specified url is incorrectly formatted. It does not contain a host")
	}

	// if original url does not have a scheme, add one
	var fixedUrl = originalUrl
	if u.Scheme == "" {
		fixedUrl = "https://" + originalUrl
	}

	f, err := url.Parse(fixedUrl)
	if err != nil || f.Host == "" || f.Scheme == "" {
		return "", errors.New("specified url is incorrectly formatted. It does not contain a url scheme")
	}

	return fixedUrl, nil
}