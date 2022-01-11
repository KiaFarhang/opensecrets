package test

import "testing"

func AssertErrorExists(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Error("Wanted error but got nil")
	}
}

func AssertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Wanted no error but got %s", err.Error())
	}
}

func AssertErrorMessage(err error, expectedMessage string, t *testing.T) {
	t.Helper()
	if err.Error() != expectedMessage {
		t.Errorf("Got error message %s but wanted %s", err.Error(), expectedMessage)
	}

}

func AssertStringMatches(got, wanted string, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Errorf("Got string %s wanted %s", got, wanted)
	}
}

func AssertIntMatches(got, wanted int, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Errorf("Got int %d wanted %d", got, wanted)
	}
}

func AssertSliceLength(got, wanted int, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Errorf("Got slice length %d wanted %d", got, wanted)
	}
}
