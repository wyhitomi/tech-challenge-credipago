package main

import (
	"fmt"
)

func main() {

	/*
		Seu script deverá ser capaz de simular o agendamento de pelo menos 3 visitas
		e imprimir o resultado no console. Não é necessário fazer a leitura destes
		dados de uma fonte externa, como um arquivo ou o console.

		Os dados podem ficar “hard-coded” no seu script, para simplificar.
	*/

	agenda := []visit{}
	visit1 := visit{1, "2023-05-25 12:00", 1, 1, "scheduled"}
	visit2 := visit{2, "2023-05-26 13:00", 2, 10, "scheduled"}
	visit3 := visit{3, "2023-05-27 14:00", 5, 16, "scheduled"}
	visit4 := newVisit(4, "2023-05-28 17:00", 9, 40, "scheduled")

	agenda = append(agenda, visit1)
	agenda = append(agenda, visit2)
	agenda = append(agenda, visit3)
	agenda = append(agenda, visit4)

	canceled1 := newVisit(4, "2023-05-28 17:00", 9, 40, "canceled")
	agenda = append(agenda, canceled1)

	// reserva
	// verificar se existe visita agendada no mesmo horário
		// comparar se no id.date há conflito
		// comparar se no id.apartment não é o mesmo apartamento
		// comparar se no id.visitor não é o mesmo visitante
	// sugerir horário seguinte ou pedir para o cliente sugerir um novo

	fmt.Println(agenda)
}
