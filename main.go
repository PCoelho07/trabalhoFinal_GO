
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


func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}


func ana(megafone chan bool) {
	flag, t_a := aprontar("Ana")

	if flag == true {
		megafone <- true
		flag_1, t_ct := calçar_tenis("Ana")
		fmt.Println("Ana demorou " + t_a + " segundos para se preparar")
	}
	
}

func maria(megafone chan bool) {
	aprontar("Maria")
}

func alarme(megafone chan bool) {


}




func main() {
	m = make(chan bool 2)
	var flag bool = false

	fmt.Println("Alo mundo!")
}
