// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { WatchtowerArgs } from "./watchtower";
export type Watchtower = import("./watchtower").Watchtower;
export const Watchtower: typeof import("./watchtower").Watchtower = null as any;
utilities.lazyLoad(exports, ["Watchtower"], () => require("./watchtower"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "svmkit:watchtower:Watchtower":
                return new Watchtower(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("svmkit", "watchtower", _module)