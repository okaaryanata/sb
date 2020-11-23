package main

import (
	"fmt"
	"sort"
	"strings"
)

// number 3
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound == -1 {
			return ""
		}
		str = str[indexFirstBracketFound:]
		indexClosingBracketFound := strings.Index(str, ")")
		if indexClosingBracketFound == -1 {
			return ""
		}
		return str[1:indexClosingBracketFound]
	}
	return ""
}

// number 4
func anagram(sliceStr []string) [][]string {
	res := [][]string{}
	m := make(map[string]bool)

	for len(sliceStr) != 0 {
		firstElement := sliceStr[0]
		slice := []string{}
		if _, ok := m[firstElement]; !ok {
			m[firstElement] = true
			slice = append(slice, firstElement)
		}
		for _, x := range sliceStr[1:] {
			if len(firstElement) == len(x) {
				if sortString(firstElement) == sortString(x) {
					if _, ok := m[x]; !ok {
						m[x] = true
						slice = append(slice, x)
					}
				}
			}
		}
		sliceStr = deleteElementByIdx(sliceStr, 0)
		if len(slice) != 0 {
			res = append(res, slice)
		}
	}

	return res
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func deleteElementByIdx(slice []string, idx int) []string {
	copy(slice[idx:], slice[idx+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	// number 3
	a := findFirstStringInBracket("adsad(okaaryanata)asad")
	b := findFirstStringInBracket("1(9)asda)as")
	c := findFirstStringInBracket(")(()")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// number 4
	sliceStr := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	// fmt.Println(deleteElementByIdx(sliceStr, 2))
	res := anagram(sliceStr)
	fmt.Println(res)
}
