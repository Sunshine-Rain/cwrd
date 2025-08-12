package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type inputConfig struct {
	startAlpha        string
	endAlpha          string
	eachWordCount     int
	eachLineWordCount int
	lineCount         int
	charSet           string
	prefix            string
	suffix            string
}

type Globals struct {
	Start     string `name:"start" short:"s" default:"a" help:"Start char in alphabet, default is a."`
	End       string `name:"end" short:"e" default:"z" help:"End char in alphabet, default is z."`
	CharCount int    `name:"charCount" short:"c" default:5 help:"Chars count in each word, default is 5."`
	WordCount int    `name:"wordCount" short:"w" default:10 help:"Words count in each word, default is 10."`
	Lines     int    `name:"lines" short:"l" default:10 help:"Total lines, default is 10."`
	CharSet   string `name:"charSet" default:"" help:"Input character set, if not empty only use character in charset, default is empty."`
	Prefix    string `name:"prefix" short:"p" default:"" help:"Output prefix, default is empty."`
	Suffix    string `name:"suffix" short:"x" default:"" help:"Output suffix, default is empty."`
}

type RandCommand struct {
}

func (r *RandCommand) Run(g *Globals) error {
	// fmt.Println(g.Start)
	// fmt.Println(g.End)
	// fmt.Println(g.CharCount)
	// fmt.Println(g.WordCount)
	// fmt.Println(g.Lines)
	// fmt.Println(g.CharSet)

	config := inputConfig{
		startAlpha:        g.Start,
		endAlpha:          g.End,
		eachWordCount:     g.CharCount,
		eachLineWordCount: g.WordCount,
		lineCount:         g.Lines,
		charSet:           g.CharSet,
		prefix:            g.Prefix,
		suffix:            g.Suffix,
	}

	if len(g.CharSet) > 0 {
		generateByCharSet(&config)
		return nil
	}

	generate(&config)
	return nil
}

func generate(config *inputConfig) {
	startInt := int(config.startAlpha[0])
	endInt := int(config.endAlpha[0])
	rangeInt := endInt - startInt
	count := config.eachWordCount * config.eachLineWordCount * config.lineCount

	var result strings.Builder
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result.WriteString(config.prefix)
	for index := range count {
		val := rand.Intn(rangeInt + 1)
		r := startInt + val

		result.WriteByte(byte(r))

		if index != 0 && (index+1)%config.eachWordCount == 0 {
			result.WriteString(" ")
		}

		if index != 0 && (index+1)%(config.eachWordCount*config.eachLineWordCount) == 0 {
			result.WriteString("\n")
		}
	}
	result.WriteString(config.suffix)

	fmt.Println(result.String())
}

func generateByCharSet(config *inputConfig) {
	charSet := config.charSet
	count := config.eachWordCount * config.eachLineWordCount * config.lineCount
	charSetCount := len(charSet)

	var result strings.Builder
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result.WriteString(config.prefix)
	for index := range count {
		val := rand.Intn(charSetCount)

		result.WriteByte(charSet[val])

		if index != 0 && (index+1)%config.eachWordCount == 0 {
			result.WriteString(" ")
		}

		if index != 0 && (index+1)%(config.eachWordCount*config.eachLineWordCount) == 0 {
			result.WriteString("\n")
		}
	}
	result.WriteString(config.suffix)

	fmt.Println(result.String())
}
