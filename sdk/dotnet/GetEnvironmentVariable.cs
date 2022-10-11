// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetEnvironmentVariable
    {
        /// <summary>
        /// `spacelift.EnvironmentVariable` defines an environment variable on the context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`), thereby allowing to pass and share various secrets and configuration with Spacelift stacks.
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
        ///     var ireland_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///         Name = "KUBECONFIG",
        ///     });
        /// 
        ///     var module_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         ModuleId = "k8s-module",
        ///         Name = "KUBECONFIG",
        ///     });
        /// 
        ///     var core_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         Name = "KUBECONFIG",
        ///         StackId = "k8s-core",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnvironmentVariableResult> InvokeAsync(GetEnvironmentVariableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentVariableResult>("spacelift:index/getEnvironmentVariable:getEnvironmentVariable", args ?? new GetEnvironmentVariableArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.EnvironmentVariable` defines an environment variable on the context (`spacelift.Context`), stack (`spacelift.Stack`) or a module (`spacelift.Module`), thereby allowing to pass and share various secrets and configuration with Spacelift stacks.
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
        ///     var ireland_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         ContextId = "prod-k8s-ie",
        ///         Name = "KUBECONFIG",
        ///     });
        /// 
        ///     var module_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         ModuleId = "k8s-module",
        ///         Name = "KUBECONFIG",
        ///     });
        /// 
        ///     var core_kubeconfig = Spacelift.GetEnvironmentVariable.Invoke(new()
        ///     {
        ///         Name = "KUBECONFIG",
        ///         StackId = "k8s-core",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEnvironmentVariableResult> Invoke(GetEnvironmentVariableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnvironmentVariableResult>("spacelift:index/getEnvironmentVariable:getEnvironmentVariable", args ?? new GetEnvironmentVariableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentVariableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the context on which the environment variable is defined
        /// </summary>
        [Input("contextId")]
        public string? ContextId { get; set; }

        /// <summary>
        /// ID of the module on which the environment variable is defined
        /// </summary>
        [Input("moduleId")]
        public string? ModuleId { get; set; }

        /// <summary>
        /// name of the environment variable
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// ID of the stack on which the environment variable is defined
        /// </summary>
        [Input("stackId")]
        public string? StackId { get; set; }

        public GetEnvironmentVariableArgs()
        {
        }
        public static new GetEnvironmentVariableArgs Empty => new GetEnvironmentVariableArgs();
    }

    public sealed class GetEnvironmentVariableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the context on which the environment variable is defined
        /// </summary>
        [Input("contextId")]
        public Input<string>? ContextId { get; set; }

        /// <summary>
        /// ID of the module on which the environment variable is defined
        /// </summary>
        [Input("moduleId")]
        public Input<string>? ModuleId { get; set; }

        /// <summary>
        /// name of the environment variable
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// ID of the stack on which the environment variable is defined
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        public GetEnvironmentVariableInvokeArgs()
        {
        }
        public static new GetEnvironmentVariableInvokeArgs Empty => new GetEnvironmentVariableInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnvironmentVariableResult
    {
        /// <summary>
        /// SHA-256 checksum of the value
        /// </summary>
        public readonly string Checksum;
        /// <summary>
        /// ID of the context on which the environment variable is defined
        /// </summary>
        public readonly string? ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the module on which the environment variable is defined
        /// </summary>
        public readonly string? ModuleId;
        /// <summary>
        /// name of the environment variable
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the stack on which the environment variable is defined
        /// </summary>
        public readonly string? StackId;
        /// <summary>
        /// value of the environment variable
        /// </summary>
        public readonly string Value;
        /// <summary>
        /// indicates whether the value can be read back outside a Run
        /// </summary>
        public readonly bool WriteOnly;

        [OutputConstructor]
        private GetEnvironmentVariableResult(
            string checksum,

            string? contextId,

            string id,

            string? moduleId,

            string name,

            string? stackId,

            string value,

            bool writeOnly)
        {
            Checksum = checksum;
            ContextId = contextId;
            Id = id;
            ModuleId = moduleId;
            Name = name;
            StackId = stackId;
            Value = value;
            WriteOnly = writeOnly;
        }
    }
}
