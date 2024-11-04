package shortlink

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"path/filepath"
)

type ShortLinkRepFiles struct {
	folderPath string
}

func NewShortLinkRepFiles(folderPath string) (*ShortLinkRepFiles, error) {
	stat, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0700)
		if err != nil {
			return nil, err
		}
	} else if !stat.IsDir() {
		return nil, errors.New("folfer path should be a directory")
	}
	shortLinkRepFiles := ShortLinkRepFiles{
		folderPath: folderPath,
	}
	return &shortLinkRepFiles, nil
}

func (p *ShortLinkRepFiles) SaveNew(shortLink *ShortLink) error {
	saveFilePathId := p.pathForIdJsonFiles(shortLink.id)
	saveFilePathSc := p.pathForShortCodeJsonFiles(shortLink.shortCode)

	auxShortLink := struct {
		LongLink  string
		ShortCode string
		Id        string
	}{
		LongLink:  shortLink.longLink.String(),
		ShortCode: shortLink.shortCode,
		Id:        shortLink.id,
	}

	data, error := json.Marshal(auxShortLink)
	if error != nil {
		return error
	}

	error = os.WriteFile(saveFilePathId, data, 0644)
	if error != nil {
		return error
	}

	error = os.WriteFile(saveFilePathSc, data, 0644)
	if error != nil {
		return error
	}
	return nil
}

func (p *ShortLinkRepFiles) SaveUpdate(shortLink *ShortLink) error {
	error := p.Delete(shortLink.id)
	if error != nil {
		return error
	}

	error = p.SaveNew(shortLink)
	if error != nil {
		return error
	}
	return nil
}

func (p *ShortLinkRepFiles) Delete(id string) error {
	shortLink, error := p.FindById(id)
	if error != nil {
		return error
	}

	error = os.Remove(p.pathForIdJsonFiles(shortLink.id))
	if error != nil {
		return error
	}

	error = os.Remove(p.pathForShortCodeJsonFiles(shortLink.shortCode))
	if error != nil {
		return error
	}
	return nil
}

// Encontrar archivos json usando el ID para buscarlos
func (p *ShortLinkRepFiles) FindById(id string) (*ShortLink, error) {
	shortLinkData, error := os.ReadFile(p.pathForIdJsonFiles(id))
	if error != nil {
		return nil, error
	}

	type AuxShortLink struct {
		LongLink  string
		ShortCode string
		Id        string
	}
	var auxShortLink AuxShortLink

	error = json.Unmarshal(shortLinkData, &auxShortLink)
	if error != nil {
		return nil, error
	}

	longLing, err := url.Parse(auxShortLink.LongLink)
	if err != nil {
		return nil, err
	}

	shortLink := ShortLink{
		longLink:  *longLing,
		id:        auxShortLink.Id,
		shortCode: auxShortLink.ShortCode,
	}
	return &shortLink, nil
}

// Encontrar archivos json usando el shortCode para buscarlos
func (p *ShortLinkRepFiles) FindByShortCode(shortCode string) (*ShortLink, error) {
	shortLinkData, error := os.ReadFile(p.pathForShortCodeJsonFiles(shortCode))
	if error != nil {
		return nil, error
	}

	type AuxShortLink struct {
		LongLink  string
		ShortCode string
		Id        string
	}
	var auxShortLink AuxShortLink

	error = json.Unmarshal(shortLinkData, &auxShortLink)
	if error != nil {
		return nil, error
	}

	longLing, err := url.Parse(auxShortLink.LongLink)
	if err != nil {
		return nil, err
	}

	shortLink := ShortLink{
		longLink:  *longLing,
		id:        auxShortLink.Id,
		shortCode: auxShortLink.ShortCode,
	}
	return &shortLink, nil
}

// Metodo para determinar el path de los archivos JSON guardados con el ID
func (p *ShortLinkRepFiles) pathForIdJsonFiles(id string) string {
	idPath := filepath.Join(p.folderPath, "_id_"+id+".json")
	return idPath
}

// Metodo para determinar el path de los archivos JSON guardados con el shortCode
func (p *ShortLinkRepFiles) pathForShortCodeJsonFiles(shortCode string) string {
	shortCodePath := filepath.Join(p.folderPath, "_short_code_"+shortCode+".json")
	return shortCodePath
}
