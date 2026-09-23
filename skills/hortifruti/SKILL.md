---
name: hortifruti
description: >-
  Order fresh produce from Hortifruti (www.hortifruti.com.br) using the
  `hortifruti` CLI. Use for fruit, vegetables, herbs and greens — the produce
  specialist, versus Zona Sul for general groceries.
allowed-tools: Bash, Read
---

# hortifruti

Order fresh produce from Hortifruti using the `hortifruti` CLI.

## Always start here

```bash
hortifruti doctor
```

Exit 0 means ordering will work. Any other exit code means it will not. Each
failed line prints the exact fix. Do that fix, or report it to the user. Do not
retry the command that failed.

## The flow

```bash
hortifruti search banana --limit 5     # 1. find a SKU — first column
hortifruti cart add 100063 --qty 2     # 2. add it
hortifruti delivery windows            # 3. pick a window number
hortifruti checkout --window 0         # 4. preview — places nothing
hortifruti checkout --window 0 --confirm   # 5. order
```

Between steps 4 and 5: **show the preview to the user and get explicit
approval.** `--confirm` spends real money.

Every result carries both `sku` and `productId`. Commands take the `sku` —
the two are separate sequences and the same number routinely appears in
both, naming two unrelated products.

```bash
hortifruti cart show
hortifruti cart update 0 --qty 3    # index from 'cart show'
hortifruti cart remove 0
hortifruti cart clear
```

## Unit vs weight — read this before quoting any price

Produce is sold both by the unit and by the kilo, and the CLI reports which in
the last column:

```
100065     Banana Prata Unidade      R$2,00  un
100063     Banana Nanica Unidade     R$1,24  kg
```

Those two numbers are not comparable. Always carry the unit when you report a
price, and never sum or compare across `un` and `kg`. `--qty 2` on a `kg` item
means two kilos, not two bananas — say so when previewing a cart.

## When to use this store

Hortifruti is the produce specialist. Zona Sul remains the default for general
groceries; route fruit, vegetables, herbs and greens here when the user wants
quality produce, and say which store you are using and why.

Produce is freshest on **Mondays**.

Do not split an order across stores without saying so — two deliveries means
two delivery windows and someone has to be home for both.

## Price comparison

`hortifruti`, `zonasul` and `prezunic` emit the same JSON shape with integer
centavo prices:

```bash
h=$(hortifruti search "tomate" --json --results-only | jq -r '.[0] | "\(.price) \(.unit)"')
z=$(zonasul    search "tomate" --json --results-only | jq -r '.[0] | "\(.price) \(.unit)"')
```

State the caveat when you report a comparison: the top hit at each store is
often a different variety, size, or unit. It is a signal, not a quote.

## Subscriptions

```bash
hortifruti subs                  # status, frequency, next delivery
hortifruti subs <id>
hortifruti subs pause <id>
hortifruti subs resume <id>
hortifruti subs skip <id>        # skip the next delivery only
hortifruti subs unskip <id>
```

Prefer `skip` to `pause`. There is no `cancel` — VTEX has no transition out of
`CANCELED`, so tell the user to cancel on the website.

## Output for scripts and agents

```bash
hortifruti search banana --json
hortifruti search banana --json --select sku,name,price,unit
hortifruti search banana --plain          # tab-separated
```

Prices in `--json` are integer centavos: `124` is R$1,24.

Data goes to stdout; progress and errors go to stderr.

## Exit codes

| Code | Meaning | What to do |
|---|---|---|
| 0 | success | continue |
| 2 | bad arguments | fix the command; do not retry it unchanged |
| 3 | empty result | tell the user nothing matched |
| 4 | login required | `hortifruti auth login --email <email>` |
| 5 | not found | the SKU is wrong; search again |
| 7, 8 | temporary | wait, then retry once |
| 9 | store rule refused | read the message; it names the rule |
| 10 | not set up | `hortifruti doctor` and follow the fixes |

Full table: `hortifruti exit-codes --json`

## Rules

- Never run `--confirm` without the user approving that exact cart and total.
- Always state the unit (`un`/`kg`) alongside any price or quantity.
- If a command fails twice the same way, stop and report it. Do not loop.
- `hortifruti doctor` diagnoses anything unexpected; its output names the fix.
