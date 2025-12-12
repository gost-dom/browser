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
