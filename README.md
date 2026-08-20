# hortifruti

CLI for [Hortifruti](https://www.hortifruti.com.br), a fresh produce grocer in Rio de
Janeiro and São Paulo. Built for humans and AI agents: data to stdout, hints to stderr,
stable exit codes, structured output.

Built on [vtexkit](https://github.com/voska/vtexkit), a shared library for Brazilian VTEX
storefronts.

## Install

```sh
brew install voska/tap/hortifruti
```

or `make build` for `bin/hortifruti`.

## Use

```sh
hortifruti doctor
hortifruti auth login --email you@example.com
hortifruti search banana --limit 5
```

```
$ hortifruti search banana --limit 3
100065     Banana Prata Unidade                                 R$2,00  un
100063     Banana Nanica Unidade                                R$1,24  kg
100066     Banana da Terra Unidade                              R$3,85  un
```

Note the last column: produce is sold by both unit (`un`) and weight (`kg`), and the two
price very differently. Check it before comparing anything.

### Ordering

```sh
hortifruti cart add 100063 --qty 2
hortifruti delivery windows
hortifruti checkout --window 0            # preview — places nothing
hortifruti checkout --window 0 --confirm  # places the order
```

Subscriptions are enabled at this store, so `hortifruti subs` works.

## Price-checking against Zona Sul and Prezunic

All three CLIs emit the same JSON shape with integer-centavo prices:

```sh
for q in "banana" "tomate" "alface"; do
  h=$(hortifruti search "$q" --json --results-only | jq -r '.[0] | "\(.price) \(.unit)"')
  z=$(zonasul    search "$q" --json --results-only | jq -r '.[0] | "\(.price) \(.unit)"')
  printf "%-10s hortifruti=%-12s zonasul=%s\n" "$q" "$h" "$z"
done
```

Compare the unit as well as the price — a per-kg price and a per-unit price are not the
same number, and the top search hit at each store is often a different product.

## What lives here

Only the store descriptor. Everything else is in vtexkit.

The one field that matters: **the VTEX account is `hortifrutibr`, not `hortifruti`.** It
cannot be derived from the domain, and with the derived name the auth cookie would be
wrong and every logged-in call would silently behave as logged out.

## License

MIT
