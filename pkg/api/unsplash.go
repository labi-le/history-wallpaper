package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Unsplash struct {
	client http.Client
}

func (u *Unsplash) Find(ctx context.Context, q string, r Resolution) (io.ReadCloser, error) {
	request, _ := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://source.unsplash.com/%s/?%s", q, r),
		nil,
	)

	resp, err := u.client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
