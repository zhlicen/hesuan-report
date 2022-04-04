package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	sum := 27*8 + 2
	done := len(hesuanList) + len(notHomeList)
	if done > sum {
		sum = done
	}

	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	fmtt := color.New(color.FgHiBlue)
	progress := float64(done) / float64(sum)
	fmt.Println()
	fmtt.Print("----------------------------------------------------------------------------\n")
	fmtt.Println("【更新时间】\t", time.Now().Format("2006-01-02 15:04:05"))
	// progress bar
	fmtt.Printf("【完成率】\t ")
	blocks := int(progress * 20)
	for i := 0; i < blocks; i++ {
		fmtt.Print("■")
	}
	for i := 0; i < 20-blocks; i++ {
		fmtt.Print("□")
	}
	fmtt.Printf(" %.2f %%\n", progress*100)
	red.Printf("【红色】未接龙\t")
	green.Printf("【绿色】已接龙\t")
	fmt.Println("【白色】不在家")
	fmtt.Print("----------------------------------------------------------------------------\n")
	floorCount := 28
	for i := 1; i <= floorCount; i++ {
		floorCounter := 0
		for j := 1; j <= 8; j++ {
			roomNo := fmt.Sprintf("%d%02d", i, j)
			if i == 28 && j > 2 {
				roomNo = "N/A"
			}
			if roomNo == "N/A" {
				fmt.Print(roomNo + "\t")
			} else if _, ok := hesuanList[roomNo]; ok {
				floorCounter++
				if _, ok := baobeiList[roomNo]; ok {
					green.Print(roomNo + "\t")
				} else {
					green.Print(roomNo + "\t")
				}
			} else {
				if _, ok := notHomeList[roomNo]; ok {
					// not home
					fmt.Print(roomNo + "\t")
					floorCounter++
				} else if _, ok := baobeiList[roomNo]; ok {
					red.Print(roomNo + "\t")
				} else {
					// not home
					fmt.Print(roomNo + "\t")
					floorCounter++
				}
			}
		}
		floorCount := 8
		if i == 28 {
			floorCount = 2
		}
		finishRate := float64(floorCounter) / float64(floorCount)
		fmtt.Printf("  | %.2f %%", finishRate*100)

		fmt.Print("\n")
	}
	fmtt.Print("----------------------------------------------------------------------------\n\n")
}
