# Spec 001 — Inicialização de Projeto DARP

Status: Draft

## Objetivo

Implementar o primeiro comando do darp-cli:

```bash
darp init
```

Este comando é responsável por transformar um diretório comum em um projeto DARP.

Esta feature estabelece a base para todas as funcionalidades futuras do CLI.

---

## Motivação

Todos os demais comandos pressupõem que existe um projeto DARP inicializado.

A inicialização deve criar a estrutura mínima necessária para permitir a evolução incremental do projeto, sem adicionar funcionalidades que ainda não existem.

---

## Escopo

Esta especificação contempla exclusivamente:

- criação da estrutura inicial do projeto
- criação do arquivo `darp.yaml`
- criação da pasta `.darp`
- cópia do arquivo `lifecycle.md`
- mensagens de progresso
- comportamento idempotente

Não faz parte desta especificação:

- criação de specs
- geração de tarefas
- integração com IA
- GitHub
- MCP
- providers
- skills
- templates inteligentes

---

## Estrutura esperada

Após a execução bem-sucedida do comando, o projeto deverá possuir:

.
├── darp.yaml
└── .darp
    ├── lifecycle.md
    ├── specs
    ├── tasks
    ├── docs
    ├── templates
    └── examples

---

## darp.yaml

O arquivo deverá possuir, inicialmente, apenas:

```yaml
name: <nome-da-pasta>
version: 1
specVersion: 0.1
```

O campo `name` deve utilizar o nome do diretório atual.

---

## lifecycle.md

O arquivo deverá ser embarcado no binário do CLI.

Durante a execução do comando, ele deverá ser copiado para:

```text
.darp/lifecycle.md
```

Não deve haver download de arquivos externos.

---

## Fluxo

### Caso 1

Projeto ainda não inicializado.

Resultado esperado:

- cria `darp.yaml`
- cria `.darp`
- cria subdiretórios
- copia `lifecycle.md`
- informa sucesso

---

### Caso 2

Projeto já inicializado.

Resultado esperado:

- não sobrescreve arquivos
- não remove arquivos
- informa que o projeto já está inicializado
- encerra normalmente

---

## Idempotência

Executar:

```bash
darp init
```

duas ou mais vezes nunca deve causar perda de dados.

---

## Mensagens

Durante a execução devem ser exibidas mensagens amigáveis.

Exemplo:

```text
✔ Inicializando projeto

✔ Criando darp.yaml

✔ Criando estrutura .darp

✔ Copiando lifecycle.md

✔ Projeto inicializado
```

As mensagens podem evoluir futuramente.

---

## Arquitetura

A implementação deve separar claramente:

- camada CLI
- regras de negócio
- acesso ao sistema de arquivos

Evitar lógica de negócio diretamente no comando CLI.

A implementação deve privilegiar baixo acoplamento e facilidade de testes.

---

## Critérios de Aceitação

A feature será considerada concluída quando:

- o projeto compilar
- todos os testes passarem
- `darp init` criar corretamente a estrutura esperada
- a execução repetida não alterar arquivos existentes
- nenhum outro comando for implementado nesta etapa

---

## Fora de Escopo

Esta especificação não contempla nenhuma funcionalidade além da inicialização do projeto.
