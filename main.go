
package main

import (
	"fmt"
	"time"
	"math/rand"
)



func aprontar(nome string) (bool, int) {
	fmt.Println(nome + " está se aprontando")
	t := random(6, 18)
	time.Sleep(time.Duration(t) * time.Second)

	return true, t
}

func calcar_tenis(nome string) (bool, int) {
	fmt.Println(nome + " começou a calçar seu tênis")
	t := random(3, 4)
	time.Sleep(time.Duration(t) * time.Second)

	return true, t
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}


func ana(megafone_1 chan int, t chan bool) {
	var flag_1 bool
	var t_ct int
	flag, t_a := aprontar("Ana")

	if flag == true {
		fmt.Println("Ana demorou",t_a,"segundos para se preparar")
		megafone_1 <- 0
		
		if <-t == false {
			flag_1, t_ct = calcar_tenis("Ana")
		}

		if flag_1 == true {
			fmt.Println("Ana demorou",t_ct,"segundos para calçar seu tenis")			
			megafone_1 <- 1		
		}
	}
	
}

func maria(megafone_2 chan int, t chan bool) {
	var flag_1 bool
	var t_ct int
	flag, t_a := aprontar("Maria")

	if flag == true {
		fmt.Println("Maria demorou",t_a,"segundos para se preparar")
		megafone_2 <- 0

		if <-t == false {
			flag_1, t_ct = calcar_tenis("Maria")
		}

		if flag_1 == true {
			fmt.Println("Maria demorou",t_ct,"segundos para calçar seu tenis")		
			megafone_2 <- 1		
		}
	}
}



// func alarme() {
// 	fmt.Println("Alarme foi acionado")
// 	fmt.Println("Alarme está em contagem regressiva")
// 	timer := time.NewTimer(time.Second * 6)
// 	<-timer.C
// }




func main() {
	m_1 := make(chan int)
	m_2 := make(chan int)
	timeout := make(chan bool)

	go ana(m_1, timeout)
	go maria(m_2, timeout)

	msg_1 := <-m_1
	msg_2 := <-m_2
	
	if msg_1 == 0 && msg_2 == 0 {
		
		go func() {
			timeout <- false
			fmt.Println("Alarme foi acionado")
			fmt.Println("Alarme está em contagem regressiva")
			time.Sleep(6 * time.Second)
			timeout <- true
		}()
	}

	msg_1 = <-m_1
	msg_2 = <-m_2

	if msg_1 == 1 && msg_2 == 1 {
		fmt.Println("Ana e Maria saíram e trancaram a porta")
	}

	select {
		case <-timeout:	
			fmt.Println("Alarme foi ativado")
	}
}
