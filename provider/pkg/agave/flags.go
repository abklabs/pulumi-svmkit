package agave

type ValidatorFlags struct {
	EntryPoint                   []string
	KnownValidator               []string
	UseSnapshotArchivesAtStartup string
	RpcPort                      string
	DynamicPortRange             string
	GossipPort                   string
	RpcBindAddress               string
	WalRecoveryMode              string
	Paths                        struct {
		Log      string
		Accounts string
		Ledger   string
	}
	LimitLedgerSize            int
	BlockProductionMethod      string
	TvuReceiveThreads          int
	FullSnapshotIntervalSlots  int
	NoWaitForVoteToStartLeader bool
	OnlyKnownRPC               bool
	PrivateRPC                 bool
}

func GenerateAgaveValidatorFlags(flags ValidatorFlags) []string {
	var l []string

	b := func(k string, v bool) {
		if v {
			l = append(l, "--"+k)
		}
	}

	s := func(k string, v interface{}) {
		if v != nil {
			l = append(l, "--"+k, v.(string))
		}
	}

	s("identity", "/home/sol/validator-keypair.json")
	s("vote-account", "/home/sol/vote-account-keypair.json")

	for _, v := range flags.EntryPoint {
		s("entrypoint", v)
	}

	for _, v := range flags.KnownValidator {
		s("known-validator", v)
	}

	s("use-snapshot-archives-at-startup", flags.UseSnapshotArchivesAtStartup)
	s("rpc-port", flags.RpcPort)
	s("dynamic-port-range", flags.DynamicPortRange)
	s("gossip-port", flags.GossipPort)
	s("rpc-bind-address", flags.RpcBindAddress)
	s("wal-recovery-mode", flags.WalRecoveryMode)
	s("log", flags.Paths.Log)
	s("accounts", flags.Paths.Accounts)
	s("ledger", flags.Paths.Ledger)
	s("limit-ledger-size", flags.LimitLedgerSize)
	s("block-production-method", flags.BlockProductionMethod)
	s("tvu-receive-threads", flags.TvuReceiveThreads)
	s("full-snapshot-interval-slots", flags.FullSnapshotIntervalSlots)

	b("no-wait-for-vote-to-start-leader", flags.NoWaitForVoteToStartLeader)
	b("only-known-rpc", flags.OnlyKnownRPC)
	b("private-rpc", flags.PrivateRPC)

	return l
}
