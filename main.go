package main

import (
	//"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"math/big"

	"github.com/joho/godotenv"

	//"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botAPI := os.Getenv("BOT_API")
  	RPC := os.Getenv("RPC")

	bot, err := tgbotapi.NewBotAPI(botAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(RPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer conn.Close()

	// Instantiate the contract and display its name
	squeeth, err := NewSqueeth(common.HexToAddress("0x64187ae08781B09368e6253F9E94951243A493D5"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the SqueethController contract: %v", err)
	}

	oracle, err := NewOracle(common.HexToAddress("0x65D66c76447ccB45dAf1e8044e918fA786A483A1"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate the Oracle contract: %v", err)
	}
	_weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	quoteCurrency := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	ethQuoteCurrencyPool := common.HexToAddress("0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8")
	TWAP_PERIOD := 420
	/* wPowerPerpPool := common.HexToAddress("0x82c427AdFDf2d245Ec51D8046b41c4ee87F0d29C") */
	INDEX_SCALE := SciNotToBigInt("1e4")
	ONE_ONE := SciNotToBigInt("1e36")
	CR_DENOMINATOR := big.NewInt(2)
	CR_NUMERATOR := big.NewInt(3)
	/*alertFac1 := big.NewInt(130)
	ALERT_DENOMINATOR := big.NewInt(100)*/

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "MarkdownV2"

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			// TODO

		case "help":

			// TODO Add /list to BotFather

			// TODO Add desc for /list

			msg.Text = "/help \\- List available commands\n/track `address` \\- Track Opyn positions of an Ethereum address\n/untrack `address` \\- Untrack tracking positions of an address"

		case "list":

			// TODO List tracking addresses

		case "track":
			address := update.Message.CommandArguments()
			// TODO Address check
			// TODO Input validity check
			// TODO Duplication check
			// TODO Position check

			// TODO Add address to track list

			response := fmt.Sprintf("Address `%s` added to tracking list\\.", address)
			msg.Text = response

			// TODO Get Vault IDs under address
			// https://etherscan.io/token/0xa653e22a963ff0026292cc8b67941c0ba7863a38?a=92#readContract
			// tokenOfOwnerByIndex

			// TODO Add vault ID to DB

		case "untrack":
			address := update.Message.CommandArguments()
			// TODO Address check
			// TODO Untracks address
			msg.Text = "Address " + address + " added to tracking list."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	// Regular check and alert
	for {
		vaultId := big.NewInt(int64(92))

		// Get vault details
		session := &SqueethSession{
			Contract: squeeth,
			/*CallOpts: bind.CallOpts{
				Pending: true,
			},*/
		}
		vault, err := session.Vaults(vaultId)
		if err != nil {
			log.Fatalf("Failed to request vault details: %v", err)
		}
		log.Printf("Operator: " + vault.Operator.Hex() + "\nNftCollateralId: " + fmt.Sprint(vault.NftCollateralId) + "\nCollateralAmount: " + vault.CollateralAmount.String() + "\nShortAmount: " + vault.ShortAmount.String())

		// Get normalizationFactor
		normFac, err := session.NormalizationFactor()
		if err != nil {
			log.Fatalf("Failed to request normalization factor: %v", err)
		}

		// Get current ETH price
		oraSes := &OracleSession{
			Contract: oracle,
		}
		twap, err := oraSes.GetTwap(
			ethQuoteCurrencyPool,
			_weth,
			quoteCurrency,
			uint32(TWAP_PERIOD),
			true,
		)
		if err != nil {
			log.Fatalf("Failed to request ETH/USD TWAP: %v", err)
		}

		// Get total collateral
		// TODO Include Uniswap NFT as collateral, https://etherscan.io/address/0x64187ae08781b09368e6253f9e94951243a493d5#code, function _getEffectiveCollateral
		totalCollateral := vault.CollateralAmount

		// Calculate liquidatable debt value
		liqDebtValueInETH := new(big.Int).Div(
			new(big.Int).Mul(
				totalCollateral,
				CR_DENOMINATOR,
			),
			CR_NUMERATOR,
		)

		// Calculate liquidatable scaledEthPrice
		liqScaledEthPrice := new(big.Int).Div(
			new(big.Int).Mul(
				liqDebtValueInETH,
				ONE_ONE,
			),
			new(big.Int).Mul(
				vault.ShortAmount,
				normFac,
			),
		)

		// Calculate liquidatable ethPrice
		liqEthPrice := new(big.Int).Mul(liqScaledEthPrice, INDEX_SCALE)

		msg.Text = "Address: " + address + "\nLiquidation Price: " + liqEthPrice.String() + "\nCurrent ETH Price: " + twap.String()

		// Compare liquidatable price with current price
		/* alertPrice1 := new(big.Int).Div(new(big.Int).Mul(twap, alertFac1), ALERT_DENOMINATOR)
		alert1 := liqEthPrice.Cmp(alertPrice1)
		if alert1 >= 0 {
			// TODO Send alert
			msg := tgbotapi.NewMessage(chatId, "")
			msg.ParseMode = "MarkdownV2"
		} */
	}
}

func SciNotToBigInt(input string) *big.Int {
	flt, _, err := big.ParseFloat(input, 10, 0, big.ToNearestEven)
	if err != nil {
		log.Fatalf("Failed to parse scientific notation to big integer: %v", err)
	}
	i := new(big.Int)
	i, _ = flt.Int(i)
	return i
}
