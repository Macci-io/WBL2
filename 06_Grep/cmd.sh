clear
cat test.txt | go run main.go -n -C 2 -B 3 "123"
sleep 3
clear
cat test.txt | go run main.go -A 2 "[[:digit:]]$"
sleep 3
clear
cat test.txt | go run main.go -n -C 2 "^[[:digit:]]{1,3}$"
