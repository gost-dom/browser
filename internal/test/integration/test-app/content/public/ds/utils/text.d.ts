import type { Modifiers } from '../engine/types';
export declare const isBoolString: (str: string) => boolean;
export declare const kebab: (str: string) => string;
export declare const camel: (str: string) => string;
export declare const snake: (str: string) => string;
export declare const pascal: (str: string) => string;
export declare const jsStrToObject: (raw: string) => any;
export declare const trimDollarSignPrefix: (str: string) => string;
export declare function modifyCasing(str: string, mods: Modifiers): string;
