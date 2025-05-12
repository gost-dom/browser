// Datastar v1.0.0-beta.9

// library/src/engine/consts.ts
var lol = /🖕JS_DS🚀/.source;
var DSP = lol.slice(0, 5);
var DSS = lol.slice(4);
var DATASTAR = "datastar";
var DATASTAR_REQUEST = "Datastar-Request";
var DefaultFragmentsSettleDurationMs = 300;
var DefaultSseRetryDurationMs = 1e3;
var DefaultExecuteScriptAttributes = "type module";
var DefaultFragmentsUseViewTransitions = false;
var DefaultMergeSignalsOnlyIfMissing = false;
var DefaultExecuteScriptAutoRemove = true;
var FragmentMergeModes = {
  // Morphs the fragment into the existing element using idiomorph.
  Morph: "morph",
  // Replaces the inner HTML of the existing element.
  Inner: "inner",
  // Replaces the outer HTML of the existing element.
  Outer: "outer",
  // Prepends the fragment to the existing element.
  Prepend: "prepend",
  // Appends the fragment to the existing element.
  Append: "append",
  // Inserts the fragment before the existing element.
  Before: "before",
  // Inserts the fragment after the existing element.
  After: "after",
  // Upserts the attributes of the existing element.
  UpsertAttributes: "upsertAttributes"
};
var DefaultFragmentMergeMode = FragmentMergeModes.Morph;
var EventTypes = {
  // An event for merging HTML fragments into the DOM.
  MergeFragments: "datastar-merge-fragments",
  // An event for merging signals.
  MergeSignals: "datastar-merge-signals",
  // An event for removing HTML fragments from the DOM.
  RemoveFragments: "datastar-remove-fragments",
  // An event for removing signals.
  RemoveSignals: "datastar-remove-signals",
  // An event for executing <script/> elements in the browser.
  ExecuteScript: "datastar-execute-script"
};

// library/src/engine/types.ts
var PluginType = /* @__PURE__ */ ((PluginType2) => {
  PluginType2[PluginType2["Attribute"] = 1] = "Attribute";
  PluginType2[PluginType2["Watcher"] = 2] = "Watcher";
  PluginType2[PluginType2["Action"] = 3] = "Action";
  return PluginType2;
})(PluginType || {});
var DATASTAR_SIGNAL_EVENT = `${DATASTAR}-signals`;

// library/src/utils/text.ts
var isBoolString = (str) => str.trim() === "true";
var kebab = (str) => str.replace(
  /[A-Z]+(?![a-z])|[A-Z]/g,
  ($, ofs) => (ofs ? "-" : "") + $.toLowerCase()
);
var camel = (str) => kebab(str).replace(/-./g, (x) => x[1].toUpperCase());
var snake = (str) => kebab(str).replace(/-/g, "_");
var pascal = (str) => camel(str).replace(/^./, (x) => x[0].toUpperCase());
var jsStrToObject = (raw) => new Function(`return Object.assign({}, ${raw})`)();
var trimDollarSignPrefix = (str) => str.startsWith("$") ? str.slice(1) : str;
var caseFns = { kebab, snake, pascal };
function modifyCasing(str, mods) {
  for (const c of mods.get("case") || []) {
    const fn = caseFns[c];
    if (fn) str = fn(str);
  }
  return str;
}

// library/src/plugins/official/core/attributes/computed.ts
var name = "computed";
var Computed = {
  type: 1 /* Attribute */,
  name,
  keyReq: 1 /* Must */,
  valReq: 1 /* Must */,
  onLoad: ({ key, mods, signals: signals2, genRX: genRX2 }) => {
    key = modifyCasing(key, mods);
    const rx = genRX2();
    signals2.setComputed(key, rx);
  }
};

// library/src/plugins/official/core/attributes/signals.ts
var Signals = {
  type: 1 /* Attribute */,
  name: "signals",
  onLoad: (ctx) => {
    const { key, mods, signals: signals2, value, genRX: genRX2 } = ctx;
    const ifMissing = mods.has("ifmissing");
    if (key !== "") {
      const k = modifyCasing(key, mods);
      const v = value === "" ? value : genRX2()();
      if (ifMissing) {
        signals2.upsertIfMissing(k, v);
      } else {
        signals2.setValue(k, v);
      }
    } else {
      const obj = jsStrToObject(ctx.value);
      ctx.value = JSON.stringify(obj);
      const rx = genRX2();
      const nv = rx();
      signals2.merge(nv, ifMissing);
    }
  }
};

// library/src/plugins/official/core/attributes/star.ts
var Star = {
  type: 1 /* Attribute */,
  name: "star",
  keyReq: 2 /* Denied */,
  valReq: 2 /* Denied */,
  onLoad: () => {
    alert("YOU ARE PROBABLY OVERCOMPLICATING IT");
  }
};

// library/src/utils/dom.ts
var Hash = class {
  #value = 0;
  #prefix;
  constructor(prefix = DATASTAR) {
    this.#prefix = prefix;
  }
  with(x) {
    if (typeof x === "string") {
      for (const c of x.split("")) {
        this.with(c.charCodeAt(0));
      }
    } else if (typeof x === "boolean") {
      this.with(1 << (x ? 7 : 3));
    } else {
      this.#value = this.#value * 33 ^ x;
    }
    return this;
  }
  get value() {
    return this.#value;
  }
  get string() {
    return this.#prefix + Math.abs(this.#value).toString(36);
  }
};
function elUniqId(el) {
  if (el.id) return el.id;
  const hash = new Hash();
  let currentEl = el;
  while (currentEl) {
    hash.with(currentEl.tagName || "");
    if (currentEl.id) {
      hash.with(currentEl.id);
      break;
    }
    const p = currentEl?.parentNode;
    if (p) hash.with([...p.children].indexOf(currentEl));
    currentEl = p;
  }
  return hash.string;
}
function attrHash(key, val) {
  return new Hash().with(key).with(val).value;
}
function walkDOM(element, callback) {
  if (!element || !(element instanceof HTMLElement || element instanceof SVGElement)) {
    return null;
  }
  const dataset = element.dataset;
  if ("starIgnore" in dataset) {
    return null;
  }
  if (!("starIgnore__self" in dataset)) {
    callback(element);
  }
  let el = element.firstElementChild;
  while (el) {
    walkDOM(el, callback);
    el = el.nextElementSibling;
  }
}

// library/src/engine/errors.ts
var url = "https://data-star.dev/errors";
function dserr(type, reason, metadata = {}) {
  const e = new Error();
  e.name = `${DATASTAR} ${type} error`;
  const r = snake(reason);
  const q = new URLSearchParams({
    metadata: JSON.stringify(metadata)
  }).toString();
  const c = JSON.stringify(metadata, null, 2);
  e.message = `${reason}
More info: ${url}/${type}/${r}?${q}
Context: ${c}`;
  return e;
}
function internalErr(from3, reason, args = {}) {
  return dserr("internal", reason, Object.assign({ from: from3 }, args));
}
function initErr(reason, ctx, metadata = {}) {
  const errCtx = {
    plugin: {
      name: ctx.plugin.name,
      type: PluginType[ctx.plugin.type]
    }
  };
  return dserr("init", reason, Object.assign(errCtx, metadata));
}
function runtimeErr(reason, ctx, metadata = {}) {
  const errCtx = {
    plugin: {
      name: ctx.plugin.name,
      type: PluginType[ctx.plugin.type]
    },
    element: {
      id: ctx.el.id,
      tag: ctx.el.tagName
    },
    expression: {
      rawKey: ctx.rawKey,
      key: ctx.key,
      value: ctx.value,
      validSignals: ctx.signals.paths(),
      fnContent: ctx.fnContent
    }
  };
  return dserr("runtime", reason, Object.assign(errCtx, metadata));
}

