# :lock: :key: Open-Closed Principle

## Fechado para modificação e aberto para ampliação.

Somente um motivo para mudança.

"Quando uma única mudança em um programa resulta em uma sucessão de mudança nos módulos dependentes, o projeto tem o mau cheiro da rigidez."

A chave para não violar o OCP é a abstração.
Uma das formas do código obedecer o OCP é usando o padrão [Strategy](../../../design-patterns/behavioral/strategy/README.md) ou utilizando [Template Method](../../../design-patterns/behavioral/template-method/README.md). Esses dois padrões representam uma clara separação entre a funcionalidade genérica e a implementação detalhada dessa funcionalidade.

### *Exemplo de Violação:*

__Descrição do cenário inicial:__

* Tenho uma lista de *__`Pessoas`__* e a partir dela crio registro de *__`Empregados`__* aos quais verifico sua contabilidade através da classe *Conta*.

* Ao final imprimo o valor da *__`Conta`__* de cada um desses empregados.

* A classe que cria a *__`Conta`__* é o __escopo__ de estudo. Ela deve ser aberta para extensão e fechada para modificação.

* Isso quer dizer que novos cenários de criação de *__`Conta`__* devem ser implementados por meio de extensão e não por modificar o que já fora implementado.

Classe *__ `Pessoa` __*:

```csharp
namespace Modelos
{
    public class Pessoa
    {
        public string Nome { get; set; }
        public string Sobrenome { get; set; }
    }
}
```

Classe *__ `Empregado` __*:

```csharp
namespace Modelos
{
    public class Empregado
    {
        public string Nome { get; set; }
        public string Sobrenome { get; set; }
        public string Email { get; set; }
    }
}
```

Classe *__ `Conta` __*:

```csharp
namespace Modelos
{
    public class Conta
    {
        public Empregado Criar(Pessoa pessoa)
        {
            return new Empregado
            {
                Nome = pessoa.Nome,
                Sobrenome = pessoa.Sobrenome,
                Email = $"{pessoa.Nome.Substring(0, 1)}{pessoa.Sobrenome}@acme.com"
            };
        }
    }
}
```

__Descrição do cenário primeira modificação:__

* Dado que surge a necessidade de termos registro de *__`Empregado`__* que possa ser ou não __Gerente__. Ainda mais se ele for gerente o comportamento da sua *Conta* é diferente de um empregado normal.

:warning: :warning: __POSSÍVEL VIOLAÇÃO:__ :warning: :warning:

* Talvez o comportamento mais instintivo seria alterar o código de *__`Empregado`__* adicionando a propriedade __EhGerente__, booleana.

* Isso de certa forma, e para o cenário que estamos propondo, seria uma __violação__ ao OCP. Afinal temos um código de produção que está sendo alterado.

* Isso trás a tona o questionamento:

### __Quando é uma boa ideia implementar ou observar o OCP, e quando não seria?__

* Resposta: durante o processo de desenvolvimento, talvez não faça sentido ser aplicado. Por quê? Porque dentro do processo de desenvolvimento é normal as coisas mudarem com frequência, e pensar em não violar o OCP num primeiro momento de descoberta pode ser difícil. O que nós realmente queremos __garantir__ é que quando o código for para produção o único motivo para mudança do código seja um :bug: __Bug__ em produção.

* No entanto, como todos os princípios, é preciso avaliar o quão rigorosos "à doutrina" vamos ser. Talvez nesse caso de adicionar a propriedade __EhGerente__ não seja uma violação impeditiva. Ainda mais se colocarmos um valor padrão, o comportamento da classe *__`Conta`__* em nada muda.

* Mudanças menores devem ser avaliadas, se não houver efeito colateral talvez valem a pena violar o princípio.

* :thinking: Pensando diferente: Se fosse o caso de:
    - Mudarmos a classe *__`Empregado`__* para receber uma interface de *__`IGerente`__* que tenha inicialmente alguns dados.
    - Isso sim poderia ser um 'breaking changes', e a violação ao OCP provavelmente deveria ser pensada com mais cuidado.

### Indo mais afundo :brain:

* Imaginando outro cenário. Digamos que vamos criar uma enumeração para a classe *__`Pessoa`__* e ela terá o *TipoDoEmpregado*. Sendo da seguinte forma:

```csharp
namespace Modelos
{
    public enum TipoDoEmpregado
    {
        Empregado,
        Gerente
    }
}
```

```csharp
namespace Modelos
{
    public class Pessoa
    {
        public string Nome { get; set; }
        public string Sobrenome { get; set; }
        public TipoDoEmpregado TipoDoEmpregado { get; set; } = TipoDoEmpregado.Empregado;
    }
}
```

* Nesse momento queremos fazer o seguinte: caso a *Pessoa* seja do tipo gerente, na hora de criar o registro de *Conta* quero que seja marcado *EhGerente* como *true*.

