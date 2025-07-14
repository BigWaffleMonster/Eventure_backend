package results

import "fmt"

func NewResultOk() Result{
	return Result{
		IsSuccess: true,
		IsFailed: false,
		Code: 200,
		Errors: nil,
	}
}

func NewResultFailed(err string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 400,
		Errors: []string{
			err,
		},
	}
}

func NewBadRequestError(err string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 400,
		Errors: []string{
			err,
		},
	}
}

func NewUnauthorizedError(err string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 401,
		Errors: []string{
			err,
		},
	}
}

func NewForbiddenError() Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 403,
		Errors: []string{
			"Access denied",
		},
	}
}

func NewNotFoundError(err string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 404,
		Errors: []string{
			err,
		},
	}
}

func NewConflictError(err string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 409,
		Errors: []string{
			err,
		},
	}
}

func NewInternalError(entity string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 500,
		Errors: []string{
			fmt.Sprintf("%s was not found", entity),
		},
	}
}

func NewInvalidDomainTypeError(domainType string, expectedType string) Result{
	return Result{
		IsSuccess: false,
		IsFailed: true,
		Code: 500,
		Errors: []string{
			fmt.Sprintf("Invalid domain type, actual: %s, expected: %s", domainType, expectedType),
		},
	}
}
