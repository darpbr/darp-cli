# Tasks 004 — Scaffold de skills de governança no `darp init`

## Status

Draft — não iniciar antes da aprovação da Spec 004 e do Plan 004.

## Related Plan

- [`plan.md`](plan.md)

## Regras para o implementador

- Marque cada checkbox somente após concluir a ação e sua validação.
- Preserve arquivos customizados e registre qualquer decisão de compatibilidade.
- Não introduza detecção de linguagem, framework ou tipologia de projeto.
- Tasks bloqueadas permanecem desmarcadas e devem registrar evidência e motivo.

## Task List

### T1 — Consolidar ativos canônicos

- [ ] Definir a fonte canônica embutida para lifecycle, quality gates e as
      quatro skills.
- [ ] Garantir que os conteúdos são determinísticos e funcionam sem rede.
- [ ] Garantir que cada skill contém as seções e limites da Spec 003.
- [ ] Remover instruções específicas de Go, Java, Spring, Python, FastAPI,
      providers ou ferramentas obrigatórias.
- [ ] Garantir que o workflow não passe a executar skills automaticamente.

### T2 — Expandir a estrutura criada pelo init

- [ ] Incluir os diretórios `architecture`, `testing` e `release` na estrutura
      criada pelo serviço de inicialização.
- [ ] Criar `SKILL.md` ausente para as quatro skills.
- [ ] Criar lifecycle e quality gates completos somente quando os arquivos
      estiverem ausentes.
- [ ] Manter a criação de workflow e templates compatível com as specs
      anteriores.

### T3 — Atualizar `darp.yml` com segurança

- [ ] Fazer projetos novos registrarem as quatro skills nos caminhos canônicos.
- [ ] Definir a estratégia de atualização aditiva para `darp.yml` existente.
- [ ] Preservar valores existentes, skills adicionais e campos desconhecidos.
- [ ] Não sobrescrever `darp.yml` quando o YAML for inválido ou a atualização
      segura não puder ser concluída.
- [ ] Cobrir entradas duplicadas, valores customizados e configuração parcial.

### T4 — Preservar ativos e idempotência

- [ ] Garantir que arquivos existentes nunca sejam substituídos integralmente;
      tratar a atualização aditiva de `darp.yml` como exceção controlada.
- [ ] Garantir que `SKILL.md` customizado seja preservado integralmente.
- [ ] Garantir que lifecycle e quality gates customizados sejam preservados.
- [ ] Garantir que a segunda execução não remova arquivos do usuário.
- [ ] Garantir que uma execução parcial possa ser reparada na execução seguinte.

### T5 — Criar testes agnósticos de projeto

- [ ] Testar projeto novo sem código.
- [ ] Testar fixture Java com Spring.
- [ ] Testar fixture Python com FastAPI.
- [ ] Testar fixture Go.
- [ ] Testar monorepo com múltiplas linguagens.
- [ ] Testar projeto parcialmente inicializado.
- [ ] Testar YAML inválido, campos desconhecidos e skills adicionais.
- [ ] Testar falhas de filesystem sem remoção de arquivos existentes.
- [ ] Testar idempotência e conteúdo mínimo de todos os contratos.

### T6 — Validar integração e documentação

- [ ] Executar `go test ./...`.
- [ ] Executar `darp doctor` em cada fixture válida.
- [ ] Confirmar que nenhum teste depende de detecção de stack para obter
      sucesso.
- [ ] Confirmar que nenhum comando do projeto-alvo é executado pelo `init`.
- [ ] Atualizar README e documentação da inicialização com a nova estrutura.
- [ ] Confirmar que `.agents/skills/security-review/` não foi alterada.
- [ ] Executar `git diff --check`.
- [ ] Revisar o diff final contra Spec 004, Plan 004 e todas as tasks.

## Validation Checklist

- [ ] Spec 004 aprovada.
- [ ] Plan 004 aprovado.
- [ ] Todas as tasks T1–T6 concluídas.
- [ ] Projetos novos recebem as quatro skills.
- [ ] Projetos existentes são reparados sem sobrescrita destrutiva.
- [ ] Fixtures de múltiplas stacks passam no `darp doctor`.
- [ ] Testes existentes passam.
- [ ] Nenhum bloqueio pendente.

## Notas e bloqueios

Registrar decisões sobre preservação de YAML, conteúdo canônico e compatibilidade
com projetos já inicializados. Uma necessidade de executar skills, workflows ou
detectar stack deve gerar uma nova especificação, não uma alteração silenciosa
nesta implementação.
