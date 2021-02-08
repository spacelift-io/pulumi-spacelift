// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetMountedFile
    {
        public static Task<GetMountedFileResult> InvokeAsync(GetMountedFileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMountedFileResult>("spacelift:index/getMountedFile:getMountedFile", args ?? new GetMountedFileArgs(), options.WithVersion());
    }


    public sealed class GetMountedFileArgs : Pulumi.InvokeArgs
    {
        [Input("contextId")]
        public string? ContextId { get; set; }

        [Input("moduleId")]
        public string? ModuleId { get; set; }

        [Input("relativePath", required: true)]
        public string RelativePath { get; set; } = null!;

        [Input("stackId")]
        public string? StackId { get; set; }

        public GetMountedFileArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMountedFileResult
    {
        public readonly string Checksum;
        public readonly string Content;
        public readonly string? ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ModuleId;
        public readonly string RelativePath;
        public readonly string? StackId;
        public readonly bool WriteOnly;

        [OutputConstructor]
        private GetMountedFileResult(
            string checksum,

            string content,

            string? contextId,

            string id,

            string? moduleId,

            string relativePath,

            string? stackId,

            bool writeOnly)
        {
            Checksum = checksum;
            Content = content;
            ContextId = contextId;
            Id = id;
            ModuleId = moduleId;
            RelativePath = relativePath;
            StackId = stackId;
            WriteOnly = writeOnly;
        }
    }
}
