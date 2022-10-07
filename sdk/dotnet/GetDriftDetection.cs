// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetDriftDetection
    {
        /// <summary>
        /// `spacelift.DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.
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
        ///         var core_infra_production_drift_detection = Output.Create(Spacelift.GetDriftDetection.InvokeAsync(new Spacelift.GetDriftDetectionArgs
        ///         {
        ///             StackId = "core-infra-production",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDriftDetectionResult> InvokeAsync(GetDriftDetectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDriftDetectionResult>("spacelift:index/getDriftDetection:getDriftDetection", args ?? new GetDriftDetectionArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.
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
        ///         var core_infra_production_drift_detection = Output.Create(Spacelift.GetDriftDetection.InvokeAsync(new Spacelift.GetDriftDetectionArgs
        ///         {
        ///             StackId = "core-infra-production",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDriftDetectionResult> Invoke(GetDriftDetectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDriftDetectionResult>("spacelift:index/getDriftDetection:getDriftDetection", args ?? new GetDriftDetectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDriftDetectionArgs : Pulumi.InvokeArgs
    {
        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetDriftDetectionArgs()
        {
        }
    }

    public sealed class GetDriftDetectionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetDriftDetectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDriftDetectionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool Reconcile;
        public readonly ImmutableArray<string> Schedules;
        public readonly string StackId;
        public readonly string Timezone;

        [OutputConstructor]
        private GetDriftDetectionResult(
            string id,

            bool reconcile,

            ImmutableArray<string> schedules,

            string stackId,

            string timezone)
        {
            Id = id;
            Reconcile = reconcile;
            Schedules = schedules;
            StackId = stackId;
            Timezone = timezone;
        }
    }
}
