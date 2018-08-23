# color
Terminal input color

![Color](https://i.imgur.com/ziC61Zp.png)

## Install
```
go get github.com/gurudin/color
```

## Examples
```
fmt.Println()

color.HighCyan("Go: ")
color.Green("Terminal ")
color.Magenta("input ")
color.White("color.")
fmt.Println()

color.Black("Black     ")
color.HighBlack("Black     ")
color.BgBlack("          ")
fmt.Println()

color.Red("Red       ")
color.HighRed("Red       ")
color.BgRed("          ")
fmt.Println()

color.Green("Green     ")
color.HighGreen("Green     ")
color.BgGreen("          ")
fmt.Println()

color.Yellow("Yellow    ")
color.HighYellow("Yellow    ")
color.BgYellow("          ")
fmt.Println()

color.Blue("Blue      ")
color.HighBlue("Blue      ")
color.BgBlue("          ")
fmt.Println()

color.Magenta("Magenta   ")
color.HighMagenta("Magenta   ")
color.BgMagenta("          ")
fmt.Println()

color.Cyan("Cyan      ")
color.HighCyan("Cyan      ")
color.BgCyan("          ")
fmt.Println()

color.White("White     ")
color.HighWhite("White     ")
color.BgWhite("          ")
fmt.Println()

fmt.Println()

left := color.Conf{High: 1, FColor: color.FWhite, BColor: color.BRed}
left.Print("===========")

center := color.Conf{High: 1, FColor: color.FWhite, BColor: color.BYellow}
center.Print(" Custom ")

right := color.Conf{High: 1, FColor: color.FWhite, BColor: color.BBlue}
right.Print("===========")

fmt.Println()
```

## Custom
```
conf := color.Conf{High: 1, FColor: color.FRed, BColor: color.BWhite}
conf.Print("=======================")
fmt.Println()
```