package advice

import (
	"fmt"
	"math/rand"
	"time"
)

var advice = [...]string{
	"To save money when cooking don't use olive oil, use the oil you would use for your car engine.",
	"During an interview ask how strict their sexual harassment policy is.",
	"To save time don't look before crossing the street.",
	"If you kick a police dog it gives you money.",
	"Girlfriend accuses you of checking out other women? Tell her you're not even attracted to thin women",
	"Feel unproductive at work? Cocaine.",
	"To test if a battery still works lick it.",
	"PEMDAS in math class means 'Please Excuse My Dope Ass Swag'.",
	"Real men don't wear pink. They wear crocs.",
	"Every good decision starts with a line of coke.",
	"It's not child abuse if you kill the child.",
	"Fight childhood obesity by beating up fat kids.",
	"Every zoo is a petting zoo.",
	"Lacking vitamin C which is essential to a healthy diet? Two teaspoons of cocaine does wonders.",
	"Stop a baby from crying by shaking it.",
	"Mix things up. Next time a cop pulls you over for speeding arrest him!",
	"Teach your dog tricks by giving him chocolate every time he does one correctly!",
	"In a fight you can't win? Get naked. No one wants to fight a naked guy.",
	"Didn't get a puppy for Christmas? Call 911 and tell them your moms a drug dealer and they'll send a dog to your house that you can play with",
	"Girlfriend telling you loads of stuff? Say 'I already knew that' after each thing, she will be impressed at how much you know.",
	"Want to run faster? Cut off your ears for aerodynamics.",
	"Call your ex, NOW",
	"If you begin any insult with 'no offense', you can say as many hurtful things as you want!",
	"On a plane? Constantly say there is no bomb to reassure the passengers everything is okay.",
	"Make jokes about your girlfriend's weight. Girls love a man with a sense of humor.",
	"If you have guests over, walk into their room every hour at night and wake them up to ask them if they need anything.",
	"If your girlfriend says 'I don't want any gifts for my birthday,' then don't get her any! It will show you're a great listener.",
	"If you ever think time is flying by too fast, get a bad haircut. Bad haircuts make time pass slowly as you wait for your hair to grow back.",
	"Want to find out who your true friends are? Kill yourself and see who goes to the funeral.",
	"Build a man a fire and he'll be warm for a day. Set a man on fire and he'll be warm for the rest of his life.",
	"If you see a group of 4 people attacking someone, stop to help out. They won't stand a chance against the 5 of you",
	"Put pedals on your wheelchair so your arms don't get tired.",
	"Install 1 or 2 viruses on your computer to boost its immune system.",
	"Want your crush to notice you? Like one of their photos on Instagram from a year ago.",
	"Steal people's parking tickets to brighten someone else's day",
	"Don't vaccinate your kids. You will save money when they don't live long enough to go to college.",
	"You can't be late to work if you don't go.",
	"Meeting a foreign person for the first time? Try to imitate their accent when talking to them, they'll feel more at home.",
	"If someone asks you to do the laundry, wash the red and whites together and you'll never have to wash clothes again!",
	"When reading a book, tear each page off as you finish reading it. Next time you open it, you'll open to the exact page you're on.",
	"Cannibalism isn't such a bad idea. Not only would it cure overpopulation, but it would also end world hunger",
	"When breaking up with someone, don't tell them. Just suddenly stop talking to them and never speak again, that way you avoid arguments.",
	"If you pour tobasco sauce in your eyes, they will turn a beautiful red color.",
	"Girl asks you to guess her age? Guess 200, if you're wrong say you thought she said guess her weight!",
	"Put a pee stain halfway down your pants to make people think you've got a longer ding dong",
	"Don't buy groceries that go out of date in February as it is the shortest month and therefore will go out of date quicker.",
	"To save money on Valentine's Day, break up with your significant other so you don't have to buy a gift. Then get back together the day after",
	"If you see a nail on the ground, be sure to stand it upright so that it's more visible for people to avoid it.",
	"You'll have more money if you stop paying bills",
	"When the bully asks you for your lunch money, tell him you left it on his mom's bedside table.",
	"Struggle remembering people's names? Call everyone bro That way you'll only have to remember one name",
	"If you see a woman walking alone in a dangerous area, follow closely behind her to make sure she arrives at her destination safely.",
	"K Y S really stands for keep yourself safe. Be sure to tell this to family members before they begin their day to ensure their safety.",
	"If you ever see a dangerous animal like a trampling elephant go near it and pet it to calm it down.",
}

func BadAdvice() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(advice))
	return fmt.Sprint(advice[n])
}
