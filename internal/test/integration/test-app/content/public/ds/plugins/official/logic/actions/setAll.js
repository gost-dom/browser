// Authors: Delaney Gillilan
// Icon: ion:checkmark-round
// Slug: Set all signals that match the signal path
// Description: Set all signals that match one or more space-separated paths in which `*` can be used as a wildcard
import { PluginType } from '../../../../engine/types';
import { getMatchingSignalPaths } from '../../../../utils/paths';
export const SetAll = {
    type: PluginType.Action,
    name: 'setAll',
    fn: ({ signals }, paths, newValue) => {
        const signalPaths = getMatchingSignalPaths(signals, paths);
        for (const path of signalPaths) {
            signals.setValue(path, newValue);
        }
    },
};
