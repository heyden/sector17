"use strict";

const sbom = require("./gotestbench-sbom.json");

const SPACER = "\t";
const LEVEL_INDICATOR = "|";
const DEP_PREFIX = "+---- ";

const rootComponentRef = sbom.metadata.component["bom-ref"];

const bomRef = sbom.dependencies.filter(
  (depRef) => depRef.ref === rootComponentRef
)[0];

bomRef.dependsOn.forEach((ref) => {
  const depth = 0;
  let prefix = "\\--- ";
  const depRef = sbom.dependencies.filter((dep) => dep.ref === ref)[0];
  console.log(`${prefix}${depRef.ref}`);
  if (depRef.dependsOn && depRef.dependsOn.length > 0)
    printLevel(depRef.dependsOn, depth);
});

function printLevel(deps, depth) {
  depth += 1;

  let prefix = `${SPACER}${DEP_PREFIX}`;
  if (depth > 1) {
    prefix = new Array(depth - 1).fill("\t").join("");
    prefix = `${prefix}${LEVEL_INDICATOR}${DEP_PREFIX}`;
  }

  deps.forEach((ref) => {
    const depRef = sbom.dependencies.filter((dep) => dep.ref === ref)[0];
    console.log(`${prefix}${ref}`);
    if (depRef.dependsOn && depRef.dependsOn.length > 0)
      printLevel(depRef.dependsOn, depth);
  });
}
