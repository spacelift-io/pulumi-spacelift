import * as pulumi from "@pulumi/pulumi";

import type { StackArgs } from "@spacelift-io/pulumi-spacelift";
import { Stack } from "@spacelift-io/pulumi-spacelift";

const stackArgs: StackArgs = {
  administrative: false,
  autodeploy: false,
  branch: "main",
  description: "A simple stack",
  name: "simple-stack",
  repository: "empty",
  projectRoot: "",
  terraformVersion: "1.3.0",
};

export const stack = new Stack("pulumi-test-stack", stackArgs);
