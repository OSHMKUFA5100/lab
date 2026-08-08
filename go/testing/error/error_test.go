package error

import (
	"errors"
	"fmt"
	"testing"
)

type MatchError struct {
	code int
	err  error
}

func (e *MatchError) Error() string {
	return fmt.Sprintf("code: %d, error: %v", e.code, e.err)
}
func GetMatchError() error {
	return &MatchError{code: 1, err: errors.New("123")}
}
func GetError() error {
	return errors.New("123")
}
func TestError(t *testing.T) {
	if err, ok := errors.AsType[*MatchError](GetMatchError()); ok {
		t.Log(err)
	} else {
		t.Errorf("expected MatchError, got %T", GetMatchError())
	}
	if err, ok := errors.AsType[*MatchError](GetError()); ok {
		t.Log(err)
	} else {
		t.Errorf("expected MatchError, got %T", GetError())
	}
}
