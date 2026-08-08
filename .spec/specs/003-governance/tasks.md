# Tasks 003 — Governança do desenvolvimento

## Status

Completed — implementation and validation completed on 2026-08-07.

## Related Plan

- [`plan.md`](plan.md)

## Regras para o implementador

- Marque cada checkbox imediatamente após concluir a ação e sua validação.
- Não marque uma task parcialmente concluída.
- Se uma task ficar bloqueada, mantenha-a desmarcada e registre o motivo abaixo.
- Não altere arquivos fora do escopo aprovado.

## Task List

### T1 — Completar o lifecycle

- [x] Remover o bloco de código externo de `.darp/lifecycle.md`, preservando o
      conteúdo válido existente.
- [x] Reconciliar `.darp/lifecycle.md` com as fases e a ordem definidas na spec.
- [x] Para cada fase, documentar propósito, entrada, saída e condição de
      passagem.
- [x] Referenciar explicitamente `architecture`, `testing`, `documentation` e
      `release` nas fases aplicáveis.
- [x] Explicitar que implementação depende de specification, plan e tasks
      aprovados.
- [x] Explicar que o lifecycle detalha o fluxo resumido da Constituição.
- [x] Substituir ou alinhar a seção `Future Evolution` para não contradizer o
      lifecycle expandido.
- [x] Verificar que todos os caminhos citados existem ou são criados por outra
      task desta lista.

### T2 — Definir quality gates

- [x] Atualizar `.darp/governance/quality-gates.md`.
- [x] Criar as seções Build, Tests, Documentation, Architecture,
      Compatibility e Release Notes.
- [x] Em cada gate, registrar objetivo, critérios, responsável, evidência e
      resultado PASS/BLOCKED ou FAIL.
- [x] Declarar que os gates são critérios de revisão e não execução automática
      pelo CLI nesta spec.

### T3 — Padronizar e completar skills

- [x] Atualizar `.agents/skills/documentation/SKILL.md` para atender o contrato
      de seções e responsabilidades da spec.
- [x] Criar `.agents/skills/architecture/SKILL.md`.
- [x] Criar `.agents/skills/testing/SKILL.md`.
- [x] Criar `.agents/skills/release/SKILL.md`.
- [x] Garantir em cada skill front matter YAML com `name` e `description`.
- [x] Garantir em cada skill as seções `When to Use`, `Responsibilities`,
      `Workflow`, `Expected Output` e `Boundaries`.
- [x] Declarar em cada skill como registrar assumptions e estados PASS,
      WARNING e BLOCKED.
- [x] Confirmar que nenhuma skill altera código ou publica artefatos fora do
      limite definido.

### T4 — Atualizar instruções de planejamento

- [x] Atualizar `.spec/templates/tasks.template.md` para que cada task seja um
      checkbox (`- [ ]`) e tenha validação associada.
- [x] Atualizar `AGENTS.md` para exigir checkboxes nas tasks e registro de
      bloqueios/evidências.
- [x] Manter a regra de que specs, plans e tasks ativos ficam em
      `.spec/specs/<number>-<short-name>/`.

### T5 — Atualizar documentação de entrada

- [x] Atualizar `README.md` com governança, skills, lifecycle e quality gates.
- [x] Alinhar o diagrama Mermaid e o fluxo textual do README à mesma sequência
      de fases do lifecycle, incluindo `Review` e `Completed`.
- [x] Documentar a responsabilidade de `.spec/`, `.darp/` e
      `.agents/skills/` sem duplicar contratos divergentes.
- [x] Remover ou corrigir qualquer afirmação de execução automática de skills
      ou gates.
- [x] Validar os links Markdown e caminhos citados no README.
- [x] Confirmar que o README não afirma que o CLI executa skills ou gates.

### T6 — Validar escopo e regressão

- [x] Confirmar que os artefatos obrigatórios da spec existem.
- [x] Confirmar front matter e seções obrigatórias das quatro skills.
- [x] Confirmar que cada quality gate possui todos os campos exigidos.
- [x] Executar `git diff --check`.
- [x] Executar `go test ./...`.
- [x] Executar `darp doctor` em modo somente leitura e registrar o resultado.
- [x] Confirmar que o `doctor` valida as quatro novas skills sem exigir
      alteração em `darp.yml`.
- [x] Confirmar com `git diff --name-only` que não houve alteração em
      `cmd/`, `internal/`, `pkg/`, testes Go, `darp.yml`, workflows ou na skill
      `security-review`.
- [x] Revisar o diff final contra `spec.md`, `plan.md` e esta lista de tasks.

## Validation Checklist

- [x] Spec 003 aprovada.
- [x] Plan 003 aprovado.
- [x] Todas as tasks T1–T6 concluídas.
- [x] Todos os quality gates documentais atendidos.
- [x] Testes existentes passando.
- [x] Nenhum bloqueio pendente.

## Notas e bloqueios

Registrar aqui dúvidas, decisões tomadas durante a implementação e qualquer
task que precise de uma nova especificação. Não resolver silenciosamente os
pontos explicitamente fora de escopo.

Validação registrada: `git diff --check`, `go test ./...` e `darp doctor`
passaram em 2026-08-07. O cache do Go foi direcionado para `/tmp` por restrição
de escrita do ambiente, sem alteração no projeto.
