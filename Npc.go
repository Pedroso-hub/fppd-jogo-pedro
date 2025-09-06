package main

import (
	"time"
)

func moverNpc(jogo *Jogo, ch chan bool) {

	for {
		for i := 0; i < 6; i++ {
			if jogo.PosY < 6 {
				jogo.Passou = true
				<-ch

			}

			if i > 2 {
				jogoMoverElemento(jogo, 21+i-3, 2, 1, 0)
				time.Sleep(time.Millisecond * 100)
				interfaceDesenharJogo(jogo)
			} else {
				jogoMoverElemento(jogo, 24-i, 2, -1, 0)
				time.Sleep(time.Millisecond * 100)
				interfaceDesenharJogo(jogo)
			}

		}

	}

}
