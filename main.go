package main

func main() {
	var programm string
	fuckFromFile("./hello_world.bf") // print Hello World
	reset()
	programm = ">+++[[>,.<-]<]"
	fuck(programm)
	reset()
	programm = ">+++++++++[<++++++++>-]<."
	fuck(programm) //print H
	reset()
	programm = ",>,>,.<.<."
	fuck(programm) //read 3, print these 3 reversed
	reset()
}
