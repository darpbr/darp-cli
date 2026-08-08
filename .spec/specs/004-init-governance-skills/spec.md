# Spec 004 — Scaffold de skills de governança no `darp init`

## Status

Draft — depende da aprovação desta especificação e do Plan 004.

## Related Specifications

- [Spec 003 — Governança do desenvolvimento](../003-governance/spec.md)
- [Spec 001 — Inicialização de projeto DARP](../001-init/spec.md)
- [Spec 002 — `darp doctor`](../002-doctor/spec.md)

## Contexto

O `darp init` atualmente cria somente a skill `documentation`, além da
estrutura básica de contratos DARP. A Spec 003 define quatro skills de
governança que devem estar disponíveis em projetos DARP:

- `documentation`;
- `architecture`;
- `testing`;
- `release`.

Sem esta evolução, um projeto inicializado posteriormente não recebe o mesmo
conjunto de instruções que existe no repositório do DARP CLI.

## Problema

O bootstrap de um projeto DARP não é completo nem consistente com o contrato de
governança. Além disso, um scaffold que assuma Go, Java, Spring, Python,
FastAPI ou qualquer stack específica não será reutilizável para todos os
projetos que o DARP pretende suportar.

## Objetivos

1. Fazer o `darp init` criar as quatro skills de governança definidas na Spec
   003.
2. Fazer o bootstrap criar contratos completos de lifecycle e quality gates,
   quando esses arquivos ainda não existirem.
3. Manter o bootstrap agnóstico a linguagem, framework, ferramenta de build,
   sistema operacional, provedor de IA e tipologia de projeto.
4. Preservar arquivos e configurações já existentes, permitindo inicialização
   repetida e atualização segura de projetos parcialmente inicializados.
5. Garantir que o projeto criado seja aceito pelo `darp doctor`.

## Não objetivos

- detectar ou classificar o tipo do projeto;
- criar arquivos de Java, Spring, Python, FastAPI, Go ou outro ecossistema;
- instalar dependências, executar comandos do projeto ou acessar a rede;
- executar skills, workflows ou quality gates;
- alterar `.darp/workflows/implement.yaml` para executar todas as skills;
- substituir integralmente `README.md`, lifecycle, quality gates ou qualquer
  `SKILL.md` existente; `darp.yml` só pode receber a atualização aditiva
  definida nesta especificação;
- implementar um mecanismo genérico de plugins ou templates configuráveis;
- modificar o comportamento de `darp doctor` além dos testes necessários para
  validar o contrato já existente.

## Contrato de scaffold

Uma execução bem-sucedida em um diretório novo deve criar, no mínimo:

```text
.
├── darp.yml
├── .darp/
│   ├── lifecycle.md
│   ├── governance/
│   │   └── quality-gates.md
│   ├── workflows/
│   │   └── implement.yaml
│   └── templates/
└── .agents/
    └── skills/
        ├── documentation/SKILL.md
        ├── architecture/SKILL.md
        ├── testing/SKILL.md
        └── release/SKILL.md
```

Os quatro `SKILL.md` devem ser os contratos agnósticos definidos pela Spec
003. O conteúdo não pode exigir uma linguagem, framework, gerenciador de
dependências, layout ou ferramenta específica.

Cada skill deve orientar o agente a:

- ler o contexto e a documentação disponíveis no projeto;
- descobrir comandos e convenções existentes antes de recomendar ações;
- declarar assumptions;
- distinguir PASS, WARNING e BLOCKED;
- não inventar requisitos nem afirmar capacidades inexistentes;
- respeitar os limites da própria skill.

## Contrato do `darp.yml`

Em um projeto novo, `darp.yml` deve registrar as quatro skills com seus
caminhos canônicos:

```yaml
skills:
  documentation: .agents/skills/documentation
  architecture: .agents/skills/architecture
  testing: .agents/skills/testing
  release: .agents/skills/release
```

Em um projeto já existente, o `init` deve fazer uma atualização aditiva e
segura:

