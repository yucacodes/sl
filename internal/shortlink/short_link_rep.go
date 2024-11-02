//Short link repositorio. Persistencia de los links acortados

package shortlink

type ShortLinkRep interface {
	SaveNew(shortlink *ShortLink) error
	FindById(id string) (*ShortLink, error)
	FindByShortCode(shortCode string) (*ShortLink, error)
	SaveUpdate(shortlink *ShortLink) error
}
