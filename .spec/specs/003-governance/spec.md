# Spec 003 — Governança do desenvolvimento

## Status

Approved and implemented — reviewed and validated on 2026-08-07.

Esta especificação define os artefatos, limites e critérios usados para a
implementação documentada neste repositório.

## Contexto atual

O repositório já possui:

- `.darp/lifecycle.md`, com conteúdo real, porém em formato embrulhado como
  código e ainda sem as fases de validação completas;
- `.darp/governance/quality-gates.md`, atualmente como placeholder;
- `.darp/workflows/implement.yaml`, atualmente com somente a skill
  `documentation`;
- `.agents/skills/documentation/SKILL.md`;
- `.agents/skills/security-review/`, uma skill existente e independente desta
  especificação;
- `darp init` e `darp doctor`, especificados respectivamente em 001 e 002.

Portanto, a implementação desta spec deve completar os artefatos documentais
de governança sem reimplementar ou alterar comandos existentes.

## Problema

O projeto não possui uma definição operacional única para as etapas posteriores
à implementação. Também não há contratos suficientemente detalhados para as
skills de arquitetura, testes, documentação e release, nem critérios de
qualidade que permitam a um agente concluir uma mudança de forma verificável.

## Objetivos

1. Definir o lifecycle completo de engenharia do DARP, incluindo as referências
   às skills responsáveis por cada revisão.
2. Definir quality gates objetivos, com critérios, responsável e resultado
   esperado.
3. Criar quatro skills de governança com instruções suficientes para execução
   por modelos de IA: `documentation`, `architecture`, `testing` e `release`.
4. Atualizar o README para explicar governança, skills, lifecycle e gates.
5. Manter cada etapa auditável, determinística e independente de provedor de IA.

## Não objetivos

- alterar comandos, flags, códigos de saída ou comportamento de `darp init`;
- alterar o comportamento de `darp doctor`;
- executar skills ou quality gates automaticamente;
- criar um novo comando de governança, workflow runner ou sistema de CI;
- substituir ou modificar a skill `security-review` existente;
- criar prompts, exemplos, scripts ou templates auxiliares sem necessidade
  demonstrada pela própria skill;
- definir regras de release do produto além do checklist documental desta spec.

## Escopo e artefatos

### Governança

Os seguintes arquivos devem existir e ser completos ou alinhados ao contrato:

```text
.darp/
├── lifecycle.md
└── governance/
    └── quality-gates.md
```

`lifecycle.md` deve ser um documento Markdown renderizável, sem um bloco de
código externo envolvendo todo o arquivo. Deve descrever, nesta ordem, as
fases:

```text
Vision → Constitution → Project Context → Specification → ADR → Plan → Tasks
→ Implementation → Tests → Architecture Review → Documentation Review
→ Quality Gates → Review → Release → Completed
```

O lifecycle deve explicar o propósito, entrada, saída e condição de passagem
de cada fase. Deve mencionar as skills de governança nas fases de review,
explicar que detalha o fluxo resumido da Constituição e afirmar que a
implementação só começa após specification, plan e tasks aprovados.

`quality-gates.md` deve definir os gates Build, Tests, Documentation,
Architecture, Compatibility e Release Notes. Para cada gate, informar:

- objetivo;
- critérios verificáveis;
- skill ou responsável pela revisão;
- evidência esperada;
- condição PASS e condição BLOCKED/FAIL.

Os gates são critérios documentais nesta spec; sua execução automática está
fora do escopo.

### Skills

Cada skill deve possuir, no mínimo:

```text
.agents/skills/<nome>/
└── SKILL.md
```

As quatro skills devem ter front matter YAML com `name` e `description`, e o
`SKILL.md` deve conter, no mínimo, as seções `When to Use`, `Responsibilities`,
`Workflow`, `Expected Output` e `Boundaries`.

Requisitos específicos:

- `documentation`: compara documentação com a implementação, atualiza
  README e documentação técnica quando necessário, valida links internos e
  não documenta comportamento inexistente; não modifica código.
- `architecture`: avalia coesão, acoplamento, responsabilidades, duplicação e
  aderência à constituição; produz recomendações e não modifica código.