:radioactive: :radioactive: __VIOLAÇÃO__ :radioactive: :radioactive:

* A classe *Conta* agora vai precisar ser modificada para adicionar o comportamento:

```csharp
namespace Modelos
{
    public class Conta
    {
        public Empregado Criar(Pessoa pessoa)
        {
            var empregado = new Empregado
            {
                Nome = pessoa.Nome,
                Sobrenome = pessoa.Sobrenome,
                Email = $"{pessoa.Nome.Substring(0, 1)}{pessoa.Sobrenome}@acme.com"
            };

            // Violação do OCP 
            if (pessoa.TipoDoEmpregado == TipoDoEmpregado.Gerente)
            {
                empregado.EhGerente = true;
            }
        }
    }
}
```

* O código só acaba tendo o requisito cumprido se nós modificarmos o comportamento interno das classes já implementadas.

* E se amanhã ou depois surgir um novo *TipoDeEmpregado*? Precisaremos modificar novamente (por exemplo: *TipoDeEmpregado. Executivo*).

### Consertando as coisas :hammer_and_wrench: :gear:

* Uma das formas de pensar uma refatoração que atenderia o requisito de negócio, e não feriria o OCP seria primeiramente parar de depender da classe concreta *Pessoa* e depender da interface.

* A interface poderia chamar *ICandidato* e a classe *Pessoa* implementá-la.

* O mesmo pode ser feito com a classe *Conta*, extraindo uma interface *IConta*.

* Agora como que faríamos para a classe *Pessoa* saber se vai criar a __conta__ do *Gerente* ou *Funcionario*. Poderia ser feita uma alteração na classe *Pessoa* para que ela soubesse a instância de *IConta* que utilizará, e não delegar essa responsabilidade para dentro de *Conta*.

* No caso do *Gerente* e outras especializações, estas implementariam a interface *ICandidato* e teriam suas especificidades bem como sua própria implementação de *IConta*.

* Consequências no código:

```csharp
using Modelos;

namespace Aplicacao
{
    public class Programa
    {
        public static void Main()
        {
            var candidatos = new List<ICandidato>
            {
                new Pessoa { Nome= "Filipe", Sobrenome = "Souza" },
                new Gerente { Nome= "Outra", Sobrenome = "Pessoa" }
            }
        }
    }
}
```

- Classe *Pessoa*:

```csharp
namespace Modelos
{
    public class Pessoa : ICandidato
    {
        public string Nome { get; set; }
        public string Sobrenome { get; set; }
        public IConta Conta { get; set; } = new Conta();
    }
}
```

- Classe *Gerente*:

```csharp
namespace Modelos
{
    public class Gerente : ICandidato
    {
        public string Nome { get; set; }
        public string Sobrenome { get; set; }
        public IConta Conta { get; set; } = new ContaDeGerente();
    }
}
```

### Resumo para o dia-a-dia

* Sempre que estiver diante de uma possível modificação de código, pare, respire e pense se vai violar o OCP.

* Uma forma boa de remover o acoplamento e extender comportamentos é por meio de interfaces.

* Nem sempre todos os módulos do software poderão obedecer o OCP, a questão é medir quando vale ou não a pena violar.

* Obdecer o OCP pode ser dispendioso. É necessário tempo e esforço para construir as abstrações adequadas. Abstrações também podem aumentar a complexidade do código.

* Podemos usar a estratégia do "Engane-me uma vez": Escrevemos o código inicalmente esperando que ele não vai mudar. Quando ocorrer uma mudança, implementamos as abstrações que nos protegem de mudanças futuras desse tipo (Toma o primeiro tiro).

* Descobrir as abstrações também é uma boa estratégia. Fazemos isso por desenvolver desde o início começando pelos testes (ajuda na descoberta e cria um bom alicerce para evolução), tendo ciclos curtos de entregar com feedback do usuário e desenvolvendo primeiro as coisas mais importantes. Esse é o cenário de desenvolvimento que vai ser o ideal para evolução de software e suas abstrações.

### Por que violar o OCP é ruim?

* Alteramos comportamentos que já funcionam e isso pode gerar bugs :bug: :bug:. Se alteramos o comportamento isso significa mexer nos testes de unidade e talvez nos lugares onde chamam este método.

* Dependendo do tipo de software que estamos lidando, alteração em uma classe existente pode desencadear uma nova compilação do módulo/namespace em que ela se encontra. Dependendo do controle de versão esse novo código alterado pode ser nocivo, seria um binário novo.

* O OCP traz os maiores benefícios de se usar linguagens OOP: flexibilidade, capacidade de reutilização e facilidade de manutenção. Lembre-se de evitar a "abstração desenfreada".

fontes: [IAmTimCorey Youtube Channel](https://youtu.be/VFlk43QGEgc), Princípios, Padrões e Práticas Ágeis em C# - Robert C. Martin e Micah Martin.
