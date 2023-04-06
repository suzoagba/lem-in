#!/usr/bin/env bash
# Run the Go program with various input files to test its behavior

# Run the program with example00.txt and wait for user input
echo "$ go run . example00"
go run . ./examples/example00.txt
read -rp "$*"

# Run the program with example01.txt and wait for user input
echo
echo "$ go run . example01"
go run . ./examples/example01.txt
read -rp "$*"

# Run the program with example02.txt and wait for user input
echo
echo "$ go run . example02"
go run . ./examples/example02.txt
read -rp "$*"

# Run the program with example03.txt and wait for user input
echo
echo "$ go run . example03"
go run . ./examples/example03.txt
read -rp "$*"

# Run the program with example04.txt and wait for user input
echo
echo "$ go run . example04"
go run . ./examples/example04.txt
read -rp "$*"

# Run the program with example05.txt and wait for user input
echo
echo "$ go run . example05"
go run . ./examples/example05.txt
read -rp "$*"

# Run the program with badexample00.txt and wait for user input
echo
echo "$ go run . badexample00"
go run . ./examples/badexample00.txt
read -rp "$*"

# Run the program with badexample01.txt and wait for user input
echo
echo "$ go run . badexample01"
go run . ./examples/badexample01.txt

# Run the program with example06.txt and wait for user input
echo
echo "$ go run . example06"
go run . ./examples/example06.txt
read -rp "$*"

# Run the program with example07.txt and wait for user input
echo
echo "$ go run . example07"
go run . ./examples/example07.txt
read -rp "$*"

# Print a message indicating that all tests are complete
echo
echo "All done!"