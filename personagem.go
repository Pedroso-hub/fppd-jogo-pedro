// personagem.go - Funções para movimentação e ações do personagem
package main

import "fmt"

// Atualiza a posição do personagem com base na tecla pressionada (WASD)
func personagemMover(tecla rune, jogo *Jogo, ch chan bool, ch2 chan bool) {
	dx, dy := 0, 0
	switch tecla {
	case 'w':
		dy = -1 // Move para cima
	case 'a':
		dx = -1 // Move para a esquerda
	case 's':
		dy = 1 // Move para baixo
	case 'd':
		dx = 1 // Move para a direita
	}
	nx, ny := jogo.PosX+dx, jogo.PosY+dy
	// Verifica se o movimento é permitido e realiza a movimentação
	if jogoPodeMoverPara(jogo, nx, ny) {
		jogoMoverElemento(jogo, jogo.PosX, jogo.PosY, dx, dy)
		jogo.PosX, jogo.PosY = nx, ny
		if (jogo.PosY > 6 || jogo.PosX > 30) && jogo.Passou {
			//libera
			jogo.Passou = false
			ch <- true
			ch2 <- true
		}
	}
}

// Define o que ocorre quando o jogador pressiona a tecla de interação
// Neste exemplo, apenas exibe uma mensagem de status
// Você pode expandir essa função para incluir lógica de interação com objetos
func personagemInteragir(jogo *Jogo) {
	// Atualmente apenas exibe uma mensagem de status

	if pertoDe(jogo, 19, 4) {
		jogo.StatusMsg = "ativou o portal"
		mudarAtivacaoAlavanca(jogo, semaforoAlavanca, true)
	} else {
		jogo.StatusMsg = fmt.Sprintf("Interagindo em (%d, %d), ultimo visitado (%c)", jogo.PosX, jogo.PosY, jogo.UltimoVisitado.simbolo)
	}
}

// Processa o evento do teclado e executa a ação correspondente
func personagemExecutarAcao(ev EventoTeclado, jogo *Jogo, ch chan bool, ch2 chan bool) bool {

	switch ev.Tipo {
	case "sair":
		// Retorna false para indicar que o jogo deve terminar
		return false
	case "interagir":
		// Executa a ação de interação
		personagemInteragir(jogo)
	case "mover":
		// Move o personagem com base na tecla
		personagemMover(ev.Tecla, jogo, ch, ch2)

	}
	return true // Continua o jogo
}
