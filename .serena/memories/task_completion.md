# Task completion checklist (shop.gyeongho.dev)

When a code task is completed:

1. Run `make format` (constitution: code quality).
2. Run `make build` and ensure it passes (constitution: quality gate).
3. Run `make test` and fix any failures.
4. Do not leave the quality gate failing.

Agents MUST run the project quality gate after any code modification (see .specify/memory/constitution.md).
