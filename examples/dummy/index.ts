import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as tls from "@pulumi/tls";
import * as command from "@pulumi/command";
import * as svmkit from "@pulumi/svmkit";

const config = new pulumi.Config();
const ami = config.require("ami");

const sshKey = new tls.PrivateKey("sshKey", {
    algorithm: "ED25519",
});

const keyPair = new aws.ec2.KeyPair("keypair", {
    publicKey: sshKey.publicKeyOpenssh,
});

const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {
    description: "Allow SSH",
    ingress: [
        {
            protocol: "tcp",
            fromPort: 22,
            toPort: 22,
            cidrBlocks: ["0.0.0.0/0"],
        },
    ],
    egress: [
        {
            protocol: "-1",
            fromPort: 0,
            toPort: 0,
            cidrBlocks: ["0.0.0.0/0"],
        },
    ],
});

const voteAccount = new svmkit.KeyPair("voteAccount");
const identity = new svmkit.KeyPair("identity");

const instance = new aws.ec2.Instance("instance", {
    ami,
    instanceType: "t3.small",
    keyName: keyPair.keyName,
    vpcSecurityGroupIds: [securityGroup.id],
});

const host = instance.publicDns.apply(dns => dns.toString());

const adminConnection: command.types.input.remote.ConnectionArgs = {
    host,
    user: "admin",
    privateKey: pulumi.output(sshKey.privateKeyOpenssh).apply(key => key.toString()),
};

const wait = new svmkit.Wait("wait", {
    connection: adminConnection,
}, { dependsOn: [instance] });


const sol = new svmkit.User("sol", {
    connection: adminConnection,
    username: "sol",
}, { dependsOn: [wait] });

const solConnection: command.types.input.remote.ConnectionArgs = {
    host,
    user: sol.name.apply(name => name.toString()),
    privateKey: pulumi.output(sol.privKey).apply(key => key.toString()),
};


const validator = new svmkit.Validator("validator", {
    connection: solConnection,
    voteAccount: voteAccount.json,
    identity: identity.json,
}, { dependsOn: [sol] });

export const publicDnsName = instance.publicDns;
export const sshPrivateKey = sshKey.privateKeyOpenssh;
export const solSshPrivateKey = sol.privKey;   