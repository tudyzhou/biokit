#! /bin/bash
rm -f main
rm -f test.txt.* test2.txt.*
go build main.go
./main unique test.txt test2.txt
