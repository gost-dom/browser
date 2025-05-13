import { type DatastarPlugin } from './types';
export declare function setAlias(value: string): void;
export declare function load(...pluginsToLoad: DatastarPlugin[]): void;
export declare function apply(): void;
