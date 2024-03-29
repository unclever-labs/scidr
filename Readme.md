# SCIDR

## Introduction

Split a large CIDR block < `/16` into smaller CIDR blocks. This is needed when using AWS WAF since the largest CIDR block they accept is a `/16`. Used in CI/CD pipelines with Jinja2 templating -> Terraform when definining WAF IP Sets (whitelists/blacklists).

## Contents

- [Install](#install)
- [Usage](#usage)

## Install

```bash
brew install unclever-labs/unclever-labs/scidr
```

## Usage

Pass an arg into it:

```bash
scidr 58.14.0.0/15
```

Pipe into it:

```bash
echo 58.14.0.0/15 | scidr
```

Here is example output:

```bash
➜  scidr git:(master) scidr 58.14.0.0/15
58.14.0.0/16
58.15.0.0/16
```

```bash
➜  scidr git:(master) scidr 42.128.0.0/12
42.128.0.0/16
42.129.0.0/16
42.130.0.0/16
42.131.0.0/16
42.132.0.0/16
42.133.0.0/16
42.134.0.0/16
42.135.0.0/16
42.136.0.0/16
42.137.0.0/16
42.138.0.0/16
42.139.0.0/16
42.140.0.0/16
42.141.0.0/16
42.142.0.0/16
42.143.0.0/16
```