- `testing`: identifica cobertura ausente, regressões, cenários de erro e
  inconsistências entre testes e implementação; registra comandos e evidências
  utilizados.
- `release`: revisa changelog/release notes, breaking changes, features,
  compatibilidade e itens que precisam de decisão antes de uma release; não
  publica artefatos.

Cada skill deve declarar que não deve inventar requisitos, deve explicitar
assumptions e deve distinguir PASS, WARNING e BLOCKED quando produzir uma
revisão.

### README

O README deve incluir links ou referências para:

- o conceito de governança;
- as quatro skills e suas responsabilidades;
- os quality gates;
- o lifecycle;
- a separação entre planejamento em `.spec/`, contratos em `.darp/` e skills
  em `.agents/skills/`.

O texto não deve afirmar que o CLI executa as skills ou gates.

### Instruções e templates

Também fazem parte do escopo desta spec:

- `AGENTS.md`, que deve refletir a ordem operacional do lifecycle e sua relação
  com a Constituição;
- `.spec/templates/tasks.template.md`, que deve exigir tasks em formato de
  checkbox com validação associada.

## Requisitos de compatibilidade

1. A mudança não altera comportamento executável. Ela pode substituir ou
   reorganizar conteúdo documental existente, preservando regras válidas.
2. Nenhum arquivo de `cmd/`, `internal/` ou `pkg/` deve ser alterado.
3. `darp.yml`, `.darp/workflows/implement.yaml`, testes Go e o comportamento
   de `darp init` e `darp doctor` não devem ser alterados.
4. Os arquivos existentes que não pertencem ao escopo devem permanecer
   inalterados, especialmente a skill `security-review`.
5. As quatro novas skills devem ser válidas para o `darp doctor`: cada uma
   deve possuir `SKILL.md` válido e nenhum recurso auxiliar inválido.
6. Links e caminhos documentados devem apontar para arquivos ou diretórios
   existentes no estado final.
7. A estrutura deve continuar compreensível para agentes que não conhecem o
   histórico das specs anteriores.

## Critérios de aceitação

- [x] Os dois artefatos de governança existem e atendem ao contrato definido.
- [x] O lifecycle contém todas as fases na ordem definida e referencia as
      skills de review.
- [x] Os seis quality gates possuem objetivo, critérios, responsável, evidência
      e resultado de passagem/bloqueio.
- [x] As quatro skills existem, têm front matter válido e todas as seções
      obrigatórias.
- [x] As skills não prometem execução automática nem alteração fora de seus
      limites.
- [x] O README documenta o fluxo final sem afirmar capacidades inexistentes.
- [x] `AGENTS.md` e o template de tasks estão alinhados ao lifecycle e exigem
      checkboxes validados.
- [x] `darp doctor` passa após a criação das quatro skills, sem alteração de
      `darp.yml` ou do workflow.
- [x] `git diff --name-only` confirma que nenhum arquivo de implementação foi
      alterado.
- [x] `go test ./...` continua passando.

## Pontos abertos e decisões registradas

Os seguintes pontos foram encontrados na revisão e ficam deliberadamente fora
da spec-003:

1. **Scaffold do `darp init`:** hoje ele cria somente a skill de documentação.
   Fazer o `init` criar as quatro skills exigirá uma alteração de implementação
   e testes próprios. Deve ser uma nova spec, pois esta spec proíbe alterar o
   comando.
2. **Execução de workflow:** `implement.yaml` não é um executor. Esta spec não
   define runner, ordem automática nem integração com modelos.
3. **Enforcement dos gates:** `darp doctor` valida contratos estruturais, mas
   não deve virar um executor de reviews por causa desta spec.
4. **Skill `security-review`:** ela já existe e não faz parte do conjunto de
   quatro skills criado aqui. Sua integração ao lifecycle poderá ser decidida
   em uma spec de segurança/review posterior.

Se o objetivo do produto for incluir qualquer um desses comportamentos, esta
especificação deve permanecer Draft até que o escopo seja alterado e os testes
necessários sejam adicionados.
