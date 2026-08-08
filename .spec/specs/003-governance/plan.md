# Plan 003 — Governança do desenvolvimento

## Status

Approved and implemented — validated on 2026-08-07.

## Related Specification

- [`spec.md`](spec.md)

## Estratégia

Implementar a governança como documentação e instruções versionadas. O trabalho
será restrito a arquivos Markdown dentro de `.darp/`, `.agents/skills/`,
`.spec/` e ao `README.md`. Não haverá mudança em Go, configuração executável,
comandos ou comportamento de runtime.

## Dependências e ordem

1. Confirmar a especificação e os limites de escopo.
2. Escrever o contrato do lifecycle e dos quality gates.
3. Criar as quatro skills com o mesmo contrato estrutural.
4. Atualizar o README e as instruções/templates de colaboração.
5. Fazer validação estrutural, documental e de regressão.

## Áreas afetadas

- `.darp/lifecycle.md`
- `.darp/governance/quality-gates.md`
- `.agents/skills/documentation/SKILL.md`
- `.agents/skills/architecture/SKILL.md`
- `.agents/skills/testing/SKILL.md`
- `.agents/skills/release/SKILL.md`
- `.spec/templates/tasks.template.md`
- `AGENTS.md`
- `README.md`

Não devem ser alterados: `cmd/`, `internal/`, `pkg/`, testes Go, `darp.yml`,
`.darp/workflows/implement.yaml`, o comportamento do `darp init` ou do
`darp doctor`, e `.agents/skills/security-review/`.

## Marcos

### M1 — Contratos de governança

Completar lifecycle e quality gates com fases, responsáveis, evidências e
condições de passagem.

### M2 — Skills

Criar quatro `SKILL.md` consistentes, específicos e sem instruções de execução
automática ou alteração de código quando isso estiver fora do papel da skill.

### M3 — Orientações para agentes

Atualizar o template de tasks e `AGENTS.md` para exigir checkboxes, dependências
explícitas e validação por task.

### M4 — Documentação de entrada

Atualizar README, links e layout sem duplicar contratos incompatíveis.

### M5 — Validação

Verificar arquivos, front matter, referências internas, diff de escopo e testes
existentes.

## Estratégia de validação

- conferir a existência de todos os arquivos obrigatórios;
- conferir front matter e seções obrigatórias de cada skill;
- conferir que lifecycle e README referenciam apenas caminhos existentes;
- conferir que lifecycle, README e `AGENTS.md` usam a mesma sequência de fases;
- conferir que cada gate possui todos os campos exigidos;
- conferir que o `darp doctor` aceita as quatro skills, mesmo sem alteração em
  `darp.yml`;
- executar `go test ./...`;
- executar `git diff --check`;
- conferir `git diff --name-only` contra a lista de áreas permitidas;
- executar `darp doctor` somente como diagnóstico, sem alterar a árvore.

## Recuperação

As mudanças são textuais e reversíveis por revisão do diff. Nenhum arquivo
existente deve ser removido ou sobrescrito sem que seu conteúdo seja preservado
ou incorporado conscientemente.

## Premissas

- A skill `security-review` permanece independente.
- Recursos opcionais (`prompts/`, `examples/`, `references/`, `scripts/` e
  `templates/`) não são necessários para cumprir a spec.
- Checkboxes são o mecanismo oficial de acompanhamento das tasks; uma task só
  pode ser marcada como concluída quando sua validação estiver registrada.