// library/src/vendored/preact-core.ts
var from = "preact-signals";
var BRAND_SYMBOL = Symbol.for("preact-signals");
var RUNNING = 1 << 0;
var NOTIFIED = 1 << 1;
var OUTDATED = 1 << 2;
var DISPOSED = 1 << 3;
var HAS_ERROR = 1 << 4;
var TRACKING = 1 << 5;
function startBatch() {
  batchDepth++;
}
function endBatch() {
  if (batchDepth > 1) {
    batchDepth--;
    return;
  }
  let error;
  let hasError = false;
  while (batchedEffect !== void 0) {
    let effect2 = batchedEffect;
    batchedEffect = void 0;
    batchIteration++;
    while (effect2 !== void 0) {
      const next = effect2._nextBatchedEffect;
      effect2._nextBatchedEffect = void 0;
      effect2._flags &= ~NOTIFIED;
      if (!(effect2._flags & DISPOSED) && needsToRecompute(effect2)) {
        try {
          effect2._callback();
        } catch (err) {
          if (!hasError) {
            error = err;
            hasError = true;
          }
        }
      }
      effect2 = next;
    }
  }
  batchIteration = 0;
  batchDepth--;
  if (hasError) throw error;
}
var evalContext = void 0;
var batchedEffect = void 0;
var batchDepth = 0;
var batchIteration = 0;
var globalVersion = 0;
function addDependency(signal) {
  if (evalContext === void 0) {
    return void 0;
  }
  let node = signal._node;
  if (node === void 0 || node._target !== evalContext) {
    node = {
      _version: 0,
      _source: signal,
      _prevSource: evalContext._sources,
      _nextSource: void 0,
      _target: evalContext,
      _prevTarget: void 0,
      _nextTarget: void 0,
      _rollbackNode: node
    };
    if (evalContext._sources !== void 0) {
      evalContext._sources._nextSource = node;
    }
    evalContext._sources = node;
    signal._node = node;
    if (evalContext._flags & TRACKING) {
      signal._subscribe(node);
    }
    return node;
  }
  if (node._version === -1) {
    node._version = 0;
    if (node._nextSource !== void 0) {
      node._nextSource._prevSource = node._prevSource;
      if (node._prevSource !== void 0) {
        node._prevSource._nextSource = node._nextSource;
      }
      node._prevSource = evalContext._sources;
      node._nextSource = void 0;
      evalContext._sources._nextSource = node;
      evalContext._sources = node;
    }
    return node;
  }
  return void 0;
}
function Signal(value) {
  this._value = value;
  this._version = 0;
  this._node = void 0;
  this._targets = void 0;
}
Signal.prototype.brand = BRAND_SYMBOL;
Signal.prototype._refresh = () => true;
Signal.prototype._subscribe = function(node) {
  if (this._targets !== node && node._prevTarget === void 0) {
    node._nextTarget = this._targets;
    if (this._targets !== void 0) {
      this._targets._prevTarget = node;
    }
    this._targets = node;
  }
};
Signal.prototype._unsubscribe = function(node) {
  if (this._targets !== void 0) {
    const prev = node._prevTarget;
    const next = node._nextTarget;
    if (prev !== void 0) {
      prev._nextTarget = next;
      node._prevTarget = void 0;
    }
    if (next !== void 0) {
      next._prevTarget = prev;
      node._nextTarget = void 0;
    }
    if (node === this._targets) {
      this._targets = next;
    }
  }
};
Signal.prototype.subscribe = function(fn) {
  return effect(() => {
    const value = this.value;
    const prevContext = evalContext;
    evalContext = void 0;
    try {
      fn(value);
    } finally {
      evalContext = prevContext;
    }
  });
};
Signal.prototype.valueOf = function() {
  return this.value;
};
Signal.prototype.toString = function() {
  return `${this.value}`;
};
Signal.prototype.toJSON = function() {
  return this.value;
};
Signal.prototype.peek = function() {
  const prevContext = evalContext;
  evalContext = void 0;
  try {
    return this.value;
  } finally {
    evalContext = prevContext;
  }
};
Object.defineProperty(Signal.prototype, "value", {
  get() {
    const node = addDependency(this);
    if (node !== void 0) {
      node._version = this._version;
    }
    return this._value;
  },
  set(value) {
    if (value !== this._value) {
      if (batchIteration > 100) {
        throw internalErr(from, "SignalCycleDetected");
      }
      const old = this._value;
      const revised = value;
      this._value = value;
      this._version++;
      globalVersion++;
      startBatch();
      try {
        for (let node = this._targets; node !== void 0; node = node._nextTarget) {
          node._target._notify();
        }
      } finally {
        endBatch();
      }
      this?._onChange({ old, revised });
    }
  }
});
function needsToRecompute(target) {
  for (let node = target._sources; node !== void 0; node = node._nextSource) {
    if (node._source._version !== node._version || !node._source._refresh() || node._source._version !== node._version) {
      return true;
    }
  }
  return false;
}
function prepareSources(target) {
  for (let node = target._sources; node !== void 0; node = node._nextSource) {
    const rollbackNode = node._source._node;
    if (rollbackNode !== void 0) {
      node._rollbackNode = rollbackNode;
    }
    node._source._node = node;
    node._version = -1;
    if (node._nextSource === void 0) {
      target._sources = node;
      break;
    }
  }
}
function cleanupSources(target) {
  let node = target._sources;
  let head = void 0;
  while (node !== void 0) {
    const prev = node._prevSource;
    if (node._version === -1) {
      node._source._unsubscribe(node);
      if (prev !== void 0) {
        prev._nextSource = node._nextSource;
      }
      if (node._nextSource !== void 0) {
        node._nextSource._prevSource = prev;
      }
    } else {
      head = node;
    }
    node._source._node = node._rollbackNode;
    if (node._rollbackNode !== void 0) {
      node._rollbackNode = void 0;
    }
    node = prev;
  }
  target._sources = head;
}
function Computed2(fn) {
  Signal.call(this, void 0);
  this._fn = fn;
  this._sources = void 0;
  this._globalVersion = globalVersion - 1;
  this._flags = OUTDATED;
}
Computed2.prototype = new Signal();
Computed2.prototype._refresh = function() {
  this._flags &= ~NOTIFIED;
  if (this._flags & RUNNING) {
    return false;
  }
  if ((this._flags & (OUTDATED | TRACKING)) === TRACKING) {
    return true;
  }
  this._flags &= ~OUTDATED;
  if (this._globalVersion === globalVersion) {
    return true;
  }
  this._globalVersion = globalVersion;
  this._flags |= RUNNING;
  if (this._version > 0 && !needsToRecompute(this)) {
    this._flags &= ~RUNNING;
    return true;
  }
  const prevContext = evalContext;
  try {
    prepareSources(this);
    evalContext = this;
    const value = this._fn();
    if (this._flags & HAS_ERROR || this._value !== value || this._version === 0) {
      this._value = value;
      this._flags &= ~HAS_ERROR;
      this._version++;
    }
  } catch (err) {
    this._value = err;
    this._flags |= HAS_ERROR;
    this._version++;
  }
  evalContext = prevContext;
  cleanupSources(this);
  this._flags &= ~RUNNING;
  return true;
};
Computed2.prototype._subscribe = function(node) {
  if (this._targets === void 0) {
    this._flags |= OUTDATED | TRACKING;
    for (let node2 = this._sources; node2 !== void 0; node2 = node2._nextSource) {
      node2._source._subscribe(node2);
    }
  }
  Signal.prototype._subscribe.call(this, node);
};
Computed2.prototype._unsubscribe = function(node) {
  if (this._targets !== void 0) {
    Signal.prototype._unsubscribe.call(this, node);
    if (this._targets === void 0) {
      this._flags &= ~TRACKING;
      for (let node2 = this._sources; node2 !== void 0; node2 = node2._nextSource) {
        node2._source._unsubscribe(node2);
      }
    }
  }
};
Computed2.prototype._notify = function() {
  if (!(this._flags & NOTIFIED)) {
    this._flags |= OUTDATED | NOTIFIED;
    for (let node = this._targets; node !== void 0; node = node._nextTarget) {
      node._target._notify();
    }
  }
};
Object.defineProperty(Computed2.prototype, "value", {
  get() {
    if (this._flags & RUNNING) {
      throw internalErr(from, "SignalCycleDetected");
    }
    const node = addDependency(this);
    this._refresh();
    if (node !== void 0) {
      node._version = this._version;
    }
    if (this._flags & HAS_ERROR) {
      throw internalErr(from, "GetComputedError", { value: this._value });
    }
    return this._value;
  }
});
function computed(fn) {
  return new Computed2(fn);
}
function cleanupEffect(effect2) {
  const cleanup = effect2._cleanup;
  effect2._cleanup = void 0;
  if (typeof cleanup === "function") {
    startBatch();
    const prevContext = evalContext;
    evalContext = void 0;
    try {
      cleanup();
    } catch (error) {
      effect2._flags &= ~RUNNING;
      effect2._flags |= DISPOSED;
      disposeEffect(effect2);
      throw internalErr(from, "CleanupEffectError", { error });
    } finally {
      evalContext = prevContext;
      endBatch();
    }
  }
}
function disposeEffect(effect2) {
  for (let node = effect2._sources; node !== void 0; node = node._nextSource) {
    node._source._unsubscribe(node);
  }
  effect2._fn = void 0;
  effect2._sources = void 0;
  cleanupEffect(effect2);
}
function endEffect(prevContext) {
  if (evalContext !== this) {
    throw internalErr(from, "EndEffectError");
  }
  cleanupSources(this);
  evalContext = prevContext;
  this._flags &= ~RUNNING;
  if (this._flags & DISPOSED) {
    disposeEffect(this);
  }
  endBatch();
}
function Effect(fn) {
  this._fn = fn;
  this._cleanup = void 0;
  this._sources = void 0;
  this._nextBatchedEffect = void 0;
  this._flags = TRACKING;
}
Effect.prototype._callback = function() {
  const finish = this._start();
  try {
    if (this._flags & DISPOSED) return;
    if (this._fn === void 0) return;
    const cleanup = this._fn();
    if (typeof cleanup === "function") {
      this._cleanup = cleanup;
    }
  } finally {
    finish();
  }
};
Effect.prototype._start = function() {
  if (this._flags & RUNNING) {
    throw internalErr(from, "SignalCycleDetected");
  }
  this._flags |= RUNNING;
  this._flags &= ~DISPOSED;
  cleanupEffect(this);
  prepareSources(this);
  startBatch();
  const prevContext = evalContext;
  evalContext = this;
  return endEffect.bind(this, prevContext);
};
Effect.prototype._notify = function() {
  if (!(this._flags & NOTIFIED)) {
    this._flags |= NOTIFIED;
    this._nextBatchedEffect = batchedEffect;
    batchedEffect = this;
  }
};
Effect.prototype._dispose = function() {
  this._flags |= DISPOSED;
  if (!(this._flags & RUNNING)) {
    disposeEffect(this);
  }
};
function effect(fn) {
  const effect2 = new Effect(fn);
  try {
    effect2._callback();
  } catch (error) {
    effect2._dispose();
    throw error;
  }
  return effect2._dispose.bind(effect2);
}

