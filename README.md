<p align="center">
  <img  src="https://raw.githubusercontent.com/helmwave/logo/main/logo.png" style="max-height:100%;" height="300px" />
</p>

<h1 align="center"> Helmwave</h1>

<p align="center">
  <a href="https://github.com/helmwave/helmwave/actions?query=workflow%3Arelease"><img src="https://github.com/helmwave/helmwave/workflows/release/badge.svg" /></a>
  <a href="https://codecov.io/gh/helmwave/helmwave"><img src="https://codecov.io/gh/helmwave/helmwave/branch/main/graph/badge.svg?token=0WXxYhIG4S"/></a> 
  <a href="https://bestpractices.coreinfrastructure.org/projects/5426"><img src="https://bestpractices.coreinfrastructure.org/projects/5426/badge"></a>
  <a href="https://www.codacy.com/gh/helmwave/helmwave/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=helmwave/helmwave&amp;utm_campaign=Badge_Grade"><img src="https://app.codacy.com/project/badge/Grade/200ca37690b7463b976f1ece36b53a4e"/></a>
  <a href="https://www.codefactor.io/repository/github/helmwave/helmwave"><img src="https://www.codefactor.io/repository/github/helmwave/helmwave/badge" alt="CodeFactor" /></a>
  <a href="https://requires.io/github/helmwave/helmwave/requirements/?branch=main"><img src="https://requires.io/github/helmwave/helmwave/requirements.svg?branch=main" alt="Requirements Status" /></a>
  <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/helmwave/helmwave">
  <img alt="GitHub" src="https://img.shields.io/github/license/zhilyaev/helmwave">
  <img alt="GitHub tag (latest SemVer)" src="https://img.shields.io/github/v/tag/zhilyaev/helmwave?label=latest">
</p>


🌊 Helmwave is **[helm](https://github.com/helm/helm/)-native** tool for deploy your Helm Charts via **GitOps**.
HelmWave is like docker-compose for helm.

- Deploy multiple environments by one step
- Separate values for environments
- Common values for apps
- Keep a directory of chart value files
- Maintain changes in version control.
- Template values
- Step by Step deployment.

Look at  the examples in our [docs](https://helmwave.github.io/docs)


## Comparison

🚀 Features  | 🌊 HelmWave   | helmfile
-------------| :------------:|:-----------:
Docker | ![Docker Image Size helmwave (latest by date)](https://img.shields.io/docker/image-size/diamon/helmwave) | ![Docker Image Size helmfile (latest by date)](https://img.shields.io/docker/image-size/chatwork/helmfile)
[Kubedog](https://github.com/werf/kubedog) |✅|❌
Without helm binary |✅|❌
All options helm|✅|partially
Helm 3 |✅|✅
Helm 2 |❌|✅
Parallel helm install/upgrade |✅|❌
Repository Skipping|✅|❌
Tags|✅| You can use labels
Store|✅| You can use labels
Planfile|✅|❌
remote values | ✅ | ❌
Sprig | ✅|✅
helm-diff  | ✅  in-compile |✅ as plugin
Call helm | via Golang Module | Shell Executor


### Run as a container ![Docker Pulls](https://img.shields.io/docker/pulls/diamon/helmwave)

```
$ docker run ghcr.io/helmwave/helmwave:scratch version
0.16.5
$ docker run --entrypoint=ash -it --rm ghcr.io/helmwave/helmwave:latest
/ # 
```

## 📖 [Documentation](https://helmwave.github.io/docs)

Documentation available at https://helmwave.github.io/docs


## Community, discussion, contribution, and support

- <a href="https://t.me/helmwave" ><img src="https://img.shields.io/badge/telegram-chat-179cde.svg?logo=telegram" /></a>
- [kanban](https://github.com/orgs/helmwave/projects/1)
- [contribution guide](https://github.com/helmwave/helmwave/blob/main/CONTRIBUTING.md)
- [security and vulnerabilities](https://github.com/helmwave/helmwave/blob/main/SECURITY.md)
