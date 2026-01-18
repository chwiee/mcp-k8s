package diag

type Result struct {
	Severity        string
	ErrorType       string
	Summary         string
	Details         []string
	Recommendations []string
}
