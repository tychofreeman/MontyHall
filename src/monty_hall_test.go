package montyhall

import (
	"testing"
)

func assertIsValidDoor(t *testing.T, door int) {
	if door != 1 && door != 2 && door != 3 {
		t.Fatalf("Winning door returned invalid door number: %v\n", door)
	}
}

func TestReturnsDoorNumber(t *testing.T) {
	for i := 0; i < 10000; i++ {
		winningDoor := GetWinningDoor()
		assertIsValidDoor(t, winningDoor)
	}
}

func TestReturnsRandomDoorNumber(t *testing.T) {
	winningDoors := make([]int, 3)
	for i := 0; i < 10000; i++ {
		winningDoor := int(GetWinningDoor())
		winningDoors[winningDoor - 1]++
	}
	validate := func(i int) {
		if winningDoors[i - 1] <= 0 {
			t.Errorf("Door %d should have been returned at least once. Found %d\n", i, winningDoors[i - 1])
		}
	}
	validate(1)
	validate(2)
	validate(3)
}

func assertIntEq(t *testing.T, actual, expected int, msg string) {
	if actual != expected {
		t.Errorf("%s\nExpected: %d\nActual: %d\n", msg, expected, actual)
	}
}

func TestWhenNonWinningDoorIsChosenOtherNonWinningDoorIsRevealed(t *testing.T) {
	assertIntEq(t, OpenDoor(1, 2), 3, "Opened wrong door")
	assertIntEq(t, OpenDoor(3, 2), 1, "Opened wrong door")
	assertIntEq(t, OpenDoor(3, 1), 2, "Opened wrong door")
	assertIntEq(t, OpenDoor(2, 1), 3, "Opened wrong door")
	assertIntEq(t, OpenDoor(2, 3), 1, "Opened wrong door")
	assertIntEq(t, OpenDoor(1, 3), 2, "Opened wrong door")
}

func TestWhenWinningDoorIsChosenRandomOtherDoorIsOpened(t *testing.T) {
	openedDoors := make([]int, 3)
	for i := 0; i < 10000; i++ {
		winningDoor := i % 3 + 1
		openedDoor := OpenDoor(winningDoor, winningDoor)
		openedDoors[openedDoor - 1]++
		assertIsValidDoor(t, openedDoor)
	}
	validate := func(i int) {
		if openedDoors[i - 1] <= 0 {
			t.Errorf("Door %d should have been returned at least once. Found %d\n", i, openedDoors[i - 1])
		}
	}
	validate(1)
	validate(2)
	validate(3)
}

func assertFalse(t *testing.T, actual bool, msg string) {
	if actual {
		t.Errorf("Expected FALSE but was TRUE\n%s\n", msg)
	}
}

func assertTrue(t *testing.T, actual bool, msg string) {
	if !actual {
		t.Errorf("Expected TRUE but was FALSE\n%s\n", msg)
	}
}

func TestUserLosesWhenTheyChooseTheRightDoorAndSwitch(t *testing.T) {
	assertFalse(t, DoesPlayerWin(1, 1, true), "Player switched to the wrong door.")
	assertFalse(t, DoesPlayerWin(2, 2, true), "Player switched to the wrong door.")
	assertFalse(t, DoesPlayerWin(3, 3, true), "Player switched to the wrong door.")
}

func TestUserWinsWhenTheyChooseTheRightDoorAndStay(t *testing.T) {
	assertTrue(t, DoesPlayerWin(1, 1, false), "Player stayed with the right door.")
	assertTrue(t, DoesPlayerWin(2, 2, false), "Player stayed with the right door.")
	assertTrue(t, DoesPlayerWin(3, 3, false), "Player stayed with the right door.")
}

func TestUserWinsWhenTheyChoseTheWrongDoorAndSwitch(t *testing.T) {
	assertTrue(t, DoesPlayerWin(3, 1, true), "Player switched to the right door.")
	assertTrue(t, DoesPlayerWin(3, 2, true), "Player switched to the right door.")
	assertTrue(t, DoesPlayerWin(2, 1, true), "Player switched to the right door.")
	assertTrue(t, DoesPlayerWin(2, 3, true), "Player switched to the right door.")
	assertTrue(t, DoesPlayerWin(1, 2, true), "Player switched to the right door.")
	assertTrue(t, DoesPlayerWin(1, 3, true), "Player switched to the right door.")
}

func TestUserLosesWhenTheyChoseTheWrongDoorAndStay(t *testing.T) {
	assertFalse(t, DoesPlayerWin(3, 1, false), "Player stayed with the wrong door.")
	assertFalse(t, DoesPlayerWin(3, 2, false), "Player stayed with the wrong door.")
	assertFalse(t, DoesPlayerWin(2, 1, false), "Player stayed with the wrong door.")
	assertFalse(t, DoesPlayerWin(2, 3, false), "Player stayed with the wrong door.")
	assertFalse(t, DoesPlayerWin(1, 2, false), "Player stayed with the wrong door.")
	assertFalse(t, DoesPlayerWin(1, 3, false), "Player stayed with the wrong door.")
}

func TestHowOftenYouWin(t *testing.T) {
	wins := 0
	plays := 0
	run := func(winningDoor, usersDoor int, userSwitches bool) {
		plays++
		if DoesPlayerWin(winningDoor, usersDoor, userSwitches) {
			wins++
		}
	}

	run(1, 1, true)
	run(2, 2, true)
	run(3, 3, true)
	run(1, 2, true)
	run(1, 3, true)
	run(2, 1, true)
	run(2, 3, true)
	run(3, 1, true)
	run(3, 2, true)

	assertTrue(t, wins == 6 && plays == 9, "Switching should give a 6/9 ratio of wins/plays.")

	wins = 0
	plays = 0
	run(1, 1, false)
	run(2, 2, false)
	run(3, 3, false)
	run(1, 2, false)
	run(1, 3, false)
	run(2, 1, false)
	run(2, 3, false)
	run(3, 1, false)
	run(3, 2, false)

	assertTrue(t, wins == 3 && plays == 9, "Not switching should give a 3/9 ratio of wins/plays.")
}
