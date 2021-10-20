// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class ModuleResource extends pulumi.CustomResource {
    /**
     * Get an existing ModuleResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ModuleResource {
        return new ModuleResource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'foobar::ModuleResource';

    /**
     * Returns true if the given object is an instance of ModuleResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModuleResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModuleResource.__pulumiType;
    }


    /**
     * Create a ModuleResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModuleResourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.plain_required_bool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plain_required_bool'");
            }
            if ((!args || args.plain_required_number === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plain_required_number'");
            }
            if ((!args || args.plain_required_string === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plain_required_string'");
            }
            if ((!args || args.required_bool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'required_bool'");
            }
            if ((!args || args.required_enum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'required_enum'");
            }
            if ((!args || args.required_number === undefined) && !opts.urn) {
                throw new Error("Missing required property 'required_number'");
            }
            if ((!args || args.required_string === undefined) && !opts.urn) {
                throw new Error("Missing required property 'required_string'");
            }
            inputs["optional_bool"] = (args ? args.optional_bool : undefined) ?? true;
            inputs["optional_enum"] = (args ? args.optional_enum : undefined) ?? 8;
            inputs["optional_number"] = (args ? args.optional_number : undefined) ?? 42;
            inputs["optional_string"] = (args ? args.optional_string : undefined) ?? "buzzer";
            inputs["plain_optional_bool"] = (args ? args.plain_optional_bool : undefined) ?? true;
            inputs["plain_optional_number"] = (args ? args.plain_optional_number : undefined) ?? 42;
            inputs["plain_optional_string"] = (args ? args.plain_optional_string : undefined) ?? "buzzer";
            inputs["plain_required_bool"] = (args ? args.plain_required_bool : undefined) ?? true;
            inputs["plain_required_number"] = (args ? args.plain_required_number : undefined) ?? 42;
            inputs["plain_required_string"] = (args ? args.plain_required_string : undefined) ?? "buzzer";
            inputs["required_bool"] = (args ? args.required_bool : undefined) ?? true;
            inputs["required_enum"] = (args ? args.required_enum : undefined) ?? 4;
            inputs["required_number"] = (args ? args.required_number : undefined) ?? 42;
            inputs["required_string"] = (args ? args.required_string : undefined) ?? "buzzer";
        } else {
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ModuleResource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModuleResource resource.
 */
export interface ModuleResourceArgs {
    optional_bool?: pulumi.Input<boolean>;
    optional_enum?: pulumi.Input<enums.EnumThing>;
    optional_number?: pulumi.Input<number>;
    optional_string?: pulumi.Input<string>;
    plain_optional_bool?: boolean;
    plain_optional_number?: number;
    plain_optional_string?: string;
    plain_required_bool: boolean;
    plain_required_number: number;
    plain_required_string: string;
    required_bool: pulumi.Input<boolean>;
    required_enum: pulumi.Input<enums.EnumThing>;
    required_number: pulumi.Input<number>;
    required_string: pulumi.Input<string>;
}
