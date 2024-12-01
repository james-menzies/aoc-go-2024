# Advent of Code Golang Starter

A quick boilerplate setup for completing Advent of Code challenges.

## How to use

1. Click the "Use this template" button in the top-right of the page to create your own repository, then clone it down to your machine.
2. Get your individualised input data from the daily challenge.
3. Copy / Paste the data into the `input.txt` file in the appropriate day folder.
4. Fill out the `solution1` and `solution2` functions to solve the challenge. Note that the same input is used for both challenges.
  The input will be a slice of strings, where each string will be one line of the input.
5. Run `go run ./dayX`, replacing 'X' with the day number you're solving. Your solution will be printed to the screen.

## Tweaking 

If you want to tweak this setup feel free to fork or raise a PR. Some under-the-hood details:
* There's a small util in `internal/read_util` which takes a file name and parses it into a slice of strings (one string per slice)
* There's a small bash script in the root of the directory which you can use to copy the "state" of the `day1` folder and propagate it to the other directories.
