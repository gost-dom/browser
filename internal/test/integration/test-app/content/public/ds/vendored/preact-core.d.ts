import type { OnRemovalFn } from '../engine/types';
declare const BRAND_SYMBOL: unique symbol;
type Node = {
    _source: Signal;
    _prevSource?: Node;
    _nextSource?: Node;
    _target: Computed | Effect;
    _prevTarget?: Node;
    _nextTarget?: Node;
    _version: number;
    _rollbackNode?: Node;
};
/**
 * Combine multiple value updates into one "commit" at the end of the provided callback.
 *
 * Batches can be nested and changes are only flushed once the outermost batch callback
 * completes.
 *
 * Accessing a signal that has been modified within a batch will reflect its updated
 * value.
 *
 * @param fn The callback function.
 * @returns The value returned by the callback.
 */
declare function batch<T>(fn: () => T): T;
/**
 * Run a callback function that can access signal values without
 * subscribing to the signal updates.
 *
 * @param fn The callback function.
 * @returns The value returned by the callback.
 */
declare function untracked<T>(fn: () => T): T;
/**
 * The base class for plain and computed signals.
 */
declare class Signal<T = any> {
    /** @internal */
    _onChange: (change: {
        old: T;
        revised: T;
    }) => void;
    /** @internal */
    _value: unknown;
    /**
     * @internal
     * Version numbers should always be >= 0, because the special value -1 is used
     * by Nodes to signify potentially unused but recyclable nodes.
     */
    _version: number;
    /** @internal */
    _node?: Node;
    /** @internal */
    _targets?: Node;
    constructor(value?: T);
    /** @internal */
    _refresh(): boolean;
    /** @internal */
    _subscribe(node: Node): void;
    /** @internal */
    _unsubscribe(node: Node): void;
    subscribe(fn: (value: T) => void): () => void;
    valueOf(): T;
    toString(): string;
    toJSON(): T;
    peek(): T;
    brand: typeof BRAND_SYMBOL;
    get value(): T;
    set value(value: T);
}
/** @internal */
declare function Signal(this: Signal, value?: unknown): void;
/**
 * Create a new plain signal.
 *
 * @param value The initial value for the signal.
 * @returns A new signal.
 */
export declare function signal<T>(value: T): Signal<T>;
export declare function signal<T = undefined>(): Signal<T | undefined>;
export declare class Computed<T = any> extends Signal<T> {
    _fn: () => T;
    _sources?: Node;
    _globalVersion: number;
    _flags: number;
    constructor(fn: () => T);
    _notify(): void;
    get value(): T;
}
export declare function Computed(this: Computed, fn: () => unknown): void;
export declare namespace Computed {
    var prototype: Computed<any>;
}
/**
 * An interface for read-only signals.
 */
interface ReadonlySignal<T = any> {
    readonly value: T;
    peek(): T;
    subscribe(fn: (value: T) => void): () => void;
    valueOf(): T;
    toString(): string;
    toJSON(): T;
    brand: typeof BRAND_SYMBOL;
}
/**
 * Create a new signal that is computed based on the values of other signals.
 *
 * The returned computed signal is read-only, and its value is automatically
 * updated when any signals accessed from within the callback function change.
 *
 * @param fn The effect callback.
 * @returns A new read-only signal.
 */
declare function computed<T>(fn: () => T): ReadonlySignal<T>;
export type EffectFn = () => OnRemovalFn | void | Promise<OnRemovalFn | void>;
declare class Effect {
    _fn?: EffectFn;
    _cleanup?: () => void;
    _sources?: Node;
    _nextBatchedEffect?: Effect;
    _flags: number;
    constructor(fn: EffectFn);
    _callback(): void;
    _start(): () => void;
    _notify(): void;
    _dispose(): void;
}
declare function Effect(this: Effect, fn: EffectFn): void;
/**
 * Create an effect to run arbitrary code in response to signal changes.
 *
 * An effect tracks which signals are accessed within the given callback
 * function `fn`, and re-runs the callback when those signals change.
 *
 * The callback may return a cleanup function. The cleanup function gets
 * run once, either when the callback is next called or when the effect
 * gets disposed, whichever happens first.
 *
 * @param fn The effect callback.
 * @returns A function for disposing the effect.
 */
declare function effect(fn: EffectFn): () => void;
export { batch, computed, effect, Signal, untracked };
export type { ReadonlySignal };
