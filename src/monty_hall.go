package montyhall

import (
	"big"
	"crypto/rand"
	"fmt"
)


func randInt(maxI int) int {
	max := big.NewInt(int64(maxI))
	door, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(fmt.Sprintf("Cannot generate random numbers! %v", err))
	}
	return int(door.Int64())
}

func randOtherInt(selected int) int {
	if selected == 1 {
		return randInt(2) + 2
	} else if selected == 2 {
		return randInt(2) * 2 + 1
	} else if selected == 3 {
		return randInt(2) + 1
	}
	return 0
}

func remainingDoor(selected, winning int) int {
	return 6 - selected - winning
}

func OpenDoor(userChoice, winningDoor int) int {
	openedDoor := 0
	if userChoice != winningDoor {
		openedDoor = remainingDoor(userChoice, winningDoor)
	} else {
		openedDoor = randOtherInt(userChoice)
	}
	return openedDoor
}

func DoesPlayerWin(winningDoor, usersDoor int, userSwitches bool) bool {
	if userSwitches {
		return usersDoor != winningDoor
	}
	return usersDoor == winningDoor
}

func GetWinningDoor() int {
	return randInt(3) + 1
}

