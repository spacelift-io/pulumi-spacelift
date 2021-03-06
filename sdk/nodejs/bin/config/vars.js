"use strict";
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
let __config = new pulumi.Config("spacelift");
/**
 * Endpoint to use when authenticating with an API key outside of Spacelift
 */
exports.apiKeyEndpoint = __config.get("apiKeyEndpoint") || utilities.getEnv("SPACELIFT_API_KEY_ENDPOINT");
/**
 * ID of the API key to use when executing outside of Spacelift
 */
exports.apiKeyId = __config.get("apiKeyId") || utilities.getEnv("SPACELIFT_API_KEY_ID");
/**
 * API key secret to use when executing outside of Spacelift
 */
exports.apiKeySecret = __config.get("apiKeySecret") || utilities.getEnv("SPACELIFT_API_KEY_SECRET");
/**
 * Spacelift token generated by a run, only useful from within Spacelift
 */
exports.apiToken = __config.get("apiToken") || utilities.getEnv("SPACELIFT_API_TOKEN");
//# sourceMappingURL=vars.js.map