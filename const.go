package cicd_feedback

const (
	StatusSkipped  Status = "skipped"
	StatusPending  Status = "pending"
	StatusRunning  Status = "running"
	StatusSuccess  Status = "success"
	StatusFailed   Status = "failed"
	StatusKilled   Status = "killed"
	StatusManual   Status = "manual"
	StatusDeclined Status = "declined"
)

const (
	// ErrorInternal: CI/CD engine encountered an unexpected internal error
	ErrorInternal ErrorType = "internal"
	// ErrorConfig: Error in the pipeline configuration or workflow definition
	ErrorConfig ErrorType = "config"
	// ErrorExternal: Error in an external service that the CI/CD engine depends on
	ErrorExternal ErrorType = "external"
	// ErrorPermission: Insufficient permissions to perform the requested action
	ErrorPermission ErrorType = "permission"
	// ErrorValidation: Input validation errors
	ErrorValidation ErrorType = "validation"
	// ErrorOther: All other unspecified error
	ErrorOther ErrorType = "other"
)

const (
	HeaderFeedback      = "CICD-Feedback"
	HeaderAuthorization = "CICD-Authorization"
)
