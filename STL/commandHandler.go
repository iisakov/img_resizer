package stl

import (
	"fmt"
	"img_resizer/config"
	"os"
	"strconv"
	"strings"
)

var Prefix_b, Prefix_s string = "b_", "s_"
var MaxWidthBigImg, MaxHightBigImg uint = 1200, 900
var MaxWidthSmallImg, MaxHightSmallImg uint = 186, 140
var PathRoot string = "."

func Handle(comands []string) {
	for _, c := range comands {
		switch c {

		case "--help":
			fallthrough
		case "-h":
			PrintMessage("Здесь будет руководство пользователя")
			Exit()

		case "--version":
			fallthrough
		case "-v":
			PrintMessage("img_resizer [by_artisan] v:" + config.Version)
			Exit()

		case "--param":
			fallthrough
		case "-p":
			setParam(comands)
		default:
		}
	}
}

func setParam(comands []string) {
	for _, c := range comands {
		cS := strings.Split(c, ":")

		if len(cS) > 1 {
			switch cS[0] {
			case "Prefix_b":
				fallthrough
			case "pb":
				Prefix_b = cS[1]

			case "Prefix_s":
				fallthrough
			case "ps":
				Prefix_s = cS[1]

			case "MaxWidthBigImg":
				fallthrough
			case "mwb":
				cSS, err := strconv.Atoi(cS[1])
				if err != nil {
					PrintError(err.Error())
					Exit()
				}
				MaxWidthBigImg = uint(cSS)

			case "MaxHightBigImg":
				fallthrough
			case "mws":
				cSS, err := strconv.Atoi(cS[1])
				if err != nil {
					PrintError(err.Error())
					Exit()
				}
				MaxHightBigImg = uint(cSS)

			case "MaxWidthSmallImg":
				fallthrough
			case "mhb":
				cSS, err := strconv.Atoi(cS[1])
				if err != nil {
					PrintError(err.Error())
					Exit()
				}
				MaxWidthSmallImg = uint(cSS)

			case "MaxHightSmallImg":
				fallthrough
			case "mhs":
				cSS, err := strconv.Atoi(cS[1])
				if err != nil {
					PrintError(err.Error())
					Exit()
				}
				MaxHightSmallImg = uint(cSS)

			case "PathRoot":
				fallthrough
			case "pr":
				PathRoot = cS[1]

			default:
				PrintMessage("Команды: " + c + " не существует.")
				Exit()
			}
		}
	}
}

func PrintMessage(messages ...string) {
	response := ""
	for _, message := range messages {
		response += message + "\n"
	}
	fmt.Println(response)
}

func Exit(messages ...string) {
	response := ""
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Спасибо, что воспользовались img_resizer от [by_artisan]"

	fmt.Println(response)
	os.Exit(0)
}

func PrintError(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Утилита закрыта."

	fmt.Println(response)
	fmt.Print("Процесс завершится после нажатия 'Enter'.")
}
