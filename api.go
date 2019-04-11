package minter_node_go_api

import (
	"encoding/json"
	"fmt"
	"github.com/MinterTeam/minter-explorer-tools/models"
	"github.com/MinterTeam/minter-node-go-api/responses"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
)

type MinterNodeApi struct {
	link string
}

func New(link string) *MinterNodeApi {
	return &MinterNodeApi{
		link: link,
	}
}

func (api *MinterNodeApi) SetLink(link string) {
	api.link = link
}

func (api *MinterNodeApi) GetLink() string {
	return api.link
}

func (api *MinterNodeApi) GetStatus() (*responses.StatusResponse, error) {
	response := responses.StatusResponse{}
	link := api.link + `/status`
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlock(height uint64) (*responses.BlockResponse, error) {
	response := responses.BlockResponse{}
	link := api.link + `/block?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	if response.Result.TxCount != "0" {
		for i, tx := range response.Result.Transactions {
			switch tx.Type {
			case models.TxTypeSend:
				var txData = models.SendTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeSellCoin:
				var txData = models.SellCoinTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeSellAllCoin:
				var txData = models.SellAllCoinTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeBuyCoin:
				var txData = models.BuyCoinTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeCreateCoin:
				var txData = models.CreateCoinTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeDeclareCandidacy:
				var txData = models.DeclareCandidacyTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeDelegate:
				var txData = models.DelegateTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeUnbound:
				var txData = models.UnbondTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeRedeemCheck:
				var txData = models.RedeemCheckTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeSetCandidateOnline:
				var txData = models.SetCandidateTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeSetCandidateOffline:
				var txData = models.SetCandidateTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeMultiSig:
				var txData = models.CreateMultisigTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeMultiSend:
				var txData = models.MultiSendTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			case models.TxTypeEditCandidate:
				var txData = models.EditCandidateTxData{}
				err = json.Unmarshal(tx.RawData, &txData)
				response.Result.Transactions[i].Data = txData
			}
		}
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlockEvents(height uint64) (*responses.EventsResponse, error) {
	response := responses.EventsResponse{}
	link := api.link + `/events?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlockValidators(height uint64) (*responses.ValidatorsResponse, error) {
	response := responses.ValidatorsResponse{}
	link := api.link + `/validators?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCandidate(pubKey string, height uint64) (*responses.CandidateResponse, error) {
	response := responses.CandidateResponse{}
	link := api.link + `/candidate?pubkey=` + pubKey + `&height=` + strconv.Itoa(int(height))
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCandidates(height uint64, stakes bool) (*responses.BlockCandidatesResponse, error) {
	response := responses.BlockCandidatesResponse{}
	link := api.link + `/candidates?height=` + strconv.Itoa(int(height))

	if stakes {
		link += `&include_stakes=true`
	}

	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCoinInfo(symbol string) (*responses.CoinInfoResponse, error) {
	response := responses.CoinInfoResponse{}
	link := api.link + `/coin_info?symbol=` + symbol
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetAddress(address string) (*responses.AddressResponse, error) {
	response := responses.AddressResponse{}
	link := api.link + `/address?address=` + strings.Title(address)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetAddresses(addresses []string, height uint64) (*responses.BalancesResponse, error) {
	response := responses.BalancesResponse{}
	queryStr := "[" + strings.Join(addresses, ",") + "]"
	link := api.link + `/addresses?addresses=` + queryStr + `&height=` + strconv.Itoa(int(height))
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateTx(tx string) (*responses.EstimateTxResponse, error) {
	response := responses.EstimateTxResponse{}
	link := api.link + `/estimate_tx_commission?tx=` + tx
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateCoinBuy(coinToSell string, coinToBuy string, value string) (*responses.EstimateCoinBuyResponse, error) {
	response := responses.EstimateCoinBuyResponse{}
	link := api.link + `/estimate_coin_buy?coin_to_sell=` + coinToSell + `&coin_to_buy=` + coinToBuy + `&value_to_buy=` + value
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateCoinSell(coinToSell string, coinToBuy string, value string) (*responses.EstimateCoinSellResponse, error) {
	response := responses.EstimateCoinSellResponse{}
	link := api.link + `/estimate_coin_sell?coin_to_sell=` + coinToSell + `&coin_to_buy=` + coinToBuy + `&value_to_sell=` + value
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateCoinSellAll(coinToSell string, coinToBuy string, value string, gasPrice string) (*responses.EstimateCoinSellAllResponse, error) {
	response := responses.EstimateCoinSellAllResponse{}
	link := api.link + `/estimate_coin_sell_all?coin_to_sell=` + coinToSell + `&coin_to_buy=` + coinToBuy + `&value_to_sell=` + value + `&gas_price=` + gasPrice
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetMinGasPrice() (*responses.GasResponse, error) {
	response := responses.GasResponse{}
	link := api.link + `/min_gas_price`
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) PushTransaction(tx string) (*responses.SendTransactionResponse, error) {
	response := responses.SendTransactionResponse{}
	link := api.link + `/send_transaction?tx=0x` + tx
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) getJson(url string, target interface{}) error {
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}