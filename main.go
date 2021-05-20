package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Ticker struct {
	Ticker      string `json:"ticker"`
	Companyname string `json:"companyName"`
	Sector      string `json:"sector"`
}

type Tickers struct {
	Response float64  `json:"response"`
	Error    string   `json:"error"`
	Message  []Ticker `json:"message"`
}

type Price struct {
	Response int    `json:"response"`
	Error    string `json:"error"`
	Message  struct {
		Ticker           string    `json:"ticker"`
		Company          string    `json:"company"`
		Latestprice      float64   `json:"latestPrice"`
		Pointchange      float64   `json:"pointChange"`
		Percentagechange float64   `json:"percentageChange"`
		Timestamp        time.Time `json:"timestamp"`
		Wtavgprice       float64   `json:"wtAvgPrice"`
		Sharestraded     int       `json:"sharesTraded"`
		Volume           int       `json:"volume"`
		Mktcap           int64     `json:"mktCap"`
	} `json:"message"`
}

type IncomeStatement struct {
	Response float64 `json:"response"`
	Error    string  `json:"error"`
	Message  struct {
		Ticker string `json:"ticker"`
		Data   []struct {
			Idcbincomestatement                              float64 `json:"idcbincomestatement"`
			Ticker                                           string  `json:"Ticker"`
			Year                                             string  `json:"Year"`
			Quarter                                          float64 `json:"Quarter"`
			Datasource                                       string  `json:"DataSource"`
			Statement                                        string  `json:"Statement"`
			Interestincome                                   float64 `json:"InterestIncome"`
			Interestexpense                                  float64 `json:"InterestExpense"`
			NetInterestincome                                float64 `json:"NetInterestIncome"`
			Feesincome                                       float64 `json:"FeesIncome"`
			Feesexpense                                      float64 `json:"FeesExpense"`
			Netfeesincome                                    float64 `json:"NetFeesIncome"`
			NetInterestfeeandcommissionincome                float64 `json:"NetInterestfeeandcommissionincome"`
			Nettradingincome                                 float64 `json:"NetTradingIncome"`
			Otheroperatingincome                             float64 `json:"OtherOperatingIncome"`
			Totaloperatingincome                             float64 `json:"TotalOperatingIncome"`
			Impairment                                       float64 `json:"Impairment"`
			Netopincome                                      float64 `json:"NetOpIncome"`
			Operatingexpense                                 float64 `json:"Operatingexpense"`
			Staffexpenses                                    float64 `json:"StaffExpenses"`
			Otheroperatingexpenses                           float64 `json:"OtherOperatingExpenses"`
			Depreciationandamortization                      float64 `json:"DepreciationandAmortization"`
			Operatingprofit                                  float64 `json:"OperatingProfit"`
			Nonoperatingincome                               float64 `json:"Nonoperatingincome"`
			Nonoperatingexpense                              float64 `json:"Nonoperatingexpense"`
			Profitbeforetax                                  float64 `json:"ProfitbeforeTax"`
			Incometax                                        float64 `json:"IncomeTax"`
			Currenttax                                       float64 `json:"CurrentTax"`
			Deferredtax                                      float64 `json:"DeferredTax"`
			Netprofitorloss                                  float64 `json:"NetProfitOrLoss"`
			Compincome                                       float64 `json:"CompIncome"`
			Totalcompincome                                  float64 `json:"TotalCompIncome"`
			Netprofitlossasperprofitorloss                   float64 `json:"NetprofitlossasperProfitorLoss"`
			Profitrequiredtobeapropriatedtostatutoryreserve  float64 `json:"ProfitrequiredtobeapropriatedtoStatutoryreserve"`
			Profitrequiredtobetransferredtoregulatoryreserve float64 `json:"ProfitrequiredtobetransferredtoRegulatoryreserve"`
			Freeprofit                                       float64 `json:"FreeProfit"`
			Prefsharediv                                     float64 `json:"PrefShareDiv"`
		} `json:"data"`
	} `json:"message"`
}

