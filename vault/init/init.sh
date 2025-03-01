#!/bin/sh

# Set environment variables
export VAULT_ADDR=https://vault-greenlight:8200
export VAULT_SKIP_VERIFY=true

echo "Starting Vault initialization..."
echo "Using Vault address: $VAULT_ADDR"

# Sleep to ensure Vault is fully ready
sleep 10

# Initialize Vault
echo "Initializing Vault..."
INIT_OUTPUT=$(vault operator init -key-shares=1 -key-threshold=1)
echo "$INIT_OUTPUT"

# Extract keys using string manipulation (no jq dependency)
UNSEAL_KEY=$(echo "$INIT_OUTPUT" | grep "Unseal Key 1" | awk '{print $NF}')
ROOT_TOKEN=$(echo "$INIT_OUTPUT" | grep "Initial Root Token" | awk '{print $NF}')

# Save keys to file
#echo "Unseal Key: $UNSEAL_KEY" > /vault/init/vault-keys.txt
#echo "Root Token: $ROOT_TOKEN" >> /vault/init/vault-keys.txt


echo "UNSEAL_KEY=$UNSEAL_KEY" > /vault/.env
echo "ROOT_TOKEN=$ROOT_TOKEN" >> /vault/.env

# Unseal Vault
echo "Unsealing Vault with key: $UNSEAL_KEY"
vault operator unseal $UNSEAL_KEY


# Login with root token
echo "Logging in with root token..."
export VAULT_TOKEN=$ROOT_TOKEN
vault login $ROOT_TOKEN

# Enable KV version 2 secrets engine
echo "Enabling KV version 2 secrets engine at path 'secret/'..."
vault secrets enable -path=secret kv-v2

# Add key-value data
echo "Adding key-value data to secret/greenlight/dbconfig..."
vault kv put secret/greenlight/dbconfig username="greenlight" password="pa55word" dbhost="localhost"

# Verify data was added successfully
echo "Verifying data was added successfully..."
vault kv get secret/greenlight/dbconfig

echo "Vault initialization and configuration complete!"

echo "Vault initialization complete!"
