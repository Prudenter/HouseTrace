#!/bin/bash
rm -rf channel-artifacts
rm -rf crypto-config
mkdir channel-artifacts

cryptogen generate --config=crypto-config.yaml



configtxgen -profile genesis -outputBlock ./channel-artifacts/genesis.block

# 生成 unionchannel 文件
 echo "---------------- Create allorgschannel.tx file BEGIN -------------------"
 configtxgen -profile allOrgsChannel -outputCreateChannelTx ./channel-artifacts/allorgschannel.tx -channelID allorgschannel
 echo "---------------- Create allorgschannel.tx file END -------------------"

# 生成 tgjchannel
 echo "---------------- Create tgjorgschannel.tx file BEGIN -------------------"
 configtxgen -profile tgjOrgsChannel -outputCreateChannelTx ./channel-artifacts/tgjorgschannel.tx -channelID tgjorgschannel
 echo "---------------- Create tgjorgschannel.tx file END -------------------"

# 生成 fgjchannel
 echo "---------------- Create fgjorgschannel.tx file BEGIN -------------------"
 configtxgen -profile fgjOrgsChannel -outputCreateChannelTx ./channel-artifacts/fgjorgschannel.tx -channelID fgjorgschannel
 echo "---------------- Create fgjorgschannel.tx file END -------------------"

# 生成 zgptchannel
 echo "---------------- Create zfptorgschannel.tx file BEGIN -------------------"
 configtxgen -profile zfptOrgsChannel -outputCreateChannelTx ./channel-artifacts/zfptorgschannel.tx -channelID zfptorgschannel
 echo "---------------- Create zfptorgschannel.tx file END -------------------"
