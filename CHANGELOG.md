# Changelog - Gost-DOM


## [0.9.3](https://github.com/gost-dom/browser/compare/v0.9.2...v0.9.3) (2025-11-26)


### Features

* `element.insertAdjacentText` ([c871936](https://github.com/gost-dom/browser/commit/c87193607e1e1e5b860e1259ccb87f2fa6a7beaa))
* `node.lastChild` ([02f4a28](https://github.com/gost-dom/browser/commit/02f4a28c49d6aff273a48c494b9e099520fb21a2))
* `Node.replaceChild` ([cd57546](https://github.com/gost-dom/browser/commit/cd5754687cbc27ecc6781b29ca0da7ba5fbc63a1))
* Allow setting `element.outerHTML` ([bd7b77b](https://github.com/gost-dom/browser/commit/bd7b77bc37afb74d7df214797c44196df363caa1))
* Create v8browser.New ([41a4856](https://github.com/gost-dom/browser/commit/41a48565703263d4a2ca2da3d8ee514c5fbe94cd))
* document.CreateElementNS ([371293a](https://github.com/gost-dom/browser/commit/371293a536e5607b0cd828cf323f9467f3cd0479))
* dummy `window.opener`. ([649c5a8](https://github.com/gost-dom/browser/commit/649c5a84014438c1a30f13047acadf662c1ea2a7))
* Implement document/element getElementsByTagName ([6f2d3e6](https://github.com/gost-dom/browser/commit/6f2d3e61a8926d0fc68701178b27fc0f49adf1fe))
* Implement outerHTML setter ([2d404fe](https://github.com/gost-dom/browser/commit/2d404fe4db9272e026ec4163f44d19d827c6e8ce))
* window.parent ([a23cd07](https://github.com/gost-dom/browser/commit/a23cd07670bf7438939fc9f826fe9431a33ce094))


### Bug Fixes

* Fix potential race condition when closing browser ([8e333d0](https://github.com/gost-dom/browser/commit/8e333d0bfc0bde348694385694df7de517e254da))
* Parse fragment error ([3ba2ce9](https://github.com/gost-dom/browser/commit/3ba2ce98e61606504207bdb825e2ecbcc4559158))

## [0.9.2](https://github.com/gost-dom/browser/compare/v0.9.1...v0.9.2) (2025-11-16)


### Features

* **dom:** implement Element.closest ([261a348](https://github.com/gost-dom/browser/commit/261a34846a3960ca89c12b19fa71fe66722e8f18))
* DOMTokenList.remove handles multiple arguments ([5e7c481](https://github.com/gost-dom/browser/commit/5e7c481d3276b253ee95a6070fe0cee4644fdeac))
* **fetch:** Add Headers support ([2fbc1ec](https://github.com/gost-dom/browser/commit/2fbc1ec7e2b97936bf562036f78ac9b5965687b1))
* **fetch:** support method and body options ([5db9672](https://github.com/gost-dom/browser/commit/5db967240681f5375bba485d7bc3e4dafcb108c7))
* implement `node.isEqualNode` ([9f1ecf4](https://github.com/gost-dom/browser/commit/9f1ecf4b68c5d5968a01f40c386fc5ee9ee636f1))
* **scripting:** Add SVGElement and MathMLElement to global scope ([c0e9c3d](https://github.com/gost-dom/browser/commit/c0e9c3dc46442541bacedfcc3c11a80208f2de8d))
* **scripting:** Log all JS callback errors ([c246025](https://github.com/gost-dom/browser/commit/c2460259d1c643b71d3c196001fd9a412e81e0f0))
* **scripting:** return explicit error on request init ([7536937](https://github.com/gost-dom/browser/commit/7536937fe7d086a55bfda063ffd1b187a01aae8f))
* **scripting:** support fetch body option ([d94882d](https://github.com/gost-dom/browser/commit/d94882d4bf352c3bf2474dbb3861472d37e12a61))
* **scripting:** support fetch method option ([5be6dc9](https://github.com/gost-dom/browser/commit/5be6dc93548d1acafb1b17e5f48e7fa124b604d0))


### Bug Fixes

* Don't panic if disconnecting flusher again ([3816f25](https://github.com/gost-dom/browser/commit/3816f251576458a75c3986c052aab99f72b272dd))
* ignore unnamed form input fields ([742f6c8](https://github.com/gost-dom/browser/commit/742f6c8c61fe2658ca46e63042b1e90d500620e2))
* use defualt timeout for `setTimeout` ([fdfd0d5](https://github.com/gost-dom/browser/commit/fdfd0d5f00338fe3c6f2737938d4b2ba45bf4d10))

## [0.9.1](https://github.com/gost-dom/browser/compare/v0.9.0...v0.9.1) (2025-10-23)


### Features

* Add "Key" to KeyboardEventInit, and support mixed case strings ([0134b49](https://github.com/gost-dom/browser/commit/0134b499ea6cd0bba990a071774e1004d2f72db2))
* Create KeyboardController to simulate keyboard input ([ad843c8](https://github.com/gost-dom/browser/commit/ad843c8269207cfb75ba32cb1c5d4c5ae9bca9b1))
* Expose correct event type, and Key to JavaScript ([08eca25](https://github.com/gost-dom/browser/commit/08eca257d0e0b6f3a824ede599e5a85d0f1041a0))
* Improve TestLogger, how it renders attribute groups ([1fb6421](https://github.com/gost-dom/browser/commit/1fb6421b5037eff8b35d87537301d7465903121a))
* Window.Navigate resolves href ([52b7521](https://github.com/gost-dom/browser/commit/52b7521984700d6ebfefb77ba8e0e6ae627e4023))


### Bug Fixes

* Browser.Open didn't set location on redirect ([e989c76](https://github.com/gost-dom/browser/commit/e989c7645146ff8ba45f10e42357319fd087779c))

## [0.9.1](https://github.com/gost-dom/browser/compare/v0.9.0...v0.9.1) (2025-07-07)


### Features

* Create KeyboardController to simulate keyboard input ([ad843c8](https://github.com/gost-dom/browser/commit/ad843c8269207cfb75ba32cb1c5d4c5ae9bca9b1))

This was in fact in 0.9 - but the the commit describing the feature was not
reachable.

## [0.9.0](https://github.com/gost-dom/browser/compare/v0.8.0...v0.9.0) (2025-07-07)


### ⚠ BREAKING CHANGES

* Remove anchor tag navigation server

This removes an internal test helper type from exported code.

* Move test server to unexported function in used package ([b8bf485](https://github.com/gost-dom/browser/commit/b8bf4854c214ff2f263fa9bbef3318760f11cbc1))


### Features

* Export a Testing.TB aware logger ([351fc7f](https://github.com/gost-dom/browser/commit/351fc7f0ccdc2282282326ce7ce1923ea1b91e31))
* FormData reads input element IDL instead of content attributes ([fe18a93](https://github.com/gost-dom/browser/commit/fe18a93fdc2580c6f99401c66fac0f8be71d6727))
* Implement name/value IDL attributes on HTMLInputElement ([3919814](https://github.com/gost-dom/browser/commit/3919814beba42251ba718a0202647193febf7925))
* Return with an error when using unfinished request option ([07634b9](https://github.com/gost-dom/browser/commit/07634b92e33ed55253bfb9c00a6b983eabf6e19d))


### Bug Fixes

* Don't write to logger when closing browser ([ec5abb0](https://github.com/gost-dom/browser/commit/ec5abb0ede4611aa755e6c8dc440f009c642deda))
* Fix TestLogger to now write after test is done ([004884b](https://github.com/gost-dom/browser/commit/004884b44c85423ec12d4b72c7000cee2ea1013e))

## [0.8.0](https://github.com/gost-dom/browser/compare/v0.7.2...v0.8.0) (2025-07-03)


### ⚠ BREAKING CHANGES

* Update minimum Go version to 1.24
* XHR.Send requres a `body` argument

### build

* Update minimum Go version to 1.24 ([75bb6f3](https://github.com/gost-dom/browser/commit/75bb6f3570155d499512c2f9e41b23dcdb3fa11b))


### Features

* Add body argument to XHR Send method ([632750c](https://github.com/gost-dom/browser/commit/632750c6fc0aafe486b34ae599ffe943f462df71))
* Add ReadableStreamDefaultReader.read() method ([8f3a918](https://github.com/gost-dom/browser/commit/8f3a9184bdc49dd47091a80e18354c0801d5f19e))
* add skeleton fetch function ([8a52970](https://github.com/gost-dom/browser/commit/8a5297075be6eecd7de220dd6119ca3f8316ce70))
* Add suport for ES modules ([4082907](https://github.com/gost-dom/browser/commit/408290748d45e9d171b5555238318b737565023b))
* Add XHR.responseType ([d9e3dd4](https://github.com/gost-dom/browser/commit/d9e3dd433c3d67023b777f6966d582b1f9849c40))
* Close browser when context is closed ([8a68449](https://github.com/gost-dom/browser/commit/8a684491b59dd8c784c1ee2a83a792fb43d76d61))
* Create EventSource ([76a0bcc](https://github.com/gost-dom/browser/commit/76a0bcce9dc76d6bea6dfb3d2a07a58ad9e91358))
* Create log error record when calling not implemented JS function ([bb49840](https://github.com/gost-dom/browser/commit/bb49840a3dcd4c7d18b88b18f0ea7b4cabd62baf))
* Create Text prototype in JS ([496098c](https://github.com/gost-dom/browser/commit/496098cd3024062cb9f3179133df72b81483aace))
* **dom:** Add AbortController.Abort() ([31e75d2](https://github.com/gost-dom/browser/commit/31e75d248217e0c81deeedf5d067e7f015246a81))
* Expose ParentNode.children to JS ([fadb9ec](https://github.com/gost-dom/browser/commit/fadb9eccc4282d28e6e244e9a5b909484d601dec))
* **goja:** Support indexed and named handlers ([2748c92](https://github.com/gost-dom/browser/commit/2748c92e91aace3c2bdbc1a5cbba7fd22f952f6b))
* Implement ChildNode.remove() ([5f1f7e5](https://github.com/gost-dom/browser/commit/5f1f7e5e395390c07d26a848b0638ebc79f05d21))
* Implement Document.importNode ([fc20d9a](https://github.com/gost-dom/browser/commit/fc20d9a9769ac529af5750ca850c3fae180d089a))
* Implement Element.InsertAdjecentElement ([aa2469e](https://github.com/gost-dom/browser/commit/aa2469e689a42946c616e7d1a90cf5c55fdf3737))
* Implement Response.body returning a ReadableStream ([edf1786](https://github.com/gost-dom/browser/commit/edf17860af8fbf93f6aa8141ac372ed51867e4bf))
* Improve error logging of JS errors ([595fe57](https://github.com/gost-dom/browser/commit/595fe573c56368016c49aec2c25585d0e54146fa))
* Override HTMLTemplateElement.SetInnerHTML ([8c2e396](https://github.com/gost-dom/browser/commit/8c2e39606177fdef25a16350c2a5836c6ba09dcc))
* **scripting:** Add Element.innerHTML attribute getter/setter ([899ae81](https://github.com/gost-dom/browser/commit/899ae81b46bcac22295c37fb0c36216c9d4d68e3))
* **scripting:** Add iterable to element attributes ([e463ed7](https://github.com/gost-dom/browser/commit/e463ed77c22b2effe468921143364c625e7ad01f))
* **scripting:** Basic window.self implementation ([52853d4](https://github.com/gost-dom/browser/commit/52853d4d1fc9a2bcde93694424af8a257ab8177b))
* **scripting:** Expose Node.parentNode ([bae0bee](https://github.com/gost-dom/browser/commit/bae0bee3a7bc62c12484d6a9f284db080a7f0fa3))
* **scripting:** Support TextDecoder through polyfill ([edaf1a4](https://github.com/gost-dom/browser/commit/edaf1a4923764bb4a26b2bf199c36e05e145876e))
* **scripting:** Support url.setSearch ([3c45e0b](https://github.com/gost-dom/browser/commit/3c45e0ba2348824dc58fcedc2fbb23bba8846043))
* Support aborting fetch request ([cc31b91](https://github.com/gost-dom/browser/commit/cc31b913411749e65a7acbc3cca7ca9cc82eb2a6))
* Support basic Datastar `GET` request ([2d29e9d](https://github.com/gost-dom/browser/commit/2d29e9d5089376b6cda9e85e3dd930654dc64c46))
* Support predictable event loop callbacks in goroutines ([ff6d160](https://github.com/gost-dom/browser/commit/ff6d1600f65e1e9ca704def89efa083373af071d))


### Bug Fixes

* Add missing return for V8 error logs ([f623045](https://github.com/gost-dom/browser/commit/f623045f15c7754c4a9cd4bc2deec776d8e69df3))
* Error when passing null as a Node in JS ([891d50a](https://github.com/gost-dom/browser/commit/891d50a1c856723c32d2441d77613efad207f495))
* **goja:** Errors in callbacks were not reported as errors ([5650df6](https://github.com/gost-dom/browser/commit/5650df6f8eac07e3a846594dfd03a2974c0b1140))

## [0.7.2](https://github.com/gost-dom/browser/compare/v0.7.1...v0.7.2) (2025-06-12)


### Features

* Add Document.CreateTextNode ([3852153](https://github.com/gost-dom/browser/commit/3852153445b82f4cb09ecbdce47cff8f728123e8))
* Add DOMStringMap.Keys() ([df25db4](https://github.com/gost-dom/browser/commit/df25db46e0864e9fd980370e832b59d5bdbea4c2))
* Add HTMLElement.dataset in JavaScript ([7348251](https://github.com/gost-dom/browser/commit/7348251000b1b3b06c2af4364526d974c83f48f6))
* **goja:** Goja script host now supports iterators ([4a60fd1](https://github.com/gost-dom/browser/commit/4a60fd18441f88e77f6ff5dea35f5da9aa6412b9))


### Bug Fixes

* **v8:** element.insertAdjacentHTML returns undefined in JS ([ab19a9d](https://github.com/gost-dom/browser/commit/ab19a9d50eeb6c5b81915724807704fe3ed20cc5))

## [0.7.1](https://github.com/gost-dom/browser/compare/v0.7.0...v0.7.1) (2025-05-27)

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
