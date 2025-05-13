// Authors: Delaney Gillilan
// Icon: mdi:clipboard
// Slug: Copy text to the clipboard
// Description: This action copies text to the clipboard using the Clipboard API.
import { runtimeErr } from '../../../../engine/errors';
import { PluginType } from '../../../../engine/types';
export const Clipboard = {
    type: PluginType.Action,
    name: 'clipboard',
    fn: (ctx, text) => {
        if (!navigator.clipboard) {
            throw runtimeErr('ClipboardNotAvailable', ctx);
        }
        navigator.clipboard.writeText(text);
    },
};
