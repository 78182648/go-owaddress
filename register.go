package owaddress

import (
	"github.com/blocktree/go-owaddress/coins/ae"
	"github.com/blocktree/go-owaddress/coins/alc"
	"github.com/blocktree/go-owaddress/coins/atom"
	"github.com/blocktree/go-owaddress/coins/bnb"
	"github.com/blocktree/go-owaddress/coins/btc"
	"github.com/blocktree/go-owaddress/coins/btx"
	"github.com/blocktree/go-owaddress/coins/cxc"
	"github.com/blocktree/go-owaddress/coins/dgb"
	"github.com/blocktree/go-owaddress/coins/dsc"
	"github.com/blocktree/go-owaddress/coins/ela"
	"github.com/blocktree/go-owaddress/coins/eth"
	"github.com/blocktree/go-owaddress/coins/fiii"
	"github.com/blocktree/go-owaddress/coins/g50"
	"github.com/blocktree/go-owaddress/coins/hc"
	"github.com/blocktree/go-owaddress/coins/hss"
	"github.com/blocktree/go-owaddress/coins/ltc"
	"github.com/blocktree/go-owaddress/coins/moac"
	"github.com/blocktree/go-owaddress/coins/nas"
	"github.com/blocktree/go-owaddress/coins/ont"
	"github.com/blocktree/go-owaddress/coins/pess"
	"github.com/blocktree/go-owaddress/coins/qtum"
	"github.com/blocktree/go-owaddress/coins/sinoc"
	truechain "github.com/blocktree/go-owaddress/coins/true"
	"github.com/blocktree/go-owaddress/coins/trx"
	"github.com/blocktree/go-owaddress/coins/tv"
	"github.com/blocktree/go-owaddress/coins/vcc"
	"github.com/blocktree/go-owaddress/coins/vds"
	"github.com/blocktree/go-owaddress/coins/vlx"
	"github.com/blocktree/go-owaddress/coins/vsys"
	"github.com/blocktree/go-owaddress/coins/wicc"
	"github.com/blocktree/go-owaddress/coins/xrp"
	"github.com/blocktree/go-owaddress/coins/xvg"
	"github.com/blocktree/go-owaddress/coins/zen"
	"reflect"
)

var AddressVerifyRegistry = make(map[string]reflect.Type)

func RegisterAddressVerify(elem interface{}, coin string) {
	t := reflect.TypeOf(elem).Elem()
	AddressVerifyRegistry[coin] = t
}

func init() {
	RegisterAddressVerify(btc.DefaultStruct, btc.CoinName)
	RegisterAddressVerify(ltc.DefaultStruct, ltc.CoinName)
	RegisterAddressVerify(qtum.DefaultStruct, qtum.CoinName)
	RegisterAddressVerify(ont.DefaultStruct, ont.CoinName)
	RegisterAddressVerify(atom.DefaultStruct, atom.CoinName)
	RegisterAddressVerify(xrp.DefaultStruct, xrp.CoinName)
	RegisterAddressVerify(trx.DefaultStruct, trx.CoinName)
	RegisterAddressVerify(wicc.DefaultStruct, wicc.CoinName)
	RegisterAddressVerify(hc.DefaultStruct, hc.CoinName)
	RegisterAddressVerify(bnb.DefaultStruct, bnb.CoinName)
	RegisterAddressVerify(vsys.DefaultStruct, vsys.CoinName)
	RegisterAddressVerify(nas.DefaultStruct, nas.CoinName)
	RegisterAddressVerify(ela.DefaultStruct, ela.CoinName)
	RegisterAddressVerify(tv.DefaultStruct, tv.CoinName)
	RegisterAddressVerify(moac.DefaultStruct, moac.CoinName)
	RegisterAddressVerify(dsc.DefaultStruct, dsc.CoinName)
	RegisterAddressVerify(fiii.DefaultStruct, fiii.CoinName)
	RegisterAddressVerify(vds.DefaultStruct, vds.CoinName)
	RegisterAddressVerify(hss.DefaultStruct, hss.CoinName)
	RegisterAddressVerify(vlx.DefaultStruct, vlx.CoinName)
	RegisterAddressVerify(btx.DefaultStruct, btx.CoinName)
	RegisterAddressVerify(cxc.DefaultStruct, cxc.CoinName)
	RegisterAddressVerify(xvg.DefaultStruct, xvg.CoinName)
	RegisterAddressVerify(zen.DefaultStruct, zen.CoinName)
	RegisterAddressVerify(dgb.DefaultStruct, dgb.CoinName)
	RegisterAddressVerify(alc.DefaultStruct, alc.CoinName)
	RegisterAddressVerify(eth.DefaultStruct, eth.CoinName)
	RegisterAddressVerify(pess.DefaultStruct, pess.CoinName)
	RegisterAddressVerify(vcc.DefaultStruct, vcc.CoinName)
	RegisterAddressVerify(truechain.DefaultStruct, truechain.CoinName)
	RegisterAddressVerify(g50.DefaultStruct, g50.CoinName)
	RegisterAddressVerify(sinoc.DefaultStruct, sinoc.CoinName)
	RegisterAddressVerify(ae.DefaultStruct, ae.CoinName)
}
