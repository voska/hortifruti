# Hortifruti CLI

Go CLI for ordering fresh produce from Hortifruti (`www.hortifruti.com.br`).

All logic lives in **`github.com/voska/vtexkit`**. This repo holds only the descriptor.

## The one thing that is not derivable

The VTEX account is **`hortifrutibr`**, not `hortifruti`. Confirmed by logging in: the
session cookie is `VtexIdclientAutCookie_hortifrutibr` and the JWT carries
`"account":"hortifrutibr"`. `store.AccountName` would derive `hortifruti` from the host,
which fails quietly — the token would be sent under a cookie the store never reads.

## Produce is sold by unit and by weight

Search results carry a `unit` field (`un` vs `kg`). Never compare or sum prices across
stores without it.

## Build & test

`make build` `make test` `make lint` `make vet` `make ci`

## Commits

Conventional Commits (`feat:`, `fix:`, `chore:`, `docs:`, `test:`, `refactor:`).
