package riddles

import (
	"math/rand"
	"time"
)

var riddles = [...]string{
	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
	"I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?",
	"What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?",
}

var answers = map[string]string{
	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?":                                "man",
	"I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?": "the letter e",
	"What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?":            "river",
}

func Ask() (q, a string) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(riddles))
	q = riddles[n]
	a = answers[q]
	return q, a
}
