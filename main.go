package test2

import "go.uber.org/zap"

func main() {
	println(zap.InfoLevel)
	println("Hello go shitty modules")
	PackageSuperPrint()
}

func PackageSuperPrint() {
	println("Super")
}
