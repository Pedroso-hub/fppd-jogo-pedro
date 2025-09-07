package main

import (
	"time"
)

func pertoDe(jogo *Jogo, ox, oy int) bool {
	//retorno := (jogo.PosX == (ox-1) || jogo.PosX == (ox+1)) && (jogo.PosY == (oy-1) || jogo.PosY == (oy+1))
	somaJogo := jogo.PosX + jogo.PosY
	somaObj := ox + oy
	retorno := somaJogo == somaObj-1 || somaJogo == somaObj+1

	return retorno
}

func AtivarAlavanca(jogo *Jogo) {
	//se o player chegar perto e apertar interagir, ativa a alavanca
	//mudarAtivacaoAlavanca(jogo, semaforoAlavanca, true)
	for {
		if acessarAtivacaoAlavanca(jogo, semaforoAlavanca) == true {
			select {
			case <-passouPortal:
				mudarAtivacaoAlavanca(jogo, semaforoAlavanca, false)
				jogo.StatusMsg = "passou do portal"
			case <-time.After(10 * time.Second):
				mudarAtivacaoAlavanca(jogo, semaforoAlavanca, false)
				jogo.StatusMsg = "demorou demais"
			}
		}
	}

}
