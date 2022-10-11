// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetMountedfile
    {
        /// <summary>
        /// `spacelift.Mountedfile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`). In principle, it's very similar to an environment variable (`spacelift.EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ireland_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///         RelativePath = "kubeconfig",
        ///     });
        /// 
        ///     var module_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         ModuleId = "k8s-module",
        ///         RelativePath = "kubeconfig",
        ///     });
        /// 
        ///     var core_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         RelativePath = "kubeconfig",
        ///         StackId = "k8s-core",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMountedfileResult> InvokeAsync(GetMountedfileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMountedfileResult>("spacelift:index/getMountedfile:getMountedfile", args ?? new GetMountedfileArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.Mountedfile` represents a file mounted in each Run's workspace that is part of a configuration of a context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`). In principle, it's very similar to an environment variable (`spacelift.EnvironmentVariable`) except that the value is written to the filesystem rather than passed to the environment.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ireland_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///         RelativePath = "kubeconfig",
        ///     });
        /// 
        ///     var module_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         ModuleId = "k8s-module",
        ///         RelativePath = "kubeconfig",
        ///     });
        /// 
        ///     var core_kubeconfig = Spacelift.GetMountedfile.Invoke(new()
        ///     {
        ///         RelativePath = "kubeconfig",
        ///         StackId = "k8s-core",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMountedfileResult> Invoke(GetMountedfileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMountedfileResult>("spacelift:index/getMountedfile:getMountedfile", args ?? new GetMountedfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMountedfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the context where the mounted file is stored
        /// </summary>
        [Input("contextId")]
        public string? ContextId { get; set; }

        /// <summary>
        /// ID of the module where the mounted file is stored
        /// </summary>
        [Input("moduleId")]
        public string? ModuleId { get; set; }

        /// <summary>
        /// relative path to the mounted file
        /// </summary>
        [Input("relativePath", required: true)]
        public string RelativePath { get; set; } = null!;

        /// <summary>
        /// ID of the stack where the mounted file is stored
        /// </summary>
        [Input("stackId")]
        public string? StackId { get; set; }

        public GetMountedfileArgs()
        {
        }
        public static new GetMountedfileArgs Empty => new GetMountedfileArgs();
    }

    public sealed class GetMountedfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the context where the mounted file is stored
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// ID of the module where the mounted file is stored
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// relative path to the mounted file
        /// </summary>
        [Input("relativePath", required: true)]
        public Input<string> RelativePath { get; set; } = null!;

        /// <summary>
        /// ID of the stack where the mounted file is stored
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public GetMountedfileInvokeArgs()
        {
        }
        public static new GetMountedfileInvokeArgs Empty => new GetMountedfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetMountedfileResult
    {
        /// <summary>
        /// SHA-256 checksum of the value
        /// </summary>
        public readonly string Checksum;
        /// <summary>
        /// content of the mounted file encoded using Base-64
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// ID of the context where the mounted file is stored
        /// </summary>
        public readonly string? ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the module where the mounted file is stored
        /// </summary>
        public readonly string? ModuleId;
        /// <summary>
        /// relative path to the mounted file
        /// </summary>
        public readonly string RelativePath;
        /// <summary>
        /// ID of the stack where the mounted file is stored
        /// </summary>
        public readonly string? StackId;
        /// <summary>
        /// indicates whether the value can be read back outside a Run
        /// </summary>
        public readonly bool WriteOnly;

        [OutputConstructor]
        private GetMountedfileResult(
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
