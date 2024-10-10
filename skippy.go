package skippy

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	defaultMaxHeight = 3
)

type SkipListOpt func(sl *skipList)

// node represents a skiplist node
type node struct {
	key        string
	value      []byte
	successors []*node
}

type skipList struct {
	// maxHeight defines the m
	maxHeight int

	// height is the current highest most level of the skiplist at any point in time.
	// it cannot be any greater than maxLevel.
	height        int
	pDistribution float32

	headers []*node

	rand *rand.Rand
}

func WithMaxHeights(height int) SkipListOpt {
	return func(sl *skipList) {
		sl.maxHeight = height
	}
}

func WithDistribution(pDistribution float32) SkipListOpt {
	return func(sl *skipList) {
		sl.pDistribution = pDistribution
	}
}

func WithRandSource(source rand.Source) SkipListOpt {
	return func(sl *skipList) {
		sl.rand = rand.New(source)
	}
}

func New(opts ...SkipListOpt) *skipList {
	sl := &skipList{
		maxHeight: defaultMaxHeight,
		height:    0,
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, opt := range opts {
		opt(sl)
	}

	sl.headers = make([]*node, sl.maxHeight)

	return sl
}

func (s *skipList) randomLevel() int {
	level := 1

	for (s.rand.Float32() <= s.pDistribution) && (level < s.maxHeight) {
		level += 1
	}

	return level
}

func (s *skipList) Insert(value int) {
	//level := s.randomLevel()
}

func (s *skipList) String() string {
	levelstrs := make([]string, s.maxHeight)

	keymap := map[string]int{}

	// process first level (height 0) first
	str := fmt.Sprintf("(%d)", 0)

	successor := s.headers[0]
	for successor != nil {
		str += fmt.Sprintf(" --> %s", successor.key)
		keymap[successor.key] = len(str) - len(successor.key) - 1

		successor = successor.successors[0]
	}

	str += " --> nil"
	nilpos := len(str) - len(" nil") - 1

	levelstrs = append(levelstrs, str)

	// now proceed to use the index position of each key to
	// print the other keys at higher levels based on those
	// indexes
	for level := 1; level < len(s.headers); level++ {
		str := fmt.Sprintf("(%d)", level)

		node := s.headers[level]
		if node != nil {
			headIndex := keymap[node.key]

			str += " "
			for i := len(str); i < headIndex-len(node.key); i++ {
				str += "-"
			}
			str += fmt.Sprintf("> %s", node.key)

			successor := node.successors[level]
			for successor != nil {
				successorIndex := keymap[successor.key]

				str += " "
				for i := len(str); i < successorIndex-len(successor.key); i++ {
					str += "-"
				}
				str += fmt.Sprintf("> %s", successor.key)

				successor = successor.successors[level]
			}
		}

		str += " "
		for i := len(str); i < nilpos; i++ {
			str += "-"
		}
		str += "> nil"

		levelstrs = append(levelstrs, str)
	}

	var finalstr string
	for i := len(levelstrs) - 1; i >= 0; i-- {
		finalstr += levelstrs[i] + "\n"
	}
	return finalstr
}
