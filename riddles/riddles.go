package riddles

import (
	"math/rand"
	"time"
)

var riddles = [...]string{
	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
	"I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?",
	"What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?",
	"The more you have of it, the less you see. What is it?",
	"I am always hungry, I must always be fed, The finger I touch, Will soon turn red",
	"What’s black when you get it, red when you use it, and white when you’re all through with it?",
	"Forward I’m heavy, but backwards I’m not. What am I?",
	"What is it that after you take away the whole, some still remains?",
	"I’m full of holes, yet I’m full of water. What am I?",
}

var answers = map[string]string{
	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?":                                "man",
	"I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?": "letter e",
	"What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?":            "river",
	"The more you have of it, the less you see. What is it?":                                                                           "darkness",
	"I am always hungry, I must always be fed, The finger I touch, Will soon turn red":                                                 "fire",
	"What’s black when you get it, red when you use it, and white when you’re all through with it?":                                    "charcoal",
	"Forward I’m heavy, but backwards I’m not. What am I?":                                                                             "ton",
	"What is it that after you take away the whole, some still remains?":                                                               "wholesome",
	"I’m full of holes, yet I’m full of water. What am I?":                                                                             "sponge",
}

func Ask() (q, a string) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(riddles))
	q = riddles[n]
	a = answers[q]
	return q, a
}
