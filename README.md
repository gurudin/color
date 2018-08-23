# color
Terminal input color

```
color.Black("Hello world")
color.HighBlack("Hello world")
color.BgBlack("Hello word")

color.Red("Hello world")
color.HighRed("Hello world")
color.BgRed("Hello world")

color.Green("Hello world")
color.HighGreen("Hello world")
color.BgGreen("Hello world")

color.Yellow("Hello world")
color.HighYellow("Hello world")
color.BgYellow("Hello world")

color.Blue("Hello world")
color.HighBlue("Hello world")
color.BgBlue("Hello world")

color.Magenta("Hello world")
color.HighMagenta("Hello world")
color.BgMagenta("Hello world")

color.Cyan("Hello world")
color.HighCyan("Hello world")
color.BgCyan("Hello world")

color.White("Hello world")
color.HighWhite("Hello world")
color.BgWhite("Hello world")

fmt.Println()

conf := color.Conf{High: 1, FColor: color.FRed}
conf.Print("=======================")
fmt.Println()
```