// Package hortifruti describes the Hortifruti storefront
// (www.hortifruti.com.br), a fresh produce grocer in Rio de Janeiro and
// São Paulo.
package hortifruti

import "github.com/voska/vtexkit/store"

// Store is the Hortifruti descriptor.
//
// Account is the field that cannot be left out, and getting it wrong fails
// quietly. store.AccountName falls back to the base URL host, which would
// yield "hortifruti" — but the real account is "hortifrutibr". Confirmed by
// logging in: VTEX returns the session as VtexIdclientAutCookie_hortifrutibr,
// and the JWT carries "account":"hortifrutibr". With the derived name the
// client would send its token under a cookie the store never reads, so every
// authenticated call would look like a logged-out one.
//
// Everything else is stock, probed live on 2026-08-20: classic password and
// emailed access code are both enabled with no OAuth provider, so no driver is
// needed; Intelligent Search REST and the catalog API both answer, so Search
// stays at SearchAuto; and subscriptions are enabled (RNS answers 401 rather
// than 404), so `hortifruti subs` works.
//
// No MinOrder, no Quirks, no Wishlist hashes — no evidence for any of them.
var Store = store.Store{
	Name:        "hortifruti",
	DisplayName: "Hortifruti",
	BaseURL:     "https://www.hortifruti.com.br",
	Account:     "hortifrutibr",
}
