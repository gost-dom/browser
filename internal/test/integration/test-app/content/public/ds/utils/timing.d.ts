import type { EventCallbackHandler, Modifiers } from '../engine/types';
export declare function debounce(callback: EventCallbackHandler, wait: number, leading?: boolean, trailing?: boolean): EventCallbackHandler;
export declare function throttle(callback: EventCallbackHandler, wait: number, leading?: boolean, trailing?: boolean): EventCallbackHandler;
export declare function modifyTiming(callback: EventCallbackHandler, mods: Modifiers): EventCallbackHandler;
