console.log("Test message");

const globalProto = Object.getPrototypeOf(globalThis);
const protoProto = Object.getPrototypeOf(globalProto);
const baseProto = Object.getPrototypeOf(protoProto);

postMessage({
  key: "instance globale",
  value: Object.getOwnPropertyNames(globalThis).join(", "),
});
postMessage({
  key: "Prototype name",
  value: globalProto.constructor.name,
});
postMessage({
  key: "Inherited name",
  value: protoProto.constructor.name,
});
postMessage({
  key: "Base proto name",
  value: baseProto.constructor.name,
});
postMessage({
  key: "Prototype globals",
  value: Object.getOwnPropertyNames(globalProto).join(", "),
});
postMessage({
  key: "Inherited globals",
  value: Object.getOwnPropertyNames(protoProto).join(", "),
});
postMessage({
  key: "Base globals",
  value: Object.getOwnPropertyNames(baseProto).join(", "),
});
