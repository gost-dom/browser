# Changelog - Gost-DOM


## [0.7.0](https://github.com/gost-dom/browser/compare/v0.6.0...v0.7.0) (2025-05-27)


### ⚠ BREAKING CHANGES

* The following Go functions have changed return values from *string to
(string, bool)

- DOMTokenList.Item()
- XMLHttpRequest.GetResponseHeader()

DOMTokenList is returned by Element.ClassList() and
HTMLAnchorElement.RelList()

* *string return vaules are returned as (string, bool) ([9f2662f](https://github.com/gost-dom/browser/commit/9f2662f50232f809c91d5197ed7590e8a73ad1f0))


### Features

* Add missing FormData.getAll in JS ([48c42f4](https://github.com/gost-dom/browser/commit/48c42f485bdddc41c3a5d3d1199298b385ea91e2))
* Use more idiomatic errors for DOMTokenList ([e4f7ad3](https://github.com/gost-dom/browser/commit/e4f7ad30048ef43a03469add7768c0065f4fc2b2))
* **v8:** Add URL.toString() ([bd09002](https://github.com/gost-dom/browser/commit/bd0900207652d241725f13827a642973386d006b))


### Bug Fixes

* Fix potential race conditions with parallel tests ([a72cee7](https://github.com/gost-dom/browser/commit/a72cee7b926a76ce5855cde8199b2686a8c4f921))
* Set correct prototype for document.documentElement ([17b695a](https://github.com/gost-dom/browser/commit/17b695abfe3a849ce0c005c8aa6bdfbc3558500c))

## [0.6.0](https://github.com/gost-dom/browser/compare/v0.5.8...v0.6.0) (2025-05-20)


### ⚠ BREAKING CHANGES

* This renamed `CustomEventInit.Details` to `CustomEventInit.Detail` as
this is the correct spelling according to the MDN documentation:
https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent/detail

As this value was not exposed to JavaScript previously, it is deemed
unlikely that this will affect client code.

### Features

* Add element.id IDL attribute ([7140aeb](https://github.com/gost-dom/browser/commit/7140aeb88a429dae9935fa6474c1653e1901eed0))
* Add element.next/previousElementSibling ([5e18652](https://github.com/gost-dom/browser/commit/5e18652f1bf30fc39b8f0fe90f1060ab7bae9ab8))
* Log errors on unhandled promise rejection ([45645f0](https://github.com/gost-dom/browser/commit/45645f0b41cba4f6d4bbfc083c69a305b2ac0e7a))
* Support CustomEvent.detail in JavaScript ([0598e74](https://github.com/gost-dom/browser/commit/0598e74b30b1f966192eb88bf94bda71cb4751b3))
* Support deferred scripts. ([a02c244](https://github.com/gost-dom/browser/commit/a02c24450ddc692e9cb7f9b86dbec63748c68abe))
* Support window.queueMicrotask ([7b307c4](https://github.com/gost-dom/browser/commit/7b307c46d5e8292ec46cd7c91d45f7beaee1e7a8))
* **v8:** Add URLSearchParams ([48031f2](https://github.com/gost-dom/browser/commit/48031f21262fadbcc5032758034654088e5fe7a5))
* **v8:** Evaluating a "wrapper" object returns the Go object ([769619c](https://github.com/gost-dom/browser/commit/769619c448ebe61aa0dc863401e2875250ac7b15))
* **v8:** Implement URLSearchParams support in V8 host ([92ec474](https://github.com/gost-dom/browser/commit/92ec474bf4fb9e9f54fe0601017aa53ff369db9b))


### Bug Fixes

* Clear node when calling Node.ReplaceChildren with no elements ([e5dd949](https://github.com/gost-dom/browser/commit/e5dd9494de9eb92b1b341b9599635e73b78edb72))

## [0.5.8](https://github.com/gost-dom/browser/compare/v0.5.7...v0.5.8) (2025-05-11)


### Features

* **html:** Add WindowOptionHTTPClient ([b50b6f8](https://github.com/gost-dom/browser/commit/b50b6f8340e0a0da8752e17e4d9658d230c4ed1a))


### Bug Fixes

* **html:** Resolve script src in context of document location ([29c8ea2](https://github.com/gost-dom/browser/commit/29c8ea2de71a6f57ebae1c3f85b6427cbcd1a8d2))

## [0.5.7](https://github.com/gost-dom/browser/compare/v0.5.6...v0.5.7) (2025-04-26)


### Bug Fixes

* **v8:** Node.insertBefore failed on missing ref ([d7dca66](https://github.com/gost-dom/browser/commit/d7dca6685ec7389f161ee3964dda3cad86a1b856))

### ⚠ DEPRECATED

* `XMLHttpRequest.send` will change signature, accepting an `io.Reader`, replacing `SendBody(io.Reader)`. For now, use `SendBody(nil)`, but eventually `Send(io.Reader)` will live, and `SendBody` be removed.
  

## [0.5.6](https://github.com/gost-dom/browser/compare/v0.5.5...v0.5.6) (2025-04-21)


### Features

* **html:** Click label forwards to "input" ([921c8f1](https://github.com/gost-dom/browser/commit/921c8f1ccd8bdc12b7beb2f665fd38c4f325e5c2))


### Bug Fixes

* Fix weird issue with parallel tests ([178f64e](https://github.com/gost-dom/browser/commit/178f64e5bf53a7993a6687ea4992d972b299a1e7))

## [0.5.5](https://github.com/gost-dom/browser/compare/v0.5.4...v0.5.5) (2025-04-11)


### Features

* **dom:** Handle checkbox input types ([992fee9](https://github.com/gost-dom/browser/commit/992fee9d8bebb2872f4ed612ab226e27b2864344))

## [0.5.4](https://github.com/gost-dom/browser/compare/v0.5.3...v0.5.4) (2025-04-11)


### Bug Fixes

* "Fix" typo in event phase const name ([cd7b4bc](https://github.com/gost-dom/browser/commit/cd7b4bc8986b99f5721bf1f5058208a31285c942))
* InsertAdjacentHTML only handled elements ([a5276cf](https://github.com/gost-dom/browser/commit/a5276cf19f6cfa64dc7e847b3230655c0d19f4ab))

## [0.5.3](https://github.com/gost-dom/browser/compare/v0.5.2...v0.5.3) (2025-03-22)

### ⚠ DEPRECATED

Constant `event.EventPhaseBubbline` deprecated because of a typo. Use
`EventPhaseBubbling` instead.

### Features

* Handle new location after redirects ([92fcabb](https://github.com/gost-dom/browser/commit/92fcabb9b132d1ddf09c594bb006d3fd8ef198bf))

## [0.5.2](https://github.com/gost-dom/browser/compare/v0.5.1...v0.5.2) (2025-03-22)


### Features

* Handle form post 307/308 status codes ([70fb78a](https://github.com/gost-dom/browser/commit/70fb78a96f58b78d9d3d0b6ad46202bb8c963157))
* Improve logging of error cases ([31b6453](https://github.com/gost-dom/browser/commit/31b645333a6ada705cc1abdb8a485ab97b15f9ce))

## [0.5.1](https://github.com/gost-dom/browser/compare/v0.5.0...v0.5.1) (2025-03-20)


### Features

* Browser-scoped logger ([a401acc](https://github.com/gost-dom/browser/commit/a401accf677b1da49a16b1a9e6782dcedb1518d0))

## [0.5.0](https://github.com/gost-dom/browser/compare/v0.4.1...v0.5.0) (2025-03-16)


### ⚠ BREAKING CHANGES

* This changes some exported names, but most clients
should be unaffected, as these are names that client code shouldn't have
been using.
* This is versioned as breaking, but the changes affect
exported names that shouldn't be used by client code, so in practice it
shouldn't be a conflict.

* Remove the requirement for a "New" entity ([71292c6](https://github.com/gost-dom/browser/commit/71292c6858c135de422c9dfc221b548f6ceb7950))


### Features

* Add HTMLElement `nonce` IDL attribute ([9a573ab](https://github.com/gost-dom/browser/commit/9a573ab9b2dc60a3810556343186a5e377acf7d2)), closes [#59](https://github.com/gost-dom/browser/issues/59) [#60](https://github.com/gost-dom/browser/issues/60)
* Add missing RemoveAttribute(string) ([46ccd78](https://github.com/gost-dom/browser/commit/46ccd78fb82b6a572ff49da8dd55fc2913e09658))
* **dom:** Implement HTMLElement.Dataset() ([9b5e7d0](https://github.com/gost-dom/browser/commit/9b5e7d05bcf2033f1ede325ddd983db6c9c1c54c))
* Focus element with autofocus attribute ([ddfbd63](https://github.com/gost-dom/browser/commit/ddfbd631e5a6a8068e62fe536d93ce9234bfa4dc))
* **html:** Add autofocus IDL attribute ([8216632](https://github.com/gost-dom/browser/commit/821663248a2a819f4557fd6188317bdc0bf16677))
* **html:** Add Focus() and Blur() on HTMLElement ([5377383](https://github.com/gost-dom/browser/commit/5377383167c889332f23dea7419b42f2f45d0898)), closes [#61](https://github.com/gost-dom/browser/issues/61)
* **html:** Add HTMLElement.Tabindex()/SetTabindex() ([110b67b](https://github.com/gost-dom/browser/commit/110b67beb730941da8098f66067f8f9de836b3d6))
* Make Event type usable as an object ([63f4e1f](https://github.com/gost-dom/browser/commit/63f4e1f389177bf31aa9d399e5bd2828f5341dbd))
* **v8:** Support most new HTMLElement IDL attributes ([58906f8](https://github.com/gost-dom/browser/commit/58906f82ea2e05c4cc69e5fe36d2169a63e90676))

## [0.4.1](https://github.com/gost-dom/browser/compare/v0.4.0...v0.4.1) (2025-03-13)

### Deprecated

* `ElementEvents` and it's methods as deprecated

## [0.4.0](https://github.com/gost-dom/browser/compare/v0.3.1...v0.4.0) (2025-03-06)

This is breaking change because of a major refactor, moving event types ([c777cc4](https://github.com/gost-dom/browser/commit/c777cc429ec40a6760eecb229bb4b241e0cee8e8)). Also see [this discussion thread](https://github.com/orgs/gost-dom/discussions/50)

### Features

* Add more event types ([311d50a](https://github.com/gost-dom/browser/commit/311d50adbe3b8495f6a832a85b483140f6b6d41a))
* Create prototype chain for `click` events ([bed3be2](https://github.com/gost-dom/browser/commit/bed3be2f02d378a1a966480010a06630fd6bb9da))
* Implement event capture phase ([29942d3](https://github.com/gost-dom/browser/commit/29942d3a59e14d4a5f1f2a19f8aac8418b39a1d7))

## [0.3.1](https://github.com/gost-dom/browser/compare/v0.3.0...v0.3.1) (2025-03-05)


### Bug Fixes

* Missing header on form posts ([9a89ca9](https://github.com/gost-dom/browser/commit/9a89ca96ae6c09867d20eb76acce8f838d13f229))

## [0.3.0](https://github.com/gost-dom/browser/compare/v0.2.0...v0.3.0) (2025-02-27)

This is a breaking change because of when callbacks will be executed. See the
commit messages for the event loop for details

### Features

* "Event loop" with delay and interval ([a132662](https://github.com/gost-dom/browser/commit/a132662bdf05067beb7bb41a795ab158dd8392f3))
* Change v8go dependency ([df1ef32](https://github.com/gost-dom/browser/commit/df1ef32fdb5d0499bb38f249956190ce7cbb533d))


### Bug Fixes

* ownerDocument when moving to new document ([590a5f0](https://github.com/gost-dom/browser/commit/590a5f0feadd49ad846d5b600024b962a11984e4))

## [0.2.0](https://github.com/gost-dom/browser/compare/v0.1.2...v0.2.0) (2025-02-19)

### Features

    * New url package (4ff3300 (https://github.com/gost-dom/browser/commit/4ff3300cbf5291e1f18dd9faea422291f49a13e6))
    * Support HTMLAnchorElement.String() (3e2efc7 (https://github.com/gost-dom/browser/commit/3e2efc7384c204db8aadc7805a9e6777b1b5eaec))

## [0.1.2](https://github.com/gost-dom/browser/compare/v0.1.1...v0.1.2) (2025-02-06)

### Bug Fixes

    * Fix ParseURLBase issues (a4ed7a5 (https://github.com/gost-dom/browser/commit/a4ed7a5c28c9a947a1a2dc50827c3b21dc9b46b2))

## [0.1.1](https://github.com/gost-dom/browser/compare/v0.1.0...v0.1.1) (2025-02-06)

### Bug Fixes

    * Passing an int to V8.NewValue (aec0653 (https://github.com/gost-dom/browser/commit/aec0653b949d40692cec9402cdf6bf995a6ba365))
