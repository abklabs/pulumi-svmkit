// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class Watchtower extends pulumi.CustomResource {
    /**
     * Get an existing Watchtower resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Watchtower {
        return new Watchtower(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'svmkit:watchtower:Watchtower';

    /**
     * Returns true if the given object is an instance of Watchtower.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Watchtower {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Watchtower.__pulumiType;
    }

    public readonly connection!: pulumi.Output<outputs.ssh.Connection>;
    public readonly environment!: pulumi.Output<outputs.solana.Environment>;
    public readonly flags!: pulumi.Output<outputs.watchtower.WatchtowerFlags>;
    public readonly notifications!: pulumi.Output<outputs.watchtower.NotificationConfig>;
    public readonly runnerConfig!: pulumi.Output<outputs.runner.Config | undefined>;

    /**
     * Create a Watchtower resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WatchtowerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connection'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.flags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flags'");
            }
            if ((!args || args.notifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifications'");
            }
            resourceInputs["connection"] = args ? (args.connection ? pulumi.output(args.connection).apply(inputs.ssh.connectionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["flags"] = args ? args.flags : undefined;
            resourceInputs["notifications"] = args ? args.notifications : undefined;
            resourceInputs["runnerConfig"] = args ? args.runnerConfig : undefined;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["flags"] = undefined /*out*/;
            resourceInputs["notifications"] = undefined /*out*/;
            resourceInputs["runnerConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Watchtower.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Watchtower resource.
 */
export interface WatchtowerArgs {
    connection: pulumi.Input<inputs.ssh.ConnectionArgs>;
    environment: pulumi.Input<inputs.solana.EnvironmentArgs>;
    flags: pulumi.Input<inputs.watchtower.WatchtowerFlagsArgs>;
    notifications: pulumi.Input<inputs.watchtower.NotificationConfigArgs>;
    runnerConfig?: pulumi.Input<inputs.runner.ConfigArgs>;
}
