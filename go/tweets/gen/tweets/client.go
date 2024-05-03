// Code generated by goa v3.15.2, DO NOT EDIT.
//
// tweets client
//
// Command:
// $ goa gen tweets/design

package tweets

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "tweets" service client.
type Client struct {
	CreateTweetEndpoint goa.Endpoint
}

// NewClient initializes a "tweets" service client given the endpoints.
func NewClient(createTweet goa.Endpoint) *Client {
	return &Client{
		CreateTweetEndpoint: createTweet,
	}
}

// CreateTweet calls the "CreateTweet" endpoint of the "tweets" service.
// CreateTweet may return the following errors:
//   - "NotFound" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) CreateTweet(ctx context.Context, p *CreateTweetPayload) (res *Tweet, err error) {
	var ires any
	ires, err = c.CreateTweetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Tweet), nil
}