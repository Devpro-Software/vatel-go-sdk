package vatel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetOrganization(ctx context.Context) (*Organization, error) {
	var out Organization
	if err := c.doJSON(ctx, http.MethodGet, "/organization", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ListLLMs(ctx context.Context) ([]string, error) {
	var out LLMStringsResponse
	if err := c.doJSON(ctx, http.MethodGet, "/llms", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return out.LLMs, nil
}

func (c *Client) ListVoices(ctx context.Context) (*VoicesListResponse, error) {
	var out VoicesListResponse
	if err := c.doJSON(ctx, http.MethodGet, "/voices", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateAgent(ctx context.Context, in AgentCreateInput) (*Agent, error) {
	var out Agent
	if err := c.doJSON(ctx, http.MethodPost, "/agents", nil, in, []int{http.StatusCreated}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetAgent(ctx context.Context, agentID string) (*Agent, error) {
	var out Agent
	path := fmt.Sprintf("/agents/%s", url.PathEscape(agentID))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateAgent(ctx context.Context, agentID string, in AgentUpdateInput) (*Agent, error) {
	var out Agent
	path := fmt.Sprintf("/agents/%s", url.PathEscape(agentID))
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, in, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteAgent(ctx context.Context, agentID string) error {
	path := fmt.Sprintf("/agents/%s", url.PathEscape(agentID))
	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, []int{http.StatusNoContent}, nil)
}

func (c *Client) ListAgentVersions(ctx context.Context, agentID string) ([]GraphVersion, error) {
	var out []GraphVersion
	path := fmt.Sprintf("/agents/%s/versions", url.PathEscape(agentID))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetAgentVersion(ctx context.Context, agentID, versionID string) (*GraphVersionDetail, error) {
	var out GraphVersionDetail
	path := fmt.Sprintf("/agents/%s/versions/%s", url.PathEscape(agentID), url.PathEscape(versionID))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PublishAgentVersion(ctx context.Context, agentID, versionID string) (*GraphVersion, error) {
	var out GraphVersion
	path := fmt.Sprintf("/agents/%s/versions/%s/publish", url.PathEscape(agentID), url.PathEscape(versionID))
	if err := c.doJSON(ctx, http.MethodPost, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DialAgent(ctx context.Context, agentID, destinationNumber string) (*DialAgentResponse, error) {
	q := url.Values{}
	q.Set("number", destinationNumber)
	path := fmt.Sprintf("/agents/%s/dial", url.PathEscape(agentID))
	var out DialAgentResponse
	if err := c.doJSON(ctx, http.MethodPost, path, q, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ListTwilioNumbers(ctx context.Context) ([]TwilioPhoneNumber, error) {
	var out []TwilioPhoneNumber
	if err := c.doJSON(ctx, http.MethodGet, "/twilio/numbers", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) ImportTwilioNumber(ctx context.Context, in TwilioPhoneNumberImportInput) (*TwilioPhoneNumber, error) {
	var out TwilioPhoneNumber
	if err := c.doJSON(ctx, http.MethodPost, "/twilio/numbers", nil, in, []int{http.StatusCreated}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PatchTwilioNumberLabel(ctx context.Context, id string, in TwilioPhoneNumberLabelPatchInput) (*TwilioPhoneNumber, error) {
	var out TwilioPhoneNumber
	path := fmt.Sprintf("/twilio/numbers/%s", url.PathEscape(id))
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, in, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) ListSipTrunks(ctx context.Context) ([]SipTrunk, error) {
	var out []SipTrunk
	if err := c.doJSON(ctx, http.MethodGet, "/sip-trunks", nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CreateSipTrunk(ctx context.Context, in SipTrunkCreateInput) (*SipTrunk, error) {
	var out SipTrunk
	if err := c.doJSON(ctx, http.MethodPost, "/sip-trunks", nil, in, []int{http.StatusCreated}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSipTrunk(ctx context.Context, id string) (*SipTrunk, error) {
	var out SipTrunk
	path := fmt.Sprintf("/sip-trunks/%s", url.PathEscape(id))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UpdateSipTrunk(ctx context.Context, id string, in SipTrunkUpdateInput) (*SipTrunk, error) {
	var out SipTrunk
	path := fmt.Sprintf("/sip-trunks/%s", url.PathEscape(id))
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, in, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteSipTrunk(ctx context.Context, id string) error {
	path := fmt.Sprintf("/sip-trunks/%s", url.PathEscape(id))
	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, []int{http.StatusNoContent}, nil)
}

func (c *Client) ListSipTrunkAssignments(ctx context.Context, sipTrunkID string) ([]SipTrunkAgentAssignment, error) {
	var out []SipTrunkAgentAssignment
	path := fmt.Sprintf("/sip-trunks/%s/assignments", url.PathEscape(sipTrunkID))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CreateSipTrunkAssignment(ctx context.Context, sipTrunkID string, in SipTrunkAgentAssignmentCreateInput) (*SipTrunkAgentAssignment, error) {
	var out SipTrunkAgentAssignment
	path := fmt.Sprintf("/sip-trunks/%s/assignments", url.PathEscape(sipTrunkID))
	if err := c.doJSON(ctx, http.MethodPost, path, nil, in, []int{http.StatusCreated}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSipTrunkAssignment(ctx context.Context, assignmentID string) (*SipTrunkAgentAssignment, error) {
	var out SipTrunkAgentAssignment
	path := fmt.Sprintf("/sip-trunks/assignments/%s", url.PathEscape(assignmentID))
	if err := c.doJSON(ctx, http.MethodGet, path, nil, nil, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PatchSipTrunkAssignment(ctx context.Context, assignmentID string, in SipTrunkAgentAssignmentPatchInput) (*SipTrunkAgentAssignment, error) {
	var out SipTrunkAgentAssignment
	path := fmt.Sprintf("/sip-trunks/assignments/%s", url.PathEscape(assignmentID))
	if err := c.doJSON(ctx, http.MethodPatch, path, nil, in, []int{http.StatusOK}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DeleteSipTrunkAssignment(ctx context.Context, assignmentID string) error {
	path := fmt.Sprintf("/sip-trunks/assignments/%s", url.PathEscape(assignmentID))
	return c.doJSON(ctx, http.MethodDelete, path, nil, nil, []int{http.StatusNoContent}, nil)
}
