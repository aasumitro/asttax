package message

import "fmt"

func HelpTextBody() string {
	youtubePlaylistURL := "https://www.youtube.com/watch?v=HavGDGUTmgs&list=PLmAMfj0qP2wwfnuRJQge2ss4sJxnhIqyt&ab_channel=Solandy%5Bsolandy.sol%5D"
	return fmt.Sprintf(`
<b><u>How do I use AsttaX?</u></b>
Check out our <a href="%s">Youtube playlist</a> where we explain it all.

<b><u>Which tokens can I trade?</u></b>
Any SPL token that is tradeable via Jupiter, including SOL and USDC pairs.

<b><u>My transaction timed out. What happened?</u></b>
Transaction timeouts can occur when there is heavy network load or instability. This is simply the nature of the current Solana network.

<b><u>What are the fees for using AsttaX?</u></b>
Transactions through AsttaX incur a fee of 1%%. We don't charge a subscription fee or pay-wall any features.

<b><u>My net profit seems wrong, why is that?</u></b>
The net profit of a trade takes into consideration the trade's transaction fees. Confirm the details of your trade on Solscan.io to verify the net profit.

<b><u>Additional questions or need support?</u></b>
send us mail hello@astta.xyz and one of our admins can assist you.
`, youtubePlaylistURL)
}
