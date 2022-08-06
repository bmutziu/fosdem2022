package kube

import "acme.com/x/infra/quote"

qr: quote.#QuoteRequest & {
    lang: quote.#EN 
    num:  3
}
