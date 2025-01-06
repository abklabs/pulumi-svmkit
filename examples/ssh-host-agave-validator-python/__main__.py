import pulumi
import pulumi_tls as tls
import pulumi_svmkit as svmkit

config = pulumi.Config('ssh')

validator_key = svmkit.KeyPair("validator-key")
vote_account_key = svmkit.KeyPair("vote-account-key")

connection = svmkit.ssh.ConnectionArgsDict({
    "host": config.require("host"),
    "user": config.require("user"),
})

ssh_private_key = config.get_secret("privateKey")

if ssh_private_key:
    connection['private_key'] = ssh_private_key

svmkit.validator.Agave(
    "validator",
    connection=connection,
    version="1.18.24-1",
    key_pairs={
        "identity": validator_key.json,
        "vote_account": vote_account_key.json,
    },
    flags={
        "entry_point": [
            "entrypoint.testnet.solana.com:8001",
            "entrypoint2.testnet.solana.com:8001",
            "entrypoint3.testnet.solana.com:8001",
        ],
        "known_validator": [
            "5D1fNXzvv5NjV1ysLjirC4WY92RNsVH18vjmcszZd8on",
            "7XSY3MrYnK8vq693Rju17bbPkCN3Z7KvvfvJx4kdrsSY",
            "Ft5fbkqNa76vnsjYNwjDZUXoTWpP7VYm3mtsaQckQADN",
            "9QxCLckBiJc783jnMvXZubK4wH86Eqqvashtrwvcsgkv",
        ],
        "expected_genesis_hash": "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY",
        "use_snapshot_archives_at_startup": "when-newest",
        "rpc_port": 8899,
        "private_rpc": True,
        "only_known_rpc": True,
        "dynamic_port_range": "8002-8020",
        "gossip_port": 8001,
        "rpc_bind_address": "0.0.0.0",
        "wal_recovery_mode": "skip_any_corrupted_record",
        "limit_ledger_size": 50000000,
        "block_production_method": "central-scheduler",
        "full_snapshot_interval_slots": 1000,
        "no_wait_for_vote_to_start_leader": True,
    }
)

pulumi.export("nodes_name", ["instance"])
pulumi.export("nodes_public_ip",
              [pulumi.Output.format(connection["host"])])
if "private_key" in connection:
    pulumi.export("nodes_private_key",
                    [pulumi.Output.format(connection["private_key"])])
