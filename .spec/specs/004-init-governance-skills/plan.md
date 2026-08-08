# Plan 004 — Scaffold de skills de governança no `darp init`

## Status

Draft — depende da aprovação da Spec 004.

## Related Specifications

- [`spec.md`](spec.md)
- [`../003-governance/spec.md`](../003-governance/spec.md)

## Estratégia

Estender o serviço de inicialização para tratar os contratos de governança como
um conjunto determinístico de ativos padrão. O serviço deve continuar usando a
abstração de filesystem existente, criando somente arquivos ausentes e
aplicando uma atualização aditiva e segura ao `darp.yml` quando necessário.

A seleção de conteúdo não deve depender de inspeção do projeto-alvo. A fonte
canônica pode ser `go:embed` ou constantes versionadas, desde que seja local,
determinística e testável.

## Dependências e ordem

1. Aprovar Spec 003 e Spec 004.
2. Consolidar os conteúdos canônicos de lifecycle, quality gates e quatro
   skills.
3. Definir e implementar a atualização segura do `darp.yml`.
4. Estender o serviço de bootstrap e preservar idempotência.
5. Adicionar testes unitários, de reparo, preservação e fixtures agnósticas.
6. Atualizar documentação e validar com `darp doctor`.

## Áreas afetadas

- `internal/project/init/`
- testes de `internal/project/init/`
- fonte dos ativos embutidos usados pelo `init`
- `README.md`
- documentação da Spec 001, se o contrato de preservação for alterado

Não devem ser alterados pelo escopo normal:

- `internal/project/doctor/`, salvo ajustes de teste estritamente necessários;
- `.darp/` e `.agents/skills/` do repositório sem atualizar também a fonte
  canônica correspondente;
- comportamento de workflows e execução de skills;
- skills específicas existentes, como `security-review`.

## Marcos

### M1 — Contratos canônicos

Definir e versionar os conteúdos completos dos quatro `SKILL.md`, lifecycle e
quality gates, verificando neutralidade de stack.

### M2 — Configuração aditiva

Implementar a inclusão das quatro entradas em `darp.yml` para projetos novos e
o reparo seguro de configurações existentes.

### M3 — Scaffold e preservação

Criar diretórios e arquivos ausentes, sem sobrescrever ativos customizados.

### M4 — Testes

Cobrir projeto novo, projeto parcial, reexecução, ativos customizados, YAML
inválido, campos desconhecidos, skills adicionais e falhas de filesystem.

### M5 — Compatibilidade e documentação

Validar `darp doctor`, atualizar documentação e confirmar que nenhuma etapa
executa comandos ou detecta o tipo do projeto.

## Estratégia de validação

- `go test ./...`;
- testes do serviço com filesystem temporário;
- comparação do conteúdo mínimo dos quatro `SKILL.md`;
- fixtures sem código, Java/Spring, Python/FastAPI, Go e monorepo;
- execução de `darp doctor` em cada fixture scaffoldada;
- segunda execução do `init` sem diferenças nos arquivos existentes;
- inspeção do diff para garantir que `security-review` não foi alterada;
- `git diff --check`.

## Recuperação

O serviço deve falhar sem apagar arquivos. Em caso de erro na atualização do
`darp.yml`, manter o arquivo original e reportar a causa. Em caso de erro após
criar alguns ativos, os ativos já criados não devem ser removidos
automaticamente; a próxima execução deve conseguir reparar os ausentes.

## Premissas

- As quatro skills da Spec 003 são os templates padrão oficiais.
- O `darp doctor` continua validando skills encontradas no diretório, além das
  entradas registradas na configuração.
- O workflow `implement.yaml` permanece sem execução automática.
- A neutralidade é obtida por descoberta contextual dentro das skills, não por
  lógica de detecção de stack no `darp init`.
