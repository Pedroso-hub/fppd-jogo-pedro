package main

var semaforoAlavanca chan bool = make(chan bool, 1)

func teleportar(jogo *Jogo, passou chan bool) {
	//jogo.StatusMsg = "ta rolando a teleportar"
	for {
		if jogo.PosX == 22 && jogo.PosY == 21 && acessarAtivacaoAlavanca(jogo, semaforoAlavanca) {
			jogo.PosX = 62
			jogo.PosY = 20
			jogo.Mapa[20][62] = Personagem
			jogo.Mapa[21][22] = Portal
			passou <- true
			interfaceDesenharJogo(jogo)
		}
	}
}

// use de semáforo binário para garantir exlusão mútua
func acessarAtivacaoAlavanca(jogo *Jogo, ch chan bool) bool {
	ch <- true
	ativou := jogo.ativouAlavanca
	<-ch
	return ativou
}

func mudarAtivacaoAlavanca(jogo *Jogo, ch chan bool, valor bool) {
	ch <- true
	jogo.ativouAlavanca = valor
	<-ch
}
