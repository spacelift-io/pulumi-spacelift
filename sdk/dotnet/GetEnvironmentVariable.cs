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
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ireland_kubeconfig = Output.Create(Spacelift.GetEnvironmentVariable.InvokeAsync(new Spacelift.GetEnvironmentVariableArgs
        ///         {
        ///             ContextId = "prod-k8s-ie",
        ///             Name = "KUBECONFIG",
        ///         }));
        ///         var module_kubeconfig = Output.Create(Spacelift.GetEnvironmentVariable.InvokeAsync(new Spacelift.GetEnvironmentVariableArgs
        ///         {
        ///             ModuleId = "k8s-module",
        ///             Name = "KUBECONFIG",
        ///         }));
        ///         var core_kubeconfig = Output.Create(Spacelift.GetEnvironmentVariable.InvokeAsync(new Spacelift.GetEnvironmentVariableArgs
        ///         {
        ///             Name = "KUBECONFIG",
        ///             StackId = "k8s-core",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// &lt;!-- schema generated by tfplugindocs --&gt;
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Schema
        /// 
        /// ### Required
        /// 
        /// - **name** (String) name of the environment variable
        /// 
        /// ### Optional
        /// 
        /// - **context_id** (String) ID of the context on which the environment variable is defined
        /// - **id** (String) The ID of this resource.
        /// - **module_id** (String) ID of the module on which the environment variable is defined
        /// - **stack_id** (String) ID of the stack on which the environment variable is defined
        /// 
        /// ### Read-Only
        /// 
        /// - **checksum** (String) SHA-256 checksum of the value
        /// - **value** (String, Sensitive) value of the environment variable
        /// - **write_only** (Boolean) indicates whether the value can be read back outside a Run
        /// </summary>
        public static Task<GetEnvironmentVariableResult> InvokeAsync(GetEnvironmentVariableArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentVariableResult>("spacelift:index/getEnvironmentVariable:getEnvironmentVariable", args ?? new GetEnvironmentVariableArgs(), options.WithVersion());
    }


    public sealed class GetEnvironmentVariableArgs : Pulumi.InvokeArgs
    {
        [Input("contextId")]
        public string? ContextId { get; set; }

        [Input("moduleId")]
        public string? ModuleId { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("stackId")]
        public string? StackId { get; set; }

        public GetEnvironmentVariableArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentVariableResult
    {
        public readonly string Checksum;
        public readonly string? ContextId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ModuleId;
        public readonly string Name;
        public readonly string? StackId;
        public readonly string Value;
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
