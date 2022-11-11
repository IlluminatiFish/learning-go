package main

func getDigitalRoot(number int) int {
  return  (number - 1) % 9 + 1
}

func main() {
	fmt.Printf("%d", getDigitalRoot(64)) // 1
	fmt.Printf("%d", getDigitalRoot(43)) // 7
}
