import type { RuntimeContext } from '../../../../engine/types';
export type SSEArgs = {
    headers?: Record<string, string>;
    openWhenHidden?: boolean;
    retryInterval?: number;
    retryScaler?: number;
    retryMaxWaitMs?: number;
    retryMaxCount?: number;
    abort?: AbortSignal;
} & ({
    contentType: 'json';
    includeLocal?: boolean;
} | {
    contentType: 'form';
    selector?: string;
});
export declare const sse: (ctx: RuntimeContext, method: string, url: string, args: SSEArgs) => Promise<void>;
