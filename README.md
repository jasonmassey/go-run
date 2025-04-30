# go-run

## Summary
The game is a simple ASCII-based obstacle avoidance game where the player controls a character (`@`) navigating through a field filled with obstacles (`#`). The objective is to avoid collisions with obstacles and accumulate points by successfully navigating through the field. The game ends after one minute or when the player collides with an obstacle.

## Installation Instructions
1. **Install Go**: Ensure you have Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the Repository**: Clone the repository containing the game source code to your local machine.
   ```sh
   git clone https://github.com/jasonmassey/go-run.git
   cd your-repo
   ```

3. **Install Dependencies**: The game uses the `termbox-go` library for terminal handling. Install it using the following command:
   ```sh
   go get github.com/nsf/termbox-go
   ```

4. **Run the Game**: Execute the game directly using the Go interpreter.
   ```sh
   go run main.go
   ```

6. **Controls**:
   - Use the arrow keys to move the player.
   - Press `q` or `Ctrl+Q` to quit the game.
   - Press `Ctrl+C` to exit the game.
