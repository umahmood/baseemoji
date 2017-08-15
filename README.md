# Base Emoji

Base emoji is a Go library which encodes/decodes text to and from emojis. For example encoding the string "i love peanut butter" in base emoji becomes "👂 🔋 👊 🔦 🎲 🐬 🔋 💾 🐬 🌳 🔥 🐸 🍟 🔋 💫 🐸 🍟 🍟 🐬 🍀"

# Installation

> $ go get github.com/umahmood/baseemoji

# Terminal Usage
```
$ baseemoji -e "hello world"

💧 🐬 👊 👊 🔦 🔋 💎 🔦 🍀 👊 💵

$ baseemoji -d "💧 🐬 👊 👊 🔦 🔋 💎 🔦 🍀 👊 💵"

hello world
```

# Library Usage
```
package main

import (
    "fmt"
    
    "github.com/umahmood/baseemoji
)

func main() {
    e := baseemoji.EncodeString("the quick brown fox jumped over the lazy dog.")
    fmt.Println(e)
    d := baseemoji.DecodeString(e)
    fmt.Println(d)  
}
```
Output:
```
🍟 💧 🐬 🔋 🏈 🐸 👂 🐕 👓 🔋 💫 🍀 🔦 💎 🔥 🔋 🍩 🔦 👻 🔋 🌍 🐸 👣 💾 🐬 💵 🔋 🔦 🎲 🐬 🍀 🔋 🍟 💧 🐬 🔋 👊 🌳 🌐 🎁 🔋 💵 🔦 🚪

the quick brown fox jumped over the lazy dog.
```

# Documentation

> http://godoc.org/github.com/umahmood/baseemoji

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
