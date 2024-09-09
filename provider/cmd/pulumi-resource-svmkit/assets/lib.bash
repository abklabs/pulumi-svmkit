if [ -z "$BASH_VERSION" ] ; then
    echo "FATAL: This library requires bash to function properly!"
    exit 1
fi

if [[ ${BASH_VERSINFO[0]} -lt 4 ]] ; then
    echo "FATAL: This library requires bash v4 or greater!"
    exit 1
fi

: "${SUDO:=sudo --preserve-env=DEBIAN_FRONTEND}"
: "${APT:=$SUDO apt-get -qy}"

# Disable interactive frontends by default.
DEBIAN_FRONTEND=noninteractive
export DEBIAN_FRONTEND

EXIT_FUNCS=()

exit::trap () {
    local func
    for func in "${EXIT_FUNCS[@]}"; do
	$func
    done
}

trap exit::trap EXIT

exit::trigger () {
    EXIT_FUNCS+=("$*")
}

TMPDIR=$(mktemp -d)
export TMPDIR

temp::cleanup () {
    log::debug cleaning up "$TMPDIR"...
    rm -rf "$TMPDIR"
}

exit::trigger temp::cleanup

# shellcheck disable=SC2120 # these options are optional.
temp::file () {
    mktemp -p "$TMPDIR" "$@"
}

CRED='' ; CGRN='' ; CYEL='' ; CBLU=''; CNONE=''

if [[ -t 1 ]] ; then
    CRED='\033[0;31m'; CGRN='\033[0;32m' ;
    CYEL='\033[0;32m' ; CBLU='\033[0;34m' ; CNONE='\033[0m'
fi

log::output () {
    local level
    level="$1"; shift

    printf "$level:\t%s\n" "$*" >&2
}

log::debug () {
    [[ -v DEBUG ]] || return 0

    log::output "${CBLU}DEBUG${CNONE}" "$@"
}

log::info () {
    log::output "${CGRN}INFO${CNONE}" "$@"
}

log::warn () {
    log::output "${CYEL}WARN${CNONE}" "$@"
}

log::error () {
    log::output "${CRED}ERROR${CNONE}" "$@"
}

log::fatal () {
    log::output "${CRED}FATAL${CNONE}" "$@"
    exit 1
}

step::run () {
    local prefix start name

    prefix=$1 ; shift

    start=""

    if [[ $# -gt 0 ]] ; then
	start="setup::$1"; shift
	log::warn "starting steps with $start..."
    fi

    while read -r name; do
	if [[ $name > $start || $name = "$start" ]] ; then
	    log::info "====> running $name..."
	    $name
	fi
    done < <(declare -F | grep "$prefix" | awk '{ print $3; }')
}

ssh::config-file () {
    local filename
    filename="$(temp::file)"
    cat > "$filename" <<EOF
# This is fine, because we're using public key crypto for our access.
LogLevel ERROR
UpdateHostKeys no
UserKnownHostsFile /dev/null
StrictHostKeyChecking off
EOF
    echo "$filename"
}

pulumi::output () {
    pulumi stack output "$@"
}

pulumi::config () {
    pulumi config "$@"
}

git::version () {
    git describe --tags --dirty 2>/dev/null || git rev-parse --short HEAD
}

env::kv () {
    for varname in "$@" ; do
	if [[ -v $varname ]] ; then
	    printf "%s='%s'\n" "$varname" "${!varname}"
	fi
    done
}

env::write () {
    local var fname mode

    varname=$1 ; shift
    fname=$1 ; shift

    if [[ -v $varname ]] ; then
	touch "$fname"

	if [[ $# -gt 0 ]] ; then
	    mode=$1 ; shift
	    chmod "$mode" "$fname"
	fi

	printf "%s" "${!varname}" > "$fname"
    fi
}

keypair::write () {
    env::write "$@" 600
}
