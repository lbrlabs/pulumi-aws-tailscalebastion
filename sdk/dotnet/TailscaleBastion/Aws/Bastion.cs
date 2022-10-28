// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.TailscaleBastion.Aws
{
    [TailscaleBastionResourceType("tailscale-bastion:aws:Bastion")]
    public partial class Bastion : global::Pulumi.ComponentResource
    {
        /// <summary>
        /// The name of the ASG that managed the bastion instances
        /// </summary>
        [Output("asgName")]
        public Output<string> AsgName { get; private set; } = null!;

        /// <summary>
        /// The SSH private key to access your bastion
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;


        /// <summary>
        /// Create a Bastion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bastion(string name, BastionArgs args, ComponentResourceOptions? options = null)
            : base("tailscale-bastion:aws:Bastion", name, args ?? new BastionArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class BastionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The EC2 instance type to use for the bastion.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The AWS region you're using.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The route you'd like to advertise via tailscale.
        /// </summary>
        [Input("route", required: true)]
        public Input<string> Route { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet Ids to launch instances in.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The VPC the Bastion should be created in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public BastionArgs()
        {
        }
        public static new BastionArgs Empty => new BastionArgs();
    }
}
