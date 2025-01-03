package solution535

// ============================================================================
// 535. Encode and Decode TinyURL
// URL: https://leetcode.com/problems/encode-and-decode-tinyurl/
// ============================================================================

import (
	"crypto/sha1"
	"encoding/hex"
)

type Codec struct {
	urls map[string]string
}

func Constructor() Codec {
	return Codec{
		urls: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	h := sha1.New()
	h.Write([]byte(longUrl))
	hash := hex.EncodeToString(h.Sum(nil))
	shortUrl := "http://tinyurl.com/" + string(hash)
	c.urls[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	url, ok := c.urls[shortUrl]
	if ok {
		return url
	}
	panic("something went wrong...")
	return ""
}
