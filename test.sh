#!/usr/bin/env bash

# Test
echo "$ go run . example00"
go run . ./examples/example00.txt
read -rp "$*"

echo
echo "$ go run . example01"
go run . ./examples/example01.txt
read -rp "$*"

echo
echo "$ go run . example02"
go run . ./examples/example02.txt
read -rp "$*"

echo
echo "$ go run . example03"
go run . ./examples/example03.txt
read -rp "$*"


echo
echo "$ go run . example04"
go run . ./examples/example04.txt
read -rp "$*"

echo
echo "$ go run . example05"
go run . ./examples/example05.txt
read -rp "$*"

echo
echo "$ go run . badexample00"
go run . ./examples/badexample00.txt
read -rp "$*"

echo
echo "$ go run . badexample01"
go run . ./examples/badexample01.txt

echo
echo "$ go run . example06"
go run . ./examples/example06.txt
read -rp "$*"

echo
echo "$ go run . example07"
go run . ./examples/example07.txt
read -rp "$*"

echo
echo "All done! If no errors, then the test was passed."