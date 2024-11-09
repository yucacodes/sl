package shortlink

import (
	"math/rand"
	"net/url"

	"github.com/google/uuid"
)

// Almacena los 62 caracteres para generar el shortCode
var base62 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// Structura que almacena la información del link acortado
type ShortLink struct {
	longLink  url.URL
	shortCode string
	id        string
}

func (sl *ShortLink) LongLink() *url.URL {
	return &sl.longLink
}

func (sl *ShortLink) ShortCode() string {
	return sl.shortCode
}

func (sl *ShortLink) ID() string {
	return sl.id
}

func (sl *ShortLink) SetLongLing(longLink *url.URL) *ShortLink {
	sl.longLink = *longLink
	return sl
}

func (sl *ShortLink) Equal(another *ShortLink) bool {
	if sl.id != another.id {
		return false
	}
	if sl.shortCode != another.shortCode {
		return false
	}
	if sl.longLink.String() != another.longLink.String() {
		return false
	}
	return true
}

func NewShortLink(longLink *url.URL) *ShortLink {
	var shortLink ShortLink
	shortCode := make([]rune, 4)
	for i := range shortCode {
		shortCode[i] = base62[rand.Intn(len(base62))]
	}
	shortLink.shortCode = string(shortCode)
	shortLink.id = uuid.New().String()
	shortLink.longLink = *longLink

	return &shortLink
}
