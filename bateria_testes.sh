#!/bin/bash

if [ -z $1 ]; then
	echo "Entre com um parâmetro válido!"
elif [ $1 -gt 0 ] 
then
	echo "teste"
	for i in `seq 1 $qtd_testes`
	do 	
		echo "---- Teste número $i ----"
		go run main.go
	done

	echo "Os testes foram finalizados!"
else
	echo "Digite um valor maior que zero"
fi



