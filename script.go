package landscape

import (
	"encoding/json"
	"fmt"
	"io"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Attachment struct {
	Filename string `json:"filename"`
	ID       int    `json:"id"`
}

type ScriptType struct {
	ID             int          `json:"id"`
	Title          string       `json:"title"`
	VersionNumber  int          `json:"version_number"`
	CreatedBy      User         `json:"created_by"`
	CreatedAt      string       `json:"created_at"`
	LastEditedBy   User         `json:"last_edited_by"`
	LastEditedAt   string       `json:"last_edited_at"`
	ScriptProfiles []any        `json:"script_profiles"`
	Status         string       `json:"status"`
	Attachments    []Attachment `json:"attachments"`
	Code           string       `json:"code"`
	Interpreter    string       `json:"interpreter"`
	AccessGroup    string       `json:"access_group"`
	TimeLimit      int          `json:"time_limit"`
	Username       string       `json:"username"`
	IsRedactable   bool         `json:"is_redactable"`
	IsEditable     bool         `json:"is_editable"`
	IsExecutable   bool         `json:"is_executable"`
}

func (c *LandscapeAPIClient) GetScript(id int) (*ScriptType, error) {
	var script ScriptType
	res, err := c.DoRequest("GET", fmt.Sprintf("scripts/%d", id), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bodyBytes, &script); err != nil {
		return nil, err
	}

	return &script, nil
}
