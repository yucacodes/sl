package shortlink_test

import (
	"testing"

	"github.com/yucacodes/sl/internal/shortlink"
)

func TestShortLinkRepFiles(t *testing.T) {
	rep, err := shortlink.NewShortLinkRepFiles("./.tests/short_link_repository_files_dir")
	if err != nil {
		t.Error(err.Error())
		return
	}

	shortlink.RunShortLinkRepTest(t, rep)
}
