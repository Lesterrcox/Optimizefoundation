#!/bin/bash

# set -x

# Define variables
GENESIS_FILE="$HOME/.ogc/config/genesis.json"
TMP_GENESIS_FILE="./tmp-genesis-update.json"

# Load the updates JSON structure (if any other updates, add here)
UPDATES_FILE="./genesis-update.json"
assetList=$(jq '.assetList' $UPDATES_FILE)
assetCount=$(jq '.assetCount' $UPDATES_FILE)

validatorList=$(jq '.validatorList' $UPDATES_FILE)
validatorCount=$(jq '.validatorCount' $UPDATES_FILE)

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "Error: jq is required but not installed. Please install jq."
    exit 1
fi

updated_json=$(jq ".app_state.asset.assetList = ${assetList}"  $GENESIS_FILE)
updated_json=$(echo "$updated_json" | jq ".app_state.asset.assetCount = ${assetCount}")

updated_json=$(echo "$updated_json" | jq ".app_state.validator.validatorList = ${validatorList}")
updated_json=$(echo "$updated_json" | jq ".app_state.validator.validatorCount = ${validatorCount}")

# Write the updated JSON object to the temp file and move it to the original location
echo "$updated_json" > $TMP_GENESIS_FILE
mv $TMP_GENESIS_FILE $GENESIS_FILE