type Financial struct {
	Response float64 `json:"response"`
	Error    string  `json:"error"`
	Message  struct {
		Ticker string `json:"ticker"`
		Data   []struct {
			Idcbkeystats                           float64 `json:"idcbKeyStats"`
			Ticker                                 string  `json:"Ticker"`
			Year                                   string  `json:"Year"`
			Quarter                                float64 `json:"Quarter"`
			Datasource                             string  `json:"DataSource"`
			Totalrevenue                           float64 `json:"TotalRevenue"`
			Growthoverpriorperiod                  float64 `json:"GrowthOverPriorPeriod"`
			Grossprofit                            float64 `json:"GrossProfit"`
			Grossprofitmargin                      float64 `json:"GrossProfitMargin"`
			Netincome                              float64 `json:"NetIncome"`
			Netincomemargin                        float64 `json:"NetIncomeMargin"`
			Returnonasset                          float64 `json:"ReturnOnAsset"`
			Returnonequity                         float64 `json:"ReturnOnEquity"`
			Epsannualized                          float64 `json:"EpsAnnualized"`
			Reportedpeannualized                   float64 `json:"ReportedPeAnnualized"`
			Bookvaluepershare                      float64 `json:"BookValuePerShare"`
			Dividendpershare                       float64 `json:"DividendPerShare"`
			Capitalfundtorwa                       float64 `json:"CapitalFundToRwa"`
			Nonperformingloannpltototalloan        float64 `json:"NonPerformingLoanNplToTotalLoan"`
			Totalloanlossprovisiontototalnpl       float64 `json:"TotalLoanLossProvisionToTotalNpl"`
			Costoffunds                            float64 `json:"CostOfFunds"`
			Creditdepositratioaspernrbcalculations float64 `json:"CreditDepositRatioAsPerNrbCalculations"`
			Baserate                               float64 `json:"BaseRate"`
			NetInterestspread                      float64 `json:"NetInterestSpread"`
			Averageyield                           float64 `json:"AverageYield"`
			Outstandingshares                      float64 `json:"OutstandingShares,omitempty"`
			Tier1Capital                           float64 `json:"Tier1Capital,omitempty"`
			Tier2Capital                           float64 `json:"Tier2Capital,omitempty"`
			Totalcapital                           float64 `json:"TotalCapital,omitempty"`
			Shareprice                             float64 `json:"SharePrice,omitempty"`
			Marketcapitalization                   float64 `json:"MarketCapitalization,omitempty"`
			Mps                                    float64 `json:"Mps"`
			Netliquidasset                         float64 `json:"NetLiquidAsset"`
		} `json:"data"`
	} `json:"message"`
}

type KeyFinancialMetrics struct {
	Ticker              string  `json:"ticker"`
	LTP                 float64 `json:"ltp"`
	DiversionFromFair   float64 `json:"divesionFromFair"`
	PE                  float64 `json:"pe"`
	Eps                 float64 `json:"eps"`
	FairValue           float64 `json:"fairValue"`
	Bvps                float64 `json:"bvps"`
	Roa                 float64 `json:"roa"`
	Roe                 float64 `json:"roe"`
	NPL                 float64 `json:"npl"`
	Listedshares        float64 `json:"listedShares"`
	Reserves            float64 `json:"reserves"`
	Mktcap              float64 `json:"mktCap"`
	DistributableProfit float64 `json:"distributableProfit"`
	PaidUpCapital       float64 `json:"paidUpCapital"`
	DividendCapacity    float64 `json:"dividendCapacity"`
}

type BalanceSheet struct {
	Response float64 `json:"response"`
	Error    string  `json:"error"`
	Message  struct {
		Ticker string `json:"ticker"`
		Data   []struct {
			Idcbbalancesheet2                 float64 `json:"idcbbalancesheet2"`
			Ticker                            string  `json:"Ticker"`
			Year                              string  `json:"Year"`
			Quarter                           float64 `json:"Quarter"`
			Datasource                        string  `json:"DataSource"`
			Statement                         string  `json:"Statement"`
			Totalassets                       float64 `json:"TotalAssets"`
			Cashandbankbalance                float64 `json:"CashAndBankBalance"`
			Moneyatcallandshortnotice         float64 `json:"MoneyAtCallAndShortNotice"`
			Duefromnrb                        float64 `json:"DueFromNrb"`
			Placementwithbfis                 float64 `json:"PlacementwithBFIs"`
			Loanandadvances                   float64 `json:"LoanAndAdvances"`
			Loanandadvancestobfis             float64 `json:"LoanandadvancestoBFIs"`
			Loansandadvancestocustomers       float64 `json:"Loansandadvancestocustomers"`
			Investments                       float64 `json:"Investments"`
			Investmentsecurities              float64 `json:"InvestmentSecurities"`
			Investmentinsubsidiaries          float64 `json:"Investmentinsubsidiaries"`
			Investmentinassociates            float64 `json:"Investmentinassociates"`
			Investmentproperty                float64 `json:"Investmentproperty"`
			Otherassetstot                    float64 `json:"OtherAssetsTot"`
			Derivativefinancialinstruments    float64 `json:"DerivativeFinancialInstruments"`
			Othertradingassets                float64 `json:"OtherTradingAssets"`
			Currenttaxassets                  float64 `json:"CurrentTaxassets"`
			Deferredtaxassets                 float64 `json:"DeferredTaxAssets"`
			Otherassets                       float64 `json:"OtherAssets"`
			Goodwill                          float64 `json:"Goodwill"`
			Propequip                         float64 `json:"PropEquip"`
			Liabilities                       float64 `json:"Liabilities"`
			Borrowingstot                     float64 `json:"BorrowingsTot"`
			Duetobankandfinancialinstitutions float64 `json:"DuetoBankandFinancialInstitutions"`
			Duetonrb                          float64 `json:"DuetoNRB"`
			Borrowings                        float64 `json:"Borrowings"`
			Debtsecuritiesissued              float64 `json:"DebtSecuritiesIssued"`
			Deposits                          float64 `json:"Deposits"`
			Othliabilitiesprov                float64 `json:"OthLiabilitiesProv"`
			Derfinancialinstruments           float64 `json:"DerFinancialInstruments"`
			Provisions                        float64 `json:"Provisions"`
			Currenttaxliabilities             float64 `json:"CurrentTaxLiabilities"`
			Deferredtaxliabilities            float64 `json:"DeferredTaxLiabilities"`
			Subordinatedliabilites            float64 `json:"SubordinatedLiabilites"`
			Otherliabilities                  float64 `json:"OtherLiabilities"`
			Equity                            float64 `json:"Equity"`
			Paidupcapital                     float64 `json:"PaidUpCapital"`
			Reservesandsurplus                float64 `json:"ReservesAndSurplus"`
			Sharepremium                      float64 `json:"SharePremium"`
			Retainedearnings                  float64 `json:"RetainedEarnings"`
			Reserves                          float64 `json:"Reserves"`
			Noncontrollingfloat64erest        float64 `json:"NonControllingfloat64erest"`
			Totalcapitalandliabilities        float64 `json:"TotalCapitalAndLiabilities"`
			Ordinaryshares                    float64 `json:"OrdinaryShares"`
			Prefshares                        float64 `json:"PrefShares"`
		} `json:"data"`
	} `json:"message"`
}

