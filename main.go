// main.go - Loop principal do jogo
package main

import "os"

func main() {
	// Inicializa a interface (termbox)
	interfaceIniciar()
	defer interfaceFinalizar()

	// Usa "mapa.txt" como arquivo padrão ou lê o primeiro argumento
	mapaFile := "mapa.txt"
	if len(os.Args) > 1 {
		mapaFile = os.Args[1]
	}

	// Inicializa o jogo
	jogo := jogoNovo()
	//jogo.PertoNpc = make(chan bool, 1)
	if err := jogoCarregarMapa(mapaFile, &jogo); err != nil {
		panic(err)
	}

	// Desenha o estado inicial do jogo
	interfaceDesenharJogo(&jogo)

	ch := make(chan bool)
	go moverNpc(&jogo, ch)

	// Loop principal de entrada
	//loop infinito
	for {

		evento := interfaceLerEventoTeclado()
		//continuar vai ser falso se o player apertar sair
		continuar := personagemExecutarAcao(evento, &jogo)

		if !continuar {
			break
		}

		interfaceDesenharJogo(&jogo)
	}
}
