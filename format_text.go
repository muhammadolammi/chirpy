package main

import "strings"

func formatString(s string) string {
	newWords := []string{}
	badWords := []string{"kerfuffle", "sharbert", "fornax", "Kerfuffle", "Sharbert", "Fornax"}
	words := strings.Fields(s)

	for _, word := range words {
		isBad := false
		for _, bw := range badWords {
			if word == bw {
				isBad = true
				break
			}
		}

		if isBad {
			newWords = append(newWords, "****")
		} else {
			newWords = append(newWords, word)
		}
	}

	formattedString := strings.Join(newWords, " ")
	return formattedString
}