// library/src/engine/signals.ts
var from2 = "namespacedSignals";
var dispatchSignalEvent = (evt) => {
  document.dispatchEvent(
    new CustomEvent(DATASTAR_SIGNAL_EVENT, {
      detail: Object.assign({ added: [], removed: [], updated: [] }, evt)
    })
  );
};
function nestedValues(signal, onlyPublic = false) {
  const kv = {};
  for (const key in signal) {
    if (Object.hasOwn(signal, key)) {
      if (onlyPublic && key.startsWith("_")) {
        continue;
      }
      const value = signal[key];
      if (value instanceof Signal) {
        kv[key] = value.value;
      } else {
        kv[key] = nestedValues(value);
      }
    }
  }
  return kv;
}
function mergeNested(target, values, onlyIfMissing = false) {
  const evt = {
    added: [],
    removed: [],
    updated: []
  };
  for (const key in values) {
    if (Object.hasOwn(values, key)) {
      if (key.match(/\_\_+/)) {
        throw internalErr(from2, "InvalidSignalKey", { key });
      }
      const value = values[key];
      if (value instanceof Object && !Array.isArray(value)) {
        if (!target[key]) {
          target[key] = {};
        }
        const subEvt = mergeNested(
          target[key],
          value,
          onlyIfMissing
        );
        evt.added.push(...subEvt.added.map((k) => `${key}.${k}`));
        evt.removed.push(...subEvt.removed.map((k) => `${key}.${k}`));
        evt.updated.push(...subEvt.updated.map((k) => `${key}.${k}`));
      } else {
        const hasKey = Object.hasOwn(target, key);
        if (hasKey) {
          if (onlyIfMissing) continue;
          const t = target[key];
          if (t instanceof Signal) {
            const oldValue = t.value;
            t.value = value;
            if (oldValue !== value) {
              evt.updated.push(key);
            }
            continue;
          }
        }
        const s = new Signal(value);
        s._onChange = () => {
          dispatchSignalEvent({ updated: [key] });
        };
        target[key] = s;
        evt.added.push(key);
      }
    }
  }
  return evt;
}
function walkNestedSignal(signal, cb) {
  for (const key in signal) {
    if (Object.hasOwn(signal, key)) {
      const value = signal[key];
      if (value instanceof Signal) {
        cb(key, value);
      } else {
        walkNestedSignal(value, (path, value2) => {
          cb(`${key}.${path}`, value2);
        });
      }
    }
  }
}
function nestedSubset(original, ...keys) {
  const subset = {};
  for (const key of keys) {
    const parts = key.split(".");
    let subOriginal = original;
    let subSubset = subset;
    for (let i = 0; i < parts.length - 1; i++) {
      const part = parts[i];
      if (!subOriginal[part]) {
        return {};
      }
      if (!subSubset[part]) {
        subSubset[part] = {};
      }
      subOriginal = subOriginal[part];
      subSubset = subSubset[part];
    }
    const last = parts[parts.length - 1];
    subSubset[last] = subOriginal[last];
  }
  return subset;
}
var SignalsRoot = class {
  #signals = {};
  exists(dotDelimitedPath) {
    return !!this.signal(dotDelimitedPath);
  }
  signal(dotDelimitedPath) {
    const parts = dotDelimitedPath.split(".");
    let subSignals = this.#signals;
    for (let i = 0; i < parts.length - 1; i++) {
      const part = parts[i];
      if (!subSignals[part]) {
        return null;
      }
      subSignals = subSignals[part];
    }
    const last = parts[parts.length - 1];
    const signal = subSignals[last];
    if (!signal)
      throw internalErr(from2, "SignalNotFound", { path: dotDelimitedPath });
    return signal;
  }
  setSignal(dotDelimitedPath, signal) {
    const parts = dotDelimitedPath.split(".");
    let subSignals = this.#signals;
    for (let i = 0; i < parts.length - 1; i++) {
      const part = parts[i];
      if (!subSignals[part]) {
        subSignals[part] = {};
      }
      subSignals = subSignals[part];
    }
    const last = parts[parts.length - 1];
    subSignals[last] = signal;
  }
  setComputed(dotDelimitedPath, fn) {
    const signal = computed(() => fn());
    this.setSignal(dotDelimitedPath, signal);
  }
  value(dotDelimitedPath) {
    const signal = this.signal(dotDelimitedPath);
    return signal?.value;
  }
  setValue(dotDelimitedPath, value) {
    const { signal } = this.upsertIfMissing(dotDelimitedPath, value);
    const oldValue = signal.value;
    signal.value = value;
    if (oldValue !== value) {
      dispatchSignalEvent({ updated: [dotDelimitedPath] });
    }
  }
  upsertIfMissing(dotDelimitedPath, defaultValue) {
    const parts = dotDelimitedPath.split(".");
    let subSignals = this.#signals;
    for (let i = 0; i < parts.length - 1; i++) {
      const part = parts[i];
      if (!subSignals[part]) {
        subSignals[part] = {};
      }
      subSignals = subSignals[part];
    }
    const last = parts[parts.length - 1];
    const current = subSignals[last];
    if (current instanceof Signal) {
      return { signal: current, inserted: false };
    }
    const signal = new Signal(defaultValue);
    signal._onChange = () => {
      dispatchSignalEvent({ updated: [dotDelimitedPath] });
    };
    subSignals[last] = signal;
    dispatchSignalEvent({ added: [dotDelimitedPath] });
    return { signal, inserted: true };
  }
  remove(...dotDelimitedPaths) {
    if (!dotDelimitedPaths.length) {
      this.#signals = {};
      return;
    }
    const removed = Array();
    for (const path of dotDelimitedPaths) {
      const parts = path.split(".");
      let subSignals = this.#signals;
      for (let i = 0; i < parts.length - 1; i++) {
        const part = parts[i];
        if (!subSignals[part]) {
          return;
        }
        subSignals = subSignals[part];
      }
      const last = parts[parts.length - 1];
      delete subSignals[last];
      removed.push(path);
    }
    dispatchSignalEvent({ removed });
  }
  merge(other, onlyIfMissing = false) {
    const evt = mergeNested(this.#signals, other, onlyIfMissing);
    if (evt.added.length || evt.removed.length || evt.updated.length) {
      dispatchSignalEvent(evt);
    }
  }
  subset(...keys) {
    return nestedSubset(this.values(), ...keys);
  }
  walk(cb) {
    walkNestedSignal(this.#signals, cb);
  }
  paths() {
    const signalNames = new Array();
    this.walk((path) => signalNames.push(path));
    return signalNames;
  }
  values(onlyPublic = false) {
    return nestedValues(this.#signals, onlyPublic);
  }
  JSON(shouldIndent = true, onlyPublic = false) {
    const values = this.values(onlyPublic);
    if (!shouldIndent) {
      return JSON.stringify(values);
    }
    return JSON.stringify(values, null, 2);
  }
  toString() {
    return this.JSON();
  }
};

// library/src/engine/engine.ts
var signals = new SignalsRoot();
var plugins = [];
var actions = {};
var watchers = [];
var removals = /* @__PURE__ */ new Map();
var mutationObserver = null;
var alias = "";
function setAlias(value) {
  alias = value;
}
function load(...pluginsToLoad) {
  for (const plugin of pluginsToLoad) {
    const ctx = {
      plugin,
      signals,
      effect: (cb) => effect(cb),
      actions,
      removals,
      applyToElement
    };
    let globalInitializer;
    switch (plugin.type) {
      case 2 /* Watcher */: {
        const wp = plugin;
        watchers.push(wp);
        globalInitializer = wp.onGlobalInit;
        break;
      }
      case 3 /* Action */: {
        actions[plugin.name] = plugin;
        break;
      }
      case 1 /* Attribute */: {
        const ap = plugin;
        plugins.push(ap);
        globalInitializer = ap.onGlobalInit;
        break;
      }
      default: {
        throw initErr("InvalidPluginType", ctx);
      }
    }
    if (globalInitializer) {
      globalInitializer(ctx);
    }
  }
  plugins.sort((a, b) => {
    const lenDiff = b.name.length - a.name.length;
    if (lenDiff !== 0) return lenDiff;
    return a.name.localeCompare(b.name);
  });
}
function apply() {
  queueMicrotask(() => {
    applyToElement(document.documentElement);
    observe();
  });
}
function applyToElement(rootElement) {
  walkDOM(rootElement, (el) => {
    const toApply = new Array();
    const elCleanups = removals.get(el.id) || /* @__PURE__ */ new Map();
    const toCleanup = new Map([...elCleanups]);
    const hashes = /* @__PURE__ */ new Map();
    for (const datasetKey of Object.keys(el.dataset)) {
      if (!datasetKey.startsWith(alias)) {
        break;
      }
      const datasetValue = el.dataset[datasetKey] || "";
      const currentHash = attrHash(datasetKey, datasetValue);
      hashes.set(datasetKey, currentHash);
      if (elCleanups.has(currentHash)) {
        toCleanup.delete(currentHash);
      } else {
        toApply.push(datasetKey);
      }
    }
    for (const [_, cleanup] of toCleanup) {
      cleanup();
    }
    for (const key of toApply) {
      const h = hashes.get(key);
      applyAttributePlugin(el, key, h);
    }
  });
}
function observe() {
  if (mutationObserver) {
    return;
  }
  mutationObserver = new MutationObserver((mutations) => {
    const toRemove = /* @__PURE__ */ new Set();
    const toApply = /* @__PURE__ */ new Set();
    for (const { target, type, addedNodes, removedNodes } of mutations) {
      switch (type) {
        case "childList":
          {
            for (const node of removedNodes) {
              toRemove.add(node);
            }
            for (const node of addedNodes) {
              toApply.add(node);
            }
          }
          break;
        case "attributes": {
          toApply.add(target);
          break;
        }
      }
    }
    for (const el of toRemove) {
      const elTracking = removals.get(el.id);
      if (elTracking) {
        for (const [hash, cleanup] of elTracking) {
          cleanup();
          elTracking.delete(hash);
        }
        if (elTracking.size === 0) {
          removals.delete(el.id);
        }
      }
    }
    for (const el of toApply) {
      applyToElement(el);
    }
  });
  mutationObserver.observe(document.body, {
    attributes: true,
    attributeOldValue: true,
    childList: true,
    subtree: true
  });
}
function applyAttributePlugin(el, camelCasedKey, hash) {
  const rawKey = camel(camelCasedKey.slice(alias.length));
  const plugin = plugins.find((p) => rawKey.startsWith(p.name));
  if (!plugin) return;
  if (!el.id.length) el.id = elUniqId(el);
  let [key, ...rawModifiers] = rawKey.slice(plugin.name.length).split(/\_\_+/);
  const hasKey = key.length > 0;
  if (hasKey) {
    key = camel(key);
  }
  const value = el.dataset[camelCasedKey] || "";
  const hasValue = value.length > 0;
  const ctx = {
    signals,
    applyToElement,
    effect: (cb) => effect(cb),
    actions,
    removals,
    genRX: () => genRX(ctx, ...plugin.argNames || []),
    plugin,
    el,
    rawKey,
    key,
    value,
    mods: /* @__PURE__ */ new Map()
  };
  const keyReq = plugin.keyReq || 0 /* Allowed */;
  if (hasKey) {
    if (keyReq === 2 /* Denied */) {
      throw runtimeErr(`${plugin.name}KeyNotAllowed`, ctx);
    }
  } else if (keyReq === 1 /* Must */) {
    throw runtimeErr(`${plugin.name}KeyRequired`, ctx);
  }
  const valReq = plugin.valReq || 0 /* Allowed */;
  if (hasValue) {
    if (valReq === 2 /* Denied */) {
      throw runtimeErr(`${plugin.name}ValueNotAllowed`, ctx);
    }
  } else if (valReq === 1 /* Must */) {
    throw runtimeErr(`${plugin.name}ValueRequired`, ctx);
  }
  if (keyReq === 3 /* Exclusive */ || valReq === 3 /* Exclusive */) {
    if (hasKey && hasValue) {
      throw runtimeErr(`${plugin.name}KeyAndValueProvided`, ctx);
    }
    if (!hasKey && !hasValue) {
      throw runtimeErr(`${plugin.name}KeyOrValueRequired`, ctx);
    }
  }
  for (const rawMod of rawModifiers) {
    const [label, ...mod] = rawMod.split(".");
    ctx.mods.set(camel(label), new Set(mod.map((t) => t.toLowerCase())));
  }
  const cleanup = plugin.onLoad(ctx) ?? (() => {
  });
  let elTracking = removals.get(el.id);
  if (!elTracking) {
    elTracking = /* @__PURE__ */ new Map();
    removals.set(el.id, elTracking);
  }
  elTracking.set(hash, cleanup);
}
function genRX(ctx, ...argNames) {
  let userExpression = "";
  const statementRe = /(\/(\\\/|[^\/])*\/|"(\\"|[^\"])*"|'(\\'|[^'])*'|`(\\`|[^`])*`|[^;])+/gm;
  const statements = ctx.value.trim().match(statementRe);
  if (statements) {
    const lastIdx = statements.length - 1;
    const last = statements[lastIdx].trim();
    if (!last.startsWith("return")) {
      statements[lastIdx] = `return (${last});`;
    }
    userExpression = statements.join(";\n");
  }
  const escaped = /* @__PURE__ */ new Map();
  const escapeRe = new RegExp(`(?:${DSP})(.*?)(?:${DSS})`, "gm");
  for (const match of userExpression.matchAll(escapeRe)) {
    const k = match[1];
    const v = new Hash("dsEscaped").with(k).string;
    escaped.set(v, k);
    userExpression = userExpression.replace(DSP + k + DSS, v);
  }
  const fnCall = /@(\w*)\(/gm;
  const matches = userExpression.matchAll(fnCall);
  const methodsCalled = /* @__PURE__ */ new Set();
  for (const match of matches) {
    methodsCalled.add(match[1]);
  }
  const actionsRe = new RegExp(`@(${Object.keys(actions).join("|")})\\(`, "gm");
  userExpression = userExpression.replaceAll(
    actionsRe,
    "ctx.actions.$1.fn(ctx,"
  );
  const signalNames = ctx.signals.paths();
  if (signalNames.length) {
    const signalsRe = new RegExp(`\\$(${signalNames.join("|")})(\\W|$)`, "gm");
    userExpression = userExpression.replaceAll(
      signalsRe,
      `ctx.signals.signal('$1').value$2`
    );
  }
  for (const [k, v] of escaped) {
    userExpression = userExpression.replace(k, v);
  }
  const fnContent = `return (() => {
${userExpression}
})()`;
  ctx.fnContent = fnContent;
  try {
    const fn = new Function("ctx", ...argNames, fnContent);
    return (...args) => {
      try {
        return fn(ctx, ...args);
      } catch (error) {
        throw runtimeErr("ExecuteExpression", ctx, {
          error: error.message
        });
      }
    };
  } catch (error) {
    throw runtimeErr("GenerateExpression", ctx, {
      error: error.message
    });
  }
}

// library/src/engine/index.ts
load(Star, Signals, Computed);

// library/src/vendored/fetch-event-source.ts
async function getBytes(stream, onChunk) {
  const reader = stream.getReader();
  let result;
  while (!(result = await reader.read()).done) {
    onChunk(result.value);
  }
}
function getLines(onLine) {
  let buffer;
  let position;
  let fieldLength;
  let discardTrailingNewline = false;
  return function onChunk(arr) {
    if (buffer === void 0) {
      buffer = arr;
      position = 0;
      fieldLength = -1;
    } else {
      buffer = concat(buffer, arr);
    }
    const bufLength = buffer.length;
    let lineStart = 0;
    while (position < bufLength) {
      if (discardTrailingNewline) {
        if (buffer[position] === 10 /* NewLine */) {
          lineStart = ++position;
        }
        discardTrailingNewline = false;
      }
      let lineEnd = -1;
      for (; position < bufLength && lineEnd === -1; ++position) {
        switch (buffer[position]) {
          case 58 /* Colon */:
            if (fieldLength === -1) {
              fieldLength = position - lineStart;
            }
            break;
          // @ts-ignore:7029 \r case below should fallthrough to \n:
          // biome-ignore lint/suspicious/noFallthroughSwitchClause: <explanation>
          case 13 /* CarriageReturn */:
            discardTrailingNewline = true;
          case 10 /* NewLine */:
            lineEnd = position;
            break;
        }
      }
      if (lineEnd === -1) {
        break;
      }
      onLine(buffer.subarray(lineStart, lineEnd), fieldLength);
      lineStart = position;
      fieldLength = -1;
    }
    if (lineStart === bufLength) {
      buffer = void 0;
    } else if (lineStart !== 0) {
      buffer = buffer.subarray(lineStart);
      position -= lineStart;
    }
  };
}
function getMessages(onId, onRetry, onMessage) {
  let message = newMessage();
  const decoder = new TextDecoder();
  return function onLine(line, fieldLength) {
    if (line.length === 0) {
      onMessage?.(message);
      message = newMessage();
    } else if (fieldLength > 0) {
      const field = decoder.decode(line.subarray(0, fieldLength));
      const valueOffset = fieldLength + (line[fieldLength + 1] === 32 /* Space */ ? 2 : 1);
      const value = decoder.decode(line.subarray(valueOffset));
      switch (field) {
        case "data":
          message.data = message.data ? `${message.data}
${value}` : value;
          break;
        case "event":
          message.event = value;
          break;
        case "id":
          onId(message.id = value);
          break;
        case "retry": {
          const retry = Number.parseInt(value, 10);
          if (!Number.isNaN(retry)) {
            onRetry(message.retry = retry);
          }
          break;
        }
      }
    }
  };
}
function concat(a, b) {
  const res = new Uint8Array(a.length + b.length);
  res.set(a);
  res.set(b, a.length);
  return res;
}
function newMessage() {
  return {
    data: "",
    event: "",
    id: "",
    retry: void 0
  };
}
var EventStreamContentType = "text/event-stream";
var LastEventId = "last-event-id";
function fetchEventSource(input, {
  signal: inputSignal,
  headers: inputHeaders,
  onopen: inputOnOpen,
  onmessage,
  onclose,
  onerror,
  openWhenHidden,
  fetch: inputFetch,
  retryInterval = 1e3,
  retryScaler = 2,
  retryMaxWaitMs = 3e4,
  retryMaxCount = 10,
  ...rest
}) {
  return new Promise((resolve, reject) => {
    let retries = 0;
    const headers = { ...inputHeaders };
    if (!headers.accept) {
      headers.accept = EventStreamContentType;
    }
    let curRequestController;
    function onVisibilityChange() {
      curRequestController.abort();
      if (!document.hidden) {
        create();
      }
    }
    if (!openWhenHidden) {
      document.addEventListener("visibilitychange", onVisibilityChange);
    }
    let retryTimer = 0;
    function dispose() {
      document.removeEventListener("visibilitychange", onVisibilityChange);
      window.clearTimeout(retryTimer);
      curRequestController.abort();
    }
    inputSignal?.addEventListener("abort", () => {
      dispose();
      resolve();
    });
    const fetch = inputFetch ?? window.fetch;
    const onopen = inputOnOpen ?? function defaultOnOpen() {
    };
    async function create() {
      curRequestController = new AbortController();
      try {
        const response = await fetch(input, {
          ...rest,
          headers,
          signal: curRequestController.signal
        });
        await onopen(response);
        await getBytes(
          response.body,
          getLines(
            getMessages(
              (id) => {
                if (id) {
                  headers[LastEventId] = id;
                } else {
                  delete headers[LastEventId];
                }
              },
              (retry) => {
                retryInterval = retry;
              },
              onmessage
            )
          )
        );
        onclose?.();
        dispose();
        resolve();
      } catch (err) {
        if (!curRequestController.signal.aborted) {
          try {
            const interval = onerror?.(err) ?? retryInterval;
            window.clearTimeout(retryTimer);
            retryTimer = window.setTimeout(create, interval);
            retryInterval *= retryScaler;
            retryInterval = Math.min(retryInterval, retryMaxWaitMs);
            retries++;
            if (retries > retryMaxCount) {
              dispose();
              reject("Max retries reached.");
            } else {
              console.error(
                `Datastar failed to reach ${input.toString()} retrying in ${interval}ms.`
              );
            }
          } catch (innerErr) {
            dispose();
            reject(innerErr);
          }
        }
      }
    }
    create();
  });
}

// library/src/plugins/official/backend/shared.ts
var DATASTAR_SSE_EVENT = `${DATASTAR}-sse`;
var SETTLING_CLASS = `${DATASTAR}-settling`;
var SWAPPING_CLASS = `${DATASTAR}-swapping`;
var STARTED = "started";
var FINISHED = "finished";
var ERROR = "error";
var RETRYING = "retrying";
function datastarSSEEventWatcher(eventType, fn) {
  document.addEventListener(
    DATASTAR_SSE_EVENT,
    (event) => {
      if (event.detail.type !== eventType) return;
      const { argsRaw } = event.detail;
      fn(argsRaw);
    }
  );
}

// library/src/plugins/official/backend/actions/sse.ts
function dispatchSSE(type, argsRaw) {
  document.dispatchEvent(
    new CustomEvent(DATASTAR_SSE_EVENT, {
      detail: { type, argsRaw }
    })
  );
}
var isWrongContent = (err) => `${err}`.includes("text/event-stream");
var sse = async (ctx, method, url2, args) => {
  const {
    el: { id: elId },
    el,
    signals: signals2
  } = ctx;
  const {
    headers: userHeaders,
    contentType,
    includeLocal,
    selector,
    openWhenHidden,
    retryInterval,
    retryScaler,
    retryMaxWaitMs,
    retryMaxCount,
    abort
  } = Object.assign(
    {
      headers: {},
      contentType: "json",
      includeLocal: false,
      selector: null,
      openWhenHidden: false,
      // will keep the request open even if the document is hidden.
      retryInterval: DefaultSseRetryDurationMs,
      // the retry interval in milliseconds
      retryScaler: 2,
      // the amount to multiply the retry interval by each time
      retryMaxWaitMs: 3e4,
      // the maximum retry interval in milliseconds
      retryMaxCount: 10,
      // the maximum number of retries before giving up
      abort: void 0
    },
    args
  );
  const action = method.toLowerCase();
  let cleanupFn = () => {
  };
  try {
    dispatchSSE(STARTED, { elId });
    if (!url2?.length) {
      throw runtimeErr("SseNoUrlProvided", ctx, { action });
    }
    const initialHeaders = {};
    initialHeaders[DATASTAR_REQUEST] = true;
    if (contentType === "json") {
      initialHeaders["Content-Type"] = "application/json";
    }
    const headers = Object.assign({}, initialHeaders, userHeaders);
    const req = {
      method,
      headers,
      openWhenHidden,
      retryInterval,
      retryScaler,
      retryMaxWaitMs,
      retryMaxCount,
      signal: abort,
      onopen: async (response) => {
        if (response.status >= 400) {
          const status = response.status.toString();
          dispatchSSE(ERROR, { status });
        }
      },
      onmessage: (evt) => {
        if (!evt.event.startsWith(DATASTAR)) {
          return;
        }
        const type = evt.event;
        const argsRawLines = {};
        const lines = evt.data.split("\n");
        for (const line of lines) {
          const colonIndex = line.indexOf(" ");
          const key = line.slice(0, colonIndex);
          let argLines = argsRawLines[key];
          if (!argLines) {
            argLines = [];
            argsRawLines[key] = argLines;
          }
          const value = line.slice(colonIndex + 1);
          argLines.push(value);
        }
        const argsRaw = {};
        for (const [key, lines2] of Object.entries(argsRawLines)) {
          argsRaw[key] = lines2.join("\n");
        }
        dispatchSSE(type, argsRaw);
      },
      onerror: (error) => {
        if (isWrongContent(error)) {
          throw runtimeErr("InvalidContentType", ctx, { url: url2 });
        }
        if (error) {
          console.error(error.message);
          dispatchSSE(RETRYING, { message: error.message });
        }
      }
    };
    const urlInstance = new URL(url2, window.location.origin);
    const queryParams = new URLSearchParams(urlInstance.search);
    if (contentType === "json") {
      const json = signals2.JSON(false, !includeLocal);
      if (method === "GET") {
        queryParams.set(DATASTAR, json);
      } else {
        req.body = json;
      }
    } else if (contentType === "form") {
      const formEl = selector ? document.querySelector(selector) : el.closest("form");
      if (formEl === null) {
        if (selector) {
          throw runtimeErr("SseFormNotFound", ctx, { action, selector });
        }
        throw runtimeErr("SseClosestFormNotFound", ctx, { action });
      }
      if (el !== formEl) {
        const preventDefault = (evt) => evt.preventDefault();
        formEl.addEventListener("submit", preventDefault);
        cleanupFn = () => formEl.removeEventListener("submit", preventDefault);
      }
      if (!formEl.checkValidity()) {
        formEl.reportValidity();
        cleanupFn();
        return;
      }
      const formData = new FormData(formEl);
      if (method === "GET") {
        const formParams = new URLSearchParams(formData);
        for (const [key, value] of formParams) {
          queryParams.set(key, value);
        }
      } else {
        req.body = formData;
      }
    } else {
      throw runtimeErr("SseInvalidContentType", ctx, { action, contentType });
    }
    urlInstance.search = queryParams.toString();
    try {
      await fetchEventSource(urlInstance.toString(), req);
    } catch (error) {
      if (!isWrongContent(error)) {
        throw runtimeErr("SseFetchFailed", ctx, { method, url: url2, error });
      }
    }
  } finally {
    dispatchSSE(FINISHED, { elId });
    cleanupFn();
  }
};

// library/src/plugins/official/backend/actions/delete.ts
var DELETE = {
  type: 3 /* Action */,
  name: "delete",
  fn: async (ctx, url2, args) => {
    return sse(ctx, "DELETE", url2, { ...args });
  }
};

// library/src/plugins/official/backend/actions/get.ts
var GET = {
  type: 3 /* Action */,
  name: "get",
  fn: async (ctx, url2, args) => {
    return sse(ctx, "GET", url2, { ...args });
  }
};

// library/src/plugins/official/backend/actions/patch.ts
var PATCH = {
  type: 3 /* Action */,
  name: "patch",
  fn: async (ctx, url2, args) => {
    return sse(ctx, "PATCH", url2, { ...args });
  }
};

// library/src/plugins/official/backend/actions/post.ts
var POST = {
  type: 3 /* Action */,
  name: "post",
  fn: async (ctx, url2, args) => {
    return sse(ctx, "POST", url2, { ...args });
  }
};

// library/src/plugins/official/backend/actions/put.ts
var PUT = {
  type: 3 /* Action */,
  name: "put",
  fn: async (ctx, url2, args) => {
    return sse(ctx, "PUT", url2, { ...args });
  }
};

// library/src/plugins/official/backend/attributes/indicator.ts
var Indicator = {
  type: 1 /* Attribute */,
  name: "indicator",
  keyReq: 3 /* Exclusive */,
  valReq: 3 /* Exclusive */,
  onLoad: ({ el, key, mods, signals: signals2, value }) => {
    const signalName = key ? modifyCasing(key, mods) : trimDollarSignPrefix(value);
    const { signal } = signals2.upsertIfMissing(signalName, false);
    const watcher = (event) => {
      const {
        type,
        argsRaw: { elId }
      } = event.detail;
      if (elId !== el.id) return;
      switch (type) {
        case STARTED:
          signal.value = true;
          break;
        case FINISHED:
          signal.value = false;
          break;
      }
    };
    document.addEventListener(DATASTAR_SSE_EVENT, watcher);
    return () => {
      signal.value = false;
      document.removeEventListener(DATASTAR_SSE_EVENT, watcher);
    };
  }
};

// library/src/plugins/official/backend/watchers/executeScript.ts
var ExecuteScript = {
  type: 2 /* Watcher */,
  name: EventTypes.ExecuteScript,
  onGlobalInit: async (ctx) => {
    datastarSSEEventWatcher(
      EventTypes.ExecuteScript,
      ({
        autoRemove: autoRemoveRaw = `${DefaultExecuteScriptAutoRemove}`,
        attributes: attributesRaw = DefaultExecuteScriptAttributes,
        script
      }) => {
        const autoRemove = isBoolString(autoRemoveRaw);
        if (!script?.length) {
          throw initErr("NoScriptProvided", ctx);
        }
        const scriptEl = document.createElement("script");
        for (const attr of attributesRaw.split("\n")) {
          const pivot = attr.indexOf(" ");
          const key = pivot ? attr.slice(0, pivot) : attr;
          const value = pivot ? attr.slice(pivot) : "";
          scriptEl.setAttribute(key.trim(), value.trim());
        }
        scriptEl.text = script;
        document.head.appendChild(scriptEl);
        if (autoRemove) {
          scriptEl.remove();
        }
      }
    );
  }
};

// library/src/utils/view-transtions.ts
var docWithViewTransitionAPI = document;
var supportsViewTransitions = !!docWithViewTransitionAPI.startViewTransition;

// library/src/vendored/idiomorph.esm.js
var Idiomorph = function() {
  "use strict";
  const noOp = () => {
  };
  const defaults = {
    morphStyle: "outerHTML",
    callbacks: {
      beforeNodeAdded: noOp,
      afterNodeAdded: noOp,
      beforeNodeMorphed: noOp,
      afterNodeMorphed: noOp,
      beforeNodeRemoved: noOp,
      afterNodeRemoved: noOp,
      beforeAttributeUpdated: noOp
    },
    head: {
      style: "merge",
      shouldPreserve: (elt) => elt.getAttribute("im-preserve") === "true",
      shouldReAppend: (elt) => elt.getAttribute("im-re-append") === "true",
      shouldRemove: noOp,
      afterHeadMorphed: noOp
    },
    restoreFocus: true
  };
  function morph(oldNode, newContent, config = {}) {
    oldNode = normalizeElement(oldNode);
    const newNode = normalizeParent(newContent);
    const ctx = createMorphContext(oldNode, newNode, config);
    const morphedNodes = saveAndRestoreFocus(ctx, () => {
      return withHeadBlocking(
        ctx,
        oldNode,
        newNode,
        /** @param {MorphContext} ctx */
        (ctx2) => {
          if (ctx2.morphStyle === "innerHTML") {
            morphChildren(ctx2, oldNode, newNode);
            return Array.from(oldNode.childNodes);
          } else {
            return morphOuterHTML(ctx2, oldNode, newNode);
          }
        }
      );
    });
    ctx.pantry.remove();
    return morphedNodes;
  }
  function morphOuterHTML(ctx, oldNode, newNode) {
    const oldParent = normalizeParent(oldNode);
    morphChildren(
      ctx,
      oldParent,
      newNode,
      // these two optional params are the secret sauce
      oldNode,
      // start point for iteration
      oldNode.nextSibling
      // end point for iteration
    );
    return Array.from(oldParent.childNodes);
  }
  function saveAndRestoreFocus(ctx, fn) {
    if (!ctx.config.restoreFocus) return fn();
    let activeElement = (
      /** @type {HTMLInputElement|HTMLTextAreaElement|null} */
      document.activeElement
    );
    if (!(activeElement instanceof HTMLInputElement || activeElement instanceof HTMLTextAreaElement)) {
      return fn();
    }
    const { id: activeElementId, selectionStart, selectionEnd } = activeElement;
    const results = fn();
    if (activeElementId && activeElementId !== document.activeElement?.id) {
      activeElement = ctx.target.querySelector(`[id="${activeElementId}"]`);
      activeElement?.focus();
    }
    if (activeElement && !activeElement.selectionEnd && selectionEnd) {
      activeElement.setSelectionRange(selectionStart, selectionEnd);
    }
    return results;
  }
  const morphChildren = /* @__PURE__ */ function() {
    function morphChildren2(ctx, oldParent, newParent, insertionPoint = null, endPoint = null) {
      if (oldParent instanceof HTMLTemplateElement && newParent instanceof HTMLTemplateElement) {
        oldParent = oldParent.content;
        newParent = newParent.content;
      }
      insertionPoint ||= oldParent.firstChild;
      for (const newChild of newParent.childNodes) {
        if (insertionPoint && insertionPoint != endPoint) {
          const bestMatch = findBestMatch(
            ctx,
            newChild,
            insertionPoint,
            endPoint
          );
          if (bestMatch) {
            if (bestMatch !== insertionPoint) {
              removeNodesBetween(ctx, insertionPoint, bestMatch);
            }
            morphNode(bestMatch, newChild, ctx);
            insertionPoint = bestMatch.nextSibling;
            continue;
          }
        }
        if (newChild instanceof Element && ctx.persistentIds.has(newChild.id)) {
          const movedChild = moveBeforeById(
            oldParent,
            newChild.id,
            insertionPoint,
            ctx
          );
          morphNode(movedChild, newChild, ctx);
          insertionPoint = movedChild.nextSibling;
          continue;
        }
        const insertedNode = createNode(
          oldParent,
          newChild,
          insertionPoint,
          ctx
        );
        if (insertedNode) {
          insertionPoint = insertedNode.nextSibling;
        }
      }
      while (insertionPoint && insertionPoint != endPoint) {
        const tempNode = insertionPoint;
        insertionPoint = insertionPoint.nextSibling;
        removeNode(ctx, tempNode);
      }
    }
    function createNode(oldParent, newChild, insertionPoint, ctx) {
      if (ctx.callbacks.beforeNodeAdded(newChild) === false) return null;
      if (ctx.idMap.has(newChild)) {
        const newEmptyChild = document.createElement(
          /** @type {Element} */
          newChild.tagName
        );
        oldParent.insertBefore(newEmptyChild, insertionPoint);
        morphNode(newEmptyChild, newChild, ctx);
        ctx.callbacks.afterNodeAdded(newEmptyChild);
        return newEmptyChild;
      } else {
        const newClonedChild = document.importNode(newChild, true);
        oldParent.insertBefore(newClonedChild, insertionPoint);
        ctx.callbacks.afterNodeAdded(newClonedChild);
        return newClonedChild;
      }
    }
    const findBestMatch = /* @__PURE__ */ function() {
      function findBestMatch2(ctx, node, startPoint, endPoint) {
        let softMatch = null;
        let nextSibling = node.nextSibling;
        let siblingSoftMatchCount = 0;
        let cursor = startPoint;
        while (cursor && cursor != endPoint) {
          if (isSoftMatch(cursor, node)) {
            if (isIdSetMatch(ctx, cursor, node)) {
              return cursor;
            }
            if (softMatch === null) {
              if (!ctx.idMap.has(cursor)) {
                softMatch = cursor;
              }
            }
          }
          if (softMatch === null && nextSibling && isSoftMatch(cursor, nextSibling)) {
            siblingSoftMatchCount++;
            nextSibling = nextSibling.nextSibling;
            if (siblingSoftMatchCount >= 2) {
              softMatch = void 0;
            }
          }
          if (cursor.contains(document.activeElement)) break;
          cursor = cursor.nextSibling;
        }
        return softMatch || null;
      }
      function isIdSetMatch(ctx, oldNode, newNode) {
        let oldSet = ctx.idMap.get(oldNode);
        let newSet = ctx.idMap.get(newNode);
        if (!newSet || !oldSet) return false;
        for (const id of oldSet) {
          if (newSet.has(id)) {
            return true;
          }
        }
        return false;
      }
      function isSoftMatch(oldNode, newNode) {
        const oldElt = (
          /** @type {Element} */
          oldNode
        );
        const newElt = (
          /** @type {Element} */
          newNode
        );
        return oldElt.nodeType === newElt.nodeType && oldElt.tagName === newElt.tagName && // If oldElt has an `id` with possible state and it doesn't match newElt.id then avoid morphing.
        // We'll still match an anonymous node with an IDed newElt, though, because if it got this far,
        // its not persistent, and new nodes can't have any hidden state.
        (!oldElt.id || oldElt.id === newElt.id);
      }
      return findBestMatch2;
    }();
    function removeNode(ctx, node) {
      if (ctx.idMap.has(node)) {
        moveBefore(ctx.pantry, node, null);
      } else {
        if (ctx.callbacks.beforeNodeRemoved(node) === false) return;
        node.parentNode?.removeChild(node);
        ctx.callbacks.afterNodeRemoved(node);
      }
    }
    function removeNodesBetween(ctx, startInclusive, endExclusive) {
      let cursor = startInclusive;
      while (cursor && cursor !== endExclusive) {
        let tempNode = (
          /** @type {Node} */
          cursor
        );
        cursor = cursor.nextSibling;
        removeNode(ctx, tempNode);
      }
      return cursor;
    }
    function moveBeforeById(parentNode, id, after, ctx) {
      const target = (
        /** @type {Element} - will always be found */
        ctx.target.id === id && ctx.target || ctx.target.querySelector(`[id="${id}"]`) || ctx.pantry.querySelector(`[id="${id}"]`)
      );
      removeElementFromAncestorsIdMaps(target, ctx);
      moveBefore(parentNode, target, after);
      return target;
    }
    function removeElementFromAncestorsIdMaps(element, ctx) {
      const id = element.id;
      while (element = element.parentNode) {
        let idSet = ctx.idMap.get(element);
        if (idSet) {
          idSet.delete(id);
          if (!idSet.size) {
            ctx.idMap.delete(element);
          }
        }
      }
    }
    function moveBefore(parentNode, element, after) {
      if (parentNode.moveBefore) {
        try {
          parentNode.moveBefore(element, after);
        } catch (e) {
          parentNode.insertBefore(element, after);
        }
      } else {
        parentNode.insertBefore(element, after);
      }
    }
    return morphChildren2;
  }();
  const morphNode = /* @__PURE__ */ function() {
    function morphNode2(oldNode, newContent, ctx) {
      if (ctx.ignoreActive && oldNode === document.activeElement) {
        return null;
      }
      if (ctx.callbacks.beforeNodeMorphed(oldNode, newContent) === false) {
        return oldNode;
      }
      if (oldNode instanceof HTMLHeadElement && ctx.head.ignore) {
      } else if (oldNode instanceof HTMLHeadElement && ctx.head.style !== "morph") {
        handleHeadElement(
          oldNode,
          /** @type {HTMLHeadElement} */
          newContent,
          ctx
        );
      } else {
        morphAttributes(oldNode, newContent, ctx);
        if (!ignoreValueOfActiveElement(oldNode, ctx)) {
          morphChildren(ctx, oldNode, newContent);
        }
      }
      ctx.callbacks.afterNodeMorphed(oldNode, newContent);
      return oldNode;
    }
    function morphAttributes(oldNode, newNode, ctx) {
      let type = newNode.nodeType;
      if (type === 1) {
        const oldElt = (
          /** @type {Element} */
          oldNode
        );
        const newElt = (
          /** @type {Element} */
          newNode
        );
        const oldAttributes = oldElt.attributes;
        const newAttributes = newElt.attributes;
        for (const newAttribute of newAttributes) {
          if (ignoreAttribute(newAttribute.name, oldElt, "update", ctx)) {
            continue;
          }
          if (oldElt.getAttribute(newAttribute.name) !== newAttribute.value) {
            oldElt.setAttribute(newAttribute.name, newAttribute.value);
          }
        }
        for (let i = oldAttributes.length - 1; 0 <= i; i--) {
          const oldAttribute = oldAttributes[i];
          if (!oldAttribute) continue;
          if (!newElt.hasAttribute(oldAttribute.name)) {
            if (ignoreAttribute(oldAttribute.name, oldElt, "remove", ctx)) {
              continue;
            }
            oldElt.removeAttribute(oldAttribute.name);
          }
        }
        if (!ignoreValueOfActiveElement(oldElt, ctx)) {
          syncInputValue(oldElt, newElt, ctx);
        }
      }
      if (type === 8 || type === 3) {
        if (oldNode.nodeValue !== newNode.nodeValue) {
          oldNode.nodeValue = newNode.nodeValue;
        }
      }
    }
    function syncInputValue(oldElement, newElement, ctx) {
      if (oldElement instanceof HTMLInputElement && newElement instanceof HTMLInputElement && newElement.type !== "file") {
        let newValue = newElement.value;
        let oldValue = oldElement.value;
        syncBooleanAttribute(oldElement, newElement, "checked", ctx);
        syncBooleanAttribute(oldElement, newElement, "disabled", ctx);
        if (!newElement.hasAttribute("value")) {
          if (!ignoreAttribute("value", oldElement, "remove", ctx)) {
            oldElement.value = "";
            oldElement.removeAttribute("value");
          }
        } else if (oldValue !== newValue) {
          if (!ignoreAttribute("value", oldElement, "update", ctx)) {
            oldElement.setAttribute("value", newValue);
            oldElement.value = newValue;
          }
        }
      } else if (oldElement instanceof HTMLOptionElement && newElement instanceof HTMLOptionElement) {
        syncBooleanAttribute(oldElement, newElement, "selected", ctx);
      } else if (oldElement instanceof HTMLTextAreaElement && newElement instanceof HTMLTextAreaElement) {
        let newValue = newElement.value;
        let oldValue = oldElement.value;
        if (ignoreAttribute("value", oldElement, "update", ctx)) {
          return;
        }
        if (newValue !== oldValue) {
          oldElement.value = newValue;
        }
        if (oldElement.firstChild && oldElement.firstChild.nodeValue !== newValue) {
          oldElement.firstChild.nodeValue = newValue;
        }
      }
    }
    function syncBooleanAttribute(oldElement, newElement, attributeName, ctx) {
      const newLiveValue = newElement[attributeName], oldLiveValue = oldElement[attributeName];
      if (newLiveValue !== oldLiveValue) {
        const ignoreUpdate = ignoreAttribute(
          attributeName,
          oldElement,
          "update",
          ctx
        );
        if (!ignoreUpdate) {
          oldElement[attributeName] = newElement[attributeName];
        }
        if (newLiveValue) {
          if (!ignoreUpdate) {
            oldElement.setAttribute(attributeName, "");
          }
        } else {
          if (!ignoreAttribute(attributeName, oldElement, "remove", ctx)) {
            oldElement.removeAttribute(attributeName);
          }
        }
      }
    }
    function ignoreAttribute(attr, element, updateType, ctx) {
      if (attr === "value" && ctx.ignoreActiveValue && element === document.activeElement) {
        return true;
      }
      return ctx.callbacks.beforeAttributeUpdated(attr, element, updateType) === false;
    }
    function ignoreValueOfActiveElement(possibleActiveElement, ctx) {
      return !!ctx.ignoreActiveValue && possibleActiveElement === document.activeElement && possibleActiveElement !== document.body;
    }
    return morphNode2;
  }();
  function withHeadBlocking(ctx, oldNode, newNode, callback) {
    if (ctx.head.block) {
      const oldHead = oldNode.querySelector("head");
      const newHead = newNode.querySelector("head");
      if (oldHead && newHead) {
        const promises = handleHeadElement(oldHead, newHead, ctx);
        return Promise.all(promises).then(() => {
          const newCtx = Object.assign(ctx, {
            head: {
              block: false,
              ignore: true
            }
          });
          return callback(newCtx);
        });
      }
    }
    return callback(ctx);
  }
  function handleHeadElement(oldHead, newHead, ctx) {
    let added = [];
    let removed = [];
    let preserved = [];
    let nodesToAppend = [];
    let srcToNewHeadNodes = /* @__PURE__ */ new Map();
    for (const newHeadChild of newHead.children) {
      srcToNewHeadNodes.set(newHeadChild.outerHTML, newHeadChild);
    }
    for (const currentHeadElt of oldHead.children) {
      let inNewContent = srcToNewHeadNodes.has(currentHeadElt.outerHTML);
      let isReAppended = ctx.head.shouldReAppend(currentHeadElt);
      let isPreserved = ctx.head.shouldPreserve(currentHeadElt);
      if (inNewContent || isPreserved) {
        if (isReAppended) {
          removed.push(currentHeadElt);
        } else {
          srcToNewHeadNodes.delete(currentHeadElt.outerHTML);
          preserved.push(currentHeadElt);
        }
      } else {
        if (ctx.head.style === "append") {
          if (isReAppended) {
            removed.push(currentHeadElt);
            nodesToAppend.push(currentHeadElt);
          }
        } else {
          if (ctx.head.shouldRemove(currentHeadElt) !== false) {
            removed.push(currentHeadElt);
          }
        }
      }
    }
    nodesToAppend.push(...srcToNewHeadNodes.values());
    let promises = [];
    for (const newNode of nodesToAppend) {
      let newElt = (
        /** @type {ChildNode} */
        document.createRange().createContextualFragment(newNode.outerHTML).firstChild
      );
      if (ctx.callbacks.beforeNodeAdded(newElt) !== false) {
        if ("href" in newElt && newElt.href || "src" in newElt && newElt.src) {
          let resolve;
          let promise = new Promise(function(_resolve) {
            resolve = _resolve;
          });
          newElt.addEventListener("load", function() {
            resolve();
          });
          promises.push(promise);
        }
        oldHead.appendChild(newElt);
        ctx.callbacks.afterNodeAdded(newElt);
        added.push(newElt);
      }
    }
    for (const removedElement of removed) {
      if (ctx.callbacks.beforeNodeRemoved(removedElement) !== false) {
        oldHead.removeChild(removedElement);
        ctx.callbacks.afterNodeRemoved(removedElement);
      }
    }
    ctx.head.afterHeadMorphed(oldHead, {
      added,
      kept: preserved,
      removed
    });
    return promises;
  }
  const createMorphContext = /* @__PURE__ */ function() {
    function createMorphContext2(oldNode, newContent, config) {
      const { persistentIds, idMap } = createIdMaps(oldNode, newContent);
      const mergedConfig = mergeDefaults(config);
      const morphStyle = mergedConfig.morphStyle || "outerHTML";
      if (!["innerHTML", "outerHTML"].includes(morphStyle)) {
        throw `Do not understand how to morph style ${morphStyle}`;
      }
      return {
        target: oldNode,
        newContent,
        config: mergedConfig,
        morphStyle,
        ignoreActive: mergedConfig.ignoreActive,
        ignoreActiveValue: mergedConfig.ignoreActiveValue,
        restoreFocus: mergedConfig.restoreFocus,
        idMap,
        persistentIds,
        pantry: createPantry(),
        callbacks: mergedConfig.callbacks,
        head: mergedConfig.head
      };
    }
    function mergeDefaults(config) {
      let finalConfig = Object.assign({}, defaults);
      Object.assign(finalConfig, config);
      finalConfig.callbacks = Object.assign(
        {},
        defaults.callbacks,
        config.callbacks
      );
      finalConfig.head = Object.assign({}, defaults.head, config.head);
      return finalConfig;
    }
    function createPantry() {
      const pantry = document.createElement("div");
      pantry.hidden = true;
      document.body.insertAdjacentElement("afterend", pantry);
      return pantry;
    }
    function findIdElements(root) {
      let elements = Array.from(root.querySelectorAll("[id]"));
      if (root.id) {
        elements.push(root);
      }
      return elements;
    }
    function populateIdMapWithTree(idMap, persistentIds, root, elements) {
      for (const elt of elements) {
        if (persistentIds.has(elt.id)) {
          let current = elt;
          while (current) {
            let idSet = idMap.get(current);
            if (idSet == null) {
              idSet = /* @__PURE__ */ new Set();
              idMap.set(current, idSet);
            }
            idSet.add(elt.id);
            if (current === root) break;
            current = current.parentElement;
          }
        }
      }
    }
    function createIdMaps(oldContent, newContent) {
      const oldIdElements = findIdElements(oldContent);
      const newIdElements = findIdElements(newContent);
      const persistentIds = createPersistentIds(oldIdElements, newIdElements);
      let idMap = /* @__PURE__ */ new Map();
      populateIdMapWithTree(idMap, persistentIds, oldContent, oldIdElements);
      const newRoot = newContent.__idiomorphRoot || newContent;
      populateIdMapWithTree(idMap, persistentIds, newRoot, newIdElements);
      return { persistentIds, idMap };
    }
    function createPersistentIds(oldIdElements, newIdElements) {
      let duplicateIds = /* @__PURE__ */ new Set();
      let oldIdTagNameMap = /* @__PURE__ */ new Map();
      for (const { id, tagName } of oldIdElements) {
        if (oldIdTagNameMap.has(id)) {
          duplicateIds.add(id);
        } else {
          oldIdTagNameMap.set(id, tagName);
        }
      }
      let persistentIds = /* @__PURE__ */ new Set();
      for (const { id, tagName } of newIdElements) {
        if (persistentIds.has(id)) {
          duplicateIds.add(id);
        } else if (oldIdTagNameMap.get(id) === tagName) {
          persistentIds.add(id);
        }
      }
      for (const id of duplicateIds) {
        persistentIds.delete(id);
      }
      return persistentIds;
    }
    return createMorphContext2;
  }();
  const { normalizeElement, normalizeParent } = /* @__PURE__ */ function() {
    const generatedByIdiomorph = /* @__PURE__ */ new WeakSet();
    function normalizeElement2(content) {
      if (content instanceof Document) {
        return content.documentElement;
      } else {
        return content;
      }
    }
    function normalizeParent2(newContent) {
      if (newContent == null) {
        return document.createElement("div");
      } else if (typeof newContent === "string") {
        return normalizeParent2(parseContent(newContent));
      } else if (generatedByIdiomorph.has(
        /** @type {Element} */
        newContent
      )) {
        return (
          /** @type {Element} */
          newContent
        );
      } else if (newContent instanceof Node) {
        if (newContent.parentNode) {
          return (
            /** @type {any} */
            new SlicedParentNode(newContent)
          );
        } else {
          const dummyParent = document.createElement("div");
          dummyParent.append(newContent);
          return dummyParent;
        }
      } else {
        const dummyParent = document.createElement("div");
        for (const elt of [...newContent]) {
          dummyParent.append(elt);
        }
        return dummyParent;
      }
    }
    class SlicedParentNode {
      /** @param {Node} node */
      constructor(node) {
        this.originalNode = node;
        this.realParentNode = /** @type {Element} */
        node.parentNode;
        this.previousSibling = node.previousSibling;
        this.nextSibling = node.nextSibling;
      }
      /** @returns {Node[]} */
      get childNodes() {
        const nodes = [];
        let cursor = this.previousSibling ? this.previousSibling.nextSibling : this.realParentNode.firstChild;
        while (cursor && cursor != this.nextSibling) {
          nodes.push(cursor);
          cursor = cursor.nextSibling;
        }
        return nodes;
      }
      /**
       * @param {string} selector
       * @returns {Element[]}
       */
      querySelectorAll(selector) {
        return this.childNodes.reduce(
          (results, node) => {
            if (node instanceof Element) {
              if (node.matches(selector)) results.push(node);
              const nodeList = node.querySelectorAll(selector);
              for (let i = 0; i < nodeList.length; i++) {
                results.push(nodeList[i]);
              }
            }
            return results;
          },
          /** @type {Element[]} */
          []
        );
      }
      /**
       * @param {Node} node
       * @param {Node} referenceNode
       * @returns {Node}
       */
      insertBefore(node, referenceNode) {
        return this.realParentNode.insertBefore(node, referenceNode);
      }
      /**
       * @param {Node} node
       * @param {Node} referenceNode
       * @returns {Node}
       */
      moveBefore(node, referenceNode) {
        return this.realParentNode.moveBefore(node, referenceNode);
      }
      /**
       * for later use with populateIdMapWithTree to halt upwards iteration
       * @returns {Node}
       */
      get __idiomorphRoot() {
        return this.originalNode;
      }
    }
    function parseContent(newContent) {
      let parser = new DOMParser();
      let contentWithSvgsRemoved = newContent.replace(
        /<svg(\s[^>]*>|>)([\s\S]*?)<\/svg>/gim,
        ""
      );
      if (contentWithSvgsRemoved.match(/<\/html>/) || contentWithSvgsRemoved.match(/<\/head>/) || contentWithSvgsRemoved.match(/<\/body>/)) {
        let content = parser.parseFromString(newContent, "text/html");
        if (contentWithSvgsRemoved.match(/<\/html>/)) {
          generatedByIdiomorph.add(content);
          return content;
        } else {
          let htmlElement = content.firstChild;
          if (htmlElement) {
            generatedByIdiomorph.add(htmlElement);
          }
          return htmlElement;
        }
      } else {
        let responseDoc = parser.parseFromString(
          "<body><template>" + newContent + "</template></body>",
          "text/html"
        );
        let content = (
          /** @type {HTMLTemplateElement} */
          responseDoc.body.querySelector("template").content
        );
        generatedByIdiomorph.add(content);
        return content;
      }
    }
    return { normalizeElement: normalizeElement2, normalizeParent: normalizeParent2 };
  }();
  return {
    morph,
    defaults
  };
}();

// library/src/plugins/official/backend/watchers/mergeFragments.ts
var MergeFragments = {
  type: 2 /* Watcher */,
  name: EventTypes.MergeFragments,
  onGlobalInit: async (ctx) => {
    const fragmentContainer = document.createElement("template");
    datastarSSEEventWatcher(
      EventTypes.MergeFragments,
      ({
        fragments: fragmentsRaw = "<div></div>",
        selector = "",
        mergeMode = DefaultFragmentMergeMode,
        settleDuration: settleDurationRaw = `${DefaultFragmentsSettleDurationMs}`,
        useViewTransition: useViewTransitionRaw = `${DefaultFragmentsUseViewTransitions}`
      }) => {
        const settleDuration = Number.parseInt(settleDurationRaw);
        const useViewTransition = isBoolString(useViewTransitionRaw);
        fragmentContainer.innerHTML = fragmentsRaw.trim();
        const fragments = [...fragmentContainer.content.children];
        for (const fragment of fragments) {
          if (!(fragment instanceof Element)) {
            throw initErr("NoFragmentsFound", ctx);
          }
          const selectorOrID = selector || `#${fragment.getAttribute("id")}`;
          const targets = [...document.querySelectorAll(selectorOrID) || []];
          if (!targets.length) {
            throw initErr("NoTargetsFound", ctx, { selectorOrID });
          }
          if (useViewTransition && supportsViewTransitions) {
            docWithViewTransitionAPI.startViewTransition(
              () => applyToTargets(ctx, mergeMode, settleDuration, fragment, targets)
            );
          } else {
            applyToTargets(ctx, mergeMode, settleDuration, fragment, targets);
          }
        }
      }
    );
  }
};
function applyToTargets(ctx, mergeMode, settleDuration, fragment, capturedTargets) {
  for (const initialTarget of capturedTargets) {
    initialTarget.classList.add(SWAPPING_CLASS);
    const originalHTML = initialTarget.outerHTML;
    const modifiedTarget = initialTarget;
    switch (mergeMode) {
      case FragmentMergeModes.Morph: {
        const fragmentWithIDs = fragment.cloneNode(true);
        walkDOM(fragmentWithIDs, (el) => {
          if (!el.id?.length && Object.keys(el.dataset).length) {
            el.id = elUniqId(el);
          }
          const elTracking = ctx.removals.get(el.id);
          if (elTracking) {
            const newElTracking = /* @__PURE__ */ new Map();
            for (const [key, cleanup] of elTracking) {
              const newKey = attrHash(key, key);
              newElTracking.set(newKey, cleanup);
              elTracking.delete(key);
            }
            ctx.removals.set(el.id, newElTracking);
          }
        });
        Idiomorph.morph(modifiedTarget, fragmentWithIDs);
        break;
      }
      case FragmentMergeModes.Inner:
        modifiedTarget.innerHTML = fragment.outerHTML;
        break;
      case FragmentMergeModes.Outer:
        modifiedTarget.replaceWith(fragment);
        break;
      case FragmentMergeModes.Prepend:
        modifiedTarget.prepend(fragment);
        break;
      case FragmentMergeModes.Append:
        modifiedTarget.append(fragment);
        break;
      case FragmentMergeModes.Before:
        modifiedTarget.before(fragment);
        break;
      case FragmentMergeModes.After:
        modifiedTarget.after(fragment);
        break;
      case FragmentMergeModes.UpsertAttributes:
        for (const attrName of fragment.getAttributeNames()) {
          const value = fragment.getAttribute(attrName);
          modifiedTarget.setAttribute(attrName, value);
        }
        break;
      default:
        throw initErr("InvalidMergeMode", ctx, { mergeMode });
    }
    const cl = modifiedTarget.classList;
    cl?.add(SWAPPING_CLASS);
    setTimeout(() => {
      initialTarget.classList.remove(SWAPPING_CLASS);
      cl?.remove(SWAPPING_CLASS);
    }, settleDuration);
    const revisedHTML = modifiedTarget.outerHTML;
    if (cl && originalHTML !== revisedHTML) {
      cl.add(SETTLING_CLASS);
      setTimeout(() => {
        cl.remove(SETTLING_CLASS);
      }, settleDuration);
    }
  }
}

// library/src/plugins/official/backend/watchers/mergeSignals.ts
var MergeSignals = {
  type: 2 /* Watcher */,
  name: EventTypes.MergeSignals,
  onGlobalInit: async (ctx) => {
    datastarSSEEventWatcher(
      EventTypes.MergeSignals,
      ({
        signals: raw = "{}",
        onlyIfMissing: onlyIfMissingRaw = `${DefaultMergeSignalsOnlyIfMissing}`
      }) => {
        const { signals: signals2 } = ctx;
        const onlyIfMissing = isBoolString(onlyIfMissingRaw);
        signals2.merge(jsStrToObject(raw), onlyIfMissing);
      }
    );
  }
};

// library/src/plugins/official/backend/watchers/removeFragments.ts
var RemoveFragments = {
  type: 2 /* Watcher */,
  name: EventTypes.RemoveFragments,
  onGlobalInit: async (ctx) => {
    datastarSSEEventWatcher(
      EventTypes.RemoveFragments,
      ({
        selector,
        settleDuration: settleDurationRaw = `${DefaultFragmentsSettleDurationMs}`,
        useViewTransition: useViewTransitionRaw = `${DefaultFragmentsUseViewTransitions}`
      }) => {
        if (!selector.length) {
          throw initErr("NoSelectorProvided", ctx);
        }
        const settleDuration = Number.parseInt(settleDurationRaw);
        const useViewTransition = isBoolString(useViewTransitionRaw);
        const removeTargets = document.querySelectorAll(selector);
        const applyToTargets2 = () => {
          for (const target of removeTargets) {
            target.classList.add(SWAPPING_CLASS);
          }
          setTimeout(() => {
            for (const target of removeTargets) {
              target.remove();
            }
          }, settleDuration);
        };
        if (useViewTransition && supportsViewTransitions) {
          docWithViewTransitionAPI.startViewTransition(() => applyToTargets2());
        } else {
          applyToTargets2();
        }
      }
    );
  }
};

// library/src/plugins/official/backend/watchers/removeSignals.ts
var RemoveSignals = {
  type: 2 /* Watcher */,
  name: EventTypes.RemoveSignals,
  onGlobalInit: async (ctx) => {
    datastarSSEEventWatcher(
      EventTypes.RemoveSignals,
      ({ paths: pathsRaw = "" }) => {
        const paths = pathsRaw.split("\n").map((p) => p.trim());
        if (!paths?.length) {
          throw initErr("NoPathsProvided", ctx);
        }
        ctx.signals.remove(...paths);
      }
    );
  }
};

// library/src/plugins/official/browser/actions/clipboard.ts
var Clipboard = {
  type: 3 /* Action */,
  name: "clipboard",
  fn: (ctx, text) => {
    if (!navigator.clipboard) {
      throw runtimeErr("ClipboardNotAvailable", ctx);
    }
    navigator.clipboard.writeText(text);
  }
};

// library/src/plugins/official/browser/attributes/customValidity.ts
var CustomValidity = {
  type: 1 /* Attribute */,
  name: "customValidity",
  keyReq: 2 /* Denied */,
  valReq: 1 /* Must */,
  onLoad: (ctx) => {
    const { el, genRX: genRX2, effect: effect2 } = ctx;
    if (!(el instanceof HTMLInputElement || el instanceof HTMLSelectElement || el instanceof HTMLTextAreaElement)) {
      throw runtimeErr("CustomValidityInvalidElement", ctx);
    }
    const rx = genRX2();
    return effect2(() => {
      const result = rx();
      if (typeof result !== "string") {
        throw runtimeErr("CustomValidityInvalidExpression", ctx, { result });
      }
      el.setCustomValidity(result);
    });
  }
};

// library/src/plugins/official/browser/attributes/intersects.ts
var ONCE = "once";
var HALF = "half";
var FULL = "full";
var Intersects = {
  type: 1 /* Attribute */,
  name: "intersects",
  keyReq: 2 /* Denied */,
  mods: /* @__PURE__ */ new Set([ONCE, HALF, FULL]),
  onLoad: ({ el, rawKey, mods, genRX: genRX2 }) => {
    const options = { threshold: 0 };
    if (mods.has(FULL)) options.threshold = 1;
    else if (mods.has(HALF)) options.threshold = 0.5;
    const rx = genRX2();
    const observer = new IntersectionObserver((entries) => {
      for (const entry of entries) {
        if (entry.isIntersecting) {
          rx();
          if (mods.has(ONCE)) {
            observer.disconnect();
            delete el.dataset[rawKey];
          }
        }
      }
    }, options);
    observer.observe(el);
    return () => observer.disconnect();
  }
};

// library/src/plugins/official/browser/attributes/persist.ts
var SESSION = "session";
var Persist = {
  type: 1 /* Attribute */,
  name: "persist",
  mods: /* @__PURE__ */ new Set([SESSION]),
  onLoad: ({ key, effect: effect2, mods, signals: signals2, value }) => {
    key = modifyCasing(key, mods);
    if (key === "") {
      key = DATASTAR;
    }
    const storage = mods.has(SESSION) ? sessionStorage : localStorage;
    let paths = value.split(/\s+/).filter((p) => p !== "");
    paths = paths.map((p) => trimDollarSignPrefix(p));
    const storageToSignals = () => {
      const data = storage.getItem(key) || "{}";
      const nestedValues2 = JSON.parse(data);
      signals2.merge(nestedValues2);
    };
    const signalsToStorage = () => {
      let nv;
      if (!paths.length) {
        nv = signals2.values();
      } else {
        nv = signals2.subset(...paths);
      }
      storage.setItem(key, JSON.stringify(nv));
    };
    storageToSignals();
    return effect2(() => {
      signalsToStorage();
    });
  }
};

// library/src/plugins/official/browser/attributes/replaceUrl.ts
var ReplaceUrl = {
  type: 1 /* Attribute */,
  name: "replaceUrl",
  keyReq: 2 /* Denied */,
  valReq: 1 /* Must */,
  onLoad: ({ effect: effect2, genRX: genRX2 }) => {
    const rx = genRX2();
    return effect2(() => {
      const url2 = rx();
      const baseUrl = window.location.href;
      const fullUrl = new URL(url2, baseUrl).toString();
      window.history.replaceState({}, "", fullUrl);
    });
  }
};

// library/src/plugins/official/browser/attributes/scrollIntoView.ts
var SMOOTH = "smooth";
var INSTANT = "instant";
var AUTO = "auto";
var HSTART = "hstart";
var HCENTER = "hcenter";
var HEND = "hend";
var HNEAREST = "hnearest";
var VSTART = "vstart";
var VCENTER = "vcenter";
var VEND = "vend";
var VNEAREST = "vnearest";
var FOCUS = "focus";
var CENTER = "center";
var START = "start";
var END = "end";
var NEAREST = "nearest";
var ScrollIntoView = {
  type: 1 /* Attribute */,
  name: "scrollIntoView",
  keyReq: 2 /* Denied */,
  valReq: 2 /* Denied */,
  mods: /* @__PURE__ */ new Set([
    SMOOTH,
    INSTANT,
    AUTO,
    HSTART,
    HCENTER,
    HEND,
    HNEAREST,
    VSTART,
    VCENTER,
    VEND,
    VNEAREST,
    FOCUS
  ]),
  onLoad: (ctx) => {
    const { el, mods, rawKey } = ctx;
    if (!el.tabIndex) el.setAttribute("tabindex", "0");
    const opts = {
      behavior: SMOOTH,
      block: CENTER,
      inline: CENTER
    };
    if (mods.has(SMOOTH)) opts.behavior = SMOOTH;
    if (mods.has(INSTANT)) opts.behavior = INSTANT;
    if (mods.has(AUTO)) opts.behavior = AUTO;
    if (mods.has(HSTART)) opts.inline = START;
    if (mods.has(HCENTER)) opts.inline = CENTER;
    if (mods.has(HEND)) opts.inline = END;
    if (mods.has(HNEAREST)) opts.inline = NEAREST;
    if (mods.has(VSTART)) opts.block = START;
    if (mods.has(VCENTER)) opts.block = CENTER;
    if (mods.has(VEND)) opts.block = END;
    if (mods.has(VNEAREST)) opts.block = NEAREST;
    if (!(el instanceof HTMLElement || el instanceof SVGElement)) {
      throw runtimeErr("ScrollIntoViewInvalidElement", ctx);
    }
    if (!el.tabIndex) {
      el.setAttribute("tabindex", "0");
    }
    el.scrollIntoView(opts);
    if (mods.has("focus")) {
      el.focus();
    }
    delete el.dataset[rawKey];
  }
};

// library/src/plugins/official/browser/attributes/show.ts
var NONE = "none";
var DISPLAY = "display";
var Show = {
  type: 1 /* Attribute */,
  name: "show",
  keyReq: 2 /* Denied */,
  valReq: 1 /* Must */,
  onLoad: ({ el: { style: s }, genRX: genRX2, effect: effect2 }) => {
    const rx = genRX2();
    return effect2(async () => {
      const shouldShow = rx();
      if (shouldShow) {
        if (s.display === NONE) {
          s.removeProperty(DISPLAY);
        }
      } else {
        s.setProperty(DISPLAY, NONE);
      }
    });
  }
};

// library/src/plugins/official/browser/attributes/viewTransition.ts
var VIEW_TRANSITION = "view-transition";
var ViewTransition = {
  type: 1 /* Attribute */,
  name: "viewTransition",
  keyReq: 2 /* Denied */,
  valReq: 1 /* Must */,
  onGlobalInit() {
    let hasViewTransitionMeta = false;
    for (const node of document.head.childNodes) {
      if (node instanceof HTMLMetaElement && node.name === VIEW_TRANSITION) {
        hasViewTransitionMeta = true;
      }
    }
    if (!hasViewTransitionMeta) {
      const meta = document.createElement("meta");
      meta.name = VIEW_TRANSITION;
      meta.content = "same-origin";
      document.head.appendChild(meta);
    }
  },
  onLoad: ({ effect: effect2, el, genRX: genRX2 }) => {
    if (!supportsViewTransitions) {
      console.error("Browser does not support view transitions");
      return;
    }
    const rx = genRX2();
    return effect2(() => {
      const name2 = rx();
      if (!name2?.length) return;
      const elVTASTyle = el.style;
      elVTASTyle.viewTransitionName = name2;
    });
  }
};

// library/src/plugins/official/dom/attributes/attr.ts
var Attr = {
  type: 1 /* Attribute */,
  name: "attr",
  valReq: 1 /* Must */,
  onLoad: ({ el, key, effect: effect2, genRX: genRX2 }) => {
    const rx = genRX2();
    if (key === "") {
      return effect2(async () => {
        const binds = rx();
        for (const [key2, val] of Object.entries(binds)) {
          if (val === false) {
            el.removeAttribute(key2);
          } else {
            el.setAttribute(key2, val);
          }
        }
      });
    }
    key = kebab(key);
    return effect2(async () => {
      let value = false;
      try {
        value = rx();
      } catch (e) {
      }
      let v;
      if (typeof value === "string") {
        v = value;
      } else {
        v = JSON.stringify(value);
      }
      if (!v || v === "false" || v === "null" || v === "undefined") {
        el.removeAttribute(key);
      } else {
        el.setAttribute(key, v);
      }
    });
  }
};

// library/src/plugins/official/dom/attributes/bind.ts
var dataURIRegex = /^data:(?<mime>[^;]+);base64,(?<contents>.*)$/;
var updateEvents = ["change", "input", "keydown"];
var Bind = {
  type: 1 /* Attribute */,
  name: "bind",
  keyReq: 3 /* Exclusive */,
  valReq: 3 /* Exclusive */,
  onLoad: (ctx) => {
    const { el, key, mods, signals: signals2, value, effect: effect2 } = ctx;
    const input = el;
    const signalName = key ? modifyCasing(key, mods) : trimDollarSignPrefix(value);
    const tnl = el.tagName.toLowerCase();
    const isInput = tnl.includes("input");
    const isSelect = tnl.includes("select");
    const type = el.getAttribute("type");
    const hasValueAttribute = el.hasAttribute("value");
    let signalDefault = "";
    const isCheckbox = isInput && type === "checkbox";
    if (isCheckbox) {
      signalDefault = hasValueAttribute ? "" : false;
    }
    const isNumber = isInput && type === "number";
    if (isNumber) {
      signalDefault = 0;
    }
    const isRadio = isInput && type === "radio";
    if (isRadio) {
      const name2 = el.getAttribute("name");
      if (!name2?.length) {
        el.setAttribute("name", signalName);
      }
    }
    const isFile = isInput && type === "file";
    const { signal, inserted } = signals2.upsertIfMissing(
      signalName,
      signalDefault
    );
    let arrayIndex = -1;
    if (Array.isArray(signal.value)) {
      if (el.getAttribute("name") === null) {
        el.setAttribute("name", signalName);
      }
      arrayIndex = [
        ...document.querySelectorAll(`[name="${signalName}"]`)
      ].findIndex((el2) => el2 === ctx.el);
    }
    const isArray = arrayIndex >= 0;
    const signalArray = () => [...signals2.value(signalName)];
    const setElementFromSignal = () => {
      let value2 = signals2.value(signalName);
      if (isArray && !isSelect) {
        value2 = value2[arrayIndex] || signalDefault;
      }
      const stringValue = `${value2}`;
      if (isCheckbox || isRadio) {
        if (typeof value2 === "boolean") {
          input.checked = value2;
        } else {
          input.checked = stringValue === input.value;
        }
      } else if (isSelect) {
        const select = el;
        if (select.multiple) {
          if (!isArray) {
            throw runtimeErr("BindSelectMultiple", ctx);
          }
          for (const opt of select.options) {
            if (opt?.disabled) return;
            const incoming = isNumber ? Number(opt.value) : opt.value;
            opt.selected = value2.includes(incoming);
          }
        } else {
          select.value = stringValue;
        }
      } else if (isFile) {
      } else if ("value" in el) {
        el.value = stringValue;
      } else {
        el.setAttribute("value", stringValue);
      }
    };
    const setSignalFromElement = async () => {
      let currentValue = signals2.value(signalName);
      if (isArray) {
        const currentArray = currentValue;
        while (arrayIndex >= currentArray.length) {
          currentArray.push(signalDefault);
        }
        currentValue = currentArray[arrayIndex] || signalDefault;
      }
      const update = (signalName2, value3) => {
        let newValue2 = value3;
        if (isArray && !isSelect) {
          newValue2 = signalArray();
          newValue2[arrayIndex] = value3;
        }
        signals2.setValue(signalName2, newValue2);
      };
      if (isFile) {
        const files = [...input?.files || []];
        const allContents = [];
        const allMimes = [];
        const allNames = [];
        await Promise.all(
          files.map((f) => {
            return new Promise((resolve) => {
              const reader = new FileReader();
              reader.onload = () => {
                if (typeof reader.result !== "string") {
                  throw runtimeErr("InvalidFileResultType", ctx, {
                    resultType: typeof reader.result
                  });
                }
                const match = reader.result.match(dataURIRegex);
                if (!match?.groups) {
                  throw runtimeErr("InvalidDataUri", ctx, {
                    result: reader.result
                  });
                }
                allContents.push(match.groups.contents);
                allMimes.push(match.groups.mime);
                allNames.push(f.name);
              };
              reader.onloadend = () => resolve(void 0);
              reader.readAsDataURL(f);
            });
          })
        );
        update(signalName, allContents);
        update(`${signalName}Mimes`, allMimes);
        update(`${signalName}Names`, allNames);
        return;
      }
      const value2 = input.value || "";
      let newValue;
      if (isCheckbox) {
        const checked = input.checked || input.getAttribute("checked") === "true";
        if (hasValueAttribute) {
          newValue = checked ? value2 : "";
        } else {
          newValue = checked;
        }
      } else if (isSelect) {
        const select = el;
        const selectedOptions = [...select.selectedOptions];
        if (isArray) {
          newValue = selectedOptions.filter((opt) => opt.selected).map((opt) => opt.value);
        } else {
          newValue = selectedOptions[0]?.value || signalDefault;
        }
      } else if (typeof currentValue === "boolean") {
        newValue = Boolean(value2);
      } else if (typeof currentValue === "number") {
        newValue = Number(value2);
      } else {
        newValue = value2 || "";
      }
      update(signalName, newValue);
    };
    if (inserted) {
      setSignalFromElement();
    }
    for (const event of updateEvents) {
      el.addEventListener(event, setSignalFromElement);
    }
    const onPageshow = (ev) => {
      if (!ev.persisted) return;
      setSignalFromElement();
    };
    window.addEventListener("pageshow", onPageshow);
    const reset = effect2(() => setElementFromSignal());
    return () => {
      reset();
      for (const event of updateEvents) {
        el.removeEventListener(event, setSignalFromElement);
      }
      window.removeEventListener("pageshow", onPageshow);
    };
  }
};

// library/src/plugins/official/dom/attributes/class.ts
var Class = {
  type: 1 /* Attribute */,
  name: "class",
  valReq: 1 /* Must */,
  onLoad: ({ el, key, mods, effect: effect2, genRX: genRX2 }) => {
    const cl = el.classList;
    const rx = genRX2();
    return effect2(() => {
      if (key === "") {
        const classes = rx();
        for (const [k, v] of Object.entries(classes)) {
          const classNames = k.split(/\s+/);
          if (v) {
            cl.add(...classNames);
          } else {
            cl.remove(...classNames);
          }
        }
      } else {
        key = modifyCasing(key, mods);
        const shouldInclude = rx();
        if (shouldInclude) {
          cl.add(key);
        } else {
          cl.remove(key);
        }
      }
    });
  }
};

// library/src/utils/tags.ts
function tagToMs(args) {
  if (!args || args.size <= 0) return 0;
  for (const arg of args) {
    if (arg.endsWith("ms")) {
      return Number(arg.replace("ms", ""));
    }
    if (arg.endsWith("s")) {
      return Number(arg.replace("s", "")) * 1e3;
    }
    try {
      return Number.parseFloat(arg);
    } catch (e) {
    }
  }
  return 0;
}
function tagHas(tags, tag, defaultValue = false) {
  if (!tags) return defaultValue;
  return tags.has(tag.toLowerCase());
}

// library/src/utils/timing.ts
function delay(callback, wait) {
  return (...args) => {
    setTimeout(() => {
      callback(...args);
    }, wait);
  };
}
function debounce(callback, wait, leading = false, trailing = true) {
  let timer = -1;
  const resetTimer = () => timer && clearTimeout(timer);
  return (...args) => {
    resetTimer();
    if (leading && !timer) {
      callback(...args);
    }
    timer = setTimeout(() => {
      if (trailing) {
        callback(...args);
      }
      resetTimer();
    }, wait);
  };
}
function throttle(callback, wait, leading = true, trailing = false) {
  let waiting = false;
  return (...args) => {
    if (waiting) return;
    if (leading) {
      callback(...args);
    }
    waiting = true;
    setTimeout(() => {
      waiting = false;
      if (trailing) {
        callback(...args);
      }
    }, wait);
  };
}

// library/src/plugins/official/dom/attributes/on.ts
var EVT = "evt";
var SIGNALS_CHANGE_PREFIX = "signalsChange";
var signalChangeKeyLength = SIGNALS_CHANGE_PREFIX.length;
var On = {
  type: 1 /* Attribute */,
  name: "on",
  keyReq: 1 /* Must */,
  valReq: 1 /* Must */,
  argNames: [EVT],
  onLoad: ({ el, key, mods, signals: signals2, effect: effect2, genRX: genRX2 }) => {
    const rx = genRX2();
    let target = el;
    if (mods.has("window")) target = window;
    let callback = (evt) => {
      if (evt) {
        if (mods.has("prevent") || key === "submit") evt.preventDefault();
        if (mods.has("stop")) evt.stopPropagation();
      }
      rx(evt);
    };
    const delayArgs = mods.get("delay");
    if (delayArgs) {
      const wait = tagToMs(delayArgs);
      callback = delay(callback, wait);
    }
    const debounceArgs = mods.get("debounce");
    if (debounceArgs) {
      const wait = tagToMs(debounceArgs);
      const leading = tagHas(debounceArgs, "leading", false);
      const trailing = !tagHas(debounceArgs, "notrail", false);
      callback = debounce(callback, wait, leading, trailing);
    }
    const throttleArgs = mods.get("throttle");
    if (throttleArgs) {
      const wait = tagToMs(throttleArgs);
      const leading = !tagHas(throttleArgs, "noleading", false);
      const trailing = tagHas(throttleArgs, "trail", false);
      callback = throttle(callback, wait, leading, trailing);
    }
    if (mods.has("viewtransition") && supportsViewTransitions) {
      const cb = callback;
      callback = (...args) => document.startViewTransition(() => cb(...args));
    }
    const evtListOpts = {
      capture: true,
      passive: false,
      once: false
    };
    if (!mods.has("capture")) evtListOpts.capture = false;
    if (mods.has("passive")) evtListOpts.passive = true;
    if (mods.has("once")) evtListOpts.once = true;
    if (key === "load") {
      setTimeout(callback, 0);
      return () => {
      };
    }
    if (key === "interval") {
      let duration = 1e3;
      const durationArgs = mods.get("duration");
      if (durationArgs) {
        duration = tagToMs(durationArgs);
        const leading = tagHas(durationArgs, "leading", false);
        if (leading) {
          callback();
        }
      }
      const intervalId = setInterval(callback, duration);
      return () => {
        clearInterval(intervalId);
      };
    }
    if (key === "raf") {
      let rafId;
      const raf = () => {
        callback();
        rafId = requestAnimationFrame(raf);
      };
      rafId = requestAnimationFrame(raf);
      return () => {
        if (rafId) {
          cancelAnimationFrame(rafId);
        }
      };
    }
    if (key === "resize") {
      let resizeObserver = new ResizeObserver(() => {
        callback();
      });
      resizeObserver.observe(el);
      return () => {
        if (resizeObserver) {
          resizeObserver.disconnect();
          resizeObserver = null;
        }
      };
    }
    if (key.startsWith(SIGNALS_CHANGE_PREFIX)) {
      if (key === SIGNALS_CHANGE_PREFIX) {
        const signalFn = (event) => callback(event);
        document.addEventListener(DATASTAR_SIGNAL_EVENT, signalFn);
        return () => {
          document.removeEventListener(DATASTAR_SIGNAL_EVENT, signalFn);
        };
      }
      const signalPath = modifyCasing(
        camel(key.slice(signalChangeKeyLength)),
        mods
      );
      const signalValues = /* @__PURE__ */ new Map();
      signals2.walk((path, signal) => {
        if (path.startsWith(signalPath)) {
          signalValues.set(signal, signal.value);
        }
      });
      return effect2(() => {
        for (const [signal, prev] of signalValues) {
          if (prev !== signal.value) {
            callback();
            signalValues.set(signal, signal.value);
          }
        }
      });
    }
    const testOutside = mods.has("outside");
    if (testOutside) {
      target = document;
      const cb = callback;
      const targetOutsideCallback = (e) => {
        const targetHTML = e?.target;
        if (!el.contains(targetHTML)) {
          cb(e);
        }
      };
      callback = targetOutsideCallback;
    }
    const eventName = modifyCasing(key, mods);
    target.addEventListener(eventName, callback, evtListOpts);
    return () => {
      target.removeEventListener(eventName, callback);
    };
  }
};

// library/src/plugins/official/dom/attributes/ref.ts
var Ref = {
  type: 1 /* Attribute */,
  name: "ref",
  keyReq: 3 /* Exclusive */,
  valReq: 3 /* Exclusive */,
  onLoad: ({ el, key, mods, signals: signals2, value }) => {
    const signalName = key ? modifyCasing(key, mods) : trimDollarSignPrefix(value);
    signals2.setValue(signalName, el);
  }
};

// library/src/plugins/official/dom/attributes/text.ts
var Text = {
  type: 1 /* Attribute */,
  name: "text",
  keyReq: 2 /* Denied */,
  valReq: 1 /* Must */,
  onLoad: (ctx) => {
    const { el, effect: effect2, genRX: genRX2 } = ctx;
    const rx = genRX2();
    if (!(el instanceof HTMLElement)) {
      runtimeErr("TextInvalidElement", ctx);
    }
    return effect2(() => {
      const res = rx(ctx);
      el.textContent = `${res}`;
    });
  }
};

// library/src/plugins/official/logic/actions/fit.ts
var { round, max, min } = Math;
var Fit = {
  type: 3 /* Action */,
  name: "fit",
  fn: (_, v, oldMin, oldMax, newMin, newMax, shouldClamp = false, shouldRound = false) => {
    let fitted = (v - oldMin) / (oldMax - oldMin) * (newMax - newMin) + newMin;
    if (shouldRound) {
      fitted = round(fitted);
    }
    if (shouldClamp) {
      fitted = max(newMin, min(newMax, fitted));
    }
    return fitted;
  }
};

// library/src/plugins/official/logic/actions/setAll.ts
var SetAll = {
  type: 3 /* Action */,
  name: "setAll",
  fn: ({ signals: signals2 }, prefix, newValue) => {
    signals2.walk((path, signal) => {
      if (!path.startsWith(prefix)) return;
      signal.value = newValue;
    });
  }
};

// library/src/plugins/official/logic/actions/toggleAll.ts
var ToggleAll = {
  type: 3 /* Action */,
  name: "toggleAll",
  fn: ({ signals: signals2 }, prefix) => {
    signals2.walk((path, signal) => {
      if (!path.startsWith(prefix)) return;
      signal.value = !signal.value;
    });
  }
};

// library/src/bundles/datastar.ts
load(
  // DOM
  Attr,
  Bind,
  Class,
  On,
  Ref,
  Show,
  Text,
  // Backend
  Indicator,
  GET,
  POST,
  PUT,
  PATCH,
  DELETE,
  MergeFragments,
  MergeSignals,
  RemoveFragments,
  RemoveSignals,
  ExecuteScript,
  // Browser
  Clipboard,
  CustomValidity,
  Intersects,
  Persist,
  ReplaceUrl,
  ScrollIntoView,
  ViewTransition,
  // Logic
  Fit,
  SetAll,
  ToggleAll
);
apply();
export {
  apply,
  load,
  setAlias
};
//# sourceMappingURL=datastar.js.map
