package main

import (
	"bsh-tfe/view"
	"github.com/nsf/termbox-go"
	"strconv"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var ch rune = 0x1037E

	termbox.SetInputMode(termbox.InputEsc)
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	draw.Box(1,1,5,36)
	termbox.Flush()


loop:
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyCtrlX:
				break loop
			case termbox.KeyArrowUp:
				ch += 1
			case termbox.KeyArrowDown:
				ch -= 1
			}
		case termbox.EventError:
			panic(event.Err)
		}

		termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
		draw.Box(1,1,5,36)
		for x := 0; x < 30; x+=2 {
			termbox.SetCell(x+2, 2, ch, termbox.ColorWhite, termbox.ColorBlack)
		}
		draw.Text(1, 10, strconv.QuoteRuneToASCII(ch))
		termbox.Flush()
	}
}
