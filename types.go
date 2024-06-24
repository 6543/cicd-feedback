package cicd_feedback

// PipelineResponse represents the structure of a successful pipeline response
type PipelineResponse struct {
	PipelineID           string     `json:"pipelineId"`
	Title                string     `json:"title"`
	Status               Status     `json:"status"`
	RequiresManualAction bool       `json:"requiresManualAction"`
	Workflows            []Workflow `json:"workflows"`
	ExternalURI          string     `json:"externalURI,omitempty"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error            ErrorType `json:"error"`
	ErrorDescription string    `json:"errorDescription,omitempty"`
}

// Status represents the possible status values
type Status string

// ErrorType represents the possible error types
type ErrorType string

type Outputs struct {
	Logs      []Log      `json:"logs,omitempty"`
	Artifacts []Artifact `json:"artifacts,omitempty"`
}

type Workflow struct {
	ID           string     `json:"id"`
	Name         string     `json:"name,omitempty"`
	Status       Status     `json:"status"`
	Dependencies []string   `json:"dependencies,omitempty"`
	Steps        []Step     `json:"steps,omitempty"`
	SubWorkflows []Workflow `json:"subWorkflows,omitempty"`
}

type Step struct {
	ID           string   `json:"id"`
	Name         string   `json:"name,omitempty"`
	Status       Status   `json:"status"`
	Inputs       Inputs   `json:"inputs,omitempty"`
	Outputs      Outputs  `json:"outputs,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}

type Inputs struct {
	Commands    []string          `json:"commands,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
}

type Artifact struct {
	Name     string `json:"name"`
	URI      string `json:"uri"`
	MimeType string `json:"mimeType"`
}

type Log struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type WellKnownResponse struct {
	Version      string `json:"version"`
	SoftwareName string `json:"softwareName"`
	Type         string `json:"type"`
}
