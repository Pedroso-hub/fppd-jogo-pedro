package main

import "time"

func moverNpc(jogo *Jogo, ch chan bool) {
	y := 2
	x := 24
	simbolo := jogo.Mapa[y][x]
	for {
		for i := 0; i < 6; i++ {
			if jogo.PosY < 6 && jogo.PosX < 30 {
				//trava
				jogo.Passou = true
				<-ch
			}
			if i > 2 {
				jogo.Mapa[y][x].simbolo = ' ' // restaura o conteúdo anterior
				x++
				jogo.Mapa[y][x] = simbolo
				time.Sleep(time.Millisecond * 100)
			} else {
				jogo.Mapa[y][x].simbolo = ' ' // restaura o conteúdo anterior
				x--
				jogo.Mapa[y][x] = simbolo
				time.Sleep(time.Millisecond * 100)
			}
			interfaceDesenharJogo(jogo)
		}
	}
}
