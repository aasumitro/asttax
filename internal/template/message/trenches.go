package message

import (
	"fmt"

	"github.com/aasumitro/asttax/internal/common"
)

func TrenchesTextBody(state string) string {
	text := ``
	switch state {
	case common.TrenchesNewPairs:
		text += fmt.Sprintf(
			`%s <b>New Pairs</b> | Recently launched tokens.

$LEMON8 | Lemon8 — 16s ago
Progress: 6.45%% | <a href="#">WEB</a> 
▰▰▱▱▱▱▱▱▱▱▱▱▱▱▱▱▱▱▱▱
New app that will integrate with TikTok application for users in the U...
TH: — | DEV: 5.11%% | H: 3
Vol: $0.51K | MC: $7.42K
<a href="https://t.me/@AsttaXBot?quick_buy=">Quick Buy</a> | <a href="#">View Chart</a>
`,
			common.PlantEmoticon)
	case common.TrenchesIgnitingEngines:
		text += fmt.Sprintf(
			`%s Igniting Engines | Tokens that are close to migrating, or in the process of migrating to Raydium.

$SUW | Swim Under Water — 1d ago
Progress: 96.49%% | <a href="https://www.suw.com/">WEB</a> | <a href="https://x.com/SUW">X</a> | <a href="#">TG</a>
▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▱
The children enjoyed diving and looking for shells.
TH: 7.77%% | DEV: — | H: 247
Vol: $32.83K | MC: $80.25K
<a href="https://t.me/@AsttaXBot?buy=token_id">Buy</a> | <a href="#">Snipe Migration</a> | <a href="#">📊</a>
`,
			common.FireEmoticon)
	case common.TrenchesGraduated:
		text += fmt.Sprintf(
			`%s Graduated | Tokens that have completed their migration to Raydium.

$LIA | lia — 1h ago
Bonded: 51m ago | <a href="http://liagirl.net/">WEB</a> | <a href="https://x.com/CaseyliaAnge">X</a> | <a href="https://vvaifu.fun/">TG</a>
a 17 year old mature teen girl- Made with vvaifu.fun Create AI charact...
TH: — | DEV: — | H: 114
Vol: $2.05M | MC: $4.81K
<a href="https://t.me/@AsttaXBot?quick_buy=token_id">Quick Buy</a> | <a href="https://dexscreener.com/solana/token_id">View Chart</a>

`,
			common.RocketEmoticon)
	}
	return text
}
