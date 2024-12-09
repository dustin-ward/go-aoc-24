package day9

import (
	"bufio"
	"io"
	"strconv"
)

func Part2(inpt io.Reader) string {
	reader := bufio.NewReader(inpt)
    fs := make([]int,0)
    type file struct {
        size int
        pos int
        id int
    }
    files := make([]file,0)
    spaces := make([]file,0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			//EOF
            break
		}
		line = line[:len(line)-1]

        isFile := true
        id := 0
        for _, ch := range line {
            if isFile {
                n := int(ch-rune('0'))
                files = append(files, file{n,len(fs),id})
                for i:=0; i<n; i++ {
                    fs = append(fs, id)
                }
                id++
            } else {
                n := int(ch-rune('0'))
                spaces = append(spaces, file{n,len(fs),-1})
                for i:=0; i<n; i++ {
                    fs = append(fs, -1)
                }
            }
            isFile = !isFile
        }
	}

    for i:=len(files)-1; i>=0; i-- {
        for j:=0; j<len(spaces); j++ {
            if spaces[j].size == 0 {
                continue
            }
            if spaces[j].pos > files[i].pos {
                goto nextFile
            }
            if spaces[j].size >= files[i].size {
                for k:=0; k<files[i].size; k++ {
                    fs[spaces[j].pos+k] = files[i].id
                    fs[files[i].pos+k] = -1
                }
                rem := spaces[j].size-files[i].size
                spaces[j].size = rem
                spaces[j].pos += files[i].size
                goto nextFile
            }
        }

nextFile:
    }

    var ans int
    for i,x := range fs {
        if x == -1 {
            continue
        }
        ans += i*x
    }

    return strconv.Itoa(ans)
}
