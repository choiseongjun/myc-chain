#!/bin/bash

# Go bin 경로를 PATH에 추가
export PATH=$PATH:$HOME/go/bin

# 체인 설정
export CHAIN_ID="mycchain"
export NODE="http://localhost:26657"

echo "Environment setup complete!"
echo "PATH에 $HOME/go/bin 추가됨"
echo ""
echo "이제 myc-chaind 명령어를 바로 사용할 수 있습니다:"
echo "  myc-chaind version"
echo ""
echo "또는 전체 경로로 사용:"
echo "  $HOME/go/bin/myc-chaind version"
