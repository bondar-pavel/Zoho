// Code generated by go generate; DO NOT EDIT.
// This Zoho CRM endpoint API was generated @ 2019-10-08 22:29:14.0940069 -0600 CST m=+0.011522901
package crm

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetNotes returns a list of all notes
// https://www.zoho.com/crm/help/api/v2/#notes-api
func (c *API) GetNotes() (data NotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &NotesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"page":     "",
			"per_page": "200",
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return NotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*NotesResponse); ok {
		return *v, nil
	}

	return NotesResponse{}, fmt.Errorf("Data returned was not 'NotesResponse'")
}

type NotesResponse struct {
	Data []struct {
		Approval struct {
			Approve  bool `json:"approve,omitempty"`
			Delegate bool `json:"delegate,omitempty"`
			Reject   bool `json:"reject,omitempty"`
		} `json:"$approval,omitempty"`
		Followed  bool   `json:"$followed,omitempty"`
		SeModule  string `json:"$se_module,omitempty"`
		CreatedBy struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"Created_By,omitempty"`
		CreatedTime string `json:"Created_Time,omitempty"`
		ModifiedBy  struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"Modified_By,omitempty"`
		ModifiedTime string `json:"Modified_Time,omitempty"`
		NoteContent  string `json:"Note_Content,omitempty"`
		NoteTitle    string `json:"Note_Title,omitempty"`
		Owner        struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"Owner,omitempty"`
		ParentID struct {
			ID   string      `json:"id,omitempty"`
			Name interface{} `json:"name,omitempty"`
		} `json:"Parent_Id,omitempty"`
		ID string `json:"id,omitempty"`
	} `json:"data,omitempty"`
	Info struct {
		Count       float64 `json:"count,omitempty"`
		MoreRecords bool    `json:"more_records,omitempty"`
		Page        float64 `json:"page,omitempty"`
		PerPage     float64 `json:"per_page,omitempty"`
	} `json:"info,omitempty"`
}

// GetNote returns the note specified by ID and module
// https://www.zoho.com/crm/help/api/v2/#get-spec-notes-data
func (c *API) GetNote(module crmModule, id string) (data NoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes", c.ZohoTLD, module, id),
		Method:       zoho.HTTPGet,
		ResponseData: &NoteResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return NoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*NoteResponse); ok {
		return *v, nil
	}

	return NoteResponse{}, fmt.Errorf("Data returned was not 'NoteResponse'")
}

type NoteResponse NotesResponse

// CreateNotes will create multiple notes provided in the request data
// https://www.zoho.com/crm/help/api/v2/#create-notes
func (c *API) CreateNotes(request CreateNoteData) (data CreateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateNoteResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateNoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateNoteResponse); ok {
		return *v, nil
	}

	return CreateNoteResponse{}, fmt.Errorf("Data returned was not 'CreateNoteResponse'")
}

type CreateNoteData struct {
	Data []struct {
		NoteContent string `json:"Note_Content,omitempty"`
		NoteTitle   string `json:"Note_Title,omitempty"`
		ParentID    string `json:"Parent_Id,omitempty"`
		SeModule    string `json:"se_module,omitempty"`
	} `json:"data,omitempty"`
}

type CreateNoteResponse struct {
	Data []struct {
		Code    string `json:"code,omitempty"`
		Details struct {
			CreatedBy struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"created_by,omitempty"`
			CreatedTime string `json:"created_time,omitempty"`
			ID          string `json:"id,omitempty"`
			ModifiedBy  struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"modified_by,omitempty"`
			ModifiedTime string `json:"modified_time,omitempty"`
		} `json:"details,omitempty"`
		Message string `json:"message,omitempty"`
		Status  string `json:"status,omitempty"`
	} `json:"data,omitempty"`
}

// UpdateNote will update the note data of the specified note on the specified record of the module
// https://www.zoho.com/crm/help/api/v2/#update-notes
func (c *API) UpdateNote(request UpdateNoteData, module crmModule, recordID string, noteID string) (data UpdateNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes/%s", c.ZohoTLD, module, recordID, noteID),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateNoteResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateNoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateNoteResponse); ok {
		return *v, nil
	}

	return UpdateNoteResponse{}, fmt.Errorf("Data returned was not 'UpdateNoteResponse'")
}

type UpdateNoteData CreateNoteData

type UpdateNoteResponse CreateNoteResponse

// DeleteNote will delete the specified note on the specified record from the module
// https://www.zoho.com/crm/help/api/v2/#delete-notes
func (c *API) DeleteNote(module crmModule, recordID string, noteID string) (data DeleteNoteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/Notes/%s", c.ZohoTLD, module, recordID, noteID),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteNoteResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteNoteResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteNoteResponse); ok {
		return *v, nil
	}

	return DeleteNoteResponse{}, fmt.Errorf("Data returned was not 'DeleteNoteResponse'")
}

type DeleteNoteResponse struct {
	Data []struct {
		Code    string `json:"code,omitempty"`
		Details struct {
			ID string `json:"id,omitempty"`
		} `json:"details,omitempty"`
		Message string `json:"message,omitempty"`
		Status  string `json:"status,omitempty"`
	} `json:"data,omitempty"`
}

// DeleteNotes will delete all notes specified in the IDs
// https://www.zoho.com/crm/help/api/v2/#delete-bulk-notes
func (c *API) DeleteNotes(IDs ...string) (data DeleteNotesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "notes",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteNotesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"IDs": "...string",
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteNotesResponse{}, fmt.Errorf("Failed to retrieve notes: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteNotesResponse); ok {
		return *v, nil
	}

	return DeleteNotesResponse{}, fmt.Errorf("Data returned was not 'DeleteNotesResponse'")
}

type DeleteNotesResponse DeleteNoteResponse