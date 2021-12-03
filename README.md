# Advent of Code 2021

[![](https://github.com/grvn/aoc2021/workflows/CI/badge.svg)](https://github.com/grvn/aoc2021/actions)

Testar att utföra uppgifterna i Advent of Code för fjärde gången.
I år så blir det att testa Go (https://go.dev/)

Jag har tidigare stött på något som kallas för [cobra](github.com/spf13/cobra) som ser ut att vara väldigt användbart här för att förenkla upprepning av kod.
Jag kommer samtidigt ta tillfället i akt och testa GitHub actions.

[Advent of Code 2021](http://adventofcode.com/2021)


       .--._.--.--.__.--.--.__.--.--.__.--.
      _(_      _Y_      _Y_      _Y_      _)_  
     [___]    [___]    [___]    [___]    [___]   
     /:' \    /:' \    /:' \    /:' \    /:' \   
    |::   |  |::   |  |::   |  |::   |  |::   | 
    \::.  /  \::.  /  \::.  /  \::.  /  \::.  / 
     \::./    \::./    \::./    \::./    \::./  
      '='      '='      '='      '='      '='

## Krav
* go 1.17

## Köra

För att köra, välj vilken problem, vilken del samt filvägen till input

Exempel: för att köra dag 1, del 1

```bash
$ go run main.go 1 1 -i day1/input.txt
Answer: 1
Took 0.2021ms
```

## Lägga till lösning

| Fil | Beskrivning |
|:----|------------:|
|`day<N>/input.txt`|puzzel input|
|`day<N>/test.txt`|test input|
|`day<N>/get.go`|Samlar ihop del 1 och 2, kopiera från tidigare dag|
|`day<N>/part1.go`|Del 1 huvudfil|
|`day<N>/part2.go`|Del 2 huvudfil|
|`cmd/cmd.go`|Plockar ihop alla dagar till Cobra, här behöver man lägga till en ny rad för varje dag man lägger till|