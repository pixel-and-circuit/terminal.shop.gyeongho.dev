<!--
Sync Impact Report
==================
Version change: (template/placeholder) → 1.0.0
Modified principles: N/A (initial fill from template)
Added sections: None
Removed sections: None
Templates: plan-template.md ✅ updated (Constitution Check aligned); spec-template.md ✅ no change needed; tasks-template.md ✅ updated (quality gate note); README.md ✅ updated (constitution link). .specify/templates/commands/ N/A (directory does not exist)
Follow-up TODOs: None. RATIFICATION_DATE set to first adoption date 2025-02-11.
-->

# Mushroom Constitution

## Core Principles

### I. Code Quality

Code MUST be written for clarity, maintainability, and consistency with the existing codebase. All contributions MUST follow project formatting and style; automated formatting (e.g. `make format`) MUST be used. Rationale: consistent style reduces cognitive load and prevents style debates in review.

### II. Testing Standards

Testing standards MUST be followed for all behavior-affecting changes. Tests MUST be deterministic and focused; coverage expectations align with feature spec and plan. Rationale: tests protect against regressions and document intended behavior.

### III. User Experience & UI Style Consistency

User-facing behavior and UI MUST align with a consistent experience and visual style. Decisions affecting UX or UI MUST preserve coherence across the product (e.g. terminal-based interaction patterns for Mushroom). Rationale: consistency builds trust and reduces user confusion.

### IV. Model-First Design

Domain logic MUST be guided by a clear model: pure Go interfaces or structs that reflect real-world objects and relationships. Implementations SHOULD depend on these abstractions to keep the codebase maintainable and testable. Rationale: a well-defined model simplifies reasoning and future changes.

### V. Quality Gates (NON-NEGOTIABLE for Agents)

After any code modification, agents MUST run the project quality gate: at minimum `make format` and `make build`. The build MUST pass before considering the change complete. Rationale: automated gates catch format and compile errors before review or merge.

## Technical Decision Governance

Principles above guide technical decisions and implementation choices:

- **Design**: Prefer solutions that satisfy Code Quality, Testing Standards, and Model-First Design.
- **UX/UI**: Any change that affects user flows or presentation MUST satisfy User Experience & UI Style Consistency.
- **Implementation**: Every code change MUST satisfy the Quality Gates (format + build) before completion.
- **Conflicts**: When trade-offs arise, document the conflict and the rationale for the choice; principle violations require explicit justification (e.g. in plan Complexity Tracking).

## Development Workflow & Quality Gates

- **Before merge**: All constitution principles apply; quality gate (`make format`, `make build`) MUST pass.
- **Reviews**: Reviews MUST verify alignment with principles; exceptions MUST be documented and justified.
- **Agents**: Any agent making code changes MUST run the quality gate after modifications and MUST NOT leave the gate failing.

## Governance

This constitution supersedes ad-hoc practices for the Mushroom project. Amendments require: (1) documentation of the change, (2) version bump per semantic versioning below, (3) update of this file and (4) propagation to dependent templates (plan, spec, tasks, and any command or guidance docs). All PRs and reviews MUST verify compliance with these principles; complexity or principle deviations MUST be justified (e.g. in plan.md Complexity Tracking or in PR description). For runtime development guidance, use README.md and specs under `specs/`.

**Version**: 1.0.0 | **Ratified**: 2025-02-11 | **Last Amended**: 2025-02-11
