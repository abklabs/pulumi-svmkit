"""An SSH Python Pulumi program"""

import pulumi
import pulumi_tls as tls
import pulumi_svmkit as svmkit

ssh_key = """-----BEGIN OPENSSH PRIVATE KEY-----
-----END OPENSSH PRIVATE KEY-----"""

validator_key = svmkit.KeyPair("validator-key")
vote_account_key = svmkit.KeyPair("vote-account-key")

connection = {
    "host": "host.name.com",
    "user": "admin",
    "privateKey": ssh_key,
}

svmkit.validator.Agave(
    "validator",
    connection=connection,
    version="1.18.24-1",
    key_pairs={
        "identity": validator_key.json,
        "voteAccount": vote_account_key.json,
    },
    flags={
        "entryPoint": [
            "entrypoint.testnet.solana.com:8001",
            "entrypoint2.testnet.solana.com:8001",
            "entrypoint3.testnet.solana.com:8001",
        ],
        "knownValidator": [
            "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
            "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
            "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
            "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
        ],
        "expectedGenesisHash": "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
        "useSnapshotArchivesAtStartup": "when-newest",
        "rpcPort": 8899,
        "privateRPC": True,
        "onlyKnownRPC": True,
        "dynamicPortRange": "8002-8020",
        "gossipPort": 8001,
        "rpcBindAddress": "0.0.0.0",
        "walRecoveryMode": "skip_any_corrupted_record",
        "limitLedgerSize": 50000000,
        "blockProductionMethod": "central-scheduler",
        "fullSnapshotIntervalSlots": 1000,
        "noWaitForVoteToStartLeader": True,
    },
#    opts=pulumi.ResourceOptions(depends_on=[instance])
)

pulumi.export("PUBLIC_DNS_NAME", pulumi.Output.format(connection["host"]))
pulumi.export("SSH_PRIVATE_KEY", pulumi.Output.format(connection["privateKey"]))
