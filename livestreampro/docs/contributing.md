<!-- /home/${USER}/livestreampro/docs/contributing.md -->
# Contributing Guide

## Branching & Conventional Commits
- Branch from `main` using descriptive names such as `feat/stream-chat` or `fix/login-bug`.
- Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/) (e.g. `feat: add chat widget`).
- Squash merge requests to keep history linear.

## Formatting & Linting
- Run `make lint` to execute `golangci-lint` for Go and `npm run lint` for the frontend.
- Code must be formatted with `gofmt` and Prettier; CI will fail otherwise.
- Install pre-commit with `pip install pre-commit` then run `pre-commit install`.

## Pull Requests & CI
- Push your branch and open a PR against `main`.
- At least one maintainer review is required before merge.
- GitHub Actions run `make lint`, `make test-backend`, and `make test-frontend` for every PR. All checks must pass.
