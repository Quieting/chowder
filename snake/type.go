package snake

import (
	"math"
	"math/rand"
	"time"
)

// 上下左右
const (
	_ = iota
	up
	right
	down
	left
)

// point 点
type point struct {
	y uint16 // 行
	x uint16 // 列
}

// same 两个点是否是同一个点
func (p *point) same(n *point) bool {
	if p.x == n.x && p.y == n.y {
		return true
	}
	return false
}

// Snake 模拟蛇
type Snake struct {
	body      []point       // 蛇的身体(包含蛇头蛇尾),第一点是蛇头,第二点是蛇尾
	direction uint8         // 前进的方向:up、right、down、left
	speed     time.Duration // 速度, 前进一格花费的时间
	length    uint16        // 蛇长

	// 左上角(0,0),右上角(widthMax,0),右下角(widthMax,heightMax),左下角(0,heightMax)
	widthMax  uint16 // 蛇活动的最大宽度
	heightMax uint16 // 蛇活动的最大高度
}

// InitSnake 获取 Snake 实例
func InitSnake(width, height uint16) *Snake {
	snake := &Snake{
		body:      []point{point{uint16(height / 2), uint16(width / 2)}},
		length:    1,
		widthMax:  width,
		heightMax: height,
	}

	// 设置方向
	rand.Seed(time.Now().Unix())
	dir := rand.Intn(left + 1)
	if dir == 0 {
		dir++
	}
	snake.direction = uint8(dir)

	// 添加一截身体
	snake.eat()

	return snake
}

// ahead 前进, 如果遇见边境和蛇身将前进失败, game over
func (s *Snake) ahead() {

	head := s.body[0]

	switch s.direction {
	case up: // 上
		head.y--
	case right: // 右
		head.x++
	case down: // 下
		head.y++
	case left: // 左
		head.x--
	}

	// 前进一格
	s.body = append([]point{head}, s.body[0:s.length-1]...)
}

// eat 蛇吃东西
func (s *Snake) eat() {

	head := s.body[0]

	switch s.direction {
	case up: // 上
		head.y--
	case right: // 右
		head.x++
	case down: // 下
		head.y++
	case left: // 左
		head.x--
	}

	// 前进一格
	s.body = append([]point{head}, s.body[0:s.length]...)
	s.length++
}

// SetSpeed 设置速度
func (s *Snake) SetSpeed(t time.Duration) {
	s.speed = t
}

// SetDirection 设置方向
func (s *Snake) SetDirection(direction uint8) {
	bew := math.Abs(float64(s.direction) - float64(direction))
	if bew == 1 || bew == 3 {
		s.direction = direction
	}
}

// Next 贪吃蛇下一次将活动的位置
func (s *Snake) Next() *point {
	head := s.body[0]

	switch s.direction {
	case up: // 上
		head.y--
	case right: // 右
		head.x++
	case down: // 下
		head.y++
	case left: // 左
		head.x--
	}
	return &head
}