type StockDetails struct {
	Response float64 `json:"response"`
	Error    string  `json:"error"`
	Message  struct {
		Keyfinancial struct {
			Ticker  string  `json:"ticker"`
			Year    string  `json:"year"`
			Quarter float64 `json:"quarter"`
			Data    []struct {
				Type         string  `json:"type"`
				Totalrevenue float64 `json:"totalRevenue"`
				Grossprofit  float64 `json:"grossProfit"`
				Netincome    float64 `json:"netIncome"`
				Eps          float64 `json:"eps"`
				Bvps         float64 `json:"bvps"`
				Roa          float64 `json:"roa"`
				Roe          float64 `json:"roe"`
			} `json:"data"`
		} `json:"keyFinancial"`
		Summary struct {
			Ticker           string  `json:"ticker"`
			Open             float64 `json:"open"`
			Avgvolume        float64 `json:"avgVolume"`
			Dayshigh         float64 `json:"daysHigh"`
			Dayslow          float64 `json:"daysLow"`
			Fiftytwoweekhigh float64 `json:"fiftyTwoWeekHigh"`
			Fiftytwoweeklow  float64 `json:"fiftyTwoWeekLow"`
			Listedshares     float64 `json:"listedShares"`
			Mktcap           float64 `json:"mktCap"`
			Epsdiluted       float64 `json:"epsDiluted"`
			Pediluted        float64 `json:"peDiluted"`
			Bvps             float64 `json:"bvps"`
			Beta             float64 `json:"beta"`
		} `json:"summary"`
	} `json:"message"`
}

func getStocks() Tickers {
	resp, err := http.Get("https://bizmandu.com/__stock/tickers/all")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var tickers Tickers

	err = json.Unmarshal(body, &tickers)

	if err != nil {
		fmt.Println("err", err)
	}
	return tickers
}

func getLatestPrice(ticker string) Price {
	resp, err := http.Get(fmt.Sprintf("https://bizmandu.com/__stock/tearsheet/header/?tkr=%s", ticker))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var price Price

	err = json.Unmarshal(body, &price)
	if err != nil {
		fmt.Println("err", err)
	}
	return price
}

func getIncomeStatement(ticker string) IncomeStatement {
	resp, err := http.Get(fmt.Sprintf("https://bizmandu.com/__stock/tearsheet/financial/incomeStatement/?tkr=%s", ticker))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var incomeStatement IncomeStatement

	err = json.Unmarshal(body, &incomeStatement)
	if err != nil {
		fmt.Println("err", err)
	}
	return incomeStatement
}

func getBalanceSheet(ticker string) BalanceSheet {
	resp, err := http.Get(fmt.Sprintf("https://bizmandu.com/__stock/tearsheet/financial/balanceSheet/?tkr=%s", ticker))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var balanceSheet BalanceSheet

	err = json.Unmarshal(body, &balanceSheet)
	if err != nil {
		fmt.Println("err", err)
	}
	return balanceSheet
}

