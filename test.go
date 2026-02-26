package main

import (
	"fmt"
	"slices"
	"strings"
)

func fixedApostrophe() {
	txt := " ' bright '  ' has it's jide. ' It  was  amazing ' lk ' it 's  "
	list := strings.Fields(txt)
	start := -1
	end := -1
	for i := 0; i < len(list); i++ {
		switch list[i] {
		case "'s", "'ve", "'t", "'re":
			if i > 0 {
				list[i-1] += list[i]
				list = slices.Delete(list, i, i+1)
				i--
			}
		case "'":
			if start >= 0 {
				end = i
				//        3           4
				list[start] += strings.Join(list[start+1:end], " ") + "'"
				list = slices.Delete(list, start+1, end+1)
				i = start
				start, end = -1, -1
			} else {
				start = i
			}
		}
	}

	fmt.Println(strings.Join(list, " "))

}
