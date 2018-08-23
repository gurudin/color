package main

import (
	"fmt"
)

const (
	FBlack = iota + 30
	FRed
	FGreen
	FYellow
	FBlue
	FMagenta
	FCyan
	FWhite
)
const (
	BBlack = iota + 40
	BRed
	BGreen
	BYellow
	BBlue
	BMagenta
	BCyan
	BWhite
)

type Conf struct {
	High   int8
	BColor int16
	FColor int16
}

func (c *Conf) Print(output string) {
	fmt.Printf(
		"%c[%d;%d;%dm%s%c[0m",
		0x1B,
		c.High,
		c.BColor,
		c.FColor,
		output,
		0x1B,
	)
}

func Print(output string, color int16) {
	conf := Conf{
		High:   0,
		BColor: 0,
		FColor: color,
	}

	conf.Print(output)
}

// Black
func Black(output string) {
	Print(output, FBlack)
}

// High Black
func HighBlack(output string) {
	conf := Conf{High: 1, FColor: FBlack, BColor: 1}

	conf.Print(output)
}

// BgBlack
func BgBlack(output string) {
	conf := Conf{FColor: 1, BColor: BBlack}

	conf.Print(output)
}

// Red
func Red(output string) {
	Print(output, FRed)
}

// High Red
func HighRed(output string) {
	conf := Conf{High: 1, FColor: FRed, BColor: 1}

	conf.Print(output)
}

// BgRed
func BgRed(output string) {
	conf := Conf{FColor: 1, BColor: BRed}

	conf.Print(output)
}

// Green
func Green(output string) {
	Print(output, FGreen)
}

// High Green
func HighGreen(output string) {
	conf := Conf{High: 1, FColor: FGreen, BColor: 1}

	conf.Print(output)
}

// BgGreen
func BgGreen(output string) {
	conf := Conf{FColor: 1, BColor: BGreen}

	conf.Print(output)
}

// Yellow
func Yellow(output string) {
	Print(output, FYellow)
}

// High Yellow
func HighYellow(output string) {
	conf := Conf{High: 1, FColor: FYellow, BColor: 1}

	conf.Print(output)
}

// BgYellow
func BgYellow(output string) {
	conf := Conf{FColor: 1, BColor: BYellow}

	conf.Print(output)
}

// Blue
func Blue(output string) {
	Print(output, FBlue)
}

// High Blue
func HighBlue(output string) {
	conf := Conf{High: 1, FColor: FBlue, BColor: 1}

	conf.Print(output)
}

// BgBlue
func BgBlue(output string) {
	conf := Conf{FColor: 1, BColor: BBlue}

	conf.Print(output)
}

// Magenta
func Magenta(output string) {
	Print(output, FMagenta)
}

// High Magenta
func HighMagenta(output string) {
	conf := Conf{High: 1, FColor: FMagenta, BColor: 1}

	conf.Print(output)
}

// BgMagenta
func BgMagenta(output string) {
	conf := Conf{FColor: 1, BColor: BMagenta}

	conf.Print(output)
}

// Cyan
func Cyan(output string) {
	Print(output, FCyan)
}

// High Cyan
func HighCyan(output string) {
	conf := Conf{High: 1, FColor: FCyan, BColor: 1}

	conf.Print(output)
}

// BgCyan
func BgCyan(output string) {
	conf := Conf{FColor: 1, BColor: BCyan}

	conf.Print(output)
}

// White
func White(output string) {
	Print(output, FWhite)
}

// High White
func HighWhite(output string) {
	conf := Conf{High: 1, FColor: FWhite, BColor: 1}

	conf.Print(output)
}

// BgWhite
func BgWhite(output string) {
	conf := Conf{FColor: 1, BColor: BWhite}

	conf.Print(output)
}
