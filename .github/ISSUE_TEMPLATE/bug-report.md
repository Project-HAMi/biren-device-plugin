---
name: Bug Report
about: Report a problem encountered while using biren-device-plugin
labels: bug
---

<!-- Please use this template while reporting a bug and provide as much info as possible. Not doing so may result in your bug not being addressed in a timely manner. Thanks!
-->

**What happened**:

**What you expected to happen**:

**How to reproduce it (as minimally and precisely as possible)**:

**Anything else we need to know?**:

- Relevant excerpts from the vendor device-management output used to inspect the affected Biren GPUs
- Relevant, time-bounded excerpts from the biren-device-plugin container logs
- Relevant, time-bounded excerpts from the HAMi scheduler and kubelet logs
- Relevant, redacted Pod, node-label, and device-plugin deployment manifests or events
- Relevant Docker or containerd configuration sections
- Relevant, time-bounded Biren driver or kernel output

Before posting, remove or mask credentials, tokens, GPU or device identifiers, Pod or workload identifiers, namespace or node names, internal image names, registry details, and other sensitive data from configuration and logs.

**Environment**:
- biren-device-plugin version, commit, or image:
- Biren GPU model:
- Biren GPU driver version:
- Kubernetes and HAMi version:
- Docker or containerd version:
- Deployment method and relevant configuration:
- Requested `birentech.com/gpu` resources:
- Kernel version from `uname -a`:
- Others:
