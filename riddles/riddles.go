package riddles

import (
	"math/rand"
	"time"
)

// var riddles = [...]string{
// 	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
// 	"I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?",
// 	"What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?",
// 	"The more you have of it, the less you see. What is it?",
// 	"I am always hungry, I must always be fed, The finger I touch, Will soon turn red",
// 	"What’s black when you get it, red when you use it, and white when you’re all through with it?",
// 	"Forward I’m heavy, but backwards I’m not. What am I?",
// 	"What is it that after you take away the whole, some still remains?",
// 	"I’m full of holes, yet I’m full of water. What am I?",
// 	"I build up castles. I tear down mountains. I make some men blind, I help others to see. What am I?",
// 	"The more you take, the more you leave behind.",
// 	"Say my name and I disappear. What am I?",
// }

var riddles = map[string]string{
	"man":       "It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
	"letter e":  "I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?",
	"river":     "What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?",
	"darkness":  "The more you have of it, the less you see. What is it?",
	"fire":      "I am always hungry, I must always be fed, The finger I touch, Will soon turn red",
	"charcoal":  "What’s black when you get it, red when you use it, and white when you’re all through with it?",
	"ton":       "Forward I’m heavy, but backwards I’m not. What am I?",
	"wholesome": "What is it that after you take away the whole, some still remains?",
	"sponge":    "I’m full of holes, yet I’m full of water. What am I?",
	"silence":   "Say my name and I disappear. What am I?",
	"footsteps": "The more you take, the more you leave behind.",
	"sand":      "I build up castles. I tear down mountains. I make some men blind, I help others to see. What am I?",
	"teapot":    "What starts with a T, ends with a T, and has T in it?",
	"shadow":    "Each morning I appear to lie at your feet, All day I will follow no matter how fast you run, Yet I nearly perish in the midday sun.",
	"few":       "I know a word of letters three. Add two, and fewer there will be!",
	"Priest":    "He has married many women, but has never been married. Who is he?",
	"Owl":       "What asks but never answers?",
	"david":     "David's father has three sons : Snap, Crackle and who?",
	"candle":    "Tall I am young, Short I am old, While with life I glow, Wind is my foe. What am I?",
	"name":      "What belongs to you, but other people use it more than you?",
	"gravity":   "Everyone is attracted to me and everybody falls for me. What am I?",
}

func Ask() (a, q string) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(riddles))
	return randomRiddle(n)
}

func randomRiddle(n int) (a, q string) {
	for key, val := range riddles {
		if n == 0 {
			return key, val
		}

		n--
	}

	return "man", riddles["man"]
}

func GetRiddle(s string) string {
	return riddles[s]
}
