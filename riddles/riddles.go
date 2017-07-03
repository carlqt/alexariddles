package riddles

import (
	"fmt"
	"math/rand"
	"time"
)

var riddles = [...]string{
	"It walks on four legs in the morning, two legs at noon and three legs in the evening. What is it?",
}

func Ask() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(riddles))
	return fmt.Sprint(riddles[n])
}
