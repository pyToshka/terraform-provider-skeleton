# terraform-provider-skeleton

## Overview

**[Terraform](https://github.com/hashicorp/terraform)** is a tool for automating infrastructure. Terraform includes the ability to provision resources at creation time through a plugin api.
It's template for creating custom terraform provider.

## Installation
```bash
go get github.com/pyToshka/terraform-provider-skeleton

```
**terraform-provider-skeleton** ships as a single binary and is compatible with **terraform**'s plugin interface. Behind the scenes, terraform plugins use https://github.com/hashicorp/go-plugin and communicate with the parent terraform process via RPC.

Once installed, a `~/.terraformrc` file is used to _enable_ the plugin.

```bash
providers {
    skeleton = "/usr/local/bin/terraform-provider-skeleton"
}
```

## Usage

Once installed, you can configure resources by including an `skeleton` provider block.



```
provider "skeleton" {

}
resource "skeleton" "my-resource" {
    key="value"
}

```

Some options
```
option- description

```

Travis build status

[![Build Status](https://travis-ci.org/pyToshka/terraform-provider-skeleton.svg?branch=master)](https://travis-ci.org/pyToshka/terraform-provider-skeleton)