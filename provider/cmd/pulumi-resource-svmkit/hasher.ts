import { createHash } from "crypto";

/**
 * Computes a SHA-256 hash for an array of buffers.
 *
 * @param {Buffer[]} buffers - An array of Buffer objects to be hashed.
 * @returns {string} - The resulting SHA-256 hash as a hexadecimal string.
 */
export function compute(buffers: Buffer[]): string {
    const hash = createHash("sha256");

    buffers.forEach((buffer) => {
        hash.update(buffer);
    });

    return hash.digest("hex");
}
