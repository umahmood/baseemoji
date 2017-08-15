/*
Package baseemoji encodes/decodes text to and from emojis.

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
*/
package baseemoji
