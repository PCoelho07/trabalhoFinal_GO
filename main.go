
package main

import (
	"fmt"
	"time"
	"math/rand"
)



func aprontar(nome string) (bool, int) {
	fmt.Println(nome + "está se aprontando")
	t := random(60, 180)
	time.Sleep(t * time.Seconds)

	return true, t
}

func calcar_tenis(nome string) (bool, int) {
	fmt.Println(nome + "começou a calçar seu tênis")
	t := random(30, 45)
	time.Sleep(t * time.Seconds)

	return true, t
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}


func ana(megafone_1 chan int) {
	flag, t_a := aprontar("Ana")

	if flag == true {
		fmt.Println("Ana demorou " + t_a + " segundos para se preparar")
		megafone_1 <- 0
		flag_1, t_ct := calçar_tenis("Ana")

		if flag_1 == true {
			fmt.Println("Ana demorou " + t_ct + " segundos para calçar seu tenis")			
			megafone_1 <- 1		
		}
	}
	
}

func maria(megafone_2 chan int) {
	flag, t_a := aprontar("Maria")

	if flag == true {
		fmt.Println("Maria demorou " + t_a + " segundos para se preparar")
		megafone_2 <- 0
		flag_1, t_ct := calçar_tenis("Maria")

		if flag_1 == true {
			fmt.Println("Maria demorou " + t_ct + " segundos para calçar seu tenis")			
			megafone_2 <- 1		
		}
	}
}

func gerenciador() {
	
}


func alarme(megafone chan bool) {


}




func main() {
	m = make(chan bool 2)
	var flag bool = false

	fmt.Println("Alo mundo!")
}
