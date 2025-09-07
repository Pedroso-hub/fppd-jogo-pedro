package main

var ch2 chan bool = make(chan bool)

func alternarSimbolo(jogo *Jogo, passouPortal chan bool, ch2 chan bool) {

	for {
		select {
		case <-ch2:
			mudarSimbolo(jogo, 18, 27, semaforoMatriz, '♡')

		case <-passouPortal:
			mudarSimbolo(jogo, 18, 27, semaforoMatriz, '✷')
		}
	}

}

func mudarSimbolo(jogo *Jogo, y, x int, semaforoMatriz chan bool, valor rune) {
	semaforoMatriz <- true
	jogo.Mapa[y][x].simbolo = valor
	<-semaforoMatriz
}
