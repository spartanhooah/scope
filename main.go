package main

import "myapp/packageone"

var myVar = "package-level var in main package"

func main() {
	// variables
	// deeclare a package-level variable for the main package
	// named 'myVar'

	// declare a block-level variable for the main function
	// called blockVar

	// declare a package-level variable in the packageone
	// package named PackageVar

	// create an exported functoni in packageone
	// called PrintMe

	// in the main function, print out the values of
	// myVar, blockVar, and PackageVar on one line
	// using the PrintMe function from packaeone

	blockVar := "block-level var in main function"

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}