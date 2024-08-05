#!/usr/bin/env bash

set -euo pipefail

usage() {
	echo "Usage: $0 [-v] [all|incremental]"
	echo "  -v          : be verbose"
	echo "  all         : Run diff tests and regen what fails"
	echo "  incremental : Use rdeps"
	exit 1
}

VERBOSE=""

all() {
	test_list_file=$(mktemp)
	test_errors_file=$(mktemp)
 	bazel query 'kind(diff_test, //lib/proto/...)' 2>/dev/null > "$test_list_file"

	if ! bazel test --test_summary none --test_output errors --target_pattern_file "$test_list_file" >"$test_errors_file" 2>/dev/null; then
		while IFS= read -r line; do
			if [[ ! -z "$VERBOSE" ]]; then
				echo "$line"
			fi
			eval "$line"
		done <<< "$(cat "$test_errors_file" | grep 'bazel run //lib/proto/' | perl -pe 's/^\s+//g' | sort | uniq)"
	fi
}

incremental() {
	echo inc
}

while getopts "v" opt; do
	case ${opt} in
		v)
			VERBOSE=1
			;;
        *)
            usage
            ;;
    esac
done

shift $((OPTIND -1))

cmd="${1:-incremental}"

case "$cmd" in
	all)
		all
		;;
	incremental)
		incremental
		;;
	*)
		usage
		;;
esac
