# GitHub Copilot Instructions

This repository follows Specification-Driven Development.

Before implementing code:

- read AGENTS.md
- read README.md
- understand the project vision
- verify whether a specification already exists

Never implement functionality directly from a prompt if no specification exists.

Prefer:

- small pull requests
- deterministic implementations
- idiomatic Go
- reusable components
- low coupling

Avoid:

- speculative abstractions
- premature optimization
- unnecessary dependencies

When generating code:

- follow Effective Go
- use the standard library whenever possible
- keep functions focused
- write meaningful comments only when needed

Always preserve architectural consistency.