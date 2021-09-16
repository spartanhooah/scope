package packageone

import "fmt"

var PackageVar = "package-level variable in packageone"

func PrintMe(first, second, third string) {
	fmt.Println(first, second, third)
}