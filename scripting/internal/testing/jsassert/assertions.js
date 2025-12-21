gost.assertEqual = (actual, expected) => {
  if (actual !== expected) {
    return gost.error(`Expected value equal to ${expected}. Got: ${actual}`);
  }
};

gost.assertInstanceOf = (actual, expected) => {
  if (!(actual instanceof expected)) {
    return gost.error(
      `Expected instanceof ${expected.prototype.constructor.name}. Got: ${actual}`,
    );
  }
};

gost.assertTypeOf = (actual, expected) => {
  if (typeof actual !== expected) {
    return gost.error(`Expected value of type ${expected}. Got: ${actual}`);
  }
};

gost.assertOk = (actual) => {
  if (!actual) {
    return gost.error(`Expected non-nullish value. Got ${actual}`);
  }
};

gost.assertTrue = (actual) => {
  if (actual !== true) {
    return gost.error(`Expected true value. Got ${actual}`);
  }
};

gost.assertFalse = (actual) => {
  if (actual !== false) {
    return gost.error(`Expected false value. Got ${actual}`);
  }
};

gost.assertUndefined = (actual) => {
  if (typeof actual !== "undefined") {
    return gost.error(`Expected an undefined value. Got ${actual}`);
  }
};

gost.assertNull = (actual) => {
  if (actual !== null) {
    return gost.error(`Expected null. Got ${actual}`);
  }
};
gost.assertNotNull = (actual) => {
  if (actual === null) {
    return gost.error(`Expected not null. Got ${actual}`);
  }
};

gost.assertHasOwnProperty = (actual, expected) => {
  if (typeof actual !== "object" || !actual) {
    return gost.error(`assetHasOwnProperty: expected an object. Got ${actual}`);
  }
  const names = Object.getOwnPropertyNames(actual);
  if (!names.includes(expected)) {
    return gost.error(
      `Expected object to have own property ${expected}. Got [${names.join(",")}]`,
    );
  }
};

gost.assertInheritsFrom = (actual, expected) => {
  if (typeof actual !== "function") {
    throw new TypeError(
      `assertInheritsFrom: actual: ${actual}: not a function. assertInherits must be called with the constructor functions`,
    );
  }
  if (typeof expected !== "function") {
    throw new TypeError(
      `assertInheritsFrom: expected: ${expected}: not a function. assertInherits must be called with the constructor functions`,
    );
  }
  const actualInherits = Object.getPrototypeOf(actual.prototype);
  const expectedProto = expected.prototype;
  if (actualInherits !== expectedProto) {
    return gost.error(
      `Expected object inherits directly from ${expected}. Inherits from ${actualInherits}`,
    );
  }
};
