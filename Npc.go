package main

import "time"

func moverNpc(jogo *Jogo, ch chan bool) {
	y := 2
	x := 24
	for {
		for i := 0; i < 6; i++ {
			if jogo.PosY < 6 && jogo.PosX < 30 {
				//trava
				jogo.Passou = true
				<-ch
			}
			jogo.Mapa[y][x] = Vazio
			if i > 2 {
				x++
			} else {
				x--
			}
			jogo.Mapa[y][x] = Npc
			time.Sleep(time.Millisecond * 100)
			interfaceDesenharJogo(jogo)
		}
	}
}
