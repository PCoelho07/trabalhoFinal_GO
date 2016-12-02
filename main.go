
package main

import (
	"fmt"
	"time"
	"math/rand"
)


/*
* Função que modela a ação 'aprontar' de uma determinada pessoa representada pelo parametro 'nome'
* return: uma flag para sinalizar a finalização da ação e o tempo gasto
*/
func aprontar(nome string) (bool, int) {
	fmt.Println(nome + " está se aprontando")
	t := random(90, 180)
	time.Sleep(time.Duration(t) * time.Second)

	return true, t
}

/*
* Função que modela a ação 'calçar tênis' de uma determinada pessoa representada pelo parametro 'nome'
* return: uma flag para sinalizar a finalização da ação e o tempo gasto
*/
func calcar_tenis(nome string) (bool, int) {
	fmt.Println(nome + " começou a calçar seu tênis")
	t := random(30, 45)
	time.Sleep(time.Duration(t) * time.Second)

	return true, t
}

/*
* Função que retorna um número aleatório entre o range passado(min, max)
*/
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

/*
* Função que modela toda a lógica de 'negócio' exigido pela aplicação
* params: @nome nome da pessoa que irá executar as ações
* 		  @c canal(sem buffer) que serve para sinalizar que as 'pessoas' terminaram de se aprontar
*		  @t canal(com buffer) que serve para que as 'pessoas' saibam que o alarme foi acionado 
*		  e que elas devem executar a ação 'calçar tenis'
*/
func pessoa(nome string, c chan int, t chan bool) {
	var flag_1 bool
	var t_ct int

	flag, t_a := aprontar(nome)

	if flag == true {
		fmt.Println(nome + " demorou",t_a,"segundos para se preparar")
		c <- 0

		if <-t == false {
			flag_1, t_ct = calcar_tenis(nome)

			if flag_1 == true {
				fmt.Println(nome + " demorou",t_ct,"segundos para calçar o tenis")
				c <- 1
			}
		}
	}

}


/*
* Essa é a nossa função principal(obviamente), que gerencia toda a aplicação e toda lógica de negócio.
*/
func main() {
	// Declarações de canais necessários para a aplicação
	m_1 := make(chan int)
	m_2 := make(chan int)
	timeout := make(chan bool)
	in := make(chan bool, 2)

	// Lançamento das go routines
	go pessoa("Ana", m_1, in)
	go pessoa("Maria", m_2, in)

	msg_1 := <-m_1
	msg_2 := <-m_2
	
	if msg_1 == 0 && msg_2 == 0 {
		
		go func() {
			fmt.Println("Alarme foi acionado")
			in <- false
			in <- false
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
