#!/bin/zsh -ue
# @(#) MyTouch sample
# @(#) Author: @yumechi

FROMFILE="template"
EXECUTOR=${HOME}"/scripts/copy_template.go"

PARAM=()
for opt in "$@"; do
    case "${opt}" in
        "-n" ) FROMFILE="$2"; shift 2 ;; # @(#) touch 
        '--' | '-' )
            shift
            PARAM+=( "$@" )
            break
            ;;
        * ) if [[ -n "$1" ]] && [[ ! "$1" =~ ^-+ ]]; then
                PARAM+=( "$1" ); shift
            fi
            ;;
    esac
done

# if OUTPUTFILE is None, Go Scripts create file by timedate.
TOFILE="${PARAM}"; PARAM=("${PARAM[@]:1}")

function my_copy() {
    local prefix=$1
    local suffix=$2
    local inputfile=$3
    local outputfile=$4
    if [ -n "$outputfile" ]; then
        go run ${EXECUTOR} \
        -prefix $prefix \
        -suffix $suffix \
        -input $inputfile \
        -output $outputfile
    else
        go run ${EXECUTOR} \
        -prefix $prefix \
        -suffix $suffix \
        -input $inputfile
    fi
}

my_copy CPP .cpp $FROMFILE $TOFILE