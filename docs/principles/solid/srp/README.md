# :wrench: :hammer: Single Responsability Principle

## Um módulo deve ser responsável por um, e apenas um, ator.

É preciso separar o código de qual atores diferentes dependam.

No nível de componentes o SRP se torna o Coomon Closure Principle (Princípio do Fechamento Comum). No nível arquitetural é o Axis of Change (Eixo da Mudança) responsável pela criação de Limites Arquiteturais.

### *Exemplo de Violação:*

__Descrição do cenário inicial:__

* Dado que tenho uma classe *__Empregado__* que tem os métodos __CalcularPagamento__, __ReportarHoras__ e __Salvar__.
* Dado que tenho vários __atores__ que tem interesse em um desses métodos. Imagine que é solicitado a uma das equipes alterar o método __CalcularPagamento__.
* No entanto este método utiliza o método *privado* __HorasTrabalhadas__, e o desenvolvedor percebe que precisa mudar o comportamento interno deste método.
* Ele não observa que este método também é chamado pelo __ReportarHoras__ e para quem utiliza este método, ele não precisaria de modificação no comportamento.
* Esse tipo de alteração causará um erro nos dados apresentados podendo causar até mesmo um prejuízo financeiro.



