import * as pulumi from "@pulumi/pulumi";
import * as svmkit from "@pulumi/svmkit";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";

const config = new pulumi.Config();

const sshKey = new tls.PrivateKey("sshKey", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: sshKey.publicKeyOpenssh,
});

const amiIdToUse = (async () => {
    let v = config.get("instance.amiId");

    if (v) {
        return v;
    }

    const amiInfo = await aws.ec2.getAmi({
        filters: [
            {
                name: "name",
                values: ["debian-12-*"],
            },
            {
                name: "architecture",
                values: ["x86_64"],
            },
        ],
        owners: ["136693071363"], // Debian
        mostRecent: true,
    });

    return amiInfo.id;
})();

const instance = new aws.ec2.Instance("instance", {
    ami: amiIdToUse,
    instanceType: "t3.large",
    keyName: keyPair.keyName,
});

const sshConnInfo: command.types.input.remote.ConnectionArgs = {
    host: instance.publicDns,
    user: "admin",
    privateKey: sshKey.privateKeyOpenssh,
};

const backendSetup = new svmkit.Backend("backendSetup", {
    connection: sshConnInfo,
    publicDNSName: instance.publicDns,
    triggers: [instance.arn],
});

export const PUBLIC_DNS_NAME = instance.publicDns;
export const SSH_PRIVATE_KEY = sshKey.privateKeyOpenssh;
