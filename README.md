<h2> Projeto Diamante - Desenvolvimento de um Sistema de Pagamentos via PIX </h2>

Esse projetoO projeto propôs o desenvolvimento de um Sistema de Pagamentos via PIX com funcionalidades de gerenciamento de usuários e transações, incluindo:

* Gerenciamento de Usuários: Foram implementadas operações CRUD (Create, Read, Update, Delete) para criar, obter, listar e atualizar usuários, bem como criar conta e chaves PIX, atualizar e remover chaves associadas a esses usuários.
* Transações PIX: Permitiu-se a realização de transferências PIX entre usuários.
* Extrato de Movimentações PIX: Possibilitou-se a listagem das movimentações PIX de um usuário, oferecendo um extrato detalhado das transações realizadas.

Sendo assim, adotou-se medidas para promover uma melhor organização do projeto em níveis teóricos e práticos. Na questão de gerenciamento de tarefas, foi criado um board no Miro para organizar ideias de arquitetura, diagramas e modelos de proto, planejamento de tarefas (TO DO e DONE), indicações de materiais de apoio, além de registros de estudos e fluxos do andamento do projeto. 

Considerando-se que Projeto Diamante teve como finalidade desenvolver um sistema de pagamentos via PIX completo, com funcionalidades de gerenciamento de usuários e transações. Para garantir organização e escalabilidade, a estrutura do projeto segue o padrão de microsserviços, dividindo o sistema em módulos independentes e coesos.

<h3> Organização em Microsserviços: </h3>
Cada microsserviço é dedicado a uma funcionalidade específica do sistema, como gerenciamento de usuários (microsserviço de profile), transações PIX ou gerenciamento de chaves (microsserviço de transaction). Essa divisão promove autonomia, facilitando o desenvolvimento, testes e manutenção de cada módulo.

<h3> Padrão de Diretórios: </h3>
Para garantir uniformidade e organização interna, cada microsserviço segue o mesmo padrão de diretórios:

```
docker:
Contém os arquivos relacionados à configuração e construção do ambiente Docker para aplicação.
Dockerfile: É um arquivo de configuração que contém as instruções para construir uma imagem Docker da aplicação. Ele descreve os passos necessários para configurar o ambiente de execução da aplicação dentro de um container Docker.
microsservico:
    cmd: Contém o código principal do microsserviço, responsável por inicializar e executar as funções principais.
    internal: Abriga os componentes internos do microsserviço, divididos em subdiretórios:
        cfg: Armazena as configurações do microsserviço, como variáveis de ambiente e parâmetros de conexão.
        errutils: Fornece funções utilitárias para lidar com erros e exceções.
        event: Implementa mecanismos de comunicação assíncrona entre microsserviços, utilizando eventos.
        utils: Contém funções utilitárias genéricas para o microsserviço.
        dto: Define os objetos de transferência de dados (DTOs) utilizados na comunicação entre o microsserviço e outros sistemas.
            model.go: Define a estrutura dos modelos de dados do microsserviço.
            repository.go: Implementa as operações de acesso ao banco de dados para os modelos do microsserviço.
            service.go: Implementa a lógica de negócio do microsserviço, utilizando os modelos e repositórios.
        webhook:Contém os arquivos relacionados ao gerenciamento de webhooks, que são mecanismos utilizados para que um aplicativo possa ser notificado automaticamente quando ocorrem              determinados eventos em outro sistema.
    platform: Integra o microsserviço com plataformas de terceiros, como:
        kafka: Implementa a comunicação com o sistema de filas Kafka para mensagens assíncronas.
        database: Contém os drivers de acesso ao banco de dados utilizado pelo microsserviço.
    proto: Define as interfaces de comunicação entre o microsserviço e outros sistemas, utilizando Protobuf.
```

<h2> ⭐ Agradecimentos ⭐</h2>
<h4>Victor Hugo Vieira Cruz</h4>

Gostaria de expressar minha gratidão ao meu tutor, Victor Hugo Vieira Cruz, pelos ensinamentos, paciência e pelo entusiasmo em me guiar na construção de um projeto do zero, mesmo com meu conhecimento limitado. A jornada ao lado dele me proporcionou um entendimento muito mais profundo sobre arquitetura de software, padrões de projeto, a linguagem Go e todas as ferramentas envolvidas.

Apesar de sua jovem idade, Victor demonstra uma impressionante amplitude de conhecimento que certamente impacta qualquer pessoa que tenha a oportunidade de trabalhar com ele. Sua orientação foi fundamental para que eu aprendesse tanto com alguém mais jovem, mas que está tão à frente em termos de habilidades técnicas e experiência de vida. Este projeto não apenas enriqueceu meu conhecimento técnico, mas também me proporcionou valiosas lições de vida para minha carreira e desenvolvimento pessoal.

Sou imensamente grata pela inspiração e sabedoria que Victor compartilhou comigo. Sua personalidade e experiências são verdadeiramente inspiradoras, e sou privilegiada por tê-lo como tutor e mentor. Este projeto foi muito mais do que uma simples jornada técnica; foi uma jornada de crescimento pessoal e profissional. Finalizo extremamente feliz de ter ganhado uma inspiração de carreira e de ser humano, quero crescer muito mais depois dessa abertura de janela e ver o Victor brilhar muito como espero que vai!

<h4>Felipe Augusto Amaral Thomaz</h4>

Gostaria de expressar minha sincera gratidão pelo olhar dedicado que você tem para cada um de seus liderados no DmPag. Desde sua chegada, sua abordagem tem sido nada menos que inspiradora. 

Sua proposta, monitoramento cuidadoso, ideias acolhedoras e conselhos  têm sido fundamentais para trazer uma nova luz, organização e apoio à equipe.
É notável o impacto que você trouxe para o nosso ambiente de trabalho. Sua presença trouxe um novo nível de eficiência e colaboração, e estou extremamente grato por isso. 

Sinto-me  realizada por ter a oportunidade de trabalhar ao seu lado e testemunhar o seu trabalho excepcionalmente bem feito, principalmente no Projeto Diamante.
Obrigado por sua dedicação contínua e pelo compromisso com o time. Você é um verdadeiro ativo para o nosso DMPag.

<h4>Cristiano da Silva Andrade</h4>

Gostaria de expressar minha gratidão pela oportunidade de aprendizado que você proporcionou e pela generosa liberação de recursos, tempo e pessoal para promover um melhor entendimento da nossa área de atuação. Este compromisso foi crucial para a evolução gradual do nosso trabalho.

É notável como sua dedicação possibilita que o DMPag se mantenha sempre vivo e inovador. Cada recurso, cada hora investida, e cada pessoa envolvida foram elementos essenciais para o progresso que alcançamos até agora.

Sou grata por sua liderança visionária e por reconhecer a importância de apoiar iniciativas de aprendizado e desenvolvimento. Este ambiente de suporte é fundamental para cultivar uma cultura de crescimento e excelência.
Mais uma vez, obrigada por tornar possível este caminho de evolução e inovação. 



