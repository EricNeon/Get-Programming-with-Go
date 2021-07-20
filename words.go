package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,"-`)
		frequency[word]++
	}
	return frequency
}

func main() {
	text := `From fairest creatures we desire increase,

That thereby beauty's rose might never die,

But as the riper should by time decease,

His tender heir might bear his memory;

But thou contracted to thine own bright eyes,

Feed'st thy light's flame with self-substantial fuel,

Making a famine where abundance lies,

Thyself thy foe, to thy sweet self too cruel.

Thou that art now the world's fresh ornament,

And only herald to the gaudy spring,

Within thine own bud buriest thy content,

And, tender churl, mak'st waste in niggarding.

 Pity the world, or else this glutton be,

 To eat the world's due, by the grave and thee.`

	frequency := countWords(text)
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%v 出现了%d次\n", word, count)
		}
	}
}
