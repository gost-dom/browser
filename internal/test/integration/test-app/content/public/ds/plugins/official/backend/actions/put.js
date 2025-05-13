// Icon: material-symbols:arrows-input
// Slug: Use a PUT request to fetch data from a server using Server-Sent Events matching the Datastar SDK interface
// Description: Remember, SSE is just a regular SSE request but with the ability to send 0-inf messages to the client.
import { PluginType, } from '../../../../engine/types';
import { sse } from './sse';
export const PUT = {
    type: PluginType.Action,
    name: 'put',
    fn: async (ctx, url, args) => {
        return sse(ctx, 'PUT', url, { ...args });
    },
};
