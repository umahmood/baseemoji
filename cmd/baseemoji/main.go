package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/umahmood/baseemoji"
)

var (
	encode string
	decode string
)

func init() {
	flag.Usage = func() {
		printUsage()
	}

	flag.StringVar(&encode, "e", "", "Encode a string of emojis.")
	flag.StringVar(&decode, "d", "", "Decode a string of emojis.")

	flag.Parse()

	if flag.NFlag() == 0 {
		os.Exit(1)
	}

	if flag.NFlag() > 1 {
		fmt.Println("to many command line flags, use '-help' for help.")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(usage)
	fmt.Println(examples)
}

const usage = `
usage:
    
    -e - Encode a string to emojis.
    -d - Decode a string of emojis.
`

const examples = `example usage:

$ baseemoji -e "hello world"

💧 🐬 👊 👊 🔦 🔋 💎 🔦 🍀 👊 💵

$ baseemoji -d "💧 🐬 👊 👊 🔦 🔋 💎 🔦 🍀 👊 💵 "

hello world

(notice the quotation marks around the emojis when decoding)
`

func main() {
	if encode != "" {
		fmt.Println(baseemoji.Encode(encode))
		return
	}

	if decode != "" {
		fmt.Println(baseemoji.Decode(decode))
		return
	}
}
