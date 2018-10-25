package isheadless

func TestHeaded(t *testing.T) {
	if IsHeadless() {
		t.Error("Expected IsHeadless to return false on Windows always.")
	}
}
