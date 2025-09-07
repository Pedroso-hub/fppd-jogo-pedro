package main

var semaforoAlavanca chan bool = make(chan bool, 1)

var passouPortal chan bool = make(chan bool)
var passouPortal2 chan bool = make(chan bool)

func teleportar(jogo *Jogo, passou, passou2 chan bool) {
	//jogo.StatusMsg = "ta rolando a teleportar"
	for {
		if jogo.PosX == 22 && jogo.PosY == 21 && acessarAtivacaoAlavanca(jogo, semaforoAlavanca) {
			jogo.PosX = 62
			jogo.PosY = 20
			mudarMatriz(jogo, 20, 62, Personagem, semaforoMatriz)

			mudarMatriz(jogo, 21, 22, Portal, semaforoMatriz)

			passou <- true
			passou2 <- true

			desenhar(jogo, semaforoDesenhar)
		}
	}
}

// use de semáforo binário para garantir exlusão mútua
func acessarAtivacaoAlavanca(jogo *Jogo, ch chan bool) bool {
	ch <- true
	ativou := jogo.ativouAlavanca
	defer liberarSemaforo(ch)
	return ativou
}

func mudarAtivacaoAlavanca(jogo *Jogo, ch chan bool, valor bool) {
	ch <- true
	jogo.ativouAlavanca = valor
	<-ch
}
