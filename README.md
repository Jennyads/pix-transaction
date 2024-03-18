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
