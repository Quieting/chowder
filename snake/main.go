package snake

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Run() {
	snake := InitSnake(12, 10)
	snake.SetSpeed(1500 * time.Millisecond)

	food := &point{3, 3}

	m := [10][13]string{}
	initSlice(&m)

	dre := make(chan uint8)
	ticker := time.NewTicker(snake.speed)
	go func() {
		for {
			select {
			case <-ticker.C:
				if food.same(snake.Next()) {
					snake.eat()
					// 刷新食物
					for {
						food.y = uint16(rand.Intn(10))
						food.x = uint16(rand.Intn(12))
						ok := true
						for i := uint16(0); i < snake.length; i++ {
							if food.same(&snake.body[i]) {
								ok = false
								break
							}
						}
						if ok {
							break
						}
					}
				} else {
					snake.ahead()
				}

				initSlice(&m)
				for i := uint16(0); i < snake.length; i++ {
					m[snake.body[i].y][snake.body[i].x] = "*"
				}
				m[food.y][food.x] = "*"
				outSlice(m)
			case d := <-dre:
				snake.direction = d
			}
		}
	}()

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
