package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Checker struct {
	task     string
	solution string
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("3 arguments are needed")
		os.Exit(1)
	}

	// cmd := exec.Command("go", "run", os.Args[1], os.Args[2], os.Args[3])

	fmt.Println("Checking The Audit Text")

	datas := Tester()

	for i, info := range datas {
		if i == 4 {
			fmt.Println("Your code successfully passed the audit test")
			fmt.Println("Checking Purecoders Test")
			fmt.Println()
		}
		err := os.WriteFile("sample.txt", []byte(info.task), 0644)

		if err != nil {
			os.Exit(1)
		}
		cmd := exec.Command("go", "run", os.Args[1], os.Args[2], os.Args[3])
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Task %d: Failed ❌\nError: %s\nCheck log file for the error details\n\n", i+1, err)
			err := os.WriteFile("log.txt", []byte(out), 0644)
			if err != nil {
				fmt.Println("Error Writing to file", err)
			}

		}

		checkFile, err := os.ReadFile(os.Args[3])

		if err != nil {
			fmt.Printf("Error Reading %s file: %s", os.Args[3], err)
			continue
		}
		if string(checkFile) == info.solution {
			fmt.Printf("Task %d: Successfully ✅\n", i+1)
			fmt.Println(string(out))
		} else {
			fmt.Printf("Task %d: Comparism Failed ❌\n", i+1)
		}

	}
}

func Tester() []Checker {
	data := []Checker{
		{
			task:     "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			solution: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			task:     "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			solution: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			task:     "Don not be sad ,because sad backwards is das . And das not good",
			solution: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			task:     "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			solution: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		}, {
			// 5. The "Lazy A/An" Test
			// Proves that tags are deleted BEFORE a/an is evaluated.
			task:     "It takes a (up) hour to eat a (cap) apple .",
			solution: "It takes AN hour to eat An apple.",
		},
		{
			// 6. The "Out of Bounds" Multi-Tag Test
			// Proves your `if index < num { num = index }` check works and doesn't panic.
			task:     "hello world (up, 50)",
			solution: "HELLO WORLD",
		},
		{
			// 7. The Quote State-Machine Test
			// Proves quotes squash spaces properly and ignore punctuation rules inside.
			task:     " ' this is a (cap) test ' said the programmer ... ",
			solution: "'this is A test' said the programmer...",
		},
		{
			// 8. The Contraction & Punctuation Cluster Test
			// Proves your CheckPunctuation loop handles multiple marks like !? properly.
			task:     "we 're going to the park ,right ? ! ? ",
			solution: "we're going to the park, right?!?",
		},
		{
			// 9. The "Beginning of Time" Shield Test
			// Proves your `index > 0` and `index == 0` shields stop panics at the start of a file.
			task:     ", hello (cap) (hex) 1A is weird but we 've seen worse .",
			solution: ", Hello (hex) 1A is weird but we've seen worse.",
		},
		{
			// 10. The "Fake Tag" Test
			// Proves that your strconv.Atoi safely ignores malformed numbers in tags.
			task:     "This (fake) tag and (up, banana) should not change .",
			solution: "This (fake) tag and (up, banana) should not change.",
		},
		{
			task:     "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair. Simply add 42 (hex) and 10 (bin) and you will see the result is 68. There is no greater agony than bearing a untold story inside you. Punctuation tests are ... kinda boring ,what do you think ? a",
			solution: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair. Simply add 66 and 2 and you will see the result is 68. There is no greater agony than bearing an untold story inside you. Punctuation tests are... kinda boring, what do you think? a",
		},
	}

	return data
}
