package main

//Hassan Hadji-Ibrahim 300126629
import (
	"fmt"
	"math/rand/v2"

	"time"
)

type Player struct {
	Type      string
	Rows, Col int
	MoveCount int
	Found     bool
	BoardCol  int
	BoarRows  int
	Message   chan string
}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	nrows := rand.IntN(191) + 10
	mcols := rand.IntN(191) + 10
	a := rand.IntN(2 * max(nrows, mcols))
	b := rand.IntN(10 * max(nrows, mcols))
	s := rand.IntN(b-a+1) + a
	
	police := Player{Type: "Police", Rows: 0, Col: 0, MoveCount: s, Found: false, BoardCol: mcols, BoarRows: nrows, Message: make(chan string)}
	thief := Player{Type: "Thief", Rows: nrows - 1, Col: mcols - 1, MoveCount: 0, BoardCol: mcols, BoarRows: nrows, Message: make(chan string)}
	go controller(&police, &thief)
	for {
		select {
		case msg := <-police.Message:
			fmt.Println("Police:", msg)
		case msg := <-thief.Message:
			fmt.Println("Thief:", msg)
		}
	}

}

func policeplays(police *Player) {

	for !police.Found && police.MoveCount > 0 {
		direction := rand.IntN(4)
		switch direction {
		case 0: //up
			if police.Rows > 0 {
				police.Rows--

			}
		case 1: //down
			if police.Rows < police.BoarRows {
				police.Rows++
			}

		case 2: // left
			if police.Col > 0 {
				police.Col--
			}
		case 3: //Right
			if police.Col < police.BoardCol {
				police.Col++
			}

		}
		police.MoveCount--

	}

}

func thiefplays(thief *Player) {

	for !thief.Found && thief.MoveCount > 0 {
		direction := rand.IntN(4)
		switch direction {
		case 0: //up
			if thief.Rows > 0 {
				thief.Rows--

			}
		case 1: //down
			if thief.Rows < thief.BoarRows {
				thief.Rows++
			}

		case 2: // left
			if thief.Col > 0 {
				thief.Col--

			}
		case 3: //Right
			if thief.Col < thief.BoardCol {
				thief.Col++

			}

		}
		thief.MoveCount++

	}

}

func controller(police, thief *Player) {

	for {
		thiefplays(thief)

		policeplays(police)

		police.Message <- "Your movement was successful, the game continues."
		thief.Message <- "Your movement was successful, the game continues."
		time.Sleep(time.Millisecond * 50)

		fmt.Printf(" The Police Position = (%d, %d), The thief Position = (%d, %d). \n", police.Rows, police.Col, thief.Rows, thief.Col)

		if police.Rows == 0 && police.Col == 0 && thief.Rows == 0 && thief.Col == 0 {
			police.Message <- "The game ends in a tie."
			thief.Message <- "The game ends in a tie."
			fmt.Printf("The game ends in a tie")
			break

		}
		if police.Rows == thief.Rows && police.Col == thief.Col {
			police.Message <- "The game ends, you won the game."
			thief.Message <- "The game, ends, you lost the game."
			police.Found = true
			thief.Found = true

			break
		}
		if police.MoveCount == 0 && (police.Rows != thief.Rows || police.Col != thief.Col) {
			police.Message <- "The game, ends, you lost the game."
			thief.Message <- "The game ends, you won the game."
			fmt.Printf("The Police ran out of moves and the Thief won the game.")
			break

		}

		if thief.Rows == 0 && thief.Col == 0 {
			police.Message <- "The game, ends, you lost the game."
			thief.Message <- "The game ends, you won the game."
			fmt.Printf("The Thief escaped and won the game.")
			break
		}

	}

}
