package funcs

import (
	"testing"
)

func TestReplaceLinks(t *testing.T) {
	got := ReplaceLinks("back to the [FrontPage] or [AnotherPage]")
	want := "back to the <a href=\"/view/FrontPage\">FrontPage</a> or <a href=\"/view/AnotherPage\">AnotherPage</a>"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
