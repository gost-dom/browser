import type { HTMLorSVGElement } from '../engine/types';
export declare class Hash {
    #private;
    constructor(prefix?: string);
    with(x: number | string | boolean): Hash;
    get value(): number;
    get string(): string;
}
export declare function elUniqId(el: Element): string;
export declare function attrHash(key: number | string, val: number | string): number;
export declare function walkDOM(element: Element | null, callback: (el: HTMLorSVGElement) => void): null | undefined;
