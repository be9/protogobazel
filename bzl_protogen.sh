#!/usr/bin/env bash

set -euo pipefail

usage() {
	echo "Usage: $0 [-v] [-b MERGE_BASE] [all|incremental]"
	echo "  -v            : be verbose"
	echo "  -b MERGE_BASE : defines merge base for incremental mode (default: origin/main)"
	echo "  all           : Run diff tests and regen what fails (default mode)"
	echo "  incremental   : Use rdeps"
	exit 1
}

VERBOSE=""
MERGE_BASE="origin/main"

all() {
	test_list_file=$(mktemp)
	test_errors_file=$(mktemp)
 	bazel query 'kind(diff_test, //lib/proto/...)' 2>/dev/null > "$test_list_file"

	[[ -n "$VERBOSE" ]] && echo "$(cat "$test_list_file" | wc -l | perl  -pe 's/^\s+//g') diff test(s) to be run"

	if ! bazel test --test_summary none --test_output errors --target_pattern_file "$test_list_file" >"$test_errors_file" 2>/dev/null; then
		while IFS= read -r line; do
			[[ -n "$VERBOSE" ]] && echo "$line"
			eval "$line"
		done <<< "$(cat "$test_errors_file" | grep 'bazel run //lib/proto/' | perl -pe 's/^\s+//g' | sort | uniq)"
	else
		[[ -n "$VERBOSE" ]] && echo "No relevant proto source updates"
	fi
}

incremental() {
	echo "Not implemented yet"
	exit 1
}

while getopts "vb:" opt; do
	case ${opt} in
		v)
			VERBOSE=1
			;;
		b)
			MERGE_BASE=${OPTARG}
			;;
        *)
            usage
            ;;
    esac
done

shift $((OPTIND -1))

case "${1:-all}" in
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
