//submitted by Rocket dev team:
// Liz Aharonian - 316584960
// Raz Shenkman - 311130777
// Ori Ben Zaken - 311492110

package main

import "fmt"

type Player struct {
	c    chan int
	left *Player
}

func main() {
	const n = 100000
	// initialize pointer to the first player - right most player
	firstPlayer := &Player{make(chan int), nil}
	currPlayer := firstPlayer
	for i := 0; i < n; i++ {
		newPlayer := &Player{make(chan int), nil}
		currPlayer.left = newPlayer
		go play(currPlayer)
		currPlayer = newPlayer
	}
	lastPlayer := currPlayer
	//send 1 to first player
	firstPlayer.c <- 1
	//take last value from last player
	fmt.Printf("Left most player say: %d", <-lastPlayer.c)
}

func play(player *Player) {
	// check if the player is not the last (left most) player
	if player.left == nil {
		return
	}
	player.left.c <- <-player.c + 1
}
