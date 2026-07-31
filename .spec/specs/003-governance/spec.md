Você é o arquiteto responsável pelo projeto darp-cli.

O projeto utiliza Spec Driven Development (SDD).

Leia atentamente:

- README.md
- .darp/lifecycle.md
- .spec/specs/

Antes de qualquer alteração, compreenda a arquitetura atual.

# Objetivo

Criar a infraestrutura de Governança do DARP.

Esta infraestrutura será utilizada por todas as futuras implementações.

Ela NÃO deve modificar funcionalidades existentes do CLI.

Ela NÃO deve implementar novos comandos.

Ela NÃO deve alterar o comportamento do `darp init`.

Ela apenas cria a organização necessária para padronizar o processo de engenharia do projeto.

--------------------------------------------------

# Motivação

Hoje o lifecycle executa apenas as etapas de implementação.

Queremos evoluí-lo para um processo completo de engenharia.

Toda implementação futura deverá passar por etapas obrigatórias de qualidade.

Exemplos:

Spec

↓

Implementação

↓

Testes

↓

Architecture Review

↓

Documentation Review

↓

Quality Gates

↓

Concluído

--------------------------------------------------

# O que criar

Criar a seguinte estrutura:

.darp/

    governance/

        lifecycle.md
        quality-gates.md

.agents/

    skills/

        documentation/
        architecture/
        testing/
        release/

Cada skill deve possuir sua própria estrutura.

Exemplo:

documentation/

    SKILL.md

    prompts/ (opcional)

    examples/ (opcional)

    templates/ (opcional)

A estrutura deve ser consistente entre todas as skills.

--------------------------------------------------

# Documentation Skill

Criar uma skill responsável por manter toda a documentação sincronizada com o código.

Ela deve possuir responsabilidades como:

- revisar README.md

- verificar inconsistências

- documentar novas features

- atualizar comandos

- atualizar exemplos

- atualizar estrutura do projeto

- validar links internos

- nunca documentar funcionalidades inexistentes

- nunca remover documentação válida

- manter README como ponto de entrada do projeto

Ela não deve modificar código.

--------------------------------------------------

# Architecture Skill

Responsável por revisar arquitetura.

Exemplos:

- responsabilidades bem definidas

- baixo acoplamento

- alta coesão

- duplicações

- violação de arquitetura

- oportunidades de refatoração

Ela apenas gera recomendações.

--------------------------------------------------

# Testing Skill

Responsável por validar:

- cobertura

- testes ausentes

- cenários não contemplados

- regressões

- consistência entre testes e implementação

--------------------------------------------------

# Release Skill

Responsável por:

- revisar CHANGELOG

- preparar release notes

- identificar breaking changes

- listar novas features

--------------------------------------------------

# Quality Gates

Criar um documento quality-gates.md definindo gates como:

✓ Build

✓ Tests

✓ Documentation

✓ Architecture

✓ Compatibility

✓ Release Notes

Cada gate deve informar:

- objetivo

- critérios

- responsável (skill)

- resultado esperado

--------------------------------------------------

# Atualizar o Lifecycle

Evoluir o lifecycle existente para incluir explicitamente:

Spec

↓

Implementação

↓

Testes

↓

Architecture Review

↓

Documentation Review

↓

Quality Gates

↓

Concluído

O lifecycle deve fazer referência às respectivas skills.

--------------------------------------------------

# Atualizar README

Atualizar README.md para documentar:

- conceito de Governança

- Skills

- Quality Gates

- Lifecycle

- nova estrutura da pasta .darp

Sem remover conteúdo existente.

Apenas complementar.

--------------------------------------------------

# Critérios

- Não implementar código do CLI.

- Não alterar comandos existentes.

- Não modificar specs já concluídas.

- Não criar funcionalidades fora do escopo.

--------------------------------------------------

# Entrega

Ao concluir:

1. Liste todos os arquivos criados.

2. Liste os arquivos modificados.

3. Explique como a Governança se integra ao Lifecycle.

4. Explique como uma futura spec poderá reutilizar automaticamente essa infraestrutura.

5. Explique quais partes poderão futuramente ser executadas automaticamente pelo comando `darp run`.

A solução deve ser extensível, modular e alinhada à visão de longo prazo do DARP como um framework de engenharia orientado por especificações e agentes de IA.
