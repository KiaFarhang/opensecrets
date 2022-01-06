package opensecrets

import "testing"

func assertErrorExists(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Fatalf("Wanted error but got nil")
	}
}

func assertErrorMessage(err error, expectedMessage string, t *testing.T) {
	t.Helper()
	if err.Error() != expectedMessage {
		t.Fatalf("Wanted error message %s but got %s", expectedMessage, err.Error())
	}

}

func assertStringMatches(got, wanted string, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Got string %s wanted string %s", got, wanted)
	}
}

func assertSliceLength(got, wanted int, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Got slice length %d wanted %d", got, wanted)
	}
}
