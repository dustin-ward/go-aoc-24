package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	// Formatter sorts them alphabetically... ew
	"github.com/dustin-ward/go-aoc-24/day1"
    "github.com/dustin-ward/go-aoc-24/day2"
	"github.com/dustin-ward/go-aoc-24/day10"
	"github.com/dustin-ward/go-aoc-24/day11"
	"github.com/dustin-ward/go-aoc-24/day12"
	"github.com/dustin-ward/go-aoc-24/day13"
	"github.com/dustin-ward/go-aoc-24/day14"
	"github.com/dustin-ward/go-aoc-24/day15"
	"github.com/dustin-ward/go-aoc-24/day16"
	"github.com/dustin-ward/go-aoc-24/day17"
	"github.com/dustin-ward/go-aoc-24/day18"
	"github.com/dustin-ward/go-aoc-24/day19"
	"github.com/dustin-ward/go-aoc-24/day20"
	"github.com/dustin-ward/go-aoc-24/day21"
	"github.com/dustin-ward/go-aoc-24/day22"
	"github.com/dustin-ward/go-aoc-24/day23"
	"github.com/dustin-ward/go-aoc-24/day24"
	"github.com/dustin-ward/go-aoc-24/day25"
	"github.com/dustin-ward/go-aoc-24/day3"
	"github.com/dustin-ward/go-aoc-24/day4"
	"github.com/dustin-ward/go-aoc-24/day5"
	"github.com/dustin-ward/go-aoc-24/day6"
	"github.com/dustin-ward/go-aoc-24/day7"
	"github.com/dustin-ward/go-aoc-24/day8"
	"github.com/dustin-ward/go-aoc-24/day9"
)

var day = flag.String("d", "", "Advent of Code puzzle to run")
var doAll = flag.Bool("a", false, "Run all days solutions. Overrules the -d value")
var doInteractive = flag.Bool("i", false, "Run the solution on stdin instead of the input file")

func run(day, part int, solution func(io.Reader) string) {
	var inpt io.Reader
	if *doInteractive {
		inpt = os.Stdin
	} else {
		inpt = getInpt(day)
	}

	startTime := time.Now()
	result := solution(inpt)
	duration := time.Now().Sub(startTime)

	log.Printf("Day %d-%d: %s - %v", day, part, result, duration)
}

