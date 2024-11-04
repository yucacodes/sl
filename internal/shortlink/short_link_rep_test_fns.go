package shortlink

import (
	"net/url"
	"testing"
)

func RunShortLinkRepTest(t *testing.T, rep ShortLinkRep) {

	t.Run("save new and then find should retrieve saved short links", func(t *testing.T) {
		url_a, _ := url.Parse("http://very.long.url/a")
		url_b, _ := url.Parse("http://very.long.url/b")
		shortLink_a := NewShortLink(url_a)
		shortLink_b := NewShortLink(url_b)

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
		if !shortLink_a.Equal(foundShortLink_a) {
			t.Error("Found short link with FindById is not equals to saved one")
			return
		}

		// Test Find By Short Code
		foundShortLink_b, err := rep.FindByShortCode(shortLink_b.ShortCode())
		if err != nil {
			t.Error(err.Error())
			return
		}
		if !shortLink_b.Equal(foundShortLink_b) {
			t.Error("Found short link with FindByShortCode is not equals to saved one")
			return
		}

	})

	t.Run("save new and then update should persist changes", func(t *testing.T) {
		url_a, _ := url.Parse("http://very.long.url/a")
		url_b, _ := url.Parse("http://very.long.url/b")
		shortLink_a := NewShortLink(url_a)

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
		if !shortLink_a.Equal(foundShortLink_a) {
			t.Error("Found short link with FindById is not equals to saved one")
			return
		}
	})

}
