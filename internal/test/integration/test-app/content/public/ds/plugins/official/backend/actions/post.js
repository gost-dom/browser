// Icon: ri:signpost-fill
// Slug: Use a POST request to fetch data from a server using Server-Sent Events matching the Datastar SDK interface
// Description: Remember, SSE is just a regular SSE request but with the ability to send 0-inf messages to the client.
import { PluginType, } from '../../../../engine/types';
import { sse } from './sse';
export const POST = {
    type: PluginType.Action,
    name: 'post',
    fn: async (ctx, url, args) => {
        return sse(ctx, 'POST', url, { ...args });
    },
};
