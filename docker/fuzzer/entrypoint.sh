# cd "$(dirname "$0")"
# echo $HOME
# NEED_INST=$1

GOMOD_DIR=$1
OUT_DIR=$2
# shift 3
# if [ "$NEED_INST" = true ]; then
/gfuzz/bin/inst --check-syntax-err --recover-syntax-err --dir $GOMOD_DIR
# fi
# Start fuzzing
/gfuzz/bin/fuzzer --gomod $GOMOD_DIR --out $OUT_DIR
