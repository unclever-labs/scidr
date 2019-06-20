# Split CIDR

## Introduction

Split a large CIDR block < `/16` into smaller CIDR blocks

## Contents

- [Usage](#usage)

## Usage

Pass an arg into it:

```bash
split-cidr 58.14.0.0/15
```

Pipe into it:

```bash
echo 58.14.0.0/15 | split-cidr
```
