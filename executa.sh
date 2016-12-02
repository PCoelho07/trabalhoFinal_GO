#!/bin/bash

echo "Digite quantos testes deseja realizar:"
read qtd_testes
	
	if [ -z $qtd_testes ]
	then
		echo "Digite um parâmetro válido!"
	elif [ $qtd_testes -gt 0  ]
	then
		./bateria_testes.sh $qtd_testes >> resultados_testes.txt
		echo "Testes finalizados \n Podem ser conferidos no arquivo 'resultados_testes.txt'"
	fi
 
