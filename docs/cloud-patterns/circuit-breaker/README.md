# :zap: :no_good: Circuit Breaker (Disjuntor)

## O que faz?

Evita processamento desnecessário, ficar esperando timeouts de comunicação entre serviços, sobrecarga nos demais
serviços e requisições inúteis quando um serviço ao qual dependemos (webapi, gateway, banco de dados, etc.)
está inoperante.

## Quando utilizar?

Quando determinada parte da nossa aplicação precisa requisitar algo de outro serviço precisamos nos preparar para evitar
falhas quando algo inesperado ocorrer durante a comunicação.

## Funcionamento:

- `Circuito` Faz a função de interagir com o serviço.
    - Tem 3 diferentes estados `Fechado`, `Aberto` e `Meio Aberto`.
        - Quando o Circuito está `Fechado` tudo está funcionando normalmente, não há erro na comunicação.
        - Quando o limite de tentativas de comunicação foi atingido e o erro permanece o Circuito fica `Aberto`.
            - Caso alguma tentativa de comunicação seja efetuada, a requisição ao serviço não é feita pois com Circuito
              está `Aberto`.
            - O Circuito fica aberto até que atingir timeout estabelecido. Quando este tempo é atingido seu estado vai
              para `Meio Aberto`.
        - Quando o Circuito está `Meio Aberto` é como se fosse um estado transitório. Na próxima chamada ao serviço é
          verificado se comunicação foi feita com sucesso. Caso sim, o Circuito volta ao estado `Fechado`, aceitando
          assim todas as novas requisições. Caso não volta para `Aberto` e o timeout zera e volta a não permitir novas
          chamadas.

## Referências:

- [Livro Cloud Native Go - Matthew A. Titmus](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/)
- [Circuit Breaker - Martin Fowler](https://martinfowler.com/bliki/CircuitBreaker.html)
