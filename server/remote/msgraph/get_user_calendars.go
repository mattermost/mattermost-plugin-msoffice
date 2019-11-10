// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package msgraph

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-msoffice/server/remote"
)

func (c *client) GetUserCalendars(userID string) ([]*remote.Calendar, error) {
	var v struct {
		Value []*remote.Calendar `json:"value"`
	}
	req := c.rbuilder.Me().Calendars().Request()
	req.Expand("children")
	err := req.JSONRequest(c.ctx, http.MethodGet, "", nil, &v)
	if err != nil {
		return nil, err
	}
	return v.Value, nil
}
