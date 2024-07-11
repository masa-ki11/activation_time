package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("奥義セット時間:")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				fmt.Println("プログラム終了。")
				break
			}
			remainingTimeInSeconds, err := convertInputToSeconds(input)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			activationTime := calculateActivationTime(remainingTimeInSeconds)
			fmt.Printf("奥義発動時間: %s\n\n", secondsToMMSS(activationTime))
		}
	}
}

func convertInputToSeconds(input string) (int, error) {
	if len(input) != 4 {
		return -1, fmt.Errorf("入力はmmss形式で4桁")
	}
	minutes, err := strconv.Atoi(input[:2])
	if err != nil {
		return -1, fmt.Errorf("分の値が無効")
	}
	seconds, err := strconv.Atoi(input[2:])
	if err != nil {
		return -1, fmt.Errorf("秒の値が無効")
	}
	if minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		return -1, fmt.Errorf("範囲外（0-59の間で入力）")
	}
	return minutes*60 + seconds, nil
}

func calculateActivationTime(currentSeconds int) int {
	return currentSeconds - 90
}

func secondsToMMSS(seconds int) string {
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d分%02d秒", minutes, seconds)
}