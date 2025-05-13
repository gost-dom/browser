import { Signal, computed } from '../vendored/preact-core';
import { internalErr } from './errors';
import { DATASTAR_SIGNAL_EVENT, } from './types';
const from = 'namespacedSignals';
const dispatchSignalEvent = (evt) => {
    document.dispatchEvent(new CustomEvent(DATASTAR_SIGNAL_EVENT, {
        detail: Object.assign({ added: [], removed: [], updated: [] }, evt),
    }));
};
// If onlyPublic is true, only signals not starting with an underscore are included
function nestedValues(signal, onlyPublic = false) {
    const kv = {};
    for (const key in signal) {
        if (Object.hasOwn(signal, key)) {
            if (onlyPublic && key.startsWith('_')) {
                continue;
            }
            const value = signal[key];
            if (value instanceof Signal) {
                kv[key] = value.value;
            }
            else {
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
        updated: [],
    };
    for (const key in values) {
        if (Object.hasOwn(values, key)) {
            if (key.match(/\_\_+/)) {
                throw internalErr(from, 'InvalidSignalKey', { key });
            }
            const value = values[key];
            if (value instanceof Object && !Array.isArray(value)) {
                if (!target[key]) {
                    target[key] = {};
                }
                const subEvt = mergeNested(target[key], value, onlyIfMissing);
                evt.added.push(...subEvt.added.map((k) => `${key}.${k}`));
                evt.removed.push(...subEvt.removed.map((k) => `${key}.${k}`));
                evt.updated.push(...subEvt.updated.map((k) => `${key}.${k}`));
            }
            else {
                const hasKey = Object.hasOwn(target, key);
                if (hasKey) {
                    if (onlyIfMissing)
                        continue;
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
            }
            else {
                walkNestedSignal(value, (path, value) => {
                    cb(`${key}.${path}`, value);
                });
            }
        }
    }
}
// Recursive function to subset a nested object, each key is a dot-delimited path
function nestedSubset(original, ...keys) {
    const subset = {};
    for (const key of keys) {
        const parts = key.split('.');
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
// Recursively walk a NestedValue with a callback and dot-delimited path
export function walkNestedValues(nv, cb) {
    for (const key in nv) {
        if (Object.hasOwn(nv, key)) {
            const value = nv[key];
            if (value instanceof Object && !Array.isArray(value)) {
                walkNestedValues(value, (path, value) => {
                    cb(`${key}.${path}`, value);
                });
            }
            else {
                cb(key, value);
            }
        }
    }
}
export class SignalsRoot {
    #signals = {};
    exists(dotDelimitedPath) {
        return !!this.signal(dotDelimitedPath);
    }
    signal(dotDelimitedPath) {
        const parts = dotDelimitedPath.split('.');
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
            throw internalErr(from, 'SignalNotFound', { path: dotDelimitedPath });
        return signal;
    }
    setSignal(dotDelimitedPath, signal) {
        const parts = dotDelimitedPath.split('.');
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
        const parts = dotDelimitedPath.split('.');
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
        return { signal: signal, inserted: true };
    }
    remove(...dotDelimitedPaths) {
        if (!dotDelimitedPaths.length) {
            this.#signals = {};
            return;
        }
        const removed = Array();
        for (const path of dotDelimitedPaths) {
            const parts = path.split('.');
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
}
