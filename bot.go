package asciibot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ErrIDLength = fmt.Errorf("id length must be 5")
	ErrIDHex    = fmt.Errorf("id must be 5 hexadecimal characters [0-9a-f]")
)

func Random() string {
	out, _ := Generate(RandomID())
	return out
}

func RandomID() string {
	out := ""
	for i := 0; i < 5; i++ {
		out += fmt.Sprintf("%x", rand.Intn(16)) // 16 because hex
	}
	return out
}

func MustGenerate(id string) string {
	out, err := Generate(id)
	if err != nil {
		panic(err)
	}
	return out
}

func Generate(id string) (string, error) {
	if len(id) != 5 {
		return "", ErrIDLength
	}

	id = strings.Map(hexOnly, id)
	if len(id) != 5 {
		return "", ErrIDHex
	}

	out := ""

	// generate body
	top, _, _ := split(templates[id[0:1]])
	_, center, _ := split(templates[id[1:2]])
	_, _, bottom := split(templates[id[2:3]])
	out += top
	out += center
	out += bottom

	// replace eyes
	out = replace(out, eyes[id[3:4]], 6, 1)

	// replace mouth
	out = replace(out, mouths[id[4:5]], 7, 2)

	return out, nil
}

func hexOnly(r rune) rune {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return r
	case 'a', 'b', 'c', 'd', 'e', 'f':
		return r
	default:
		return -1
	}
}

// split splits template into top, center and bottom parts
func split(template string) (top, center, bottom string) {
	s := strings.Split(template, "\n")
	top = strings.Join(s[0:3], "\n") + "\n"
	center = strings.Join(s[3:5], "\n") + "\n"
	bottom = strings.Join(s[5:7], "\n")
	return
}

// replace replaces body parts at position x,y
func replace(body string, replace string, x, y int) string {
	lines := strings.Split(body, "\n")
	lines[y] = lines[y][0:x] + replace + lines[y][x+len(replace):]
	return strings.Join(lines, "\n")
}
