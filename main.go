package main

import (
	"fmt"
	"sort"
	"time"
)

func binarySearch(arr []string, uri string) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == uri {
			return mid
		} else if arr[mid] < uri {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Elemento não encontrado
}

func main() {
	urls := []string{
		"/propostas",
		"/propostas/{id_proposta}",
		"/propostas/{id_proposta}/contas",
		"/propostas/{id_proposta}/contas/{id_conta}",
		"/cursos",
		"/cursos/{id_curso}",
		"/cursos/{id_curso}/materias",
		"/cursos/{id_curso}/materias/{id_materia}",
		"/clientes",
		"/clientes/{id_cliente}",
		"/clientes/{id_cliente}/pedidos",
		"/clientes/{id_cliente}/pedidos/{id_pedido}",
		"/produtos",
		"/produtos/{id_produto}",
		"/produtos/{id_produto}/avaliacoes",
		"/produtos/{id_produto}/avaliacoes/{id_avaliacao}",
		"/vendas",
		"/vendas/{id_venda}",
		"/vendas/{id_venda}/itens",
		"/vendas/{id_venda}/itens/{id_item}",
		"/fornecedores",
		"/fornecedores/{id_fornecedor}",
		"/fornecedores/{id_fornecedor}/produtos",
		"/fornecedores/{id_fornecedor}/produtos/{id_produto}",
		"/estoques",
		"/estoques/{id_estoque}",
		"/estoques/{id_estoque}/produtos",
		"/estoques/{id_estoque}/produtos/{id_produto}",
		"/pedidos",
		"/pedidos/{id_pedido}",
		"/pedidos/{id_pedido}/itens",
		"/pedidos/{id_pedido}/itens/{id_item}",
		"/pagamentos",
		"/pagamentos/{id_pagamento}",
		"/pagamentos/{id_pagamento}/transacoes",
		"/pagamentos/{id_pagamento}/transacoes/{id_transacao}",
		"/categorias",
		"/categorias/{id_categoria}",
		"/categorias/{id_categoria}/subcategorias",
		"/categorias/{id_categoria}/subcategorias/{id_subcategoria}",
		"/usuarios",
		"/usuarios/{id_usuario}",
		"/usuarios/{id_usuario}/permissoes",
		"/usuarios/{id_usuario}/permissoes/{id_permissao}",
		"/grupos",
		"/grupos/{id_grupo}",
		"/grupos/{id_grupo}/membros",
		"/grupos/{id_grupo}/membros/{id_membro}",
		"/projetos",
		"/projetos/{id_projeto}",
		"/projetos/{id_projeto}/tarefas",
		"/projetos/{id_projeto}/tarefas/{id_tarefa}",
		"/eventos",
		"/eventos/{id_evento}",
		"/eventos/{id_evento}/participantes",
		"/eventos/{id_evento}/participantes/{id_participante}",
		"/locais",
		"/locais/{id_local}",
		"/locais/{id_local}/salas",
		"/locais/{id_local}/salas/{id_sala}",
		"/equipamentos",
		"/equipamentos/{id_equipamento}",
		"/equipamentos/{id_equipamento}/manutencoes",
		"/equipamentos/{id_equipamento}/manutencoes/{id_manutencao}",
		"/relatorios",
		"/relatorios/{id_relatorio}",
		"/relatorios/{id_relatorio}/secoes",
		"/relatorios/{id_relatorio}/secoes/{id_secao}",
		"/configuracoes",
		"/configuracoes/{id_configuracao}",
		"/configuracoes/{id_configuracao}/opcoes",
		"/configuracoes/{id_configuracao}/opcoes/{id_opcao}",
		"/logs",
		"/logs/{id_log}",
		"/logs/{id_log}/eventos",
		"/logs/{id_log}/eventos/{id_evento}",
		"/notificacoes",
		"/notificacoes/{id_notificacao}",
		"/notificacoes/{id_notificacao}/destinatarios",
		"/notificacoes/{id_notificacao}/destinatarios/{id_destinatario}",
		"/mensagens",
		"/mensagens/{id_mensagem}",
		"/mensagens/{id_mensagem}/respostas",
		"/mensagens/{id_mensagem}/respostas/{id_resposta}",
		"/alertas",
		"/alertas/{id_alerta}",
		"/alertas/{id_alerta}/acoes",
		"/alertas/{id_alerta}/acoes/{id_acao}",
		"/documentos",
		"/documentos/{id_documento}",
		"/documentos/{id_documento}/versoes",
		"/documentos/{id_documento}/versoes/{id_versao}",
		"/formularios",
		"/formularios/{id_formulario}",
		"/formularios/{id_formulario}/campos",
		"/formularios/{id_formulario}/campos/{id_campo}",
		"/tabelas",
		"/tabelas/{id_tabela}",
		"/tabelas/{id_tabela}/colunas",
		"/tabelas/{id_tabela}/colunas/{id_coluna}",
		"/graficos",
		"/graficos/{id_grafico}",
		"/graficos/{id_grafico}/dados",
		"/graficos/{id_grafico}/dados/{id_dado}",
		"/dashboards",
		"/dashboards/{id_dashboard}",
		"/dashboards/{id_dashboard}/widgets",
		"/dashboards/{id_dashboard}/widgets/{id_widget}",
		"/mapas",
		"/mapas/{id_mapa}",
		"/mapas/{id_mapa}/camadas",
		"/mapas/{id_mapa}/camadas/{id_camada}",
		"/agendamentos",
		"/agendamentos/{id_agendamento}",
		"/agendamentos/{id_agendamento}/recursos",
		"/agendamentos/{id_agendamento}/recursos/{id_recurso}",
		"/assinaturas",
		"/assinaturas/{id_assinatura}",
		"/assinaturas/{id_assinatura}/planos",
		"/assinaturas/{id_assinatura}/planos/{id_plano}",
		"/preferencias",
		"/preferencias/{id_preferencia}",
		"/preferencias/{id_preferencia}/opcoes",
		"/preferencias/{id_preferencia}/opcoes/{id_opcao}",
		"/reservas",
		"/reservas/{id_reserva}",
		"/reservas/{id_reserva}/recursos",
		"/reservas/{id_reserva}/recursos/{id_recurso}",
		"/transacoes",
		"/transacoes/{id_transacao}",
		"/transacoes/{id_transacao}/detalhes",
		"/transacoes/{id_transacao}/detalhes/{id_detalhe}",
		"/configuracoes",
		"/configuracoes/{id_configuracao}",
		"/configuracoes/{id_configuracao}/opcoes",
		"/configuracoes/{id_configuracao}/opcoes/{id_opcao}",
		"/relatorios",
		"/relatorios/{id_relatorio}",
		"/relatorios/{id_relatorio}/secoes",
		"/relatorios/{id_relatorio}/secoes/{id_secao}",
	}
	// Criar um array com URLs
	// Ordenar o array de URLs para garantir que a busca binária funcione corretamente
	sort.Strings(urls)

	uri := "/cursos/{id_curso}/materias" // URL a ser encontrada
	// uri := "/propostas" // URL a ser encontrada

	start := time.Now()
	index := binarySearch(urls, uri)
	duration := time.Since(start)

	if index != -1 {
		fmt.Printf("URL encontrada no índice %d\n", index)
	} else {
		fmt.Println("URL não encontrada")
	}
	fmt.Printf("Tempo de execução: %v\n", duration)
}
