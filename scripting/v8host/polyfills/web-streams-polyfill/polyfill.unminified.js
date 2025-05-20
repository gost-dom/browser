(function () {
  "use strict";
  function e() {}
  function t(e) {
    return (typeof e == "object" && e !== null) || typeof e == "function";
  }
  function o(e, t) {
    try {
      Object.defineProperty(e, "name", { value: t, configurable: true });
    } catch (e) {}
  }
  function u(e) {
    return new n(e);
  }
  function c(e) {
    return u((t) => t(e));
  }
  function d(e) {
    return l(e);
  }
  function f(e, t, r) {
    return i.call(e, t, r);
  }
  function b(e, t, o) {
    f(f(e, t, o), void 0, r);
  }
  function h(e, t) {
    b(e, t);
  }
  function m(e, t) {
    b(e, void 0, t);
  }
  function _(e, t, r) {
    return f(e, t, r);
  }
  function p(e) {
    f(e, void 0, r);
  }
  function S(e, t, r) {
    if (typeof e != "function") {
      throw new TypeError("Argument is not a function");
    }
    return Function.prototype.apply.call(e, t, r);
  }
  function g(e, t, r) {
    try {
      return c(S(e, t, r));
    } catch (e) {
      return d(e);
    }
  }
  function q(e, t) {
    e._ownerReadableStream = t;
    t._reader = e;
    if (t._state === "readable") {
      O(e);
    } else if (t._state === "closed") {
      (function (e) {
        O(e);
        A(e);
      })(e);
    } else {
      j(e, t._storedError);
    }
  }
  function E(e, t) {
    return Br(e._ownerReadableStream, t);
  }
  function W(e) {
    const t = e._ownerReadableStream;
    if (t._state === "readable") {
      k(
        e,
        new TypeError(
          "Reader was released and can no longer be used to monitor the stream's closedness"
        )
      );
    } else {
      (function (e, t) {
        j(e, t);
      })(
        e,
        new TypeError(
          "Reader was released and can no longer be used to monitor the stream's closedness"
        )
      );
    }
    t._readableStreamController[P]();
    t._reader = void 0;
    e._ownerReadableStream = void 0;
  }
  function B(e) {
    return new TypeError("Cannot " + e + " a stream using a released reader");
  }
  function O(e) {
    e._closedPromise = u((t, r) => {
      e._closedPromise_resolve = t;
      e._closedPromise_reject = r;
    });
  }
  function j(e, t) {
    O(e);
    k(e, t);
  }
  function k(e, t) {
    if (e._closedPromise_reject !== void 0) {
      p(e._closedPromise);
      e._closedPromise_reject(t);
      e._closedPromise_resolve = void 0;
      e._closedPromise_reject = void 0;
    }
  }
  function A(e) {
    if (e._closedPromise_resolve !== void 0) {
      e._closedPromise_resolve(void 0);
      e._closedPromise_resolve = void 0;
      e._closedPromise_reject = void 0;
    }
  }
  function L(e, t) {
    if (e !== void 0 && typeof (r = e) != "object" && typeof r != "function") {
      throw new TypeError(`${t} is not an object.`);
    }
    var r;
  }
  function F(e, t) {
    if (typeof e != "function") {
      throw new TypeError(`${t} is not a function.`);
    }
  }
  function I(e, t) {
    if (
      !(function (e) {
        return (typeof e == "object" && e !== null) || typeof e == "function";
      })(e)
    ) {
      throw new TypeError(`${t} is not an object.`);
    }
  }
  function $(e, t, r) {
    if (e === void 0) {
      throw new TypeError(`Parameter ${t} is required in '${r}'.`);
    }
  }
  function M(e, t, r) {
    if (e === void 0) {
      throw new TypeError(`${t} is required in '${r}'.`);
    }
  }
  function Y(e) {
    return Number(e);
  }
  function Q(e) {
    if (e === 0) {
      return 0;
    } else {
      return e;
    }
  }
  function x(e, t) {
    const r = Number.MAX_SAFE_INTEGER;
    let o = Number(e);
    o = Q(o);
    if (!z(o)) {
      throw new TypeError(`${t} is not a finite number`);
    }
    o = (function () {
      var e = o;
      return Q(D(e));
    })();
    if (o < 0 || o > r) {
      throw new TypeError(
        `${t} is outside the accepted range of 0 to ${r}, inclusive`
      );
    }
    if (z(o) && o !== 0) {
      return o;
    } else {
      return 0;
    }
  }
  function N(e, t) {
    if (!Er(e)) {
      throw new TypeError(`${t} is not a ReadableStream.`);
    }
  }
  function H(e) {
    return new ReadableStreamDefaultReader(e);
  }
  function V(e, t) {
    e._reader._readRequests.push(t);
  }
  function U(e, t, r) {
    const o = e._reader._readRequests.shift();
    if (r) {
      o._closeSteps();
    } else {
      o._chunkSteps(t);
    }
  }
  function G(e) {
    return e._reader._readRequests.length;
  }
  function X(e) {
    const t = e._reader;
    return t !== void 0 && !!J(t);
  }
  function J(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_readRequests") &&
      e instanceof ReadableStreamDefaultReader
    );
  }
  function K(e, t) {
    const r = e._ownerReadableStream;
    r._disturbed = true;
    if (r._state === "closed") {
      t._closeSteps();
    } else if (r._state === "errored") {
      t._errorSteps(r._storedError);
    } else {
      r._readableStreamController[C](t);
    }
  }
  function Z(e, t) {
    const r = e._readRequests;
    e._readRequests = new v();
    r.forEach((e) => {
      e._errorSteps(t);
    });
  }
  function ee(e) {
    return new TypeError(
      `ReadableStreamDefaultReader.prototype.${e} can only be used on a ReadableStreamDefaultReader`
    );
  }
  function ne(e) {
    return e.slice();
  }
  function ae(e, t, r, o, n) {
    new Uint8Array(e).set(new Uint8Array(r, o, n), t);
  }
  function se(e, t, r) {
    if (e.slice) {
      return e.slice(t, r);
    }
    const o = r - t;
    const n = new ArrayBuffer(o);
    ae(n, 0, e, t, o);
    return n;
  }
  function ue(e, t) {
    const r = e[t];
    if (r != null) {
      if (typeof r != "function") {
        throw new TypeError(`${String(t)} is not a function`);
      }
      return r;
    }
  }
  function ce(e) {
    try {
      const t = e.done;
      const r = e.value;
      return f(s(r), (e) => ({ done: t, value: e }));
    } catch (e) {
      return d(e);
    }
  }
  function fe(e, r = "sync", o) {
    if (o === void 0) {
      if (r === "async") {
        if ((o = ue(e, de)) === void 0) {
          return (function (e) {
            const r = {
              next() {
                let t;
                try {
                  t = be(e);
                } catch (e) {
                  return d(e);
                }
                return ce(t);
              },
              return(r) {
                let o;
                try {
                  const t = ue(e.iterator, "return");
                  if (t === void 0) {
                    return c({ done: true, value: r });
                  }
                  o = S(t, e.iterator, [r]);
                } catch (e) {
                  return d(e);
                }
                if (t(o)) {
                  return ce(o);
                } else {
                  return d(
                    new TypeError(
                      "The iterator.return() method must return an object"
                    )
                  );
                }
              },
            };
            return { iterator: r, nextMethod: r.next, done: false };
          })(fe(e, "sync", ue(e, Symbol.iterator)));
        }
      } else {
        o = ue(e, Symbol.iterator);
      }
    }
    if (o === void 0) {
      throw new TypeError("The object is not iterable");
    }
    const n = S(o, e, []);
    if (!t(n)) {
      throw new TypeError("The iterator method must return an object");
    }
    return { iterator: n, nextMethod: n.next, done: false };
  }
  function be(e) {
    const r = S(e.nextMethod, e.iterator, []);
    if (!t(r)) {
      throw new TypeError("The iterator.next() method must return an object");
    }
    return r;
  }
  function _e(e) {
    if (!t(e)) {
      return false;
    }
    if (!Object.prototype.hasOwnProperty.call(e, "_asyncIteratorImpl")) {
      return false;
    }
    try {
      return e._asyncIteratorImpl instanceof he;
    } catch (e) {
      return false;
    }
  }
  function pe(e) {
    return new TypeError(
      `ReadableStreamAsyncIterator.${e} can only be used on a ReadableSteamAsyncIterator`
    );
  }
  function Se(e) {
    const t = se(e.buffer, e.byteOffset, e.byteOffset + e.byteLength);
    return new Uint8Array(t);
  }
  function ge(e) {
    const t = e._queue.shift();
    e._queueTotalSize -= t.size;
    if (e._queueTotalSize < 0) {
      e._queueTotalSize = 0;
    }
    return t.value;
  }
  function ve(e, t, r) {
    if (typeof (o = r) != "number" || ye(o) || o < 0 || r === 1 / 0) {
      throw new RangeError(
        "Size must be a finite, non-NaN, non-negative number."
      );
    }
    var o;
    e._queue.push({ value: t, size: r });
    e._queueTotalSize += r;
  }
  function we(e) {
    e._queue = new v();
    e._queueTotalSize = 0;
  }
  function Re(e) {
    return e === DataView;
  }
  function Te(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(
        e,
        "_controlledReadableByteStream"
      ) &&
      e instanceof ReadableByteStreamController
    );
  }
  function Ce(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(
        e,
        "_associatedReadableByteStreamController"
      ) &&
      e instanceof ReadableStreamBYOBRequest
    );
  }
  function Pe(e) {
    const t = (function (e) {
      const t = e._controlledReadableByteStream;
      if (t._state !== "readable") {
        return false;
      }
      if (e._closeRequested) {
        return false;
      }
      if (!e._started) {
        return false;
      }
      if (X(t) && G(t) > 0) {
        return true;
      }
      if (nt(t) && ot(t) > 0) {
        return true;
      }
      const r = Ue(e);
      if (r > 0) {
        return true;
      }
      return false;
    })(e);
    if (!t) {
      return;
    }
    if (e._pulling) {
      e._pullAgain = true;
      return;
    }
    e._pulling = true;
    b(
      e._pullAlgorithm(),
      () => (
        (e._pulling = false),
        e._pullAgain && ((e._pullAgain = false), Pe(e)),
        null
      ),
      (t) => (Ne(e, t), null)
    );
  }
  function qe(e) {
    Le(e);
    e._pendingPullIntos = new v();
  }
  function Ee(e, t) {
    let r = false;
    if (e._state === "closed") {
      r = true;
    }
    const o = Be(t);
    if (t.readerType === "default") {
      U(e, o, r);
    } else {
      (function (e, t, r) {
        const o = e._reader;
        const n = o._readIntoRequests.shift();
        if (r) {
          n._closeSteps(t);
        } else {
          n._chunkSteps(t);
        }
      })(e, o, r);
    }
  }
  function We(e, t) {
    for (let r = 0; r < t.length; ++r) {
      Ee(e, t[r]);
    }
  }
  function Be(e) {
    const t = e.bytesFilled;
    const r = e.elementSize;
    return new e.viewConstructor(e.buffer, e.byteOffset, t / r);
  }
  function Oe(e, t, r, o) {
    e._queue.push({ buffer: t, byteOffset: r, byteLength: o });
    e._queueTotalSize += o;
  }
  function je(e, t, r, o) {
    let n;
    try {
      n = se(t, r, r + o);
    } catch (t) {
      Ne(e, t);
      throw t;
    }
    Oe(e, n, 0, o);
  }
  function ke(e, t) {
    if (t.bytesFilled > 0) {
      je(e, t.buffer, t.byteOffset, t.bytesFilled);
    }
    Me(e);
  }
  function Ae(e, t) {
    const r = Math.min(e._queueTotalSize, t.byteLength - t.bytesFilled);
    const o = t.bytesFilled + r;
    let n = r;
    let a = false;
    const i = o - (o % t.elementSize);
    if (i >= t.minimumFill) {
      n = i - t.bytesFilled;
      a = true;
    }
    const l = e._queue;
    while (n > 0) {
      const r = l.peek();
      const o = Math.min(n, r.byteLength);
      const a = t.byteOffset + t.bytesFilled;
      ae(t.buffer, a, r.buffer, r.byteOffset, o);
      if (r.byteLength === o) {
        l.shift();
      } else {
        r.byteOffset += o;
        r.byteLength -= o;
      }
      e._queueTotalSize -= o;
      ze(e, o, t);
      n -= o;
    }
    return a;
  }
  function ze(e, t, r) {
    r.bytesFilled += t;
  }
  function De(e) {
    if (e._queueTotalSize === 0 && e._closeRequested) {
      Ye(e);
      Or(e._controlledReadableByteStream);
    } else {
      Pe(e);
    }
  }
  function Le(e) {
    if (e._byobRequest !== null) {
      e._byobRequest._associatedReadableByteStreamController = void 0;
      e._byobRequest._view = null;
      e._byobRequest = null;
    }
  }
  function Fe(e) {
    const t = [];
    while (e._pendingPullIntos.length > 0 && e._queueTotalSize !== 0) {
      const r = e._pendingPullIntos.peek();
      if (Ae(e, r)) {
        Me(e);
        t.push(r);
      }
    }
    return t;
  }
  function Ie(e, t, r, o) {
    const n = e._controlledReadableByteStream;
    const a = t.constructor;
    const i = (function () {
      var e = a;
      if (Re(e)) {
        return 1;
      } else {
        return e.BYTES_PER_ELEMENT;
      }
    })();
    const { byteOffset: l, byteLength: s } = t;
    const u = r * i;
    let c;
    try {
      c = ie(t.buffer);
    } catch (e) {
      o._errorSteps(e);
      return;
    }
    const d = {
      buffer: c,
      bufferByteLength: c.byteLength,
      byteOffset: l,
      byteLength: s,
      bytesFilled: 0,
      minimumFill: u,
      elementSize: i,
      viewConstructor: a,
      readerType: "byob",
    };
    if (e._pendingPullIntos.length > 0) {
      e._pendingPullIntos.push(d);
      rt(n, o);
      return;
    }
    if (n._state === "closed") {
      const e = new a(d.buffer, d.byteOffset, 0);
      o._closeSteps(e);
    } else {
      if (e._queueTotalSize > 0) {
        if (Ae(e, d)) {
          const t = Be(d);
          De(e);
          o._chunkSteps(t);
          return;
        }
        if (e._closeRequested) {
          const t = new TypeError(
            "Insufficient bytes to fill elements in the given buffer"
          );
          Ne(e, t);
          o._errorSteps(t);
          return;
        }
      }
      e._pendingPullIntos.push(d);
      rt(n, o);
      Pe(e);
    }
  }
  function $e(e, t) {
    const r = e._pendingPullIntos.peek();
    Le(e);
    if (e._controlledReadableByteStream._state === "closed") {
      (function (e, t) {
        if (t.readerType === "none") {
          Me(e);
        }
        const r = e._controlledReadableByteStream;
        if (nt(r)) {
          const t = [];
          for (let o = 0; o < ot(r); ++o) {
            t.push(Me(e));
          }
          We(r, t);
        }
      })(e, r);
    } else {
      (function (e, t, r) {
        ze(0, t, r);
        if (r.readerType === "none") {
          ke(e, r);
          const t = Fe(e);
          We(e._controlledReadableByteStream, t);
          return;
        }
        if (r.bytesFilled < r.minimumFill) {
          return;
        }
        Me(e);
        const o = r.bytesFilled % r.elementSize;
        if (o > 0) {
          const t = r.byteOffset + r.bytesFilled;
          je(e, r.buffer, t - o, o);
        }
        r.bytesFilled -= o;
        const n = Fe(e);
        Ee(e._controlledReadableByteStream, r);
        We(e._controlledReadableByteStream, n);
      })(e, t, r);
    }
    Pe(e);
  }
  function Me(e) {
    return e._pendingPullIntos.shift();
  }
  function Ye(e) {
    e._pullAlgorithm = void 0;
    e._cancelAlgorithm = void 0;
  }
  function Qe(e) {
    const t = e._controlledReadableByteStream;
    if (!e._closeRequested && t._state === "readable") {
      if (e._queueTotalSize > 0) {
        e._closeRequested = true;
      } else {
        if (e._pendingPullIntos.length > 0) {
          const t = e._pendingPullIntos.peek();
          if (t.bytesFilled % t.elementSize != 0) {
            const t = new TypeError(
              "Insufficient bytes to fill elements in the given buffer"
            );
            Ne(e, t);
            throw t;
          }
        }
        Ye(e);
        Or(t);
      }
    }
  }
  function xe(e, t) {
    const r = e._controlledReadableByteStream;
    if (e._closeRequested || r._state !== "readable") {
      return;
    }
    const { buffer: o, byteOffset: n, byteLength: a } = t;
    if (le(o)) {
      throw new TypeError(
        "chunk's buffer is detached and so cannot be enqueued"
      );
    }
    const i = ie(o);
    if (e._pendingPullIntos.length > 0) {
      const t = e._pendingPullIntos.peek();
      if (le(t.buffer)) {
        throw new TypeError(
          "The BYOB request's buffer has been detached and so cannot be filled with an enqueued chunk"
        );
      }
      Le(e);
      t.buffer = ie(t.buffer);
      if (t.readerType === "none") {
        ke(e, t);
      }
    }
    if (X(r)) {
      (function (e) {
        const t = e._controlledReadableByteStream._reader;
        while (t._readRequests.length > 0) {
          if (e._queueTotalSize === 0) {
            return;
          }
          He(e, t._readRequests.shift());
        }
      })(e);
      if (G(r) === 0) {
        Oe(e, i, n, a);
      } else {
        if (e._pendingPullIntos.length > 0) {
          Me(e);
        }
        U(r, new Uint8Array(i, n, a), false);
      }
    } else if (nt(r)) {
      Oe(e, i, n, a);
      const t = Fe(e);
      We(e._controlledReadableByteStream, t);
    } else {
      Oe(e, i, n, a);
    }
    Pe(e);
  }
  function Ne(e, t) {
    const r = e._controlledReadableByteStream;
    if (r._state === "readable") {
      qe(e);
      we(e);
      Ye(e);
      jr(r, t);
    }
  }
  function He(e, t) {
    const r = e._queue.shift();
    e._queueTotalSize -= r.byteLength;
    De(e);
    const o = new Uint8Array(r.buffer, r.byteOffset, r.byteLength);
    t._chunkSteps(o);
  }
  function Ve(e) {
    if (e._byobRequest === null && e._pendingPullIntos.length > 0) {
      const t = e._pendingPullIntos.peek();
      const r = new Uint8Array(
        t.buffer,
        t.byteOffset + t.bytesFilled,
        t.byteLength - t.bytesFilled
      );
      const o = Object.create(ReadableStreamBYOBRequest.prototype);
      (function (e, t, r) {
        e._associatedReadableByteStreamController = t;
        e._view = r;
      })(o, e, r);
      e._byobRequest = o;
    }
    return e._byobRequest;
  }
  function Ue(e) {
    const t = e._controlledReadableByteStream._state;
    if (t === "errored") {
      return null;
    } else if (t === "closed") {
      return 0;
    } else {
      return e._strategyHWM - e._queueTotalSize;
    }
  }
  function Ge(e, t) {
    const r = e._pendingPullIntos.peek();
    if (e._controlledReadableByteStream._state === "closed") {
      if (t !== 0) {
        throw new TypeError(
          "bytesWritten must be 0 when calling respond() on a closed stream"
        );
      }
    } else {
      if (t === 0) {
        throw new TypeError(
          "bytesWritten must be greater than 0 when calling respond() on a readable stream"
        );
      }
      if (r.bytesFilled + t > r.byteLength) {
        throw new RangeError("bytesWritten out of range");
      }
    }
    r.buffer = ie(r.buffer);
    $e(e, t);
  }
  function Xe(e, t) {
    const r = e._pendingPullIntos.peek();
    if (e._controlledReadableByteStream._state === "closed") {
      if (t.byteLength !== 0) {
        throw new TypeError(
          "The view's length must be 0 when calling respondWithNewView() on a closed stream"
        );
      }
    } else if (t.byteLength === 0) {
      throw new TypeError(
        "The view's length must be greater than 0 when calling respondWithNewView() on a readable stream"
      );
    }
    if (r.byteOffset + r.bytesFilled !== t.byteOffset) {
      throw new RangeError(
        "The region specified by view does not match byobRequest"
      );
    }
    if (r.bufferByteLength !== t.buffer.byteLength) {
      throw new RangeError(
        "The buffer of view has different capacity than byobRequest"
      );
    }
    if (r.bytesFilled + t.byteLength > r.byteLength) {
      throw new RangeError(
        "The region specified by view is larger than byobRequest"
      );
    }
    const o = t.byteLength;
    r.buffer = ie(t.buffer);
    $e(e, o);
  }
  function Je(e, t, r, o, n, a, i) {
    t._controlledReadableByteStream = e;
    t._pullAgain = false;
    t._pulling = false;
    t._byobRequest = null;
    t._queue = t._queueTotalSize = void 0;
    we(t);
    t._closeRequested = false;
    t._started = false;
    t._strategyHWM = a;
    t._pullAlgorithm = o;
    t._cancelAlgorithm = n;
    t._autoAllocateChunkSize = i;
    t._pendingPullIntos = new v();
    e._readableStreamController = t;
    b(
      c(r()),
      () => ((t._started = true), Pe(t), null),
      (e) => (Ne(t, e), null)
    );
  }
  function Ke(e) {
    return new TypeError(
      `ReadableStreamBYOBRequest.prototype.${e} can only be used on a ReadableStreamBYOBRequest`
    );
  }
  function Ze(e) {
    return new TypeError(
      `ReadableByteStreamController.prototype.${e} can only be used on a ReadableByteStreamController`
    );
  }
  function et(e, t) {
    if ((e = `${e}`) !== "byob") {
      throw new TypeError(
        `${t} '${e}' is not a valid enumeration value for ReadableStreamReaderMode`
      );
    }
    return e;
  }
  function tt(e) {
    return new ReadableStreamBYOBReader(e);
  }
  function rt(e, t) {
    e._reader._readIntoRequests.push(t);
  }
  function ot(e) {
    return e._reader._readIntoRequests.length;
  }
  function nt(e) {
    const t = e._reader;
    return t !== void 0 && !!at(t);
  }
  function at(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_readIntoRequests") &&
      e instanceof ReadableStreamBYOBReader
    );
  }
  function it(e, t, r, o) {
    const n = e._ownerReadableStream;
    n._disturbed = true;
    if (n._state === "errored") {
      o._errorSteps(n._storedError);
    } else {
      Ie(n._readableStreamController, t, r, o);
    }
  }
  function lt(e, t) {
    const r = e._readIntoRequests;
    e._readIntoRequests = new v();
    r.forEach((e) => {
      e._errorSteps(t);
    });
  }
  function st(e) {
    return new TypeError(
      `ReadableStreamBYOBReader.prototype.${e} can only be used on a ReadableStreamBYOBReader`
    );
  }
  function ut(e, t) {
    const { highWaterMark: r } = e;
    if (r === void 0) {
      return t;
    }
    if (ye(r) || r < 0) {
      throw new RangeError("Invalid highWaterMark");
    }
    return r;
  }
  function ct(e) {
    const { size: t } = e;
    return t || (() => 1);
  }
  function dt(e, t) {
    L(e, t);
    const r = e == null ? void 0 : e.highWaterMark;
    const o = e == null ? void 0 : e.size;
    return {
      highWaterMark: r === void 0 ? void 0 : Y(r),
      size: o === void 0 ? void 0 : ft(o, `${t} has member 'size' that`),
    };
  }
  function ft(e, t) {
    F(e, t);
    return (t) => Y(e(t));
  }
  function bt(e, t, r) {
    F(e, r);
    return (r) => g(e, t, [r]);
  }
  function ht(e, t, r) {
    F(e, r);
    return () => g(e, t, []);
  }
  function mt(e, t, r) {
    F(e, r);
    return (r) => S(e, t, [r]);
  }
  function _t(e, t, r) {
    F(e, r);
    return (r, o) => g(e, t, [r, o]);
  }
  function pt(e, t) {
    if (!gt(e)) {
      throw new TypeError(`${t} is not a WritableStream.`);
    }
  }
  function yt(e) {
    return new WritableStreamDefaultWriter(e);
  }
  function St(e) {
    e._state = "writable";
    e._storedError = void 0;
    e._writer = void 0;
    e._writableStreamController = void 0;
    e._writeRequests = new v();
    e._inFlightWriteRequest = void 0;
    e._closeRequest = void 0;
    e._inFlightCloseRequest = void 0;
    e._pendingAbortRequest = void 0;
    e._backpressure = false;
  }
  function gt(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_writableStreamController") &&
      e instanceof WritableStream
    );
  }
  function vt(e) {
    return e._writer !== void 0;
  }
  function wt(e, t) {
    var r;
    if (e._state === "closed" || e._state === "errored") {
      return c(void 0);
    }
    e._writableStreamController._abortReason = t;
    if (
      (r = e._writableStreamController._abortController) !== null &&
      r !== void 0
    ) {
      r.abort(t);
    }
    const o = e._state;
    if (o === "closed" || o === "errored") {
      return c(void 0);
    }
    if (e._pendingAbortRequest !== void 0) {
      return e._pendingAbortRequest._promise;
    }
    let n = false;
    if (o === "erroring") {
      n = true;
      t = void 0;
    }
    const a = u((r, o) => {
      e._pendingAbortRequest = {
        _promise: void 0,
        _resolve: r,
        _reject: o,
        _reason: t,
        _wasAlreadyErroring: n,
      };
    });
    e._pendingAbortRequest._promise = a;
    if (!n) {
      Ct(e, t);
    }
    return a;
  }
  function Rt(e) {
    const t = e._state;
    if (t === "closed" || t === "errored") {
      return d(
        new TypeError(
          `The stream (in ${t} state) is not in the writable state and cannot be closed`
        )
      );
    }
    const r = u((t, r) => {
      const o = { _resolve: t, _reject: r };
      e._closeRequest = o;
    });
    const o = e._writer;
    var n;
    if (o !== void 0 && e._backpressure && t === "writable") {
      or(o);
    }
    ve((n = e._writableStreamController), Dt, 0);
    Mt(n);
    return r;
  }
  function Tt(e, t) {
    if (e._state === "writable") {
      Ct(e, t);
    } else {
      Pt(e);
    }
  }
  function Ct(e, t) {
    const r = e._writableStreamController;
    e._state = "erroring";
    e._storedError = t;
    const o = e._writer;
    if (o !== void 0) {
      kt(o, t);
    }
    if (
      !(function (e) {
        if (
          e._inFlightWriteRequest === void 0 &&
          e._inFlightCloseRequest === void 0
        ) {
          return false;
        }
        return true;
      })(e) &&
      r._started
    ) {
      Pt(e);
    }
  }
  function Pt(e) {
    e._state = "errored";
    e._writableStreamController[R]();
    const t = e._storedError;
    e._writeRequests.forEach((e) => {
      e._reject(t);
    });
    e._writeRequests = new v();
    if (e._pendingAbortRequest === void 0) {
      Et(e);
      return;
    }
    const r = e._pendingAbortRequest;
    e._pendingAbortRequest = void 0;
    if (r._wasAlreadyErroring) {
      r._reject(t);
      Et(e);
      return;
    }
    b(
      e._writableStreamController[w](r._reason),
      () => (r._resolve(), Et(e), null),
      (t) => (r._reject(t), Et(e), null)
    );
  }
  function qt(e) {
    return e._closeRequest !== void 0 || e._inFlightCloseRequest !== void 0;
  }
  function Et(e) {
    if (e._closeRequest !== void 0) {
      e._closeRequest._reject(e._storedError);
      e._closeRequest = void 0;
    }
    const t = e._writer;
    if (t !== void 0) {
      Jt(t, e._storedError);
    }
  }
  function Wt(e, t) {
    const r = e._writer;
    if (r !== void 0 && t !== e._backpressure) {
      if (t) {
        (function () {
          var e = r;
          Zt(e);
        })();
      } else {
        or(r);
      }
    }
    e._backpressure = t;
  }
  function Bt(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_ownerWritableStream") &&
      e instanceof WritableStreamDefaultWriter
    );
  }
  function Ot(e) {
    return Rt(e._ownerWritableStream);
  }
  function jt(e, t) {
    if (e._closedPromiseState === "pending") {
      Jt(e, t);
    } else {
      (function (e, t) {
        Xt(e, t);
      })(e, t);
    }
  }
  function kt(e, t) {
    if (e._readyPromiseState === "pending") {
      rr(e, t);
    } else {
      (function (e, t) {
        er(e, t);
      })(e, t);
    }
  }
  function At(e) {
    const t = e._ownerWritableStream;
    const r = new TypeError(
      "Writer was released and can no longer be used to monitor the stream's closedness"
    );
    kt(e, r);
    jt(e, r);
    t._writer = void 0;
    e._ownerWritableStream = void 0;
  }
  function zt(e, t) {
    const r = e._ownerWritableStream;
    const o = r._writableStreamController;
    const n = (function (e, t) {
      if (e._strategySizeAlgorithm === void 0) {
        return 1;
      }
      try {
        return e._strategySizeAlgorithm(t);
      } catch (t) {
        Yt(e, t);
        return 1;
      }
    })(o, t);
    if (r !== e._ownerWritableStream) {
      return d(Ut("write to"));
    }
    const a = r._state;
    if (a === "errored") {
      return d(r._storedError);
    }
    if (qt(r) || a === "closed") {
      return d(
        new TypeError(
          "The stream is closing or closed and cannot be written to"
        )
      );
    }
    if (a === "erroring") {
      return d(r._storedError);
    }
    const i = (function () {
      var e = r;
      return u((t, r) => {
        const o = { _resolve: t, _reject: r };
        e._writeRequests.push(o);
      });
    })();
    (function (e, t, r) {
      try {
        ve(e, t, r);
      } catch (t) {
        Yt(e, t);
        return;
      }
      const o = e._controlledWritableStream;
      if (!qt(o) && o._state === "writable") {
        Wt(o, Qt(e));
      }
      Mt(e);
    })(o, t, n);
    return i;
  }
  function Lt(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_controlledWritableStream") &&
      e instanceof WritableStreamDefaultController
    );
  }
  function Ft(e, t, r, o, n, a, i, l) {
    t._controlledWritableStream = e;
    e._writableStreamController = t;
    t._queue = void 0;
    t._queueTotalSize = void 0;
    we(t);
    t._abortReason = void 0;
    t._abortController = (function () {
      if (typeof AbortController == "function") {
        return new AbortController();
      }
    })();
    t._started = false;
    t._strategySizeAlgorithm = l;
    t._strategyHWM = i;
    t._writeAlgorithm = o;
    t._closeAlgorithm = n;
    t._abortAlgorithm = a;
    const s = Qt(t);
    Wt(e, s);
    b(
      c(r()),
      () => ((t._started = true), Mt(t), null),
      (r) => ((t._started = true), Tt(e, r), null)
    );
  }
  function It(e) {
    e._writeAlgorithm = void 0;
    e._closeAlgorithm = void 0;
    e._abortAlgorithm = void 0;
    e._strategySizeAlgorithm = void 0;
  }
  function $t(e) {
    return e._strategyHWM - e._queueTotalSize;
  }
  function Mt(e) {
    const t = e._controlledWritableStream;
    if (!e._started) {
      return;
    }
    if (t._inFlightWriteRequest !== void 0) {
      return;
    }
    if (t._state === "erroring") {
      Pt(t);
      return;
    }
    if (e._queue.length === 0) {
      return;
    }
    const r = e._queue.peek().value;
    if (r === Dt) {
      (function (e) {
        const t = e._controlledWritableStream;
        (function () {
          var e = t;
          e._inFlightCloseRequest = e._closeRequest;
          e._closeRequest = void 0;
        })();
        ge(e);
        const r = e._closeAlgorithm();
        It(e);
        b(
          r,
          () => (
            (function (e) {
              e._inFlightCloseRequest._resolve(void 0);
              e._inFlightCloseRequest = void 0;
              if (e._state === "erroring") {
                e._storedError = void 0;
                if (e._pendingAbortRequest !== void 0) {
                  e._pendingAbortRequest._resolve();
                  e._pendingAbortRequest = void 0;
                }
              }
              e._state = "closed";
              const t = e._writer;
              if (t !== void 0) {
                Kt(t);
              }
            })(t),
            null
          ),
          (e) => (
            (function (e, t) {
              e._inFlightCloseRequest._reject(t);
              e._inFlightCloseRequest = void 0;
              if (e._pendingAbortRequest !== void 0) {
                e._pendingAbortRequest._reject(t);
                e._pendingAbortRequest = void 0;
              }
              Tt(e, t);
            })(t, e),
            null
          )
        );
      })(e);
    } else {
      (function (e, t) {
        const r = e._controlledWritableStream;
        (function () {
          var e = r;
          e._inFlightWriteRequest = e._writeRequests.shift();
        })();
        const o = e._writeAlgorithm(t);
        b(
          o,
          () => {
            (function () {
              var e = r;
              e._inFlightWriteRequest._resolve(void 0);
              e._inFlightWriteRequest = void 0;
            })();
            const t = r._state;
            ge(e);
            if (!qt(r) && t === "writable") {
              const t = Qt(e);
              Wt(r, t);
            }
            Mt(e);
            return null;
          },
          (t) => (
            r._state === "writable" && It(e),
            (function (e, t) {
              e._inFlightWriteRequest._reject(t);
              e._inFlightWriteRequest = void 0;
              Tt(e, t);
            })(r, t),
            null
          )
        );
      })(e, r);
    }
  }
  function Yt(e, t) {
    if (e._controlledWritableStream._state === "writable") {
      xt(e, t);
    }
  }
  function Qt(e) {
    return $t(e) <= 0;
  }
  function xt(e, t) {
    const r = e._controlledWritableStream;
    It(e);
    Ct(r, t);
  }
  function Nt(e) {
    return new TypeError(
      `WritableStream.prototype.${e} can only be used on a WritableStream`
    );
  }
  function Ht(e) {
    return new TypeError(
      `WritableStreamDefaultController.prototype.${e} can only be used on a WritableStreamDefaultController`
    );
  }
  function Vt(e) {
    return new TypeError(
      `WritableStreamDefaultWriter.prototype.${e} can only be used on a WritableStreamDefaultWriter`
    );
  }
  function Ut(e) {
    return new TypeError("Cannot " + e + " a stream using a released writer");
  }
  function Gt(e) {
    e._closedPromise = u((t, r) => {
      e._closedPromise_resolve = t;
      e._closedPromise_reject = r;
      e._closedPromiseState = "pending";
    });
  }
  function Xt(e, t) {
    Gt(e);
    Jt(e, t);
  }
  function Jt(e, t) {
    if (e._closedPromise_reject !== void 0) {
      p(e._closedPromise);
      e._closedPromise_reject(t);
      e._closedPromise_resolve = void 0;
      e._closedPromise_reject = void 0;
      e._closedPromiseState = "rejected";
    }
  }
  function Kt(e) {
    if (e._closedPromise_resolve !== void 0) {
      e._closedPromise_resolve(void 0);
      e._closedPromise_resolve = void 0;
      e._closedPromise_reject = void 0;
      e._closedPromiseState = "resolved";
    }
  }
  function Zt(e) {
    e._readyPromise = u((t, r) => {
      e._readyPromise_resolve = t;
      e._readyPromise_reject = r;
    });
    e._readyPromiseState = "pending";
  }
  function er(e, t) {
    Zt(e);
    rr(e, t);
  }
  function tr(e) {
    Zt(e);
    or(e);
  }
  function rr(e, t) {
    if (e._readyPromise_reject !== void 0) {
      p(e._readyPromise);
      e._readyPromise_reject(t);
      e._readyPromise_resolve = void 0;
      e._readyPromise_reject = void 0;
      e._readyPromiseState = "rejected";
    }
  }
  function or(e) {
    if (e._readyPromise_resolve !== void 0) {
      e._readyPromise_resolve(void 0);
      e._readyPromise_resolve = void 0;
      e._readyPromise_reject = void 0;
      e._readyPromiseState = "fulfilled";
    }
  }
  function ir(t, r, o, n, a, i) {
    const l = H(t);
    const s = yt(r);
    t._disturbed = true;
    let _ = false;
    let y = c(void 0);
    return u((S, g) => {
      function C() {
        const e = y;
        return f(y, () => (e !== y ? C() : void 0));
      }
      function P(e, t, r) {
        if (e._state === "errored") {
          r(e._storedError);
        } else {
          m(t, r);
        }
      }
      function q(e, t, o) {
        function n() {
          b(
            e(),
            () => B(t, o),
            (e) => B(true, e)
          );
          return null;
        }
        if (!_) {
          _ = true;
          if (r._state !== "writable" || qt(r)) {
            n();
          } else {
            h(C(), n);
          }
        }
      }
      function E(e, t) {
        if (!_) {
          _ = true;
          if (r._state !== "writable" || qt(r)) {
            B(e, t);
          } else {
            h(C(), () => B(e, t));
          }
        }
      }
      function B(e, t) {
        At(s);
        W(l);
        if (i !== void 0) {
          i.removeEventListener("abort", v);
        }
        if (e) {
          g(t);
        } else {
          S(void 0);
        }
        return null;
      }
      let v;
      if (i !== void 0) {
        v = () => {
          const e =
            i.reason !== void 0 ? i.reason : new ar("Aborted", "AbortError");
          const o = [];
          if (!n) {
            o.push(() => (r._state === "writable" ? wt(r, e) : c(void 0)));
          }
          if (!a) {
            o.push(() => (t._state === "readable" ? Br(t, e) : c(void 0)));
          }
          q(() => Promise.all(o.map((e) => e())), true, e);
        };
        if (i.aborted) {
          v();
          return;
        }
        i.addEventListener("abort", v);
      }
      P(
        t,
        l._closedPromise,
        (e) => (n ? E(true, e) : q(() => wt(r, e), true, e), null)
      );
      P(
        r,
        s._closedPromise,
        (e) => (a ? E(true, e) : q(() => Br(t, e), true, e), null)
      );
      var w = t;
      var R = l._closedPromise;
      T = () => (
        o
          ? E()
          : q(() =>
              (function () {
                var e = s;
                const t = e._ownerWritableStream;
                const r = t._state;
                if (qt(t) || r === "closed") {
                  return c(void 0);
                } else if (r === "errored") {
                  return d(t._storedError);
                } else {
                  return Ot(e);
                }
              })()
            ),
        null
      );
      if (w._state === "closed") {
        T();
      } else {
        h(R, T);
      }
      if (qt(r) || r._state === "closed") {
        const e = new TypeError(
          "the destination writable stream closed before all data could be piped to it"
        );
        if (a) {
          E(true, e);
        } else {
          q(() => Br(t, e), true, e);
        }
      }
      p(
        u((t, r) => {
          (function o() {
            var n = false;
            if (n) {
              t();
            } else {
              f(
                _
                  ? c(true)
                  : f(s._readyPromise, () =>
                      u((t, r) => {
                        K(l, {
                          _chunkSteps: (r) => {
                            y = f(zt(s, r), void 0, e);
                            t(false);
                          },
                          _closeSteps: () => t(true),
                          _errorSteps: r,
                        });
                      })
                    ),
                o,
                r
              );
            }
          })();
        })
      );
    });
  }
  function lr(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_controlledReadableStream") &&
      e instanceof ReadableStreamDefaultController
    );
  }
  function sr(e) {
    if (!ur(e)) {
      return;
    }
    if (e._pulling) {
      e._pullAgain = true;
      return;
    }
    e._pulling = true;
    b(
      e._pullAlgorithm(),
      () => (
        (e._pulling = false),
        e._pullAgain && ((e._pullAgain = false), sr(e)),
        null
      ),
      (t) => (br(e, t), null)
    );
  }
  function ur(e) {
    const t = e._controlledReadableStream;
    if (!mr(e)) {
      return false;
    }
    if (!e._started) {
      return false;
    }
    if (Wr(t) && G(t) > 0) {
      return true;
    }
    return hr(e) > 0;
  }
  function cr(e) {
    e._pullAlgorithm = void 0;
    e._cancelAlgorithm = void 0;
    e._strategySizeAlgorithm = void 0;
  }
  function dr(e) {
    if (!mr(e)) {
      return;
    }
    const t = e._controlledReadableStream;
    e._closeRequested = true;
    if (e._queue.length === 0) {
      cr(e);
      Or(t);
    }
  }
  function fr(e, t) {
    if (!mr(e)) {
      return;
    }
    const r = e._controlledReadableStream;
    if (Wr(r) && G(r) > 0) {
      U(r, t, false);
    } else {
      let r;
      try {
        r = e._strategySizeAlgorithm(t);
      } catch (t) {
        br(e, t);
        throw t;
      }
      try {
        ve(e, t, r);
      } catch (t) {
        br(e, t);
        throw t;
      }
    }
    sr(e);
  }
  function br(e, t) {
    const r = e._controlledReadableStream;
    if (r._state === "readable") {
      we(e);
      cr(e);
      jr(r, t);
    }
  }
  function hr(e) {
    const t = e._controlledReadableStream._state;
    if (t === "errored") {
      return null;
    } else if (t === "closed") {
      return 0;
    } else {
      return e._strategyHWM - e._queueTotalSize;
    }
  }
  function mr(e) {
    const t = e._controlledReadableStream._state;
    return !e._closeRequested && t === "readable";
  }
  function _r(e, t, r, o, n, a, i) {
    t._controlledReadableStream = e;
    t._queue = void 0;
    t._queueTotalSize = void 0;
    we(t);
    t._started = false;
    t._closeRequested = false;
    t._pullAgain = false;
    t._pulling = false;
    t._strategySizeAlgorithm = i;
    t._strategyHWM = a;
    t._pullAlgorithm = o;
    t._cancelAlgorithm = n;
    e._readableStreamController = t;
    b(
      c(r()),
      () => ((t._started = true), sr(t), null),
      (e) => (br(t, e), null)
    );
  }
  function pr(e) {
    return new TypeError(
      `ReadableStreamDefaultController.prototype.${e} can only be used on a ReadableStreamDefaultController`
    );
  }
  function yr(e, t) {
    if (Te(e._readableStreamController)) {
      return (function (e) {
        function _(e) {
          m(
            e._closedPromise,
            (t) => (
              e !== i ||
                (Ne(o._readableStreamController, t),
                Ne(n._readableStreamController, t),
                (f && b) || a(void 0)),
              null
            )
          );
        }
        function p() {
          if (at(i)) {
            W(i);
            i = H(e);
            _(i);
          }
          K(i, {
            _chunkSteps: (t) => {
              y(() => {
                s = false;
                d = false;
                const r = t;
                let i = t;
                if (!f && !b) {
                  try {
                    i = Se(t);
                  } catch (t) {
                    Ne(o._readableStreamController, t);
                    Ne(n._readableStreamController, t);
                    a(Br(e, t));
                    return;
                  }
                }
                if (!f) {
                  xe(o._readableStreamController, r);
                }
                if (!b) {
                  xe(n._readableStreamController, i);
                }
                l = false;
                if (s) {
                  g();
                } else if (d) {
                  v();
                }
              });
            },
            _closeSteps: () => {
              l = false;
              if (!f) {
                Qe(o._readableStreamController);
              }
              if (!b) {
                Qe(n._readableStreamController);
              }
              if (o._readableStreamController._pendingPullIntos.length > 0) {
                Ge(o._readableStreamController, 0);
              }
              if (n._readableStreamController._pendingPullIntos.length > 0) {
                Ge(n._readableStreamController, 0);
              }
              if (!f || !b) {
                a(void 0);
              }
            },
            _errorSteps: () => {
              l = false;
            },
          });
        }
        function S(t, r) {
          if (J(i)) {
            W(i);
            i = tt(e);
            _(i);
          }
          const u = r ? n : o;
          const c = r ? o : n;
          it(i, t, 1, {
            _chunkSteps: (t) => {
              y(() => {
                s = false;
                d = false;
                const o = r ? b : f;
                if (r ? f : b) {
                  if (!o) {
                    Xe(u._readableStreamController, t);
                  }
                } else {
                  let r;
                  try {
                    r = Se(t);
                  } catch (t) {
                    Ne(u._readableStreamController, t);
                    Ne(c._readableStreamController, t);
                    a(Br(e, t));
                    return;
                  }
                  if (!o) {
                    Xe(u._readableStreamController, t);
                  }
                  xe(c._readableStreamController, r);
                }
                l = false;
                if (s) {
                  g();
                } else if (d) {
                  v();
                }
              });
            },
            _closeSteps: (e) => {
              l = false;
              const t = r ? b : f;
              const o = r ? f : b;
              if (!t) {
                Qe(u._readableStreamController);
              }
              if (!o) {
                Qe(c._readableStreamController);
              }
              if (e !== void 0) {
                if (!t) {
                  Xe(u._readableStreamController, e);
                }
                if (
                  !o &&
                  c._readableStreamController._pendingPullIntos.length > 0
                ) {
                  Ge(c._readableStreamController, 0);
                }
              }
              if (!t || !o) {
                a(void 0);
              }
            },
            _errorSteps: () => {
              l = false;
            },
          });
        }
        function g() {
          if (l) {
            s = true;
            return c(void 0);
          }
          l = true;
          const e = Ve(o._readableStreamController);
          if (e === null) {
            p();
          } else {
            S(e._view, false);
          }
          return c(void 0);
        }
        function v() {
          if (l) {
            d = true;
            return c(void 0);
          }
          l = true;
          const e = Ve(n._readableStreamController);
          if (e === null) {
            p();
          } else {
            S(e._view, true);
          }
          return c(void 0);
        }
        function w(o) {
          f = true;
          t = o;
          if (b) {
            const o = ne([t, r]);
            const n = Br(e, o);
            a(n);
          }
          return h;
        }
        function R(o) {
          b = true;
          r = o;
          if (f) {
            const o = ne([t, r]);
            const n = Br(e, o);
            a(n);
          }
          return h;
        }
        function T() {}
        let t;
        let r;
        let o;
        let n;
        let a;
        let i = H(e);
        let l = false;
        let s = false;
        let d = false;
        let f = false;
        let b = false;
        const h = u((e) => {
          a = e;
        });
        o = Pr(T, g, w);
        n = Pr(T, v, R);
        _(i);
        return [o, n];
      })(e);
    } else {
      return (function (e, t) {
        function _() {
          if (s) {
            d = true;
            return c(void 0);
          }
          s = true;
          K(r, {
            _chunkSteps: (e) => {
              y(() => {
                d = false;
                const t = e;
                const r = e;
                if (!f) {
                  fr(a._readableStreamController, t);
                }
                if (!b) {
                  fr(i._readableStreamController, r);
                }
                s = false;
                if (d) {
                  _();
                }
              });
            },
            _closeSteps: () => {
              s = false;
              if (!f) {
                dr(a._readableStreamController);
              }
              if (!b) {
                dr(i._readableStreamController);
              }
              if (!f || !b) {
                l(void 0);
              }
            },
            _errorSteps: () => {
              s = false;
            },
          });
          return c(void 0);
        }
        function p(t) {
          f = true;
          o = t;
          if (b) {
            const t = ne([o, n]);
            const r = Br(e, t);
            l(r);
          }
          return h;
        }
        function S(t) {
          b = true;
          n = t;
          if (f) {
            const t = ne([o, n]);
            const r = Br(e, t);
            l(r);
          }
          return h;
        }
        function g() {}
        const r = H(e);
        let o;
        let n;
        let a;
        let i;
        let l;
        let s = false;
        let d = false;
        let f = false;
        let b = false;
        const h = u((e) => {
          l = e;
        });
        a = Cr(g, _, p);
        i = Cr(g, _, S);
        m(
          r._closedPromise,
          (e) => (
            br(a._readableStreamController, e),
            br(i._readableStreamController, e),
            (f && b) || l(void 0),
            null
          )
        );
        return [a, i];
      })(e);
    }
  }
  function Sr(r) {
    if (t((o = r)) && o.getReader !== void 0) {
      return (function (r) {
        function n() {
          let e;
          try {
            e = r.read();
          } catch (e) {
            return d(e);
          }
          return _(e, (e) => {
            if (!t(e)) {
              throw new TypeError(
                "The promise returned by the reader.read() method must fulfill with an object"
              );
            }
            if (e.done) {
              dr(o._readableStreamController);
            } else {
              const t = e.value;
              fr(o._readableStreamController, t);
            }
          });
        }
        function a(e) {
          try {
            return c(r.cancel(e));
          } catch (e) {
            return d(e);
          }
        }
        let o;
        o = Cr(e, n, a, 0);
        return o;
      })(r.getReader());
    } else {
      return (function (r) {
        function a() {
          let e;
          try {
            e = be(n);
          } catch (e) {
            return d(e);
          }
          return _(c(e), (e) => {
            if (!t(e)) {
              throw new TypeError(
                "The promise returned by the iterator.next() method must fulfill with an object"
              );
            }
            if (e.done) {
              dr(o._readableStreamController);
            } else {
              const t = e.value;
              fr(o._readableStreamController, t);
            }
          });
        }
        function i(e) {
          const r = n.iterator;
          let o;
          try {
            o = ue(r, "return");
          } catch (e) {
            return d(e);
          }
          if (o === void 0) {
            return c(void 0);
          }
          return _(g(o, r, [e]), (e) => {
            if (!t(e)) {
              throw new TypeError(
                "The promise returned by the iterator.return() method must fulfill with an object"
              );
            }
          });
        }
        let o;
        const n = fe(r, "async");
        o = Cr(e, a, i, 0);
        return o;
      })(r);
    }
    var o;
  }
  function gr(e, t, r) {
    F(e, r);
    return (r) => g(e, t, [r]);
  }
  function vr(e, t, r) {
    F(e, r);
    return (r) => g(e, t, [r]);
  }
  function wr(e, t, r) {
    F(e, r);
    return (r) => S(e, t, [r]);
  }
  function Rr(e, t) {
    if ((e = `${e}`) !== "bytes") {
      throw new TypeError(
        `${t} '${e}' is not a valid enumeration value for ReadableStreamType`
      );
    }
    return e;
  }
  function Tr(e, t) {
    L(e, t);
    const r = e == null ? void 0 : e.preventAbort;
    const o = e == null ? void 0 : e.preventCancel;
    const n = e == null ? void 0 : e.preventClose;
    const a = e == null ? void 0 : e.signal;
    if (a !== void 0) {
      (function (e, t) {
        if (
          !(function (e) {
            if (typeof e != "object" || e === null) {
              return false;
            }
            try {
              return typeof e.aborted == "boolean";
            } catch (e) {
              return false;
            }
          })(e)
        ) {
          throw new TypeError(`${t} is not an AbortSignal.`);
        }
      })(a, `${t} has member 'signal' that`);
    }
    return {
      preventAbort: Boolean(r),
      preventCancel: Boolean(o),
      preventClose: Boolean(n),
      signal: a,
    };
  }
  function Cr(e, t, r, o = 1, n = () => 1) {
    const a = Object.create(ReadableStream.prototype);
    qr(a);
    _r(
      a,
      Object.create(ReadableStreamDefaultController.prototype),
      e,
      t,
      r,
      o,
      n
    );
    return a;
  }
  function Pr(e, t, r) {
    const o = Object.create(ReadableStream.prototype);
    qr(o);
    Je(
      o,
      Object.create(ReadableByteStreamController.prototype),
      e,
      t,
      r,
      0,
      void 0
    );
    return o;
  }
  function qr(e) {
    e._state = "readable";
    e._reader = void 0;
    e._storedError = void 0;
    e._disturbed = false;
  }
  function Er(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_readableStreamController") &&
      e instanceof ReadableStream
    );
  }
  function Wr(e) {
    return e._reader !== void 0;
  }
  function Br(t, r) {
    t._disturbed = true;
    if (t._state === "closed") {
      return c(void 0);
    }
    if (t._state === "errored") {
      return d(t._storedError);
    }
    Or(t);
    const o = t._reader;
    if (o !== void 0 && at(o)) {
      const e = o._readIntoRequests;
      o._readIntoRequests = new v();
      e.forEach((e) => {
        e._closeSteps(void 0);
      });
    }
    return _(t._readableStreamController[T](r), e);
  }
  function Or(e) {
    e._state = "closed";
    const t = e._reader;
    if (t !== void 0) {
      A(t);
      if (J(t)) {
        const e = t._readRequests;
        t._readRequests = new v();
        e.forEach((e) => {
          e._closeSteps();
        });
      }
    }
  }
  function jr(e, t) {
    e._state = "errored";
    e._storedError = t;
    const r = e._reader;
    if (r !== void 0) {
      k(r, t);
      if (J(r)) {
        Z(r, t);
      } else {
        lt(r, t);
      }
    }
  }
  function kr(e) {
    return new TypeError(
      `ReadableStream.prototype.${e} can only be used on a ReadableStream`
    );
  }
  function Ar(e, t) {
    L(e, t);
    const r = e == null ? void 0 : e.highWaterMark;
    M(r, "highWaterMark", "QueuingStrategyInit");
    return { highWaterMark: Y(r) };
  }
  function Dr(e) {
    return new TypeError(
      `ByteLengthQueuingStrategy.prototype.${e} can only be used on a ByteLengthQueuingStrategy`
    );
  }
  function Lr(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(
        e,
        "_byteLengthQueuingStrategyHighWaterMark"
      ) &&
      e instanceof ByteLengthQueuingStrategy
    );
  }
  function Ir(e) {
    return new TypeError(
      `CountQueuingStrategy.prototype.${e} can only be used on a CountQueuingStrategy`
    );
  }
  function $r(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(
        e,
        "_countQueuingStrategyHighWaterMark"
      ) &&
      e instanceof CountQueuingStrategy
    );
  }
  function Mr(e, t, r) {
    F(e, r);
    return (r) => g(e, t, [r]);
  }
  function Yr(e, t, r) {
    F(e, r);
    return (r) => S(e, t, [r]);
  }
  function Qr(e, t, r) {
    F(e, r);
    return (r, o) => g(e, t, [r, o]);
  }
  function xr(e, t, r) {
    F(e, r);
    return (r) => g(e, t, [r]);
  }
  function Nr(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_transformStreamController") &&
      e instanceof TransformStream
    );
  }
  function Hr(e, t) {
    br(e._readable._readableStreamController, t);
    Vr(e, t);
  }
  function Vr(e, t) {
    Jr(e._transformStreamController);
    Yt(e._writable._writableStreamController, t);
    Ur(e);
  }
  function Ur(e) {
    if (e._backpressure) {
      Gr(e, false);
    }
  }
  function Gr(e, t) {
    if (e._backpressureChangePromise !== void 0) {
      e._backpressureChangePromise_resolve();
    }
    e._backpressureChangePromise = u((t) => {
      e._backpressureChangePromise_resolve = t;
    });
    e._backpressure = t;
  }
  function Xr(e) {
    return (
      !!t(e) &&
      !!Object.prototype.hasOwnProperty.call(e, "_controlledTransformStream") &&
      e instanceof TransformStreamDefaultController
    );
  }
  function Jr(e) {
    e._transformAlgorithm = void 0;
    e._flushAlgorithm = void 0;
    e._cancelAlgorithm = void 0;
  }
  function Kr(e, t) {
    const r = e._controlledTransformStream;
    const o = r._readable._readableStreamController;
    if (!mr(o)) {
      throw new TypeError(
        "Readable side is not in a state that permits enqueue"
      );
    }
    try {
      fr(o, t);
    } catch (e) {
      Vr(r, e);
      throw r._readable._storedError;
    }
    const n = (function () {
      var e = o;
      return !ur(e);
    })();
    if (n !== r._backpressure) {
      Gr(r, true);
    }
  }
  function Zr(e, t) {
    return _(e._transformAlgorithm(t), void 0, (t) => {
      Hr(e._controlledTransformStream, t);
      throw t;
    });
  }
  function eo(e) {
    return new TypeError(
      `TransformStreamDefaultController.prototype.${e} can only be used on a TransformStreamDefaultController`
    );
  }
  function to(e) {
    if (e._finishPromise_resolve !== void 0) {
      e._finishPromise_resolve();
      e._finishPromise_resolve = void 0;
      e._finishPromise_reject = void 0;
    }
  }
  function ro(e, t) {
    if (e._finishPromise_reject !== void 0) {
      p(e._finishPromise);
      e._finishPromise_reject(t);
      e._finishPromise_resolve = void 0;
      e._finishPromise_reject = void 0;
    }
  }
  function oo(e) {
    return new TypeError(
      `TransformStream.prototype.${e} can only be used on a TransformStream`
    );
  }
  const r = e;
  const n = Promise;
  const a = Promise.resolve.bind(n);
  const i = Promise.prototype.then;
  const l = Promise.reject.bind(n);
  const s = a;
  let y = (e) => {
    if (typeof queueMicrotask == "function") {
      y = queueMicrotask;
    } else {
      const e = c(void 0);
      y = (t) => f(e, t);
    }
    return y(e);
  };
  class v {
    constructor() {
      this._cursor = 0;
      this._size = 0;
      this._front = { _elements: [], _next: void 0 };
      this._back = this._front;
      this._cursor = 0;
      this._size = 0;
    }
    get length() {
      return this._size;
    }
    push(e) {
      const t = this._back;
      let r = t;
      if (t._elements.length === 16383) {
        r = { _elements: [], _next: void 0 };
      }
      t._elements.push(e);
      if (r !== t) {
        this._back = r;
        t._next = r;
      }
      ++this._size;
    }
    shift() {
      const e = this._front;
      let t = e;
      const r = this._cursor;
      let o = r + 1;
      const n = e._elements;
      const a = n[r];
      if (o === 16384) {
        t = e._next;
        o = 0;
      }
      --this._size;
      this._cursor = o;
      if (e !== t) {
        this._front = t;
      }
      n[r] = void 0;
      return a;
    }
    forEach(e) {
      let t = this._cursor;
      let r = this._front;
      let o = r._elements;
      while (
        (t !== o.length || r._next !== void 0) &&
        (t !== o.length ||
          !((r = r._next), (o = r._elements), (t = 0), o.length === 0))
      ) {
        e(o[t]);
        ++t;
      }
    }
    peek() {
      const e = this._front;
      const t = this._cursor;
      return e._elements[t];
    }
  }
  const w = Symbol("[[AbortSteps]]");
  const R = Symbol("[[ErrorSteps]]");
  const T = Symbol("[[CancelSteps]]");
  const C = Symbol("[[PullSteps]]");
  const P = Symbol("[[ReleaseSteps]]");
  const z =
    Number.isFinite ||
    function (e) {
      return typeof e == "number" && isFinite(e);
    };
  const D =
    Math.trunc ||
    function (e) {
      if (e < 0) {
        return Math.ceil(e);
      } else {
        return Math.floor(e);
      }
    };
  class ReadableStreamDefaultReader {
    constructor(e) {
      $(e, 1, "ReadableStreamDefaultReader");
      N(e, "First parameter");
      if (Wr(e)) {
        throw new TypeError(
          "This stream has already been locked for exclusive reading by another reader"
        );
      }
      q(this, e);
      this._readRequests = new v();
    }
    get closed() {
      if (J(this)) {
        return this._closedPromise;
      } else {
        return d(ee("closed"));
      }
    }
    cancel(e = void 0) {
      if (J(this)) {
        if (this._ownerReadableStream === void 0) {
          return d(B("cancel"));
        } else {
          return E(this, e);
        }
      } else {
        return d(ee("cancel"));
      }
    }
    read() {
      if (!J(this)) {
        return d(ee("read"));
      }
      if (this._ownerReadableStream === void 0) {
        return d(B("read from"));
      }
      let e;
      let t;
      const r = u((r, o) => {
        e = r;
        t = o;
      });
      K(this, {
        _chunkSteps: (t) => e({ value: t, done: false }),
        _closeSteps: () => e({ value: void 0, done: true }),
        _errorSteps: (e) => t(e),
      });
      return r;
    }
    releaseLock() {
      if (!J(this)) {
        throw ee("releaseLock");
      }
      if (this._ownerReadableStream !== void 0) {
        (function (e) {
          W(e);
          const t = new TypeError("Reader was released");
          Z(e, t);
        })(this);
      }
    }
  }
  var te;
  var re;
  var oe;
  Object.defineProperties(ReadableStreamDefaultReader.prototype, {
    cancel: { enumerable: true },
    read: { enumerable: true },
    releaseLock: { enumerable: true },
    closed: { enumerable: true },
  });
  o(ReadableStreamDefaultReader.prototype.cancel, "cancel");
  o(ReadableStreamDefaultReader.prototype.read, "read");
  o(ReadableStreamDefaultReader.prototype.releaseLock, "releaseLock");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ReadableStreamDefaultReader.prototype,
      Symbol.toStringTag,
      { value: "ReadableStreamDefaultReader", configurable: true }
    );
  }
  let ie = (e) => (
    (ie =
      typeof e.transfer == "function"
        ? (e) => e.transfer()
        : typeof structuredClone == "function"
        ? (e) => structuredClone(e, { transfer: [e] })
        : (e) => e),
    ie(e)
  );
  let le = (e) => (
    (le =
      typeof e.detached == "boolean"
        ? (e) => e.detached
        : (e) => e.byteLength === 0),
    le(e)
  );
  const de =
    (oe =
      (te = Symbol.asyncIterator) !== null && te !== void 0
        ? te
        : (re = Symbol.for) === null || re === void 0
        ? void 0
        : re.call(Symbol, "Symbol.asyncIterator")) !== null && oe !== void 0
      ? oe
      : "@@asyncIterator";
  class he {
    constructor(e, t) {
      this._ongoingPromise = void 0;
      this._isFinished = false;
      this._reader = e;
      this._preventCancel = t;
    }
    next() {
      const e = () => this._nextSteps();
      this._ongoingPromise = this._ongoingPromise
        ? _(this._ongoingPromise, e, e)
        : e();
      return this._ongoingPromise;
    }
    return(e) {
      const t = () => this._returnSteps(e);
      this._ongoingPromise = this._ongoingPromise
        ? _(this._ongoingPromise, t, t)
        : t();
      return this._ongoingPromise;
    }
    _nextSteps() {
      if (this._isFinished) {
        return Promise.resolve({ value: void 0, done: true });
      }
      const e = this._reader;
      let t;
      let r;
      const o = u((e, o) => {
        t = e;
        r = o;
      });
      K(e, {
        _chunkSteps: (e) => {
          this._ongoingPromise = void 0;
          y(() => t({ value: e, done: false }));
        },
        _closeSteps: () => {
          this._ongoingPromise = void 0;
          this._isFinished = true;
          W(e);
          t({ value: void 0, done: true });
        },
        _errorSteps: (t) => {
          this._ongoingPromise = void 0;
          this._isFinished = true;
          W(e);
          r(t);
        },
      });
      return o;
    }
    _returnSteps(e) {
      if (this._isFinished) {
        return Promise.resolve({ value: e, done: true });
      }
      this._isFinished = true;
      const t = this._reader;
      if (!this._preventCancel) {
        const r = E(t, e);
        W(t);
        return _(r, () => ({ value: e, done: true }));
      }
      W(t);
      return c({ value: e, done: true });
    }
  }
  const me = {
    next() {
      if (_e(this)) {
        return this._asyncIteratorImpl.next();
      } else {
        return d(pe("next"));
      }
    },
    return(e) {
      if (_e(this)) {
        return this._asyncIteratorImpl.return(e);
      } else {
        return d(pe("return"));
      }
    },
    [de]() {
      return this;
    },
  };
  Object.defineProperty(me, de, { enumerable: false });
  const ye =
    Number.isNaN ||
    function (e) {
      return e != e;
    };
  class ReadableStreamBYOBRequest {
    constructor() {
      throw new TypeError("Illegal constructor");
    }
    get view() {
      if (!Ce(this)) {
        throw Ke("view");
      }
      return this._view;
    }
    respond(e) {
      if (!Ce(this)) {
        throw Ke("respond");
      }
      $(e, 1, "respond");
      e = x(e, "First parameter");
      if (this._associatedReadableByteStreamController === void 0) {
        throw new TypeError("This BYOB request has been invalidated");
      }
      if (le(this._view.buffer)) {
        throw new TypeError(
          "The BYOB request's buffer has been detached and so cannot be used as a response"
        );
      }
      Ge(this._associatedReadableByteStreamController, e);
    }
    respondWithNewView(e) {
      if (!Ce(this)) {
        throw Ke("respondWithNewView");
      }
      $(e, 1, "respondWithNewView");
      if (!ArrayBuffer.isView(e)) {
        throw new TypeError("You can only respond with array buffer views");
      }
      if (this._associatedReadableByteStreamController === void 0) {
        throw new TypeError("This BYOB request has been invalidated");
      }
      if (le(e.buffer)) {
        throw new TypeError(
          "The given view's buffer has been detached and so cannot be used as a response"
        );
      }
      Xe(this._associatedReadableByteStreamController, e);
    }
  }
  Object.defineProperties(ReadableStreamBYOBRequest.prototype, {
    respond: { enumerable: true },
    respondWithNewView: { enumerable: true },
    view: { enumerable: true },
  });
  o(ReadableStreamBYOBRequest.prototype.respond, "respond");
  o(
    ReadableStreamBYOBRequest.prototype.respondWithNewView,
    "respondWithNewView"
  );
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ReadableStreamBYOBRequest.prototype,
      Symbol.toStringTag,
      { value: "ReadableStreamBYOBRequest", configurable: true }
    );
  }
  class ReadableByteStreamController {
    constructor() {
      throw new TypeError("Illegal constructor");
    }
    get byobRequest() {
      if (!Te(this)) {
        throw Ze("byobRequest");
      }
      return Ve(this);
    }
    get desiredSize() {
      if (!Te(this)) {
        throw Ze("desiredSize");
      }
      return Ue(this);
    }
    close() {
      if (!Te(this)) {
        throw Ze("close");
      }
      if (this._closeRequested) {
        throw new TypeError(
          "The stream has already been closed; do not close it again!"
        );
      }
      const e = this._controlledReadableByteStream._state;
      if (e !== "readable") {
        throw new TypeError(
          `The stream (in ${e} state) is not in the readable state and cannot be closed`
        );
      }
      Qe(this);
    }
    enqueue(e) {
      if (!Te(this)) {
        throw Ze("enqueue");
      }
      $(e, 1, "enqueue");
      if (!ArrayBuffer.isView(e)) {
        throw new TypeError("chunk must be an array buffer view");
      }
      if (e.byteLength === 0) {
        throw new TypeError("chunk must have non-zero byteLength");
      }
      if (e.buffer.byteLength === 0) {
        throw new TypeError("chunk's buffer must have non-zero byteLength");
      }
      if (this._closeRequested) {
        throw new TypeError("stream is closed or draining");
      }
      const t = this._controlledReadableByteStream._state;
      if (t !== "readable") {
        throw new TypeError(
          `The stream (in ${t} state) is not in the readable state and cannot be enqueued to`
        );
      }
      xe(this, e);
    }
    error(e = void 0) {
      if (!Te(this)) {
        throw Ze("error");
      }
      Ne(this, e);
    }
    [T](e) {
      qe(this);
      we(this);
      const t = this._cancelAlgorithm(e);
      Ye(this);
      return t;
    }
    [C](e) {
      const t = this._controlledReadableByteStream;
      if (this._queueTotalSize > 0) {
        He(this, e);
        return;
      }
      const r = this._autoAllocateChunkSize;
      if (r !== void 0) {
        let t;
        try {
          t = new ArrayBuffer(r);
        } catch (t) {
          e._errorSteps(t);
          return;
        }
        const o = {
          buffer: t,
          bufferByteLength: r,
          byteOffset: 0,
          byteLength: r,
          bytesFilled: 0,
          minimumFill: 1,
          elementSize: 1,
          viewConstructor: Uint8Array,
          readerType: "default",
        };
        this._pendingPullIntos.push(o);
      }
      V(t, e);
      Pe(this);
    }
    [P]() {
      if (this._pendingPullIntos.length > 0) {
        const e = this._pendingPullIntos.peek();
        e.readerType = "none";
        this._pendingPullIntos = new v();
        this._pendingPullIntos.push(e);
      }
    }
  }
  Object.defineProperties(ReadableByteStreamController.prototype, {
    close: { enumerable: true },
    enqueue: { enumerable: true },
    error: { enumerable: true },
    byobRequest: { enumerable: true },
    desiredSize: { enumerable: true },
  });
  o(ReadableByteStreamController.prototype.close, "close");
  o(ReadableByteStreamController.prototype.enqueue, "enqueue");
  o(ReadableByteStreamController.prototype.error, "error");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ReadableByteStreamController.prototype,
      Symbol.toStringTag,
      { value: "ReadableByteStreamController", configurable: true }
    );
  }
  class ReadableStreamBYOBReader {
    constructor(e) {
      $(e, 1, "ReadableStreamBYOBReader");
      N(e, "First parameter");
      if (Wr(e)) {
        throw new TypeError(
          "This stream has already been locked for exclusive reading by another reader"
        );
      }
      if (!Te(e._readableStreamController)) {
        throw new TypeError(
          "Cannot construct a ReadableStreamBYOBReader for a stream not constructed with a byte source"
        );
      }
      q(this, e);
      this._readIntoRequests = new v();
    }
    get closed() {
      if (at(this)) {
        return this._closedPromise;
      } else {
        return d(st("closed"));
      }
    }
    cancel(e = void 0) {
      if (at(this)) {
        if (this._ownerReadableStream === void 0) {
          return d(B("cancel"));
        } else {
          return E(this, e);
        }
      } else {
        return d(st("cancel"));
      }
    }
    read(e, t = {}) {
      if (!at(this)) {
        return d(st("read"));
      }
      if (!ArrayBuffer.isView(e)) {
        return d(new TypeError("view must be an array buffer view"));
      }
      if (e.byteLength === 0) {
        return d(new TypeError("view must have non-zero byteLength"));
      }
      if (e.buffer.byteLength === 0) {
        return d(new TypeError("view's buffer must have non-zero byteLength"));
      }
      if (le(e.buffer)) {
        return d(new TypeError("view's buffer has been detached"));
      }
      let r;
      try {
        r = (function (e, t) {
          var r;
          L(e, t);
          return {
            min: x(
              (r = e == null ? void 0 : e.min) !== null && r !== void 0 ? r : 1,
              `${t} has member 'min' that`
            ),
          };
        })(t, "options");
      } catch (e) {
        return d(e);
      }
      const o = r.min;
      if (o === 0) {
        return d(new TypeError("options.min must be greater than 0"));
      }
      if (
        (function (e) {
          return Re(e.constructor);
        })(e)
      ) {
        if (o > e.byteLength) {
          return d(
            new RangeError(
              "options.min must be less than or equal to view's byteLength"
            )
          );
        }
      } else if (o > e.length) {
        return d(
          new RangeError(
            "options.min must be less than or equal to view's length"
          )
        );
      }
      if (this._ownerReadableStream === void 0) {
        return d(B("read from"));
      }
      let n;
      let a;
      const i = u((e, t) => {
        n = e;
        a = t;
      });
      it(this, e, o, {
        _chunkSteps: (e) => n({ value: e, done: false }),
        _closeSteps: (e) => n({ value: e, done: true }),
        _errorSteps: (e) => a(e),
      });
      return i;
    }
    releaseLock() {
      if (!at(this)) {
        throw st("releaseLock");
      }
      if (this._ownerReadableStream !== void 0) {
        (function (e) {
          W(e);
          const t = new TypeError("Reader was released");
          lt(e, t);
        })(this);
      }
    }
  }
  Object.defineProperties(ReadableStreamBYOBReader.prototype, {
    cancel: { enumerable: true },
    read: { enumerable: true },
    releaseLock: { enumerable: true },
    closed: { enumerable: true },
  });
  o(ReadableStreamBYOBReader.prototype.cancel, "cancel");
  o(ReadableStreamBYOBReader.prototype.read, "read");
  o(ReadableStreamBYOBReader.prototype.releaseLock, "releaseLock");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ReadableStreamBYOBReader.prototype,
      Symbol.toStringTag,
      { value: "ReadableStreamBYOBReader", configurable: true }
    );
  }
  class WritableStream {
    constructor(e = {}, t = {}) {
      if (e === void 0) {
        e = null;
      } else {
        I(e, "First parameter");
      }
      const r = dt(t, "Second parameter");
      const o = (function (e, t) {
        L(e, t);
        const r = e == null ? void 0 : e.abort;
        const o = e == null ? void 0 : e.close;
        const n = e == null ? void 0 : e.start;
        const a = e == null ? void 0 : e.type;
        const i = e == null ? void 0 : e.write;
        return {
          abort:
            r === void 0 ? void 0 : bt(r, e, `${t} has member 'abort' that`),
          close:
            o === void 0 ? void 0 : ht(o, e, `${t} has member 'close' that`),
          start:
            n === void 0 ? void 0 : mt(n, e, `${t} has member 'start' that`),
          write:
            i === void 0 ? void 0 : _t(i, e, `${t} has member 'write' that`),
          type: a,
        };
      })(e, "First parameter");
      St(this);
      if (o.type !== void 0) {
        throw new RangeError("Invalid type is specified");
      }
      const n = ct(r);
      (function (e, t, r, o) {
        const n = Object.create(WritableStreamDefaultController.prototype);
        let a;
        let i;
        let l;
        let s;
        a = t.start !== void 0 ? () => t.start(n) : () => {};
        i = t.write !== void 0 ? (e) => t.write(e, n) : () => c(void 0);
        l = t.close !== void 0 ? () => t.close() : () => c(void 0);
        s = t.abort !== void 0 ? (e) => t.abort(e) : () => c(void 0);
        Ft(e, n, a, i, l, s, r, o);
      })(this, o, ut(r, 1), n);
    }
    get locked() {
      if (!gt(this)) {
        throw Nt("locked");
      }
      return vt(this);
    }
    abort(e = void 0) {
      if (gt(this)) {
        if (vt(this)) {
          return d(
            new TypeError("Cannot abort a stream that already has a writer")
          );
        } else {
          return wt(this, e);
        }
      } else {
        return d(Nt("abort"));
      }
    }
    close() {
      if (gt(this)) {
        if (vt(this)) {
          return d(
            new TypeError("Cannot close a stream that already has a writer")
          );
        } else if (qt(this)) {
          return d(new TypeError("Cannot close an already-closing stream"));
        } else {
          return Rt(this);
        }
      } else {
        return d(Nt("close"));
      }
    }
    getWriter() {
      if (!gt(this)) {
        throw Nt("getWriter");
      }
      return yt(this);
    }
  }
  Object.defineProperties(WritableStream.prototype, {
    abort: { enumerable: true },
    close: { enumerable: true },
    getWriter: { enumerable: true },
    locked: { enumerable: true },
  });
  o(WritableStream.prototype.abort, "abort");
  o(WritableStream.prototype.close, "close");
  o(WritableStream.prototype.getWriter, "getWriter");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(WritableStream.prototype, Symbol.toStringTag, {
      value: "WritableStream",
      configurable: true,
    });
  }
  class WritableStreamDefaultWriter {
    constructor(e) {
      $(e, 1, "WritableStreamDefaultWriter");
      pt(e, "First parameter");
      if (vt(e)) {
        throw new TypeError(
          "This stream has already been locked for exclusive writing by another writer"
        );
      }
      this._ownerWritableStream = e;
      e._writer = this;
      const t = e._state;
      if (t === "writable") {
        if (!qt(e) && e._backpressure) {
          Zt(this);
        } else {
          tr(this);
        }
        Gt(this);
      } else if (t === "erroring") {
        er(this, e._storedError);
        Gt(this);
      } else if (t === "closed") {
        tr(this);
        Gt((r = this));
        Kt(r);
      } else {
        const t = e._storedError;
        er(this, t);
        Xt(this, t);
      }
      var r;
    }
    get closed() {
      if (Bt(this)) {
        return this._closedPromise;
      } else {
        return d(Vt("closed"));
      }
    }
    get desiredSize() {
      if (!Bt(this)) {
        throw Vt("desiredSize");
      }
      if (this._ownerWritableStream === void 0) {
        throw Ut("desiredSize");
      }
      return (function (e) {
        const t = e._ownerWritableStream;
        const r = t._state;
        if (r === "errored" || r === "erroring") {
          return null;
        }
        if (r === "closed") {
          return 0;
        }
        return $t(t._writableStreamController);
      })(this);
    }
    get ready() {
      if (Bt(this)) {
        return this._readyPromise;
      } else {
        return d(Vt("ready"));
      }
    }
    abort(e = void 0) {
      if (Bt(this)) {
        if (this._ownerWritableStream === void 0) {
          return d(Ut("abort"));
        } else {
          return (function (e, t) {
            return wt(e._ownerWritableStream, t);
          })(this, e);
        }
      } else {
        return d(Vt("abort"));
      }
    }
    close() {
      if (!Bt(this)) {
        return d(Vt("close"));
      }
      const e = this._ownerWritableStream;
      if (e === void 0) {
        return d(Ut("close"));
      } else if (qt(e)) {
        return d(new TypeError("Cannot close an already-closing stream"));
      } else {
        return Ot(this);
      }
    }
    releaseLock() {
      if (!Bt(this)) {
        throw Vt("releaseLock");
      }
      if (this._ownerWritableStream !== void 0) {
        At(this);
      }
    }
    write(e = void 0) {
      if (Bt(this)) {
        if (this._ownerWritableStream === void 0) {
          return d(Ut("write to"));
        } else {
          return zt(this, e);
        }
      } else {
        return d(Vt("write"));
      }
    }
  }
  Object.defineProperties(WritableStreamDefaultWriter.prototype, {
    abort: { enumerable: true },
    close: { enumerable: true },
    releaseLock: { enumerable: true },
    write: { enumerable: true },
    closed: { enumerable: true },
    desiredSize: { enumerable: true },
    ready: { enumerable: true },
  });
  o(WritableStreamDefaultWriter.prototype.abort, "abort");
  o(WritableStreamDefaultWriter.prototype.close, "close");
  o(WritableStreamDefaultWriter.prototype.releaseLock, "releaseLock");
  o(WritableStreamDefaultWriter.prototype.write, "write");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      WritableStreamDefaultWriter.prototype,
      Symbol.toStringTag,
      { value: "WritableStreamDefaultWriter", configurable: true }
    );
  }
  const Dt = {};
  class WritableStreamDefaultController {
    constructor() {
      throw new TypeError("Illegal constructor");
    }
    get abortReason() {
      if (!Lt(this)) {
        throw Ht("abortReason");
      }
      return this._abortReason;
    }
    get signal() {
      if (!Lt(this)) {
        throw Ht("signal");
      }
      if (this._abortController === void 0) {
        throw new TypeError(
          "WritableStreamDefaultController.prototype.signal is not supported"
        );
      }
      return this._abortController.signal;
    }
    error(e = void 0) {
      if (!Lt(this)) {
        throw Ht("error");
      }
      if (this._controlledWritableStream._state === "writable") {
        xt(this, e);
      }
    }
    [w](e) {
      const t = this._abortAlgorithm(e);
      It(this);
      return t;
    }
    [R]() {
      we(this);
    }
  }
  Object.defineProperties(WritableStreamDefaultController.prototype, {
    abortReason: { enumerable: true },
    signal: { enumerable: true },
    error: { enumerable: true },
  });
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      WritableStreamDefaultController.prototype,
      Symbol.toStringTag,
      { value: "WritableStreamDefaultController", configurable: true }
    );
  }
  const nr =
    typeof globalThis != "undefined"
      ? globalThis
      : typeof self != "undefined"
      ? self
      : typeof global != "undefined"
      ? global
      : void 0;
  const ar =
    (function () {
      const e = nr == null ? void 0 : nr.DOMException;
      if (
        (function (e) {
          if (typeof e != "function" && typeof e != "object") {
            return false;
          }
          if (e.name !== "DOMException") {
            return false;
          }
          try {
            new e();
            return true;
          } catch (e) {
            return false;
          }
        })(e)
      ) {
        return e;
      } else {
        return;
      }
    })() ||
    (function () {
      const e = function (e, t) {
        this.message = e || "";
        this.name = t || "Error";
        if (Error.captureStackTrace) {
          Error.captureStackTrace(this, this.constructor);
        }
      };
      o(e, "DOMException");
      e.prototype = Object.create(Error.prototype);
      Object.defineProperty(e.prototype, "constructor", {
        value: e,
        writable: true,
        configurable: true,
      });
      return e;
    })();
  class ReadableStreamDefaultController {
    constructor() {
      throw new TypeError("Illegal constructor");
    }
    get desiredSize() {
      if (!lr(this)) {
        throw pr("desiredSize");
      }
      return hr(this);
    }
    close() {
      if (!lr(this)) {
        throw pr("close");
      }
      if (!mr(this)) {
        throw new TypeError("The stream is not in a state that permits close");
      }
      dr(this);
    }
    enqueue(e = void 0) {
      if (!lr(this)) {
        throw pr("enqueue");
      }
      if (!mr(this)) {
        throw new TypeError(
          "The stream is not in a state that permits enqueue"
        );
      }
      return fr(this, e);
    }
    error(e = void 0) {
      if (!lr(this)) {
        throw pr("error");
      }
      br(this, e);
    }
    [T](e) {
      we(this);
      const t = this._cancelAlgorithm(e);
      cr(this);
      return t;
    }
    [C](e) {
      const t = this._controlledReadableStream;
      if (this._queue.length > 0) {
        const r = ge(this);
        if (this._closeRequested && this._queue.length === 0) {
          cr(this);
          Or(t);
        } else {
          sr(this);
        }
        e._chunkSteps(r);
      } else {
        V(t, e);
        sr(this);
      }
    }
    [P]() {}
  }
  Object.defineProperties(ReadableStreamDefaultController.prototype, {
    close: { enumerable: true },
    enqueue: { enumerable: true },
    error: { enumerable: true },
    desiredSize: { enumerable: true },
  });
  o(ReadableStreamDefaultController.prototype.close, "close");
  o(ReadableStreamDefaultController.prototype.enqueue, "enqueue");
  o(ReadableStreamDefaultController.prototype.error, "error");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ReadableStreamDefaultController.prototype,
      Symbol.toStringTag,
      { value: "ReadableStreamDefaultController", configurable: true }
    );
  }
  class ReadableStream {
    constructor(e = {}, t = {}) {
      if (e === void 0) {
        e = null;
      } else {
        I(e, "First parameter");
      }
      const r = dt(t, "Second parameter");
      const o = (function (e, t) {
        L(e, t);
        const r = e;
        const o = r == null ? void 0 : r.autoAllocateChunkSize;
        const n = r == null ? void 0 : r.cancel;
        const a = r == null ? void 0 : r.pull;
        const i = r == null ? void 0 : r.start;
        const l = r == null ? void 0 : r.type;
        return {
          autoAllocateChunkSize:
            o === void 0
              ? void 0
              : x(o, `${t} has member 'autoAllocateChunkSize' that`),
          cancel:
            n === void 0 ? void 0 : gr(n, r, `${t} has member 'cancel' that`),
          pull: a === void 0 ? void 0 : vr(a, r, `${t} has member 'pull' that`),
          start:
            i === void 0 ? void 0 : wr(i, r, `${t} has member 'start' that`),
          type: l === void 0 ? void 0 : Rr(l, `${t} has member 'type' that`),
        };
      })(e, "First parameter");
      qr(this);
      if (o.type === "bytes") {
        if (r.size !== void 0) {
          throw new RangeError(
            "The strategy for a byte stream cannot have a size function"
          );
        }
        (function (e, t, r) {
          const o = Object.create(ReadableByteStreamController.prototype);
          let n;
          let a;
          let i;
          n = t.start !== void 0 ? () => t.start(o) : () => {};
          a = t.pull !== void 0 ? () => t.pull(o) : () => c(void 0);
          i = t.cancel !== void 0 ? (e) => t.cancel(e) : () => c(void 0);
          const l = t.autoAllocateChunkSize;
          if (l === 0) {
            throw new TypeError("autoAllocateChunkSize must be greater than 0");
          }
          Je(e, o, n, a, i, r, l);
        })(this, o, ut(r, 0));
      } else {
        const e = ct(r);
        (function (e, t, r, o) {
          const n = Object.create(ReadableStreamDefaultController.prototype);
          let a;
          let i;
          let l;
          a = t.start !== void 0 ? () => t.start(n) : () => {};
          i = t.pull !== void 0 ? () => t.pull(n) : () => c(void 0);
          l = t.cancel !== void 0 ? (e) => t.cancel(e) : () => c(void 0);
          _r(e, n, a, i, l, r, o);
        })(this, o, ut(r, 1), e);
      }
    }
    get locked() {
      if (!Er(this)) {
        throw kr("locked");
      }
      return Wr(this);
    }
    cancel(e = void 0) {
      if (Er(this)) {
        if (Wr(this)) {
          return d(
            new TypeError("Cannot cancel a stream that already has a reader")
          );
        } else {
          return Br(this, e);
        }
      } else {
        return d(kr("cancel"));
      }
    }
    getReader(e = void 0) {
      if (!Er(this)) {
        throw kr("getReader");
      }
      if (
        (function (e, t) {
          L(e, t);
          const r = e == null ? void 0 : e.mode;
          return {
            mode: r === void 0 ? void 0 : et(r, `${t} has member 'mode' that`),
          };
        })(e, "First parameter").mode === void 0
      ) {
        return H(this);
      } else {
        return tt(this);
      }
    }
    pipeThrough(e, t = {}) {
      if (!Er(this)) {
        throw kr("pipeThrough");
      }
      $(e, 1, "pipeThrough");
      const r = (function (e, t) {
        L(e, t);
        const r = e == null ? void 0 : e.readable;
        M(r, "readable", "ReadableWritablePair");
        N(r, `${t} has member 'readable' that`);
        const o = e == null ? void 0 : e.writable;
        M(o, "writable", "ReadableWritablePair");
        pt(o, `${t} has member 'writable' that`);
        return { readable: r, writable: o };
      })(e, "First parameter");
      const o = Tr(t, "Second parameter");
      if (Wr(this)) {
        throw new TypeError(
          "ReadableStream.prototype.pipeThrough cannot be used on a locked ReadableStream"
        );
      }
      if (vt(r.writable)) {
        throw new TypeError(
          "ReadableStream.prototype.pipeThrough cannot be used on a locked WritableStream"
        );
      }
      p(
        ir(
          this,
          r.writable,
          o.preventClose,
          o.preventAbort,
          o.preventCancel,
          o.signal
        )
      );
      return r.readable;
    }
    pipeTo(e, t = {}) {
      if (!Er(this)) {
        return d(kr("pipeTo"));
      }
      if (e === void 0) {
        return d("Parameter 1 is required in 'pipeTo'.");
      }
      if (!gt(e)) {
        return d(
          new TypeError(
            "ReadableStream.prototype.pipeTo's first argument must be a WritableStream"
          )
        );
      }
      let r;
      try {
        r = Tr(t, "Second parameter");
      } catch (e) {
        return d(e);
      }
      if (Wr(this)) {
        return d(
          new TypeError(
            "ReadableStream.prototype.pipeTo cannot be used on a locked ReadableStream"
          )
        );
      } else if (vt(e)) {
        return d(
          new TypeError(
            "ReadableStream.prototype.pipeTo cannot be used on a locked WritableStream"
          )
        );
      } else {
        return ir(
          this,
          e,
          r.preventClose,
          r.preventAbort,
          r.preventCancel,
          r.signal
        );
      }
    }
    tee() {
      if (!Er(this)) {
        throw kr("tee");
      }
      return ne(yr(this));
    }
    values(e = void 0) {
      if (!Er(this)) {
        throw kr("values");
      }
      return (function (e, t) {
        const r = H(e);
        const o = new he(r, t);
        const n = Object.create(me);
        n._asyncIteratorImpl = o;
        return n;
      })(
        this,
        (function (e, t) {
          L(e, t);
          const r = e == null ? void 0 : e.preventCancel;
          return { preventCancel: Boolean(r) };
        })(e, "First parameter").preventCancel
      );
    }
    [de](e) {
      return this.values(e);
    }
    static from(e) {
      return Sr(e);
    }
  }
  Object.defineProperties(ReadableStream, { from: { enumerable: true } });
  Object.defineProperties(ReadableStream.prototype, {
    cancel: { enumerable: true },
    getReader: { enumerable: true },
    pipeThrough: { enumerable: true },
    pipeTo: { enumerable: true },
    tee: { enumerable: true },
    values: { enumerable: true },
    locked: { enumerable: true },
  });
  o(ReadableStream.from, "from");
  o(ReadableStream.prototype.cancel, "cancel");
  o(ReadableStream.prototype.getReader, "getReader");
  o(ReadableStream.prototype.pipeThrough, "pipeThrough");
  o(ReadableStream.prototype.pipeTo, "pipeTo");
  o(ReadableStream.prototype.tee, "tee");
  o(ReadableStream.prototype.values, "values");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(ReadableStream.prototype, Symbol.toStringTag, {
      value: "ReadableStream",
      configurable: true,
    });
  }
  Object.defineProperty(ReadableStream.prototype, de, {
    value: ReadableStream.prototype.values,
    writable: true,
    configurable: true,
  });
  const zr = (e) => e.byteLength;
  o(zr, "size");
  class ByteLengthQueuingStrategy {
    constructor(e) {
      $(e, 1, "ByteLengthQueuingStrategy");
      e = Ar(e, "First parameter");
      this._byteLengthQueuingStrategyHighWaterMark = e.highWaterMark;
    }
    get highWaterMark() {
      if (!Lr(this)) {
        throw Dr("highWaterMark");
      }
      return this._byteLengthQueuingStrategyHighWaterMark;
    }
    get size() {
      if (!Lr(this)) {
        throw Dr("size");
      }
      return zr;
    }
  }
  Object.defineProperties(ByteLengthQueuingStrategy.prototype, {
    highWaterMark: { enumerable: true },
    size: { enumerable: true },
  });
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      ByteLengthQueuingStrategy.prototype,
      Symbol.toStringTag,
      { value: "ByteLengthQueuingStrategy", configurable: true }
    );
  }
  const Fr = () => 1;
  o(Fr, "size");
  class CountQueuingStrategy {
    constructor(e) {
      $(e, 1, "CountQueuingStrategy");
      e = Ar(e, "First parameter");
      this._countQueuingStrategyHighWaterMark = e.highWaterMark;
    }
    get highWaterMark() {
      if (!$r(this)) {
        throw Ir("highWaterMark");
      }
      return this._countQueuingStrategyHighWaterMark;
    }
    get size() {
      if (!$r(this)) {
        throw Ir("size");
      }
      return Fr;
    }
  }
  Object.defineProperties(CountQueuingStrategy.prototype, {
    highWaterMark: { enumerable: true },
    size: { enumerable: true },
  });
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(CountQueuingStrategy.prototype, Symbol.toStringTag, {
      value: "CountQueuingStrategy",
      configurable: true,
    });
  }
  class TransformStream {
    constructor(e = {}, t = {}, r = {}) {
      if (e === void 0) {
        e = null;
      }
      const o = dt(t, "Second parameter");
      const n = dt(r, "Third parameter");
      const a = (function (e, t) {
        L(e, t);
        const r = e == null ? void 0 : e.cancel;
        const o = e == null ? void 0 : e.flush;
        const n = e == null ? void 0 : e.readableType;
        const a = e == null ? void 0 : e.start;
        const i = e == null ? void 0 : e.transform;
        const l = e == null ? void 0 : e.writableType;
        return {
          cancel:
            r === void 0 ? void 0 : xr(r, e, `${t} has member 'cancel' that`),
          flush:
            o === void 0 ? void 0 : Mr(o, e, `${t} has member 'flush' that`),
          readableType: n,
          start:
            a === void 0 ? void 0 : Yr(a, e, `${t} has member 'start' that`),
          transform:
            i === void 0
              ? void 0
              : Qr(i, e, `${t} has member 'transform' that`),
          writableType: l,
        };
      })(e, "First parameter");
      if (a.readableType !== void 0) {
        throw new RangeError("Invalid readableType specified");
      }
      if (a.writableType !== void 0) {
        throw new RangeError("Invalid writableType specified");
      }
      const i = ut(n, 0);
      const l = ct(n);
      const s = ut(o, 1);
      const f = ct(o);
      let h;
      (function (e, t, r, o, n, a) {
        function i() {
          return t;
        }
        function l(t) {
          return (function (e, t) {
            const r = e._transformStreamController;
            if (e._backpressure) {
              return _(e._backpressureChangePromise, () => {
                const o = e._writable;
                if (o._state === "erroring") {
                  throw o._storedError;
                }
                return Zr(r, t);
              });
            }
            return Zr(r, t);
          })(e, t);
        }
        function s(t) {
          return (function (e, t) {
            const r = e._transformStreamController;
            if (r._finishPromise !== void 0) {
              return r._finishPromise;
            }
            const o = e._readable;
            r._finishPromise = u((e, t) => {
              r._finishPromise_resolve = e;
              r._finishPromise_reject = t;
            });
            const n = r._cancelAlgorithm(t);
            Jr(r);
            b(
              n,
              () => (
                o._state === "errored"
                  ? ro(r, o._storedError)
                  : (br(o._readableStreamController, t), to(r)),
                null
              ),
              (e) => (br(o._readableStreamController, e), ro(r, e), null)
            );
            return r._finishPromise;
          })(e, t);
        }
        function c() {
          return (function (e) {
            const t = e._transformStreamController;
            if (t._finishPromise !== void 0) {
              return t._finishPromise;
            }
            const r = e._readable;
            t._finishPromise = u((e, r) => {
              t._finishPromise_resolve = e;
              t._finishPromise_reject = r;
            });
            const o = t._flushAlgorithm();
            Jr(t);
            b(
              o,
              () => (
                r._state === "errored"
                  ? ro(t, r._storedError)
                  : (dr(r._readableStreamController), to(t)),
                null
              ),
              (e) => (br(r._readableStreamController, e), ro(t, e), null)
            );
            return t._finishPromise;
          })(e);
        }
        function d() {
          return (function (e) {
            Gr(e, false);
            return e._backpressureChangePromise;
          })(e);
        }
        function f(t) {
          return (function (e, t) {
            const r = e._transformStreamController;
            if (r._finishPromise !== void 0) {
              return r._finishPromise;
            }
            const o = e._writable;
            r._finishPromise = u((e, t) => {
              r._finishPromise_resolve = e;
              r._finishPromise_reject = t;
            });
            const n = r._cancelAlgorithm(t);
            Jr(r);
            b(
              n,
              () => (
                o._state === "errored"
                  ? ro(r, o._storedError)
                  : (Yt(o._writableStreamController, t), Ur(e), to(r)),
                null
              ),
              (t) => (Yt(o._writableStreamController, t), Ur(e), ro(r, t), null)
            );
            return r._finishPromise;
          })(e, t);
        }
        e._writable = (function (e, t, r, o, n = 1, a = () => 1) {
          const i = Object.create(WritableStream.prototype);
          St(i);
          Ft(
            i,
            Object.create(WritableStreamDefaultController.prototype),
            e,
            t,
            r,
            o,
            n,
            a
          );
          return i;
        })(i, l, c, s, r, o);
        e._readable = Cr(i, d, f, n, a);
        e._backpressure = void 0;
        e._backpressureChangePromise = void 0;
        e._backpressureChangePromise_resolve = void 0;
        Gr(e, true);
        e._transformStreamController = void 0;
      })(
        this,
        u((e) => {
          h = e;
        }),
        s,
        f,
        i,
        l
      );
      (function (e, t) {
        const r = Object.create(TransformStreamDefaultController.prototype);
        let o;
        let n;
        let a;
        o =
          t.transform !== void 0
            ? (e) => t.transform(e, r)
            : (e) => {
                try {
                  Kr(r, e);
                  return c(void 0);
                } catch (e) {
                  return d(e);
                }
              };
        n = t.flush !== void 0 ? () => t.flush(r) : () => c(void 0);
        a = t.cancel !== void 0 ? (e) => t.cancel(e) : () => c(void 0);
        (function (e, t, r, o, n) {
          t._controlledTransformStream = e;
          e._transformStreamController = t;
          t._transformAlgorithm = r;
          t._flushAlgorithm = o;
          t._cancelAlgorithm = n;
          t._finishPromise = void 0;
          t._finishPromise_resolve = void 0;
          t._finishPromise_reject = void 0;
        })(e, r, o, n, a);
      })(this, a);
      if (a.start === void 0) {
        h(void 0);
      } else {
        h(a.start(this._transformStreamController));
      }
    }
    get readable() {
      if (!Nr(this)) {
        throw oo("readable");
      }
      return this._readable;
    }
    get writable() {
      if (!Nr(this)) {
        throw oo("writable");
      }
      return this._writable;
    }
  }
  Object.defineProperties(TransformStream.prototype, {
    readable: { enumerable: true },
    writable: { enumerable: true },
  });
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(TransformStream.prototype, Symbol.toStringTag, {
      value: "TransformStream",
      configurable: true,
    });
  }
  class TransformStreamDefaultController {
    constructor() {
      throw new TypeError("Illegal constructor");
    }
    get desiredSize() {
      if (!Xr(this)) {
        throw eo("desiredSize");
      }
      return hr(
        this._controlledTransformStream._readable._readableStreamController
      );
    }
    enqueue(e = void 0) {
      if (!Xr(this)) {
        throw eo("enqueue");
      }
      Kr(this, e);
    }
    error(e = void 0) {
      if (!Xr(this)) {
        throw eo("error");
      }
      var t = e;
      Hr(this._controlledTransformStream, t);
    }
    terminate() {
      if (!Xr(this)) {
        throw eo("terminate");
      }
      (function (e) {
        const t = e._controlledTransformStream;
        dr(t._readable._readableStreamController);
        const r = new TypeError("TransformStream terminated");
        Vr(t, r);
      })(this);
    }
  }
  Object.defineProperties(TransformStreamDefaultController.prototype, {
    enqueue: { enumerable: true },
    error: { enumerable: true },
    terminate: { enumerable: true },
    desiredSize: { enumerable: true },
  });
  o(TransformStreamDefaultController.prototype.enqueue, "enqueue");
  o(TransformStreamDefaultController.prototype.error, "error");
  o(TransformStreamDefaultController.prototype.terminate, "terminate");
  if (typeof Symbol.toStringTag == "symbol") {
    Object.defineProperty(
      TransformStreamDefaultController.prototype,
      Symbol.toStringTag,
      { value: "TransformStreamDefaultController", configurable: true }
    );
  }
  const no = {
    ReadableStream: ReadableStream,
    ReadableStreamDefaultController: ReadableStreamDefaultController,
    ReadableByteStreamController: ReadableByteStreamController,
    ReadableStreamBYOBRequest: ReadableStreamBYOBRequest,
    ReadableStreamDefaultReader: ReadableStreamDefaultReader,
    ReadableStreamBYOBReader: ReadableStreamBYOBReader,
    WritableStream: WritableStream,
    WritableStreamDefaultController: WritableStreamDefaultController,
    WritableStreamDefaultWriter: WritableStreamDefaultWriter,
    ByteLengthQueuingStrategy: ByteLengthQueuingStrategy,
    CountQueuingStrategy: CountQueuingStrategy,
    TransformStream: TransformStream,
    TransformStreamDefaultController: TransformStreamDefaultController,
  };
  console.log("nr: " + String(nr === window), nr);
  for (const e in no) {
    console.log("ITER: " + e);
    if (Object.prototype.hasOwnProperty.call(no, e)) {
      Object.defineProperty(nr, e, {
        value: no[e],
        writable: true,
        configurable: true,
      });
    }
    console.log(`Window has: ${e} - ${[e]}`);
  }
})();
