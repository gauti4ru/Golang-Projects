package controller

import (
	"crypto/sha256"
	"encoding/hex"
)

type UrlMapper struct {
	LongShortMap map[string]string
	ShortLongMap map[string]string
}

func (urlMap *UrlMapper) Post(url string) string {
	longurl, isPresnet := urlMap.LongShortMap[url]
	if !isPresnet {
		urlMap.ShortLongMap[url] = convertinHash(url)
		urlMap.LongShortMap[urlMap.ShortLongMap[url]] = url
		longurl = urlMap.ShortLongMap[url]
	}
	return longurl
}

func convertinHash(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
