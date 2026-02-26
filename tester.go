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

	cmd := exec.Command("go", "run", os.Args[1], os.Args[2], os.Args[3])

	fmt.Println("Checking The Audit Text")

	datas := []Checker{
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
		},
	}

	for i, task := range datas {

	}

	err := os.WriteFile(os.Args[1], []byte(""))

	// Capture the output
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Output from other file:\n%s", string(out))
}
