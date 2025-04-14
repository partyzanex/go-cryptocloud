package cryptocloud

import (
	"fmt"
	"net/url"

	"resty.dev/v3"
)

const (
	authorizationHeader   = "Authorization"
	contentTypeHeader     = "Content-Type"
	tokenPrefix           = "Token "
	contentTypeHeaderJSON = "application/json"
	localeQueryParam      = "locale"
)

type config struct {
	apiKey string
	rawUrl string
	locale Locale
	client *resty.Client
}

type Option func(cfg *config)

func WithAPIKey(apiKey string) Option {
	return func(cfg *config) {
		cfg.apiKey = apiKey
	}
}

func WithAPIUrl(apiUrl string) Option {
	return func(cfg *config) {
		cfg.rawUrl = apiUrl
	}
}

func WithLocale(locale Locale) Option {
	return func(cfg *config) {
		cfg.locale = locale
	}
}

func WithClient(client *resty.Client) Option {
	return func(cfg *config) {
		cfg.client = client
	}
}

func newConfig(options []Option) (*config, error) {
	cfg := new(config)

	for _, option := range options {
		option(cfg)
	}

	if cfg.apiKey == "" {
		return nil, ErrRequiredAPIKey
	}

	if cfg.rawUrl == "" {
		return nil, ErrRequiredAPIUrl
	}

	uri, err := url.Parse(cfg.rawUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	if cfg.client == nil {
		cfg.client = resty.New()
	}

	if cfg.locale == "" {
		cfg.locale = LocaleEN
	}

	cfg.client.SetHeader(authorizationHeader, tokenPrefix+cfg.apiKey)
	cfg.client.SetHeader(contentTypeHeader, contentTypeHeaderJSON)
	cfg.client.SetBaseURL(uri.String())
	cfg.client.SetQueryParam(localeQueryParam, string(cfg.locale))

	return cfg, nil
}