func getFinancialDetails(ticker string) Financial {
	resp, err := http.Get(fmt.Sprintf("https://bizmandu.com/__stock/tearsheet/financial/keyStats/?tkr=%s", ticker))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	var financial Financial

	err = json.Unmarshal(body, &financial)
	if err != nil {
		fmt.Println("err", err)
	}
	return financial
}

func getStockDetails(ticker string) StockDetails {
	resp, err := http.Get(fmt.Sprintf("https://bizmandu.com/__stock/tearsheet/summary/?tkr=%s", ticker))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var detail StockDetails

	err = json.Unmarshal(body, &detail)
	if err != nil {
		fmt.Println("err", err)
	}

	return detail
}

func getBanks() []Ticker {
	tickers := getStocks()

	var ticks []Ticker

	for _, tick := range tickers.Message {
		if tick.Sector == "Commercial Banks" {
			if !strings.Contains(tick.Companyname, "Promoter") {
				if tick.Ticker != "JBNL" {
					ticks = append(ticks, Ticker{Ticker: tick.Ticker, Companyname: tick.Companyname, Sector: tick.Sector})
				}
			}

		}

	}
	return ticks
}

func CalculateGrahamValue(eps, bookValue float64) float64 {
	return math.Sqrt(22.5 * eps * bookValue)
}

func main() {
	banks := getBanks()
	var keys []KeyFinancialMetrics
	for _, bank := range banks {
		var key KeyFinancialMetrics

		detail := getStockDetails(bank.Ticker)
		financial := getFinancialDetails(bank.Ticker)
		balancesheet := getBalanceSheet(bank.Ticker)
		incomeStatement := getIncomeStatement(bank.Ticker)
		price := getLatestPrice(bank.Ticker)

		key.Eps = detail.Message.Summary.Epsdiluted
		key.PE = detail.Message.Summary.Pediluted
		key.LTP = detail.Message.Summary.Open
		key.Bvps = detail.Message.Summary.Bvps
		key.Ticker = detail.Message.Keyfinancial.Ticker
		key.Listedshares = detail.Message.Summary.Listedshares
		key.LTP = price.Message.Latestprice
		key.Mktcap = detail.Message.Summary.Mktcap

		for _, quarter := range detail.Message.Keyfinancial.Data {
			if quarter.Type == "CURRENT" {
				key.Roa = quarter.Roa * 100
				key.Roe = quarter.Roe * 100
			}
		}

		if len(financial.Message.Data) != 0 {
			key.NPL = financial.Message.Data[0].Nonperformingloannpltototalloan * 100
		}

		if len(balancesheet.Message.Data) != 0 {
			key.PaidUpCapital = float64(balancesheet.Message.Data[0].Paidupcapital)
			key.Reserves = float64(balancesheet.Message.Data[0].Reserves)
		}

		if len(incomeStatement.Message.Data) != 0 {
			key.DistributableProfit = float64(incomeStatement.Message.Data[0].Freeprofit)
		}

		if len(balancesheet.Message.Data) != 0 && len(incomeStatement.Message.Data) != 0 {
			key.DividendCapacity = (key.DistributableProfit / key.PaidUpCapital) * 100
		}
		key.FairValue = CalculateGrahamValue(key.Eps, key.Bvps)
		key.DiversionFromFair = ((key.LTP - key.FairValue) / (key.FairValue)) * 100

		keys = append(keys, key)
	}

	categories := map[string]string{
		"A1": "Ticker", "B1": "LTP", "C1": "%Fair", "D1": "P/E", "E1": "EPS", "F1": "FairValue",
		"G1": "BookValue", "H1": "ROA", "I1": "ROE", "J1": "NPL", "K1": "TotalShare", "L1": "Reserve",
		"M1": "MarketCap", "N1": "DisProfit", "O1": "paidUp", "P1": "ExepectedDividend",
	}
	var excelVals []map[string]interface{}

	for k, v := range keys {
		excelVal := map[string]interface{}{
			getColumn("A", k): v.Ticker, getColumn("B", k): v.LTP, getColumn("C", k): v.DiversionFromFair,
			getColumn("D", k): v.PE, getColumn("E", k): v.Eps, getColumn("F", k): v.FairValue, getColumn("G", k): v.Bvps,
			getColumn("H", k): v.Roa, getColumn("I", k): v.Roe, getColumn("J", k): v.NPL,
			getColumn("K", k): v.Listedshares, getColumn("L", k): v.Reserves, getColumn("M", k): v.Mktcap,
			getColumn("N", k): v.DistributableProfit, getColumn("O", k): v.PaidUpCapital, getColumn("P", k): v.DividendCapacity,
		}

		excelVals = append(excelVals, excelVal)
	}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}

	for _, vals := range excelVals {
		for k, v := range vals {
			f.SetCellValue("Sheet1", k, v)
		}
	}

	if err := f.SaveAs("Banking.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func getColumn(column string, num int) string {
	return fmt.Sprintf("%s%d", column, num+2)
}
