# -*- mode: shell-script -*-
# shellcheck shell=bash
source "./lib.bash"

step::00::setup-abklabs-apt() {
  apt::abk
}

step::10::install-base-software() {
  $SUDO apt-get update
  $APT install logrotate ufw
}

step::00::create-sol-user() {
  id sol >/dev/null 2>&1 || $SUDO adduser --disabled-password --gecos "" sol
  $SUDO mkdir -p "/home/sol"
  $SUDO chown -R sol:sol "/home/sol"
  $SUDO chmod -R 755 "/home/sol"
}

step::10::copy-validator-keys() {
  keypairs::write "sol" "${IDENTITY_KEYPAIR}:validator-keypair" "${VOTE_ACCOUNT_KEYPAIR}:vote-account-keypair"
}

step::15::configure-sysctl() {
  cat <<EOF | $SUDO tee /etc/sysctl.d/21-solana-validator.conf >/dev/null
# Increase UDP buffer sizes
net.core.rmem_default = 134217728
net.core.rmem_max = 134217728
net.core.wmem_default = 134217728
net.core.wmem_max = 134217728
# Increase memory mapped files limit
vm.max_map_count = 1000000
# Increase number of allowed open file descriptors
fs.nr_open = 1000000
vm.swappiness=1
EOF

  $SUDO sysctl -p /etc/sysctl.d/21-solana-validator.conf
}

step::30::configure-firewall() {
  $SUDO ufw allow 53
  $SUDO ufw allow ssh
  $SUDO ufw allow 8000:8020/tcp
  $SUDO ufw allow 8000:8020/udp
  # TODO: Only open for RPC nodes
  $SUDO ufw allow 8899/tcp
  $SUDO ufw allow 8899/udp
  $SUDO ufw --force enable
}

step::35::setup-logrotate() {
  cat <<EOF | $SUDO tee /etc/logrotate.d/solana >/dev/null
/home/sol/solana-validator.log {
su sol sol
daily
rotate 1
missingok
postrotate
    systemctl kill -s USR1 sol.service
endscript
}
EOF

  $SUDO systemctl restart logrotate
}

step::40::install-validator() {
  $APT install zuma-agave-validator zuma-solana-cli
}

step::50::setup-validator-startup() {
  if systemctl list-unit-files svmkit-agave-validator.service >/dev/null; then
    $SUDO systemctl stop svmkit-agave-validator.service || true
  fi

  cat <<EOF | $SUDO tee /home/sol/run-validator >/dev/null
#!/usr/bin/env bash
exec agave-validator $VALIDATOR_FLAGS
EOF
  $SUDO chmod 755 /home/sol/run-validator
  $SUDO chown sol:sol /home/sol/run-validator

  cat <<EOF | $SUDO tee /etc/systemd/system/svmkit-agave-validator.service >/dev/null
[Unit]
Description=SVMkit Agave validator

[Service]
User=sol
Group=sol
ExecStart=/home/sol/run-validator
LimitNOFILE=1000000

[Install]
WantedBy=default.target
EOF
  $SUDO systemctl daemon-reload
  $SUDO systemctl enable svmkit-agave-validator.service
  $SUDO systemctl start svmkit-agave-validator.service
}