func getInpt(day int) io.Reader {
	filepath := fmt.Sprintf("day%d/inpt.txt", day)
	file, err := os.Open(filepath)
	if errors.Is(err, os.ErrNotExist) {
        log.Println("Input file not on disk. attempting to fetch...")

		aocCookie := os.Getenv("AOC_COOKIE")
		if aocCookie == "" {
			log.Fatal("$AOC_COOKIE not set. Unable to retrieve puzzle input")
		}

		puzzleURL := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
        client := http.Client{}
        req, _ := http.NewRequest("GET", puzzleURL, nil)
        req.Header.Set("Cookie", "session="+aocCookie)
        res, err := client.Do(req)
		if err != nil {
			log.Fatal("Error retrieving puzzle input: ", err)
		}

		if res.StatusCode >= 300 {
			log.Fatal("Error retrieving puzzle input: ", res.Status)
		}

		file, err = os.Create(filepath)
		if err != nil {
			log.Fatal("Error creating file: ", err)
		}

		io.Copy(file, res.Body)
        file.Close()
		res.Body.Close()

        file, err = os.Open(filepath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
	}

	return file
}

func main() {
	flag.Parse()

	type SolutionInfo struct {
		Day      int
		Part     int
		Solution func(io.Reader) string
	}
	solutions := make([]SolutionInfo, 0)

    masterList := []SolutionInfo{
        {1,1,day1.Part1},
        {1,2,day1.Part2},
        {2,1,day2.Part1},
        {2,2,day2.Part2},
        {3,1,day3.Part1},
        {3,2,day3.Part2},
        {4,1,day4.Part1},
        {4,2,day4.Part2},
        {5,1,day5.Part1},
        {5,2,day5.Part2},
        {6,1,day6.Part1},
        {6,2,day6.Part2},
        {7,1,day7.Part1},
        {7,2,day7.Part2},
        {8,1,day8.Part1},
        {8,2,day8.Part2},
        {9,1,day9.Part1},
        {9,2,day9.Part2},
        {10,1,day10.Part1},
        {10,2,day10.Part2},
        {11,1,day11.Part1},
        {11,2,day11.Part2},
        {12,1,day12.Part1},
        {12,2,day12.Part2},
        {13,1,day13.Part1},
        {13,2,day13.Part2},
        {14,1,day14.Part1},
        {14,2,day14.Part2},
        {15,1,day15.Part1},
        {15,2,day15.Part2},
        {16,1,day16.Part1},
        {16,2,day16.Part2},
        {17,1,day17.Part1},
        {17,2,day17.Part2},
        {18,1,day18.Part1},
        {18,2,day18.Part2},
        {19,1,day19.Part1},
        {19,2,day19.Part2},
        {20,1,day20.Part1},
        {20,2,day20.Part2},
        {21,1,day21.Part1},
        {21,2,day21.Part2},
        {22,1,day22.Part1},
        {22,2,day22.Part2},
        {23,1,day23.Part1},
        {23,2,day23.Part2},
        {24,1,day24.Part1},
        {24,2,day24.Part2},
        {25,1,day25.Part1},
        {25,2,day25.Part2},
    }

	if *doAll {
        solutions = masterList
	} else {
        dayRe := regexp.MustCompile(`^([1-9]|1[0-9]|2[0-5])(?:-([12]))?$`)
        match := dayRe.FindAllStringSubmatch(*day, 1)
        if match == nil {
            log.Fatal("Invalid option passed for -d (day). Format should be one of <dayNum> or <dayNum>-<partNum>")
        }

        var day, part int
        both := true

        day, _ = strconv.Atoi(match[0][1])
        if match[0][2] != "" {
            both = false
            part, _ = strconv.Atoi(match[0][2])
        }

        switch day {
        case 1:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{1,1,day1.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{1,2,day1.Part2})
            }
        case 2:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{2,1,day2.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{2,2,day2.Part2})
            }
        case 3:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{3,1,day3.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{3,2,day3.Part2})
            }
        case 4:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{4,1,day4.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{4,2,day4.Part2})
            }
        case 5:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{5,1,day5.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{5,2,day5.Part2})
            }
        case 6:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{6,1,day6.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{6,2,day6.Part2})
            }
        case 7:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{7,1,day7.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{7,2,day7.Part2})
            }
        case 8:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{8,1,day8.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{8,2,day8.Part2})
            }
        case 9:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{9,1,day9.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{9,2,day9.Part2})
            }
        case 10:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{10,1,day10.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{10,2,day10.Part2})
            }
        case 11:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{11,1,day11.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{11,2,day11.Part2})
            }
        case 12:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{12,1,day12.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{12,2,day12.Part2})
            }
        case 13:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{13,1,day13.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{13,2,day13.Part2})
            }
        case 14:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{14,1,day14.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{14,2,day14.Part2})
            }
        case 15:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{15,1,day15.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{15,2,day15.Part2})
            }
        case 16:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{16,1,day16.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{16,2,day16.Part2})
            }
        case 17:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{17,1,day17.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{17,2,day17.Part2})
            }
        case 18:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{18,1,day18.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{18,2,day18.Part2})
            }
        case 19:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{19,1,day19.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{19,2,day19.Part2})
            }
        case 20:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{20,1,day20.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{20,2,day20.Part2})
            }
        case 21:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{21,1,day21.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{21,2,day21.Part2})
            }
        case 22:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{22,1,day22.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{22,2,day22.Part2})
            }
        case 23:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{23,1,day23.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{23,2,day23.Part2})
            }
        case 24:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{24,1,day24.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{24,2,day24.Part2})
            }
        case 25:
            if both || part == 1 {
                solutions = append(solutions, SolutionInfo{25,1,day25.Part1})
            }
            if both || part == 2 {
                solutions = append(solutions, SolutionInfo{25,2,day25.Part2})
            }
        }
    }

	for _, solution := range solutions {
		run(solution.Day, solution.Part, solution.Solution)
	}
}
