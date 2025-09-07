package main

import "time"

var yNpc int = 2
var xNpc int = 24

var semaforoDesenhar chan bool = make(chan bool, 1)
var semaforoMatriz chan bool = make(chan bool, 1)

func moverNpc(jogo *Jogo, ch chan bool) {
	for {
		for i := 0; i < 6; i++ {
			if jogo.PosY < 6 && jogo.PosX < 30 {
				//trava
				jogo.Passou = true
				<-ch
			}
			mudarMatriz(jogo, yNpc, xNpc, Vazio, semaforoMatriz)

			if i > 2 {
				xNpc++
			} else {
				xNpc--
			}
			mudarMatriz(jogo, yNpc, xNpc, Npc, semaforoMatriz)

			time.Sleep(time.Millisecond * 100)
			desenhar(jogo, semaforoDesenhar)
		}
	}
}

func desenhar(jogo *Jogo, semaforoDesenhar chan bool) {
	semaforoDesenhar <- true
	interfaceDesenharJogo(jogo)
	<-semaforoDesenhar
}

func mudarMatriz(jogo *Jogo, y, x int, elem Elemento, semaforoMatriz chan bool) {
	semaforoMatriz <- true
	jogo.Mapa[y][x] = elem
	<-semaforoMatriz
}

func acessarMatriz(jogo *Jogo, y, x int, semaforoMatriz chan bool) Elemento {
	semaforoMatriz <- true
	elem := jogo.Mapa[y][x]
	defer liberarSemaforo(semaforoMatriz)
	return elem
}

func liberarSemaforo(sem chan bool) {
	<-sem
}