- adicionar somente entradas ausentes das quatro skills;
- preservar valores existentes para as mesmas chaves;
- preservar campos desconhecidos do YAML;
- preservar configurações de skills adicionais;
- rejeitar ou reportar YAML inválido sem substituir o arquivo original;
- não remover comentários ou conteúdo por uma simples reexecução, quando o
  mecanismo de atualização escolhido permitir sua preservação.

Se a atualização segura não puder ser feita, o comando deve preservar o
`darp.yml` original, reportar claramente o problema e ainda evitar sobrescrever
qualquer outro ativo existente.

## Regras de idempotência e preservação

1. Arquivos ausentes devem ser criados.
2. Diretórios ausentes devem ser criados.
3. Arquivos existentes nunca devem ser substituídos integralmente. A única
   exceção é a inclusão aditiva e segura das entradas de skills em `darp.yml`.
4. Uma skill parcialmente criada deve receber somente os arquivos ausentes.
5. A execução repetida não deve remover arquivos do usuário.
6. Um lifecycle ou quality gate customizado deve ser preservado.
7. Um `SKILL.md` customizado deve ser preservado integralmente.
8. Falhas parciais devem produzir erro explícito sem apagar o que já existia.

## Agnosticidade de projeto

O `darp init` não deve executar detecção de stack. O mesmo resultado estrutural
deve ser produzido em diretórios que contenham, por exemplo:

- uma aplicação Java com Spring;
- uma aplicação Python com FastAPI;
- um projeto Go;
- um repositório sem código ainda;
- um monorepo com múltiplas linguagens.

As skills podem recomendar a leitura de `README`, manifestos, scripts e CI do
projeto, mas não podem pressupor que qualquer desses arquivos exista.

## Contrato de conteúdo

Os conteúdos padrão devem ter uma única fonte canônica no código do CLI ou em
ativos versionados incorporados ao binário. O mecanismo escolhido deve:

- funcionar sem rede;
- produzir conteúdo determinístico;
- manter os mesmos nomes e caminhos em todos os projetos;
- permitir testes exatos do conteúdo mínimo;
- não depender do diretório atual do repositório fonte.

O lifecycle e o quality gate fornecidos pelo `init` devem ser compatíveis com a
Spec 003. O workflow existente continua sendo somente um contrato YAML e não
passa a executar skills automaticamente por causa desta spec.

## Critérios de aceitação

- [ ] Projeto novo recebe `darp.yml`, lifecycle, quality gates, workflow e as
      quatro skills.
- [ ] O `darp.yml` novo registra as quatro skills nos caminhos canônicos.
- [ ] Projeto já inicializado recebe skills e contratos ausentes sem perder
      arquivos ou conteúdo existente.
- [ ] Configurações existentes de `darp.yml`, incluindo campos desconhecidos e
      skills adicionais, são preservadas.
- [ ] Arquivos customizados de lifecycle, quality gates e skills não são
      sobrescritos.
- [ ] O comando não detecta stack nem executa ferramentas do projeto.
- [ ] O mesmo scaffold funciona em fixtures Java/Spring, Python/FastAPI, Go,
      monorepo e diretório vazio.
- [ ] A execução repetida é idempotente.
- [ ] `darp doctor` passa nos projetos scaffoldados.
- [ ] Testes unitários e de integração cobrem criação, reparo, preservação,
      YAML inválido e falhas parciais.
- [ ] Nenhuma skill criada contém instruções específicas de uma linguagem,
      framework ou provedor.

## Riscos e decisões

### Atualização de `darp.yml`

Adicionar entradas em um YAML existente pode alterar formatação ou comentários.
O plano deve escolher uma estratégia que preserve o máximo possível do arquivo
e testar campos desconhecidos, skills extras, duplicatas e YAML inválido.

### Conteúdo duplicado entre Spec 003 e CLI

Os templates do CLI precisam ser atualizados quando o contrato de governança
mudar. A implementação deve documentar a fonte canônica escolhida e incluir
testes que detectem conteúdo mínimo incompatível.

### Projetos parcialmente inicializados

O comando deve reparar apenas o que falta. Não deve transformar a reparação em
uma migração destrutiva nem substituir decisões locais do projeto.
