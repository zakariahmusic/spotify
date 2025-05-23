package spotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// UserHasTracks checks if one or more tracks are saved to the current user's
// "Your Music" library.
func (c *Client) UserHasTracks(ctx context.Context, ids ...ID) ([]bool, error) {
	if l := len(ids); l == 0 || l > 50 {
		return nil, errors.New("spotify: UserHasTracks supports 1 to 50 IDs per call")
	}
	spotifyURL := fmt.Sprintf("%sme/tracks/contains?ids=%s", c.baseURL, strings.Join(toStringSlice(ids), ","))

	var result []bool

	err := c.get(ctx, spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// AddTracksToLibrary saves one or more tracks to the current user's
// "Your Music" library.  This call requires the ScopeUserLibraryModify scope.
// A track can only be saved once; duplicate IDs are ignored.
func (c *Client) AddTracksToLibrary(ctx context.Context, ids ...ID) error {
	return c.modifyLibraryTracks(ctx, true, ids...)
}

// RemoveTracksFromLibrary removes one or more tracks from the current user's
// "Your Music" library.  This call requires the ScopeUserModifyLibrary scope.
// Trying to remove a track when you do not have the user's authorization
// results in a `spotify.Error` with the status code set to http.StatusUnauthorized.
func (c *Client) RemoveTracksFromLibrary(ctx context.Context, ids ...ID) error {
	return c.modifyLibraryTracks(ctx, false, ids...)
}

func (c *Client) modifyLibraryTracks(ctx context.Context, add bool, ids ...ID) error {
	if l := len(ids); l == 0 || l > 50 {
		return errors.New("spotify: this call supports 1 to 50 IDs per call")
	}
	spotifyURL := fmt.Sprintf("%sme/tracks?ids=%s", c.baseURL, strings.Join(toStringSlice(ids), ","))
	method := "DELETE"
	if add {
		method = "PUT"
	}
	req, err := http.NewRequestWithContext(ctx, method, spotifyURL, nil)
	if err != nil {
		return err
	}
	err = c.execute(req, nil)
	if err != nil {
		return err
	}
	return nil
}
