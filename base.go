package baseemoji

import (
	"encoding/hex"
	"strings"
)

// Encode a string to emojis
func Encode(plain string) (emojis string) {
	e := []string{}
	for _, b := range []byte(plain) {
		e = append(e, emojiList[b].char)
	}
	return strings.Join(e, " ")
}

// Decode a string of emoji characters
func Decode(emojis string) (plain string) {
	text := []byte{}
	for _, c := range emojis {
		c := string(c)
		for i := 0; i < len(emojiList); i++ {
			if c == emojiList[i].char {
				text = append(text, byte(i))
			}
		}
	}
	decoded, err := hex.DecodeString(hex.EncodeToString(text))
	if err != nil {
		panic(err)
	}
	return string(decoded)
}
