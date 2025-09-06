package main

func teleportar(jogo *Jogo) {
	//jogo.StatusMsg = "ta rolando a teleportar"
	for {
		if jogo.PosX == 22 && jogo.PosY == 21 {
			jogo.PosX = 62
			jogo.PosY = 20
			jogo.Mapa[20][62] = Personagem
			jogo.Mapa[21][22] = Portal

			interfaceDesenharJogo(jogo)
		}

	}
}
