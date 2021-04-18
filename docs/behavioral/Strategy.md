# Strategy

* Características:

- Contexto
    Tem uma referência do strategy e invoca ele.
- IStrategy
    Define a interface para ser dada ao Strategy. Se trata do contrato.
- Strategy
    Implementação concreta do Strategy.

    - Exemplo:
        - Contexto: Chama IStrategy.CriarCobranca(...)
        - IStrategy: Define o contrato. CriarCobranca(Pedido pedido)
        - Strategy: PDFCobrancaStrategy, EmailCobrancaStrategy,
          ImpressaoCobrancaStrategy.

- A implementação é selecionada à partir dos dados inputados pelo usuário sem
  necessidade de extender a classe.

- O objetivo principal de usar o Strategy é fazer com que o código seja mais
  extensível e mais testável.

- A interface define o contrato que cada classe concreta do Strategy usará.

- O contexto não deve saber detalhes da implementação apenas utilizar a
  interface.

- O Strategy é um padrão de comportamento, então na realidade qual estratégia
  será escolhida dependerá dos dados informados pelo usuário.

- Classes Strategy não necessariamente precisam ter o sufixo 'Strategy'.

- Strategy nos permite criar aplicações extensíveis e dinâmicas que podem ser
  facilmente mudadas a qualquer momento.

- Mas precisamos tomar cuidado para não deixar as aplicações mais complexas do
  que deveriam.

- Strategy utilizando interfaces fica muito mais fácil o teste dos componentes
  de software.

- Sempre que você usa de uma Interface e injeta para permitir mudança de
  comportamento, você está utilizando o padrão Strategy.

* Quando utilizar?

* Resumo

- Um dos padrões mais comuns.

- Desacoplo o contexto de uso da implementação concreta.

- Permite uma implementação mais limpa no contexto que chama.

- Facilmente extensível, permite a adição de novos Strategys sem afetar os que
  já existem.

- Faz os testes bem mais fáceis, é mais simples de mockar e injetar as
  interfaces.
