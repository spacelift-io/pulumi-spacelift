// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Spacelift
{
    public static class GetStacks
    {
        /// <summary>
        /// `spacelift.getStacks` represents all the stacks in the Spacelift account visible to the API user, matching predicates.
        /// </summary>
        public static Task<GetStacksResult> InvokeAsync(GetStacksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStacksResult>("spacelift:index/getStacks:getStacks", args ?? new GetStacksArgs(), options.WithDefaults());

        /// <summary>
        /// `spacelift.getStacks` represents all the stacks in the Spacelift account visible to the API user, matching predicates.
        /// </summary>
        public static Output<GetStacksResult> Invoke(GetStacksInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStacksResult>("spacelift:index/getStacks:getStacks", args ?? new GetStacksInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStacksArgs : global::Pulumi.InvokeArgs
    {
        [Input("administrative")]
        public Inputs.GetStacksAdministrativeArgs? Administrative { get; set; }

        [Input("branch")]
        public Inputs.GetStacksBranchArgs? Branch { get; set; }

        /// <summary>
        /// Require stacks to be on one of the commits
        /// </summary>
        [Input("commit")]
        public Inputs.GetStacksCommitArgs? Commit { get; set; }

        [Input("labels")]
        private List<Inputs.GetStacksLabelArgs>? _labels;
        public List<Inputs.GetStacksLabelArgs> Labels
        {
            get => _labels ?? (_labels = new List<Inputs.GetStacksLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Require stacks to be locked
        /// </summary>
        [Input("locked")]
        public Inputs.GetStacksLockedArgs? Locked { get; set; }

        [Input("name")]
        public Inputs.GetStacksNameArgs? Name { get; set; }

        [Input("projectRoot")]
        public Inputs.GetStacksProjectRootArgs? ProjectRoot { get; set; }

        [Input("repository")]
        public Inputs.GetStacksRepositoryArgs? Repository { get; set; }

        /// <summary>
        /// Require stacks to have one of the states
        /// </summary>
        [Input("state")]
        public Inputs.GetStacksStateArgs? State { get; set; }

        /// <summary>
        /// Require stacks to use one of the IaC vendors
        /// </summary>
        [Input("vendor")]
        public Inputs.GetStacksVendorArgs? Vendor { get; set; }

        /// <summary>
        /// Require stacks to use one of the worker pools
        /// </summary>
        [Input("workerPool")]
        public Inputs.GetStacksWorkerPoolArgs? WorkerPool { get; set; }

        public GetStacksArgs()
        {
        }
        public static new GetStacksArgs Empty => new GetStacksArgs();
    }

    public sealed class GetStacksInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("administrative")]
        public Input<Inputs.GetStacksAdministrativeInputArgs>? Administrative { get; set; }

        [Input("branch")]
        public Input<Inputs.GetStacksBranchInputArgs>? Branch { get; set; }

        /// <summary>
        /// Require stacks to be on one of the commits
        /// </summary>
        [Input("commit")]
        public Input<Inputs.GetStacksCommitInputArgs>? Commit { get; set; }

        [Input("labels")]
        private InputList<Inputs.GetStacksLabelInputArgs>? _labels;
        public InputList<Inputs.GetStacksLabelInputArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.GetStacksLabelInputArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Require stacks to be locked
        /// </summary>
        [Input("locked")]
        public Input<Inputs.GetStacksLockedInputArgs>? Locked { get; set; }

        [Input("name")]
        public Input<Inputs.GetStacksNameInputArgs>? Name { get; set; }

        [Input("projectRoot")]
        public Input<Inputs.GetStacksProjectRootInputArgs>? ProjectRoot { get; set; }

        [Input("repository")]
        public Input<Inputs.GetStacksRepositoryInputArgs>? Repository { get; set; }

        /// <summary>
        /// Require stacks to have one of the states
        /// </summary>
        [Input("state")]
        public Input<Inputs.GetStacksStateInputArgs>? State { get; set; }

        /// <summary>
        /// Require stacks to use one of the IaC vendors
        /// </summary>
        [Input("vendor")]
        public Input<Inputs.GetStacksVendorInputArgs>? Vendor { get; set; }

        /// <summary>
        /// Require stacks to use one of the worker pools
        /// </summary>
        [Input("workerPool")]
        public Input<Inputs.GetStacksWorkerPoolInputArgs>? WorkerPool { get; set; }

        public GetStacksInvokeArgs()
        {
        }
        public static new GetStacksInvokeArgs Empty => new GetStacksInvokeArgs();
    }


    [OutputType]
    public sealed class GetStacksResult
    {
        /// <summary>
        /// Require stacks to be administrative or not
        /// </summary>
        public readonly Outputs.GetStacksAdministrativeResult? Administrative;
        /// <summary>
        /// Require stacks to be on one of the branches
        /// </summary>
        public readonly Outputs.GetStacksBranchResult? Branch;
        /// <summary>
        /// Require stacks to be on one of the commits
        /// </summary>
        public readonly Outputs.GetStacksCommitResult? Commit;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Require stacks to have one of the labels
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStacksLabelResult> Labels;
        /// <summary>
        /// Require stacks to be locked
        /// </summary>
        public readonly Outputs.GetStacksLockedResult? Locked;
        /// <summary>
        /// Require stacks to have one of the names
        /// </summary>
        public readonly Outputs.GetStacksNameResult? Name;
        /// <summary>
        /// Require stacks to be in one of the project roots
        /// </summary>
        public readonly Outputs.GetStacksProjectRootResult? ProjectRoot;
        /// <summary>
        /// Require stacks to be in one of the repositories
        /// </summary>
        public readonly Outputs.GetStacksRepositoryResult? Repository;
        /// <summary>
        /// List of stacks matching the predicates
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStacksStackResult> Stacks;
        /// <summary>
        /// Require stacks to have one of the states
        /// </summary>
        public readonly Outputs.GetStacksStateResult? State;
        /// <summary>
        /// Require stacks to use one of the IaC vendors
        /// </summary>
        public readonly Outputs.GetStacksVendorResult? Vendor;
        /// <summary>
        /// Require stacks to use one of the worker pools
        /// </summary>
        public readonly Outputs.GetStacksWorkerPoolResult? WorkerPool;

        [OutputConstructor]
        private GetStacksResult(
            Outputs.GetStacksAdministrativeResult? administrative,

            Outputs.GetStacksBranchResult? branch,

            Outputs.GetStacksCommitResult? commit,

            string id,

            ImmutableArray<Outputs.GetStacksLabelResult> labels,

            Outputs.GetStacksLockedResult? locked,

            Outputs.GetStacksNameResult? name,

            Outputs.GetStacksProjectRootResult? projectRoot,

            Outputs.GetStacksRepositoryResult? repository,

            ImmutableArray<Outputs.GetStacksStackResult> stacks,

            Outputs.GetStacksStateResult? state,

            Outputs.GetStacksVendorResult? vendor,

            Outputs.GetStacksWorkerPoolResult? workerPool)
        {
            Administrative = administrative;
            Branch = branch;
            Commit = commit;
            Id = id;
            Labels = labels;
            Locked = locked;
            Name = name;
            ProjectRoot = projectRoot;
            Repository = repository;
            Stacks = stacks;
            State = state;
            Vendor = vendor;
            WorkerPool = workerPool;
        }
    }
}
