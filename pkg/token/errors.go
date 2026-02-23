package token

import "errors"

type tokenErrors struct {
	InvalidToken error
	MakingToken  error
	//Internal     error
}

var Errors = tokenErrors{

	InvalidToken: errors.New("Invalid token"),
	MakingToken:  errors.New("Error making token"),
	//Internal:     errors.New("Internal error"),
}
