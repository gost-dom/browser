export const isBoolString = (str) => str.trim() === 'true';
export const kebab = (str) => str.replace(/[A-Z]+(?![a-z])|[A-Z]/g, ($, ofs) => (ofs ? '-' : '') + $.toLowerCase());
export const camel = (str) => kebab(str).replace(/-./g, (x) => x[1].toUpperCase());
export const snake = (str) => kebab(str).replace(/-/g, '_');
export const pascal = (str) => camel(str).replace(/^./, (x) => x[0].toUpperCase());
export const jsStrToObject = (raw) => new Function(`return Object.assign({}, ${raw})`)();
export const trimDollarSignPrefix = (str) => str.startsWith('$') ? str.slice(1) : str;
const caseFns = { kebab, snake, pascal };
export function modifyCasing(str, mods) {
    for (const c of mods.get('case') || []) {
        const fn = caseFns[c];
        if (fn)
            str = fn(str);
    }
    return str;
}
