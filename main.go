package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min := 1
	max := 100
	fmt.Println("Marhaba aklımdan 1 ile 100 arasında bir sayı tuttum. Bunu 10 tahmin hakkında bilmeye çalışmalısın")
	target := randomNumber(min, max)
	attemps:=0

	for ; attemps<10;attemps++ {
		fmt.Println(10-attemps,"hakkın kaldı")
		reader := bufio.NewReader(os.Stdin)
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if num < target {
			fmt.Println("Yukarı")
		} else if num > target {
			fmt.Println("Aşağı")
		} else {
			fmt.Println("Tebrikler!", attemps,"denemede sonucu buldunuz.")
			break
		}

	}

	if attemps==10{
		fmt.Println("Maalesef sayıyı tahmin edemedin. Hedef sayı:", target)
	}
	
	
	// fmt.Println("Maalesef tahmin edemedin. Bulman gereken sayı", number,"idi" )
	

}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}