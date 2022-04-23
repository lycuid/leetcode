// https://leetcode.com/problems/encode-and-decode-tinyurl/
package main

import "fmt"

type Codec struct {
	pattern string
	urls    []string
}

func Constructor() Codec {
	return Codec{urls: make([]string, 0), pattern: "https://tiny.url/%d"}
}

func (c *Codec) encode(longUrl string) string {
	c.urls = append(c.urls, longUrl)
	return fmt.Sprintf(c.pattern, len(c.urls)-1)
}

func (c *Codec) decode(shortUrl string) string {
	var index int
	fmt.Sscanf(shortUrl, c.pattern, &index)
	return c.urls[index]
}

func main() {}
