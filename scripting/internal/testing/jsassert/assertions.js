gost.assertEqual = (actual, expected) => {
  if (actual !== expected) {
    return gost.error(`Expected value equal to ${expected}. Got: ${actual}`);
  }
};
