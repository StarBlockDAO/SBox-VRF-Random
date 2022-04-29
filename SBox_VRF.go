package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func main() {

	nftWinnerList := []int{}

	// SBOXRandomSeedGenerator: https://etherscan.io/address/0x8a664be336304385b2886e706828e4ee1e549d44#code
	var contractRandomSeed = "GET_SEED_FROM_SBOXRandomSeedGenerator_CONTRACT"
	var nftCount = 20

	// we are converting seed to md5 then int64
	md5Seed := md5.New()
	_, _ = io.WriteString(md5Seed, contractRandomSeed)
	var seed = binary.BigEndian.Uint64(md5Seed.Sum(nil))
	rand.Seed(int64(seed))
	
	
	// this ignore list is team/investor owned sboxes.
	// this list can be updated (please watch this repo)
	ignoreList := []int{}

	for len(nftWinnerList) < nftCount {
		// get random sbox id between 0, 19
		randBox := randInt(0, 19)
		// add this sbox to winnerList
		if !contains(nftWinnerList, randBox) && !contains(ignoreList, randBox) {
			nftWinnerList = append(nftWinnerList, randBox)
		}
	}

	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nftWinnerList)), ","), "[]"))

}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
