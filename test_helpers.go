package opensecrets

import "testing"

func assertErrorExists(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Fatalf("Wanted error but got nil")
	}
}

func assertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Fatalf("Wanted no error but got %s", err.Error())
	}
}

func assertErrorMessage(err error, expectedMessage string, t *testing.T) {
	t.Helper()
	if err.Error() != expectedMessage {
		t.Fatalf("Got error message %s but wanted %s", err.Error(), expectedMessage)
	}

}

func assertStringMatches(got, wanted string, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Got string %s wanted %s", got, wanted)
	}
}

func assertIntMatches(got, wanted int, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Got int %d wanted %d", got, wanted)
	}
}

func assertSliceLength(got, wanted int, t *testing.T) {
	t.Helper()
	if got != wanted {
		t.Fatalf("Got slice length %d wanted %d", got, wanted)
	}
}
