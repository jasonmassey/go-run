package main

// numberToASCII converts a number to its ASCII art representation.
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

    // Convert the number to a string
    numStr := string(rune(num + '0'))
    // Get the ASCII art for the number
    asciiNum := asciiArt[num]

    // Return the ASCII art
    return []string{asciiNum}
}