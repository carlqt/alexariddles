package riddles

import (
	"math/rand"
	"time"
)

var riddles = map[string]string{
	"man":           "It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
	"letter e":      "I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?",
	"river":         "What always runs but never walks, often murmurs, never talks, has a bed but never sleeps, has a mouth but never eats?",
	"darkness":      "The more you have of it, the less you see. What is it?",
	"fire":          "I am always hungry, I must always be fed, The finger I touch, Will soon turn red",
	"charcoal":      "What’s black when you get it, red when you use it, and white when you’re all through with it?",
	"ton":           "Forward I’m heavy, but backwards I’m not. What am I?",
	"wholesome":     "What is it that after you take away the whole, some still remains?",
	"sponge":        "I’m full of holes, yet I’m full of water. What am I?",
	"silence":       "Say my name and I disappear. What am I?",
	"footsteps":     "The more you take, the more you leave behind.",
	"sand":          "I build up castles. I tear down mountains. I make some men blind, I help others to see. What am I?",
	"teapot":        "What starts with a T, ends with a T, and has T in it?",
	"shadow":        "Each morning I appear to lie at your feet, All day I will follow no matter how fast you run, Yet I nearly perish in the midday sun.",
	"few":           "I know a word of letters three. Add two, and fewer there will be!",
	"Priest":        "He has married many women, but has never been married. Who is he?",
	"Owl":           "What asks but never answers?",
	"david":         "David's father has three sons : Snap, Crackle and who?",
	"candle":        "Tall I am young, Short I am old, While with life I glow, Wind is my foe. What am I?",
	"name":          "What belongs to you, but other people use it more than you?",
	"gravity":       "Everyone is attracted to me and everybody falls for me. What am I?",
	"gum":           "I go in hard. I come out soft. You blow me hard. What am I?",
	"sleep":         "I weaken all men for hours each day. I show you strange visions while you are away. I take you by night, by day take you back, None suffer to have me, but do from my lack. What am I?",
	"chalkboard":    "What gets whiter the dirtier it gets?",
	"trouble":       "What is easy to get into, but hard to get out of?",
	"clouds":        "I fly without wings, I cry without eyes. What am I?",
	"alphabet":      "Which word contains 26 letters but only three syllables?",
	"match":         "Tear one off and scratch my head, what once was red is black instead!",
	"bubble":        "I am lighter than air but a million men cannot lift me up, What am I?",
	"history":       "You will always find me in the past. I can be created in the present, But the future can never taint me. What am I?",
	"corn":          "You throw away the outside and cook the inside. Then you eat the outside and throw away the inside. What did you eat?",
	"eye":           "What body part is pronounced as one letter but written with three, only two different letters are used?",
	"ruler":         "I am a king who’s good at measuring stuff. What am I?",
	"mailbox":       "A seven letter word containing thousands of letters",
	"bookkeeper":    "What English word has three consecutive double letters?",
	"noon":          "What time of day, when written in a capital letters, is the same forwards, backwards and upside down?",
	"heroine":       "There is a word in the English language in which the first two letters signify a male, the first three letters signify a female, the first four signify a great man, and the whole word, a great woman. What is the word?",
	"friendship":    "I am a ship that can be made to ride the greatest waves. I am not built by tool, but built by hearts and minds. What am I?",
	"map":           "What has cities, but no houses; forests, but no trees; and water, but no fish?",
	"lounger":       "What 7 letter word becomes longer when the third letter is removed?",
	"comb":          "I have many teeth and sometimes they're fine, First I'm by your head, then I'm down your spine. What am I? What am I?",
	"a barber":      "A man shaves several times a day, yet he still has a beard. Who is this man?",
	"pea":           "I am a seed with three letters in my name. Take away the last two and I still sound the same. What am I?",
	"garbage truck": "What has four wheels and flies?",
	"clock":         "When I take five and add six, I get eleven, but when I take six and add seven, I get one. What am I?",
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
