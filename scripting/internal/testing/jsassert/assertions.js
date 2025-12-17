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
