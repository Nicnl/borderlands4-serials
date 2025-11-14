package codex

import (
	"borderlands_4_serials/b4s/b85"
	"borderlands_4_serials/b4s/serial"
	"borderlands_4_serials/b4s/serial_datatypes/part"
	"borderlands_4_serials/b4s/serial_tokenizer"
	"fmt"
	"testing"
	"time"
)

func TestDeserializeBenchLagAttack(t *testing.T) {
	const (
		numIterations = 25
		numDeserials  = 100
		bigSerial     = "@Ugr$xKm/)}}!f;8TMir`*VL&qs4O(Gn&>q8p&X=I`jRT!87wCMqK<7&ZI^QeM`TBv*HxqQem!R{71f6d-=zKKVM4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+CfY=sXcKLsO/*$N(I(nNn`jelqD{1kHqj>9M4M<6ZK6%Ii8j$D+C-aZ6K$eRw23y+1^@"
	)

	measurements := make([]time.Duration, 0, numIterations)

	for range numIterations {
		start := time.Now()
		for range numDeserials {

			//data, _ := b85.Decode(bigSerial)
			//_, _, _ = Deserialize(data)
			Deserialize(bigSerial, true)
		}
		duration := time.Since(start)
		fmt.Println("4000:", duration)
		measurements = append(measurements, duration)
	}

	total := time.Duration(0)
	for _, d := range measurements {
		total += d
	}
	avg := total / time.Duration(len(measurements))
	fmt.Println("Average for 4000 deserializations:", avg)
}

func TestCodesMachin(t *testing.T) {
	// Collect all parts
	mapParts := make(map[string]part.Part)
	for _, jsonItem := range Codex.JsonItems {
		pos := 1
		for {
			p := jsonItem.Item.FindPartAtPos(pos, true)
			pos++
			if p == nil {
				break
			}

			if p.SubType != part.SUBTYPE_NONE {
				continue
			}

			mapParts[p.String()] = *p
		}
	}

	// Group parts by barrel base
	baseToParts := make(map[BaseBarrel]map[uint32]bool)
	for _, jsonItem := range Codex.JsonItems {
		base, found := jsonItem.Item.BaseBarrel()
		if !found {
			//fmt.Println("no base barrel found for " + jsonItem.Name)
			continue
		}

		if _, exists := baseToParts[base.BaseBarrel]; !exists {
			baseToParts[base.BaseBarrel] = make(map[uint32]bool)
		}

		pos := 1
		for {
			p := jsonItem.Item.FindPartAtPos(pos, true)
			pos++
			if p == nil {
				break
			}

			if p.SubType != part.SUBTYPE_NONE {
				continue
			}

			if p.Index == base.BaseBarrel.BarrelIndex {
				// Skip the barrel part itself
				continue
			}

			if _, exists := baseToParts[base.BaseBarrel][p.Index]; !exists {
				baseToParts[base.BaseBarrel][p.Index] = true
			}
		}
	}

	combinations := 0
	fmt.Println("Total parts", len(mapParts))
	fmt.Println("Total bases", len(baseToParts))
	fmt.Println()
	fmt.Print("ALl parts =")
	for _, part := range mapParts {
		fmt.Print(" ", part.String())
	}
	fmt.Println()

	fmt.Println()
	for baseBarrel, parts := range baseToParts {
		combinations += len(parts)
		infos := Barrels[baseBarrel]
		fmt.Println("Base:", infos.Name, infos.BaseBarrel.ManufacturerIndex, infos.BaseBarrel.BaseIndex, infos.BaseBarrel.BarrelIndex)

		generatedSerials := make([]string, 0)
		for partIndex := range parts {
			encoded := b85.Encode(serial.Serialize([]serial.Block{
				{Token: serial_tokenizer.TOK_VARINT, Value: baseBarrel.ManufacturerIndex},
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 0}, // Unknown, always zero
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 1}, // Unknown, always one before the level
				{Token: serial_tokenizer.TOK_SEP2},
				{Token: serial_tokenizer.TOK_VARINT, Value: 50}, // Level 50
				{Token: serial_tokenizer.TOK_SEP1},
				{Token: serial_tokenizer.TOK_SEP1},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BaseIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BarrelIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: partIndex, SubType: part.SUBTYPE_NONE}},
				{Token: serial_tokenizer.TOK_SEP1},
			}))
			generatedSerials = append(generatedSerials, encoded)
		}

		_serialsToYaml(generatedSerials)
	}

	fmt.Println("Total combinations:", combinations)
}

func TestAddSamePartALlBases(t *testing.T) {
	generatedSerials := make([]string, 0)
	for baseBarrel, _ := range Barrels {

		encoded := b85.Encode(serial.Serialize([]serial.Block{
			{Token: serial_tokenizer.TOK_VARINT, Value: baseBarrel.ManufacturerIndex},
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 0}, // Unknown, always zero
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 1}, // Unknown, always one before the level
			{Token: serial_tokenizer.TOK_SEP2},
			{Token: serial_tokenizer.TOK_VARINT, Value: 50}, // Level 50
			{Token: serial_tokenizer.TOK_SEP1},
			{Token: serial_tokenizer.TOK_SEP1},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BaseIndex, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: baseBarrel.BarrelIndex, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: 25, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_PART, Part: part.Part{Index: 29, SubType: part.SUBTYPE_NONE}},
			{Token: serial_tokenizer.TOK_SEP1},
		}))
		generatedSerials = append(generatedSerials, encoded)
	}

	_serialsToYaml(generatedSerials)
}
