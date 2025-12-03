package year2025

import (
	"fmt"
	"strconv"
	"strings"
)

func Lobby(input string) (res int) {
	banks := strings.Split(input, "\n")

	for _, bank := range banks {
		fl, sl := '0', '0'
		for i, batterie := range bank {
			if batterie > fl && i != len(bank)-1 {
				fl = batterie
				sl = '0'
			} else if batterie > sl {
				sl = batterie
			}
		}

		max, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(fl), string(sl)))
		res += max

	}
	return res
}

func LobbyP2(input string) (res int) {
	banks := strings.Split(input, "\n")
	var MAX_BATTERIES = 12

	for _, bank := range banks {
		joltageS := make([]byte, MAX_BATTERIES)

		pi := 0
		for i := 0; i < MAX_BATTERIES; i++ {
			for s := 0 + pi; s < len(bank)-MAX_BATTERIES+i+1; s++ {
				if bank[s] > joltageS[i] {
					joltageS[i] = bank[s]
					pi = s + 1
				}
			}
		}

		max, _ := strconv.Atoi(string(joltageS))
		fmt.Println(max)
		res += max

	}
	return res
}
