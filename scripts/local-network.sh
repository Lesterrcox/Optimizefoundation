#!/bin/bash

# debug
# set -x

source ./common.sh
HOME_DIR=$HOME/.ogc

platform=$(get_platform)

update_test_genesis() {
    cat $HOME_DIR/config/genesis.json | jq "$1" >$HOME_DIR/config/tmp_genesis.json && mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
}

# Check if screen installed
check_screen_installed

# Set the sed command
sed_command="sed"

if [ "$platform" == "darwin" ]; then
  # Check if gsed is installed
  check_gsed_installed

  sed_command="gsed"
fi

# Check if any of the named screens already exist and prompt to stop them
prompt_to_stop_screens genesis

# Delete home folder
rm -rf $HOME_DIR

# Get the current working directory
current_dir=$(pwd)

UPDATE_GENESIS_SCRIPT="./update-genesis.sh"
# Check if update-genesis.sh exists
if [ ! -f "$UPDATE_GENESIS_SCRIPT" ]; then
    echo "Error: update-genesis.sh not found."
    exit 1
fi

echo "Initializing genesis node..."

ogcd init --chain-id ogc_testnet validator --home $HOME_DIR
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | ogcd keys add genesis --keyring-backend test --recover --home $HOME_DIR
echo "faculty head please solid picnic benefit hurt gloom flag transfer thrive zebra" | ogcd keys add test --keyring-backend test --recover --home $HOME_DIR
ogcd genesis add-genesis-account genesis 110000000000000uogc --keyring-backend test --home $HOME_DIR
ogcd genesis add-genesis-account test 10000000000000uogc --keyring-backend test --home $HOME_DIR
ogcd genesis gentx genesis 110000000000uogc --keyring-backend test --chain-id ogc_testnet
ogcd genesis collect-gentxs --home $HOME_DIR
$sed_command -i 's/stake/uogc/g' ~/.ogc/config/genesis.json
$sed_command -i 's/pruning = "default"/pruning = "custom"/g' $HOME/.ogc/config/config.toml
$sed_command -i 's/pruning-keep-recent = "0"/pruning-keep-recent = "2"/g' $HOME/.ogc/config/app.toml
$sed_command -i 's/pruning-interval = "0"/pruning-interval = "100"/g' $HOME/.ogc/config/app.toml
$sed_command -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0uogc"/g' $HOME/.ogc/config/app.toml
$sed_command -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOME/.ogc/config/config.toml
$sed_command -i 's/cors_allowed_origins\s*=\s*\[\]/cors_allowed_origins = ["*",]/g' $HOME/.ogc/config/config.toml
$sed_command -i 's/prometheus = false/prometheus = true/g' $HOME/.ogc/config/config.toml
$sed_command -i 's/timeout_commit = "5s"/timeout_commit = "4s"/g' $HOME/.ogc/config/config.toml

ogcd config keyring-backend test --home "$HOME/.ogc" >/dev/null 2>&1
ogcd config broadcast-mode sync --home "$HOME/.ogc" >/dev/null 2>&1

echo "Updating genesis paramsters."
# Update gov params
update_test_genesis '.app_state["gov"]["params"]["voting_period"]="300s"'
update_test_genesis '.app_state["gov"]["params"]["max_deposit_period"]="200s"'

cd $current_dir
# Execute update-genesis.sh
"$UPDATE_GENESIS_SCRIPT"

echo "Starting genesis node in a screen session..."

screen -dmS genesis ogcd start --rpc.laddr "tcp://0.0.0.0:26657"

# sleep 5s to wait ogc L1 RPC gets ready
echo "ogc L1 network started. sleep 5s to wait ogc blockchain rpc gets ready"
