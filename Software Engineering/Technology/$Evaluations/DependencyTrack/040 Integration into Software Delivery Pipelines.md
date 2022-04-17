#dependency-track, #devsecops, #source-composition-analysis

# Integration into Software Delivery Pipelines

This serves as a simple example to implement into a software delivery pipeline.


## Workflow

### Prerequisites

A toolchain that can output a `CycloneDX` or `SPDX` SBOM like [cyclonedx-bom](https://github.com/CycloneDX/cyclonedx-node-module) is required.

*See [Appendix A - SBOM Toolchains](#appendix-a-sbom-toolchains) for available toolchains.*

### High Level Flow

This acts as a high level flow for an SCA (Source Composition Analysis) stage of a software delivery pipeline.

1. Execute SBOM toolchain to output an SBOM.
2. Upload the SBOM to Dependency Track
3. Poll the processing status endpoint until SBOM processing is done.
4. Retrieve vulnerabilities for project @ version.
5. Process the vulnerability output for policy failures and suppressions.
6. Display insights in the pipeline.
7. Optionally use a risk based policy gate for remediation or disposition.


## Appendix

### Appendix A - SBOM Toolchains

List of toolchains that can output CycloneDX SBOMs. [CycloneDX](https://github.com/CycloneDX) publishes for a variety of stacks.

#### Go

- [cycleondx-gomod](https://github.com/CycloneDX/cyclonedx-gomod)

#### NPM

- [cyclonedx-bom](https://github.com/CycloneDX/cyclonedx-node-module) 