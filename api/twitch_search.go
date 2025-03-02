package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type SearchChannel struct {
	UserID          string    `json:"id"`
	UserLogin       string    `json:"broadcaster_login"`
	UserDisplayName string    `json:"display_name"`
	GameID          string    `json:"game_id"`
	GameName        string    `json:"game_name"`
	Title           string    `json:"title"`
	Tags            []string  `json:"tags"`
	Language        string    `json:"broadcaster_language"`
	ThumbnailURL    string    `json:"thumbnail_url"`
	IsLive          bool      `json:"is_live"`
	StartedAt       time.Time `json:"started_at"`
}

type SearchChannelsResource struct {
	client *Client
}

func NewSearchResource(client *Client) *SearchChannelsResource {
	return &SearchChannelsResource{client}
}

type SearchChannelsListCall struct {
	resource *SearchChannelsResource
	opts     []RequestOption
}

type SearchChannelsListResponse struct {
	Header http.Header
	Data   []SearchChannel
	Cursor string
}

// List creates a request to list streams based on the specified criteria.
//
// Requires an app or user access token. No scope is required.
func (r *SearchChannelsResource) List() *SearchChannelsListCall {
	return &SearchChannelsListCall{resource: r}
}

// Query - the URI-encoded search string.
//
// For example, encode search strings like angel of death as angel%20of%20death.
func (c *SearchChannelsListCall) Query(opts string) *SearchChannelsListCall {
	c.opts = append(c.opts, SetQueryParameter("query", opts))
	return c
}

// LiveOnly - a Boolean value that determines whether
// the response includes only channels that are currently streaming live.
//
// Set to true to get only channels that are streaming live; otherwise,
// false to get live and offline channels. The default is false.
func (c *SearchChannelsListCall) LiveOnly(isLiveOnly bool) *SearchChannelsListCall {
	c.opts = append(c.opts, SetQueryParameter("live_only", fmt.Sprint(isLiveOnly)))
	return c
}

// First limits the number of results to the specified amount.
//
// Maximum: 100 (default: 20)
func (c *SearchChannelsListCall) First(n int) *SearchChannelsListCall {
	c.opts = append(c.opts, SetQueryParameter("first", fmt.Sprint(n)))
	return c
}

// After filters the results to streams that started after the specified cursor.
func (c *SearchChannelsListCall) After(cursor string) *SearchChannelsListCall {
	c.opts = append(c.opts, SetQueryParameter("after", cursor))
	return c
}

// Do executes the request.
func (c *SearchChannelsListCall) Do(ctx context.Context, opts ...RequestOption) (*SearchChannelsListResponse, error) {
	res, err := c.resource.client.doRequest(ctx, http.MethodGet, "/streams", nil, append(opts, c.opts...)...)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := decodeResponse[SearchChannel](res)
	if err != nil {
		return nil, err
	}

	return &SearchChannelsListResponse{
		Header: res.Header,
		Data:   data.Data,
		Cursor: data.Pagination.Cursor,
	}, nil
}
