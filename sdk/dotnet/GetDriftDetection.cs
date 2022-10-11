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
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Spacelift = Pulumi.Spacelift;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var core_infra_production_drift_detection = Spacelift.GetDriftDetection.Invoke(new()
        ///     {
        ///         StackId = "core-infra-production",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDriftDetectionResult> InvokeAsync(GetDriftDetectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDriftDetectionResult>("spacelift:index/getDriftDetection:getDriftDetection", args ?? new GetDriftDetectionArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.DriftDetection` represents a Drift Detection configuration for a Stack. It will trigger a proposed run on the given schedule, which you can listen for using run state webhooks. If reconcile is true, then a tracked run will be triggered when drift is detected.
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
        ///     var core_infra_production_drift_detection = Spacelift.GetDriftDetection.Invoke(new()
        ///     {
        ///         StackId = "core-infra-production",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDriftDetectionResult> Invoke(GetDriftDetectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDriftDetectionResult>("spacelift:index/getDriftDetection:getDriftDetection", args ?? new GetDriftDetectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDriftDetectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the stack for which to set up drift detection
        /// </summary>
        [Input("stackId", required: true)]
        public string StackId { get; set; } = null!;

        public GetDriftDetectionArgs()
        {
        }
        public static new GetDriftDetectionArgs Empty => new GetDriftDetectionArgs();
    }

    public sealed class GetDriftDetectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the stack for which to set up drift detection
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        public GetDriftDetectionInvokeArgs()
        {
        }
        public static new GetDriftDetectionInvokeArgs Empty => new GetDriftDetectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDriftDetectionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether a tracked run should be triggered when drift is detected.
        /// </summary>
        public readonly bool Reconcile;
        /// <summary>
        /// List of cron schedule expressions based on which drift detection should be triggered.
        /// </summary>
        public readonly ImmutableArray<string> Schedules;
        /// <summary>
        /// ID of the stack for which to set up drift detection
        /// </summary>
        public readonly string StackId;
        /// <summary>
        /// Timezone in which the schedule is expressed
        /// </summary>
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
