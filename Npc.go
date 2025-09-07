package main

import "time"

var yNpc int = 2
var xNpc int = 24

func moverNpc(jogo *Jogo, ch chan bool) {
	for {
		for i := 0; i < 6; i++ {
			if jogo.PosY < 6 && jogo.PosX < 30 {
				//trava
				jogo.Passou = true
				<-ch
			}
			jogo.Mapa[yNpc][xNpc] = Vazio
			if i > 2 {
				xNpc++
			} else {
				xNpc--
			}
			jogo.Mapa[yNpc][xNpc] = Npc
			time.Sleep(time.Millisecond * 100)
			interfaceDesenharJogo(jogo)
		}
	}
}
