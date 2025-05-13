// Icon: material-symbols:delete-outline
// Slug: Use a DELETE request to fetch data from a server using Server-Sent Events matching the Datastar SDK interface
// Description: Remember, SSE is just a regular SSE request but with the ability to send 0-inf messages to the client.
import { PluginType, } from '../../../../engine/types';
import { sse } from './sse';
export const DELETE = {
    type: PluginType.Action,
    name: 'delete',
    fn: async (ctx, url, args) => {
        return sse(ctx, 'DELETE', url, { ...args });
    },
};
