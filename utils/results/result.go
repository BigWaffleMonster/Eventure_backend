package results

type Result struct {
	IsFailed bool
	IsSuccess bool
	Code int
	Errors []string
}

func (r *Result) AddError(err string) {
	r.Errors = append(r.Errors, err)

	r.IsFailed = true
	r.IsSuccess = false

	r.Code = 400
}

func (r *Result) AddErrorAndNewCode(err string, code int) {
	r.Errors = append(r.Errors, err)

	r.IsFailed = true
	r.IsSuccess = false

	r.Code = code
}

func (r *Result) Merge(result Result) Result{
	return Result{
		IsSuccess: r.IsSuccess && result.IsSuccess,
		IsFailed: r.IsFailed && result.IsFailed,
		Code: max(r.Code, result.Code),
		Errors: append(r.Errors, result.Errors...),
	}
}