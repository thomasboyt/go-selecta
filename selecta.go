package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"go-selecta/selecta"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	app := cli.NewApp()
	app.Name = "selecta"
	app.Usage = "fuzzy find whatever you want"
	app.Action = func(c *cli.Context) {
		// parse choices
		bytes, _ := ioutil.ReadAll(os.Stdin)
		choices := strings.Split(string(bytes), "\n")

		// create a search
		search := selecta.BlankSearch(choices, "", 0)

		// set up termbox
		err := termbox.Init()
		if err != nil {
			panic(err)
		}

		termbox.SetInputMode(termbox.InputEsc)

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		DrawApp(search)
		EventLoop(search)

		termbox.Close()

		fmt.Fprint(os.Stdout, search.SelectedChoice())
	}

	app.Run(os.Args)
}

func EventLoop(s *selecta.Search) {
loop:
	for !s.Done {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC:
				break loop
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				s.Backspace()
			case termbox.KeyEnter:
				s.Done = true
			case termbox.KeyCtrlN:
				s.Down()
			case termbox.KeyCtrlP:
				s.Up()
			case termbox.KeyCtrlW:
				s.DeleteWord()
			default:
				char := rune(ev.Ch)
				if !unicode.IsControl(char) {
					s.AppendQuery(string(char))
				} else if ev.Key == termbox.KeySpace {
					s.AppendQuery(" ")
				}
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			DrawApp(s)
		}
	}
}

// Rendering

func DrawApp(s *selecta.Search) {
	WriteLine(0, "> "+s.Query, false)
	termbox.SetCursor(2+len(s.Query), 0)

	for i, match := range s.Matches {
		choice := match.Value
		highlight := false
		if s.Index == i {
			highlight = true
		}
		WriteLine(i+1, choice, highlight)
	}

	termbox.Flush()
}

func WriteLine(row int, str string, highlight bool) {
	bgColor := termbox.ColorDefault
	width, _ := termbox.Size()
	if highlight {
		bgColor = termbox.ColorBlack
	}

	for col := 0; col < width; col++ {
		if col < len(str) {
			termbox.SetCell(col, row, rune(str[col]), termbox.ColorDefault, bgColor)
		} else {
			termbox.SetCell(col, row, ' ', termbox.ColorDefault, bgColor)
		}
	}
}
