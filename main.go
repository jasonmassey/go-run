package main

import (
	"os"
	"os/exec"
	"os/signal"
	"math/rand"
	"syscall"
	"time"
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	width  = 80
	height = 20
)

// numberToASCII is defined in asciiart.go
func numberToASCII(num int) []string {
	// Basic ASCII art for numbers 0-9
	asciiArt := []string{
		" 000 ",
		"0   0",
		"0   0",
		"0   0",
		" 000 ",
		" 1  ",
		" 1  ",
		" 1  ",
		" 1  ",
		" 1  ",
		" 222",
		"2   2",
		" 222",
		"2   2",
		" 222",
		" 333",
		"3   3",
		" 333",
		"3   3",
		" 333",
		" 4 4",
		" 4 4",
		" 444",
		" 4 4",
		" 4 4",
		" 555",
		"5   5",
		" 555",
		"5   5",
		" 555",
		" 666",
		"6   6",
		" 666",
		"6   6",
		" 666",
		" 777",
		"7   7",
		" 7 7",
		" 7 7",
		" 7 7",
		" 888",
		"8   8",
		" 888",
		"8   8",
		" 888",
		" 999",
		"9   9",
		" 999",
		"9   9",
		" 999",
	}

	// Get the ASCII art for the number
	asciiNum := asciiArt[num]

	// Return the ASCII art
	return []string{asciiNum}
}

type Game struct {
	playerX   int
	playerY   int
	score     int
	obstacles []struct {
		x int
		y int
	}
}

func (g *Game) init() {
	g.playerX = 1
	g.playerY = height / 2
	g.score = 0
	g.obstacles = []struct {
		x int
		y int
	}{
		{x: 6, y: height / 2}, // Initial obstacle 5 spaces away from the player
	}
}

func (g *Game) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == g.playerX && y == g.playerY {
				termbox.SetCell(x, y, '@', termbox.ColorGreen, termbox.ColorDefault)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
	for _, obstacle := range g.obstacles {
		termbox.SetCell(obstacle.x, obstacle.y, '#', termbox.ColorRed, termbox.ColorDefault)
	}

	// Draw the score as ASCII art
	// scoreASCII := numberToASCII(g.score)
	// scoreLines := []string{}
	// for _, line := range scoreASCII {
	// 	scoreLines = append(scoreLines, line)
	// }

	// // Display the score at the top of the screen
	// for i, line := range scoreLines {
	// 	for j, char := range line {
	// 		termbox.SetCell(j, i, char, termbox.ColorYellow, termbox.ColorDefault)
	// 	}
	// }

	// Display the score as a simple three-digit number
	scoreStr := fmt.Sprintf("%03d", g.score)
	for i, char := range scoreStr {
		termbox.SetCell(i, 0, char, termbox.ColorYellow, termbox.ColorDefault)
	}

	termbox.Flush()
}

func (g *Game) handleInput(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowUp:
		if g.playerY > 0 {
			g.playerY--
		}
	case termbox.KeyArrowDown:
		if g.playerY < height-1 {
			g.playerY++
		}
	case termbox.KeyArrowLeft:
		if g.playerX > 0 {
			g.playerX--
		}
	case termbox.KeyArrowRight:
		if g.playerX < width-1 {
			g.playerX++
		}
	case termbox.KeyCtrlQ:
		exitGame()
	case termbox.KeyCtrlC:
		exitGame()
	}
	if ev.Ch == 'q' {
		exitGame()
	}
}

func (g *Game) update() {
	for i, obstacle := range g.obstacles {
		obstacle.x--
		if obstacle.x < 0 {
			g.obstacles = append(g.obstacles[:i], g.obstacles[i+1:]...)
			g.score++
		}
	}
	// Check for collision with obstacles
	for _, obstacle := range g.obstacles {
		if g.playerX == obstacle.x && g.playerY == obstacle.y {
			g.score++ // Increment the score when a collision is detected
			// Handle collision, e.g., reduce score or end game
			// Double the number of obstacles
			if len(g.obstacles) < 200 { // Limit the number of obstacles
				// Add one new obstacle at a random position
				randX := rand.Intn(width)
				randY := rand.Intn(height)
				g.obstacles = append(g.obstacles, struct {
					x int
					y int
				}{x: randX, y: randY})
			}
			// Optionally, you can add a game over condition here
			// exitGame()
		}
	}
	if len(g.obstacles) < 5 {
		g.obstacles = append(g.obstacles, struct {
			x int
			y int
		}{x: width - 1, y: height / 2})
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer func() {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()
		termbox.Close()
	}()

	game := &Game{}
	game.init()

	// Handle Ctrl+C signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Run signal handler in a goroutine
	go func() {
		<-signalChan
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()
		exitGame()
	}()

	// Start a timer for one minute
	timeout := time.After(1 * time.Minute)

	for {
		select {
		case <-timeout:
			exitGame()
		default:
			game.draw()
			game.update()

			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey {
				game.handleInput(ev)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}

}

func exitGame() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Exit(0)
}
