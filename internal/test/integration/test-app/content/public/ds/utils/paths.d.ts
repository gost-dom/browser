import type { SignalsRoot } from '../engine/signals';
export declare function pathMatchesPattern(path: string, pattern: string): boolean;
export declare function getMatchingSignalPaths(signals: SignalsRoot, paths: string): string[];
