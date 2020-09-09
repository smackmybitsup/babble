package babble

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Babbler struct {
	Count     int
	Separator string
	Words     []string
	Wordfield int
}

func NewBabbler() (b Babbler) {
	b.Count = 2
	b.Wordfield = 10
	b.Separator = "-"
	b.Words = readAvailableDictionary()
	return
}

func (this Babbler) Babble() string {
	pieces := []string{}

	for i := 0; i < this.Count; i++ {
		if len(pieces) > 5 {
			break
		}
		pieces = append(pieces, this.Words[rand.Int()%len(this.Words)])
	}

	return strings.Join(pieces, this.Separator)
}
