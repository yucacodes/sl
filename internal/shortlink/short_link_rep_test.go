package shortlink_test

import (
	"net/url"
	"testing"

	"github.com/yucacodes/sl/internal/shortlink"
)

func deepEquals(a *shortlink.ShortLink, b *shortlink.ShortLink) bool {
	if a.ID() != b.ID() {
		return false
	}
	if a.ShortCode() != b.ShortCode() {
		return false
	}
	if a.LongLink().String() != b.LongLink().String() {
		return false
	}
	return true
}

func RunShortLinkRepTest(t *testing.T, rep shortlink.ShortLinkRep) {

	t.Run("save new and then find should retrieve saved short links", func(t *testing.T) {
		url_a, _ := url.Parse("http://very.long.url/a")
		url_b, _ := url.Parse("http://very.long.url/b")
		shortLink_a := shortlink.NewShortLink(url_a)
		shortLink_b := shortlink.NewShortLink(url_b)

		// Test Save New
		err := rep.SaveNew(shortLink_a)
		if err != nil {
			t.Error(err.Error())
			return
		}
		err = rep.SaveNew(shortLink_b)
		if err != nil {
			t.Error(err.Error())
			return
		}

		// Test Find By Id
		foundShortLink_a, err := rep.FindById(shortLink_a.ID())
		if err != nil {
			t.Error(err.Error())
			return
		}
		if !deepEquals(shortLink_a, foundShortLink_a) {
			t.Error("Found short link with FindById is not equals to saved one")
			return
		}

		// Test Find By Short Code
		foundShortLink_b, err := rep.FindByShortCode(shortLink_b.ShortCode())
		if err != nil {
			t.Error(err.Error())
			return
		}
		if !deepEquals(shortLink_b, foundShortLink_b) {
			t.Error("Found short link with FindByShortCode is not equals to saved one")
			return
		}

	})

	t.Run("save new and then update should persist changes", func(t *testing.T) {
		url_a, _ := url.Parse("http://very.long.url/a")
		url_b, _ := url.Parse("http://very.long.url/b")
		shortLink_a := shortlink.NewShortLink(url_a)

		// Test Save New
		err := rep.SaveNew(shortLink_a)
		if err != nil {
			t.Error(err.Error())
			return
		}

		shortLink_a.SetLongLing(url_b)

		// Test Save Update
		err = rep.SaveUpdate(shortLink_a)
		if err != nil {
			t.Error(err.Error())
			return
		}

		// Test Find By Id
		foundShortLink_a, err := rep.FindById(shortLink_a.ID())
		if err != nil {
			t.Error(err.Error())
			return
		}
		if !deepEquals(shortLink_a, foundShortLink_a) {
			t.Error("Found short link with FindById is not equals to saved one")
			return
		}
	})

}
