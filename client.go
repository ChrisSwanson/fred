package fred

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

var (
	baseApiUrl = "https://api.stlouisfed.org/"
)

type FredClient struct {
	apiKey      string
	apiUrl      *url.URL
	HttpClient  *http.Client
	RateLimiter *rate.Limiter
	Logger      *zerolog.Logger
}

func NewFredClient(apiKey string, rateLimitTime time.Duration, rateLimitCount int, logger *zerolog.Logger) (c *FredClient, err error) {

	if logger == nil {
		lgr := zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger = &lgr
	}

	logger.Debug().Interface("time", rateLimitTime).Interface("count", rateLimitCount).Msg("creating new sec client with rate limiter")

	rateLimiter := rate.NewLimiter(rate.Every(rateLimitTime), rateLimitCount)

	baseURL, err := url.Parse(baseApiUrl)
	if err != nil {
		logger.Error().Err(err).Msg("error parsing base url")
	}

	return &FredClient{
		apiKey:      apiKey,
		apiUrl:      baseURL,
		HttpClient:  http.DefaultClient,
		RateLimiter: rateLimiter,
		Logger:      &zerolog.Logger{},
	}, nil
}

func (f *FredClient) Request(urlPath string, params map[string]interface{}) (results *Results, err error) {

	values := url.Values{
		"api_key":   {f.apiKey},
		"file_type": {"json"},
	}

	for param := range params {
		values.Add(param, params[param].(string))
	}

	urlObj := f.apiUrl
	urlObj.Path = urlPath
	urlObj.RawQuery = values.Encode()
	url := urlObj.String()

	f.Logger.Debug().Str("url", url).Msg("calling API")

	ctx := context.Background()
	err = f.RateLimiter.Wait(ctx)
	if err != nil {
		f.Logger.Error().Err(err).Msg("rate limiter error")
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		f.Logger.Error().Err(err).Msg("error with http request")
		return nil, err
	}

	resp, err := f.HttpClient.Do(req)
	if err != nil {
		f.Logger.Error().Err(err).Msg("error with http call")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		f.Logger.Error().Err(err).Msg("error reading response body")
		return nil, err
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		f.Logger.Error().Err(err).Msg("error unmarshalling response body")
		return nil, err
	}

	return results, err
}
