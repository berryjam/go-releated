package main

func h() int { return 1 }

func main() {
	f := h
	f()
}