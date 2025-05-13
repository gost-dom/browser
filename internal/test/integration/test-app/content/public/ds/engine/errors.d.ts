import { type InitContext, type RuntimeContext } from './types';
export declare function internalErr(from: string, reason: string, args?: {}): Error;
export declare function initErr(reason: string, ctx: InitContext, metadata?: {}): Error;
export declare function runtimeErr(reason: string, ctx: RuntimeContext, metadata?: {}): Error;
