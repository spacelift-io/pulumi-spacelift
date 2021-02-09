// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./awsRole";
export * from "./context";
export * from "./contextAttachment";
export * from "./environmentVariable";
export * from "./gcpServiceAccount";
export * from "./getAwsRole";
export * from "./getContext";
export * from "./getContextAttachment";
export * from "./getCurrentStack";
export * from "./getEnvironmentVariable";
export * from "./getGcpServiceAccount";
export * from "./getIps";
export * from "./getModule";
export * from "./getMountedFile";
export * from "./getPolicy";
export * from "./getStack";
export * from "./getStackAwsRole";
export * from "./getStackGcpServiceAccount";
export * from "./getWebhook";
export * from "./getWorkerPool";
export * from "./getWorkerPools";
export * from "./module";
export * from "./mountedFile";
export * from "./policy";
export * from "./policyAttachment";
export * from "./provider";
export * from "./stack";
export * from "./stackAwsRole";
export * from "./stackGcpServiceAccount";
export * from "./webhook";
export * from "./workerPool";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { AwsRole } from "./awsRole";
import { Context } from "./context";
import { ContextAttachment } from "./contextAttachment";
import { EnvironmentVariable } from "./environmentVariable";
import { GcpServiceAccount } from "./gcpServiceAccount";
import { Module } from "./module";
import { MountedFile } from "./mountedFile";
import { Policy } from "./policy";
import { PolicyAttachment } from "./policyAttachment";
import { Stack } from "./stack";
import { StackAwsRole } from "./stackAwsRole";
import { StackGcpServiceAccount } from "./stackGcpServiceAccount";
import { Webhook } from "./webhook";
import { WorkerPool } from "./workerPool";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "spacelift:index/awsRole:AwsRole":
                return new AwsRole(name, <any>undefined, { urn })
            case "spacelift:index/context:Context":
                return new Context(name, <any>undefined, { urn })
            case "spacelift:index/contextAttachment:ContextAttachment":
                return new ContextAttachment(name, <any>undefined, { urn })
            case "spacelift:index/environmentVariable:EnvironmentVariable":
                return new EnvironmentVariable(name, <any>undefined, { urn })
            case "spacelift:index/gcpServiceAccount:GcpServiceAccount":
                return new GcpServiceAccount(name, <any>undefined, { urn })
            case "spacelift:index/module:Module":
                return new Module(name, <any>undefined, { urn })
            case "spacelift:index/mountedFile:MountedFile":
                return new MountedFile(name, <any>undefined, { urn })
            case "spacelift:index/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "spacelift:index/policyAttachment:PolicyAttachment":
                return new PolicyAttachment(name, <any>undefined, { urn })
            case "spacelift:index/stack:Stack":
                return new Stack(name, <any>undefined, { urn })
            case "spacelift:index/stackAwsRole:StackAwsRole":
                return new StackAwsRole(name, <any>undefined, { urn })
            case "spacelift:index/stackGcpServiceAccount:StackGcpServiceAccount":
                return new StackGcpServiceAccount(name, <any>undefined, { urn })
            case "spacelift:index/webhook:Webhook":
                return new Webhook(name, <any>undefined, { urn })
            case "spacelift:index/workerPool:WorkerPool":
                return new WorkerPool(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("spacelift", "index/awsRole", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/context", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/contextAttachment", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/environmentVariable", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/gcpServiceAccount", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/module", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/mountedFile", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/policy", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/policyAttachment", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/stack", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/stackAwsRole", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/stackGcpServiceAccount", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/webhook", _module)
pulumi.runtime.registerResourceModule("spacelift", "index/workerPool", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("spacelift", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:spacelift") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});