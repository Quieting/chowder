package snake

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSnake(t *testing.T) {
	snake := InitSnake(12, 10)

	m := [10][13]string{}
	initSlice(&m)

	// 测试 increase 方法
	t.Run("increase", func(t *testing.T) {

	})

	// 测试 ahead 方法
	t.Run("ahead", func(t *testing.T) {

		initSlice(&m)
		for i := uint16(0); i < snake.length; i++ {
			m[snake.body[i].y][snake.body[i].x] = "*"
		}
		outSlice(m)

		// 走两步
		snake.SetDirection(left)
		snake.ahead()
		snake.ahead()

		initSlice(&m)
		for i := uint16(0); i < snake.length; i++ {
			m[snake.body[i].y][snake.body[i].x] = "*"
		}

		outSlice(m)
	})
}

func initSlice(res *[10][13]string) {
	for i := uint16(0); i < 10; i++ {
		for j := uint16(0); j < 13; j++ {
			res[i][j] = "."
		}
		res[i][12] = "\n"
	}
}
func outSlice(res [10][13]string) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 13; j++ {
			fmt.Printf(res[i][j])
		}
	}
	fmt.Println()
}

func TestGameStart(t *testing.T) {
	snake := InitSnake(12, 10)
	snake.SetSpeed(500 * time.Millisecond)

	m := [10][13]string{}
	initSlice(&m)

	dre := make(chan uint8)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		char := scanner.Text()
		switch char {
		case "w":
			dre <- up
		case "s":
			dre <- down
		case "a":
			dre <- left
		case "d":
			dre <- right
		}
	}

	ticker := time.NewTicker(snake.speed)
	for {
		select {
		case <-ticker.C:
			snake.ahead()
			initSlice(&m)
			for i := uint16(0); i < snake.length; i++ {
				m[snake.body[i].y][snake.body[i].x] = "*"
			}

			outSlice(m)
		case d := <-dre:
			snake.direction = d
		}
	}
}
