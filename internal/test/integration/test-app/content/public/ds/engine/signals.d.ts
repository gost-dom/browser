import { Signal } from '../vendored/preact-core';
import { type NestedValues } from './types';
export declare function walkNestedValues(nv: NestedValues, cb: (path: string, value: any) => void): void;
export declare class SignalsRoot {
    #private;
    exists(dotDelimitedPath: string): boolean;
    signal<T>(dotDelimitedPath: string): Signal<T> | null;
    setSignal<T extends Signal<T>>(dotDelimitedPath: string, signal: T): void;
    setComputed<T>(dotDelimitedPath: string, fn: () => T): void;
    value<T>(dotDelimitedPath: string): T;
    setValue<T>(dotDelimitedPath: string, value: T): void;
    upsertIfMissing<T>(dotDelimitedPath: string, defaultValue: T): {
        signal: Signal<T>;
        inserted: boolean;
    };
    remove(...dotDelimitedPaths: string[]): void;
    merge(other: NestedValues, onlyIfMissing?: boolean): void;
    subset(...keys: string[]): NestedValues;
    walk(cb: (name: string, signal: Signal<any>) => void): void;
    paths(): string[];
    values(onlyPublic?: boolean): NestedValues;
    JSON(shouldIndent?: boolean, onlyPublic?: boolean): string;
    toString(): string;
}
