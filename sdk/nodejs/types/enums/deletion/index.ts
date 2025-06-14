// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Policy = {
    /**
     * Do not delete files upon destroying their associated component
     */
    Keep: "keep",
    /**
     * Delete files upon destroying their associated component; if the files exist before the component is created, will throw an error
     */
    Delete: "delete",
    /**
     * Delete files upon destroying their associated component, and overwrite existing files on creation — use with caution
     */
    Delete_and_force_creation: "delete-and-force-creation",
} as const;

export type Policy = (typeof Policy)[keyof typeof Policy];
